package service

import (
	"context"
	"errors"

	"github.com/ariefsn/upwork/helper"
	"github.com/ariefsn/upwork/logger"
	"github.com/ariefsn/upwork/models"
	"github.com/ariefsn/upwork/notification"
	"github.com/ariefsn/upwork/validator"
)

type userService struct {
	scrapeService models.ScrapeService
	userRepo      models.UserRepository
	notif         notification.Notification
}

// GetIds implements models.UserService.
func (u *userService) GetIds(ctx context.Context) ([]string, error) {
	return u.userRepo.GetIds(ctx)
}

// SendDeleteToken implements models.UserService.
func (u *userService) SendDeleteToken(ctx context.Context, userID string) (*models.UserProfile, error) {
	err := validator.ValidateVar(userID, "required")
	if err != nil {
		return nil, err
	}

	exists, _ := u.Get(ctx, userID)

	if exists == nil {
		logger.Error(err, models.M{
			"file": "user_service",
			"func": "SendDeleteToken",
			"line": "exists.check",
		})
		return nil, errors.New("invalid user id")
	}

	code := helper.RandomString(16)

	exists.DeleteToken = code

	res, err := u.userRepo.Upsert(ctx, *exists)
	if err != nil {
		return nil, err
	}

	tmp, err := helper.Template("deletion-code.mjml")
	if err != nil {
		logger.Error(err)
	}

	go u.notif.SendEmail(
		notification.SendEmailPayload{
			Subject:        "Hello",
			RecipientEmail: res.Email,
			RecipientName:  res.FullName,
			Body:           tmp,
			Variables: models.M{
				"fullName": res.FullName,
				"code":     code,
			},
		},
	)

	return res, nil
}

// Delete implements models.UserService.
func (u *userService) Delete(ctx context.Context, input models.DeleteUserInput) (*models.UserProfile, error) {
	err := validator.ValidateStruct(input)
	if err != nil {
		return nil, err
	}

	userID := input.ID
	token := input.Code

	exists, _ := u.userRepo.Get(ctx, userID)
	if exists == nil {
		logger.Error(err, models.M{
			"file": "user_service",
			"func": "Delete",
			"line": "exists.check",
		})
		return nil, errors.New("invalid user id")
	}

	if exists.DeleteToken != token {
		return nil, errors.New("invalid token")
	}

	return exists, u.userRepo.Delete(ctx, userID)
}

// Upsert implements models.UserService.
func (u *userService) Upsert(ctx context.Context, userID, email string) (*models.UserProfile, error) {
	err := validator.ValidateVar(userID, "required")
	if err != nil {
		return nil, err
	}

	err = validator.ValidateVar(email, "required,email")
	if err != nil {
		return nil, err
	}

	exists, _ := u.Get(ctx, userID)
	isNewUser := false

	if exists == nil {
		// Fetch from Upwork
		upwork, err := u.scrapeService.GetProfile(ctx, userID)
		if err != nil {
			logger.Error(err, models.M{
				"file": "user_service",
				"func": "Upsert",
				"line": "upwork.check",
			})
			return nil, errors.New("invalid user id")
		}

		exists = &models.UserProfile{
			ID:       userID,
			Email:    email,
			FullName: upwork.FullName,
			City:     upwork.City,
			Country:  upwork.Country,
			Title:    upwork.Title,
		}

		isNewUser = true
	}

	code := helper.RandomString(16)

	if isNewUser {
		exists.DeleteToken = code
		exists.Email = email
	}

	res, err := u.userRepo.Upsert(ctx, *exists)
	if err != nil {
		return nil, err
	}

	if isNewUser {
		tmp, err := helper.Template("deletion-code.mjml")
		if err != nil {
			logger.Error(err)
		}

		go u.notif.SendEmail(
			notification.SendEmailPayload{
				Subject:        "Hello",
				RecipientEmail: exists.Email,
				RecipientName:  exists.FullName,
				Body:           tmp,
				Variables: models.M{
					"fullName": exists.FullName,
					"code":     code,
				},
			},
		)
	}

	return res, err
}

// Get implements models.UserService.
func (u *userService) Get(ctx context.Context, userID string) (*models.UserProfile, error) {
	err := validator.ValidateVar(userID, "required")
	if err != nil {
		return nil, err
	}

	return u.userRepo.Get(ctx, userID)
}

func New(scrapeService models.ScrapeService, userRepo models.UserRepository, notif notification.Notification) models.UserService {
	return &userService{
		scrapeService: scrapeService,
		userRepo:      userRepo,
		notif:         notif,
	}
}
