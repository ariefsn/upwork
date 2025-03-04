package resolvers

import (
	"github.com/ariefsn/upwork/models"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ScrapeService   models.ScrapeService
	UserService     models.UserService
	EarningsService models.EarningsService
}
