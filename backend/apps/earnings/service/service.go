package service

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ariefsn/upwork/helper"
	"github.com/ariefsn/upwork/logger"
	"github.com/ariefsn/upwork/models"
	"github.com/ariefsn/upwork/validator"
	"github.com/redis/go-redis/v9"
)

type earningsService struct {
	earningsRepo models.EarningsRepository
	rdb          *redis.Client
}

// SubscribeEarningUsers implements models.EarningsService.
func (e *earningsService) SubscribeEarningUsers(ctx context.Context, input int) (<-chan []*models.EarningsUserPerYear, error) {
	ch := make(chan []*models.EarningsUserPerYear)

	go func() {
		defer close(ch)

		for {
			time.Sleep(1 * time.Second)

			t, _ := e.EarningUsersYearly(ctx, input)

			select {
			case <-ctx.Done(): // This runs when context gets cancelled. Subscription closes.
				logger.Info("Subscription Closed", models.M{
					"file": "earnings.resolvers",
					"func": "SubEarningUsers",
					"variables": models.M{
						"input": input,
					},
				})
				return // Remember to return to end the routine.

			case ch <- t: // This is the actual send.
				// Our message went through, do nothing
			}
		}
	}()

	// We return the channel and no error.
	return ch, nil
}

// SubscribeEarnings implements models.EarningsService.
func (e *earningsService) SubscribeEarnings(ctx context.Context, input models.EarningsUserPerYearInput) (<-chan []*models.EarningsDataMonthly, error) {
	ch := make(chan []*models.EarningsDataMonthly)

	go func() {
		defer close(ch)

		for {
			time.Sleep(1 * time.Second)

			t, _ := e.GetByUserIDPerYear(ctx, input)

			select {
			case <-ctx.Done(): // This runs when context gets cancelled. Subscription closes.
				logger.Info("Subscription Closed", models.M{
					"file": "earnings.resolvers",
					"func": "SubEarningUsers",
					"variables": models.M{
						"input": input,
					},
				})
				return // Remember to return to end the routine.

			case ch <- t: // This is the actual send.
				// Our message went through, do nothing
			}
		}
	}()

	// We return the channel and no error.
	return ch, nil
}

// SubscribeOnEarningUpdated implements models.EarningsService.
func (e *earningsService) SubscribeOnEarningUpdated(ctx context.Context, input models.EarningsUserPerYearInput) (<-chan []*models.EarningsDataMonthly, error) {
	ch := make(chan []*models.EarningsDataMonthly)

	pubsub := e.rdb.Subscribe(ctx, fmt.Sprintf("earnings-%s-%d", input.UserId, input.Year))

	go func() {
		rdbCh := pubsub.Channel()

		defer func() {
			close(ch)
			pubsub.Close()
		}()

		for res := range rdbCh {
			var t []*models.EarningsDataMonthly

			if res.Payload != "" {
				parsed, err := helper.FromBytes[[]*models.EarningsDataMonthly]([]byte(res.Payload))
				if err == nil && parsed != nil {
					t = parsed
				}
			}

			select {
			case <-ctx.Done(): // This runs when context gets cancelled. Subscription closes.
				logger.Info("Subscription Closed", models.M{
					"file": "earnings.resolvers",
					"func": "SubEarningUsers",
					"variables": models.M{
						"input": input,
					},
				})
				return // Remember to return to end the routine.

			case ch <- t: // This is the actual send.
				// Our message went through, do nothing
			}
		}
	}()

	// We return the channel and no error.
	return ch, nil
}

// GetYearsByUserID implements models.EarningsService.
func (e *earningsService) GetYearsByUserID(ctx context.Context, userID string) ([]int, error) {
	err := validator.ValidateVar(userID, "required")
	if err != nil {
		return nil, err
	}

	return e.earningsRepo.GetYearsByUserID(ctx, userID)
}

// EarningUsersYearly implements models.EarningsService.
func (e *earningsService) EarningUsersYearly(ctx context.Context, year int) ([]*models.EarningsUserPerYear, error) {
	err := validator.ValidateVar(year, "required")
	if err != nil {
		return nil, err
	}

	return e.earningsRepo.EarningUsersYearly(ctx, year)
}

// ParseCsv implements models.EarningsService.
func (e *earningsService) ParseCsv(ctx context.Context, input models.EarningsInput) ([]*models.EarningsData, error) {
	err := validator.ValidateStruct(input)
	if err != nil {
		return nil, err
	}

	isCsv := input.File.ContentType == "text/csv"
	if !isCsv {
		return nil, fmt.Errorf("please upload csv file")
	}

	file := input.File.File
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	headers := records[0]

	for _, h := range models.EarningFileHeaders {
		found := false
		for _, currentH := range headers {
			if string(h) == currentH {
				found = true
				break
			}
		}

		if !found {
			return nil, fmt.Errorf("invalid file header - %s", h)
		}
	}

	contents := records[1:]

	data := []*models.EarningsData{}

	for _, v := range contents {
		earningType := models.EarningType(v[2])
		if models.AllowedEarningTypes[earningType] {
			dateStr := v[0] // Jan 25, 2025
			dateParsed, err := time.Parse("Jan 2, 2006", dateStr)
			if err != nil {
				fmt.Println(dateStr, dateParsed, v)
				return nil, errors.New("invalid date")
			}
			year, month, day := dateParsed.Date()
			amount, _ := strconv.ParseFloat(v[9], 64)
			refID := v[1]
			earningID := fmt.Sprintf("%s-%s", dateParsed.Format("20060102"), refID)
			description := v[3]
			team := v[6]

			fee := float64(0)
			for _, elFee := range contents {
				elFeeEarningType := models.EarningType(elFee[2])
				elFeeDescription := elFee[3]
				if strings.Contains(elFeeDescription, refID) && elFeeEarningType == models.EarningTypeServiceFee {
					elFeeAmount, _ := strconv.ParseFloat(elFee[9], 64)
					fee = elFeeAmount
					if fee < 0 {
						fee *= -1
					}
				}
			}

			newData := &models.EarningsData{
				ID:          earningID,
				UserID:      input.UserId,
				Day:         day,
				Month:       int(month),
				Year:        year,
				RefID:       refID,
				Type:        earningType,
				Description: description,
				Team:        team,
				Amount:      amount,
				Fee:         fee,
			}
			data = append(data, newData)

			_, err = e.Upsert(ctx, *newData)
			if err != nil {
				return nil, err
			}
		}
	}

	defer func() {
		years, _ := e.GetYearsByUserID(ctx, input.UserId)
		if len(years) > 0 {
			for _, year := range years {
				res, _ := e.GetByUserIDPerYear(ctx, models.EarningsUserPerYearInput{UserId: input.UserId, Year: year})
				if len(res) > 0 {
					e.rdb.Publish(ctx, fmt.Sprintf("earnings-%s-%d", input.UserId, year), string(helper.ToBytes(res)))
				}
			}
		}
	}()

	return data, nil
}

// GetByUserIDPerYear implements models.EarningsService.
func (e *earningsService) GetByUserIDPerYear(ctx context.Context, input models.EarningsUserPerYearInput) ([]*models.EarningsDataMonthly, error) {
	err := validator.ValidateStruct(input)
	if err != nil {
		return nil, err
	}

	return e.earningsRepo.GetByUserIDPerYear(ctx, input.UserId, input.Year)
}

// DeleteByUserID implements models.EarningsService.
func (e *earningsService) DeleteByUserID(ctx context.Context, userID string) (*int64, error) {
	err := validator.ValidateVar(userID, "required")
	if err != nil {
		return nil, err
	}

	return e.earningsRepo.DeleteByUserID(ctx, userID)
}

// Upsert implements models.EarningsService.
func (e *earningsService) Upsert(ctx context.Context, data models.EarningsData) (*models.EarningsData, error) {
	err := validator.ValidateStruct(data)
	if err != nil {
		return nil, err
	}

	return e.earningsRepo.Upsert(ctx, data)
}

func New(earningsRepo models.EarningsRepository, rdb *redis.Client) models.EarningsService {
	return &earningsService{
		earningsRepo: earningsRepo,
		rdb:          rdb,
	}
}
