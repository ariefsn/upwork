package models

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

type EarningField string

const (
	EarningFieldDate                  EarningField = "Date"
	EarningFieldRefID                 EarningField = "Ref ID"
	EarningFieldType                  EarningField = "Type"
	EarningFieldDescription           EarningField = "Description"
	EarningFieldAgency                EarningField = "Agency"
	EarningFieldFreelancer            EarningField = "Freelancer"
	EarningFieldTeam                  EarningField = "Team"
	EarningFieldAccountName           EarningField = "Account Name"
	EarningFieldPO                    EarningField = "PO"
	EarningFieldAmount                EarningField = "Amount"
	EarningFieldAmountInLocalCurrency EarningField = "Amount in local currency"
	EarningFieldCurrency              EarningField = "Currency"
	EarningFieldBalance               EarningField = "Balance"
)

var EarningFileHeaders = []EarningField{
	EarningFieldDate,
	EarningFieldRefID,
	EarningFieldType,
	EarningFieldDescription,
	EarningFieldAgency,
	EarningFieldFreelancer,
	EarningFieldTeam,
	EarningFieldAccountName,
	EarningFieldPO,
	EarningFieldAmount,
	EarningFieldAmountInLocalCurrency,
	EarningFieldCurrency,
	EarningFieldBalance,
}

type EarningType string

const (
	EarningTypeSubscription  EarningType = "Subscription"
	EarningTypeWithdrawal    EarningType = "Withdrawal"
	EarningTypeMembershipFee EarningType = "Membership Fee"
	EarningTypeWithdrawalFee EarningType = "WithdrawalFee"
	EarningTypeServiceFee    EarningType = "Service Fee"
	EarningTypeFixedPrice    EarningType = "Fixed Price"
	EarningTypeHourly        EarningType = "Hourly"
	EarningTypeConnects      EarningType = "Connects"
	EarningTypeAdjustment    EarningType = "Adjustment"
	EarningTypePayment       EarningType = "Payment"
)

var AllowedEarningTypes = map[EarningType]bool{
	EarningTypeFixedPrice: true,
	EarningTypeHourly:     true,
}

type EarningsData struct {
	ID          string      `bson:"_id" json:"id" validate:"required"`
	UserID      string      `bson:"userID" json:"userID" validate:"required"`
	Day         int         `bson:"day" json:"day" validate:"required,numeric"`
	Month       int         `bson:"month" json:"month" validate:"required,numeric"`
	Year        int         `bson:"year" json:"year" validate:"required,numeric"`
	RefID       string      `bson:"refID" json:"refID" validate:"required"`
	Type        EarningType `bson:"type" json:"type" validate:"required,oneof='Fixed Price' Hourly"`
	Description string      `bson:"description" json:"description" validate:"required"`
	Team        string      `bson:"team" json:"team" validate:"required"`
	Amount      float64     `bson:"amount" json:"amount" validate:"required"`
	Fee         float64     `bson:"fee" json:"fee"`
}

func (m *EarningsData) TableName() string {
	return "earnings"
}

type EarningsDataMonthlyItem struct {
	Type   EarningType `bson:"type" json:"type" validate:"required,oneof='Fixed Price' Hourly"`
	Amount float64     `bson:"amount" json:"amount" validate:"required"`
	Fee    float64     `bson:"fee" json:"fee"`
}

type EarningsDataMonthly struct {
	UserID      string                    `bson:"userID" json:"userID" validate:"required"`
	Month       int                       `bson:"month" json:"month" validate:"required,numeric"`
	Year        int                       `bson:"year" json:"year" validate:"required,numeric"`
	TotalAmount float64                   `bson:"totalAmount" json:"totalAmount" validate:"required"`
	TotalFee    float64                   `bson:"totalFee" json:"totalFee" validate:"required"`
	Items       []EarningsDataMonthlyItem `bson:"items" json:"items"`
}

func (m *EarningsDataMonthly) TableName() string {
	return "earnings"
}

type EarningsInput struct {
	UserId string         `json:"userID" validate:"required"`
	File   graphql.Upload `json:"file" validate:"required"`
	Email  string         `json:"email" validate:"required,email"`
}

type EarningsUserPerYearInput struct {
	UserId string `json:"userID" validate:"required"`
	Year   int    `json:"year" validate:"required"`
}

type EarningsUserPerYear struct {
	User   UserData `json:"user" validate:"required"`
	Amount float64  `bson:"amount" json:"amount" validate:"required"`
	Fee    float64  `bson:"fee" json:"fee"`
}

type EarningsService interface {
	ParseCsv(ctx context.Context, input EarningsInput) ([]*EarningsData, error)
	GetByUserIDPerYear(ctx context.Context, input EarningsUserPerYearInput) ([]*EarningsDataMonthly, error)
	Upsert(ctx context.Context, data EarningsData) (*EarningsData, error)
	DeleteByUserID(ctx context.Context, userID string) (*int64, error)
	EarningUsersYearly(ctx context.Context, year int) ([]*EarningsUserPerYear, error)
	GetYearsByUserID(ctx context.Context, userID string) ([]int, error)
	SubscribeEarnings(ctx context.Context, input EarningsUserPerYearInput) (<-chan []*EarningsDataMonthly, error)
	SubscribeEarningUsers(ctx context.Context, input int) (<-chan []*EarningsUserPerYear, error)
	SubscribeOnEarningUpdated(ctx context.Context, input EarningsUserPerYearInput) (<-chan []*EarningsDataMonthly, error)
}

type EarningsRepository interface {
	GetByUserIDPerYear(ctx context.Context, userID string, year int) ([]*EarningsDataMonthly, error)
	Upsert(ctx context.Context, data EarningsData) (*EarningsData, error)
	DeleteByUserID(ctx context.Context, userID string) (*int64, error)
	EarningUsersYearly(ctx context.Context, year int) ([]*EarningsUserPerYear, error)
	GetYearsByUserID(ctx context.Context, userID string) ([]int, error)
}
