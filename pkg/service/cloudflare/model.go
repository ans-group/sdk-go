//go:generate go run ../../gen/model_response/main.go -package cloudflare -source model.go -destination model_response_generated.go
//go:generate go run ../../gen/model_paginated/main.go -package cloudflare -source model.go -destination model_paginated_generated.go

package cloudflare

import "github.com/ukfast/sdk-go/pkg/connection"

// Account represents a Cloudflare account
// +genie:model_response
// +genie:model_paginated
type Account struct {
	ID                  string              `json:"id"`
	Status              string              `json:"status"`
	Name                string              `json:"name"`
	CloudflareAccountID string              `json:"cloudflare_account_id"`
	CreatedAt           connection.DateTime `json:"created_at"`
	UpdatedAt           connection.DateTime `json:"updated_at"`
}

// AccountMember represents a Cloudflare account member
// +genie:model_response
// +genie:model_paginated
type AccountMember struct {
	EmailAddress string `json:"email_address"`
}

// SpendPlan represents a Cloudflare spend plan
// +genie:model_response
// +genie:model_paginated
type SpendPlan struct {
	ID        string              `json:"id"`
	Amount    float32             `json:"amount"`
	StartedAt connection.DateTime `json:"started_at"`
	EndedAt   connection.DateTime `json:"ended_at"`
	CreatedAt connection.DateTime `json:"created_at"`
	UpdatedAt connection.DateTime `json:"updated_at"`
}

// Subscription represents a Cloudflare subscription
// +genie:model_response
// +genie:model_paginated
type Subscription struct {
	ID                   string              `json:"id"`
	Name                 string              `json:"name"`
	Type                 string              `json:"type"`
	Description          string              `json:"description"`
	Price                float32             `json:"price"`
	CloudflareRatePlanID string              `json:"cloudflare_rate_plan_id"`
	CreatedAt            connection.DateTime `json:"created_at"`
	UpdatedAt            connection.DateTime `json:"updated_at"`
}

// Zone represents a Cloudflare zone
// +genie:model_response
// +genie:model_paginated
type Zone struct {
	ID                 string              `json:"id"`
	AccountID          string              `json:"account_id"`
	Name               string              `json:"name"`
	PlanSubscriptionID string              `json:"plan_subscription_id"`
	CloudflareZoneID   string              `json:"cloudflare_zone_id"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

// TotalSpend represents total spend
// +genie:model_response
type TotalSpend struct {
	SpendPlanAmount float32 `json:"spend_plan_amount"`
	TotalSpend      float32 `json:"total_spend"`
}
