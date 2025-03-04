package models

import "context"

type UpworkProfile struct {
	ID       string `json:"id" bson:"_id"`
	FullName string `json:"fullName" bson:"fullName"`
	City     string `json:"city" bson:"city"`
	Country  string `json:"country" bson:"country"`
	Title    string `json:"title" bson:"title"`
}

type ScrapeService interface {
	InstallBrowser() error
	GetProfile(ctx context.Context, userID string) (*UpworkProfile, error)
}
