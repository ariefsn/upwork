package models

import "context"

type UserData struct {
	ID       string `json:"id" bson:"_id" validate:"required"`
	FullName string `json:"fullName" bson:"fullName" validate:"required"`
	City     string `json:"city" bson:"city"`
	Country  string `json:"country" bson:"country"`
	Title    string `json:"title" bson:"title"`
}

func (m *UserData) TableName() string {
	return "users"
}

type UserProfile struct {
	ID          string `json:"id" bson:"_id" validate:"required"`
	Email       string `json:"email" bson:"email" validate:"required,email"`
	FullName    string `json:"fullName" bson:"fullName" validate:"required"`
	City        string `json:"city" bson:"city"`
	Country     string `json:"country" bson:"country"`
	Title       string `json:"title" bson:"title"`
	DeleteToken string `json:"deleteToken" bson:"deleteToken" validate:"required"`
}

func (m *UserProfile) TableName() string {
	return "users"
}

type DeleteUserInput struct {
	ID   string `json:"id"`
	Code string `json:"code"`
}

type UserService interface {
	Get(ctx context.Context, userID string) (*UserProfile, error)
	Upsert(ctx context.Context, userID, email string) (*UserProfile, error)
	Delete(ctx context.Context, input DeleteUserInput) (*UserProfile, error)
	SendDeleteToken(ctx context.Context, userID string) (*UserProfile, error)
	GetIds(ctx context.Context) ([]string, error)
}

type UserRepository interface {
	Get(ctx context.Context, userID string) (*UserProfile, error)
	Upsert(ctx context.Context, userData UserProfile) (*UserProfile, error)
	Delete(ctx context.Context, userID string) error
	GetIds(ctx context.Context) ([]string, error)
}
