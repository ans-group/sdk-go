package cloudflare

// CreateAccountRequest represents a request to create an account
type CreateAccountRequest struct {
	Name string `json:"name"`
}

// CreateAccountMemberRequest represents a request to create an account member
type CreateAccountMemberRequest struct {
	EmailAddress string `json:"email_address"`
}

// CreateOrchestrationRequest represents a request to create new orchestration
type CreateOrchestrationRequest struct {
	ZoneName                  string `json:"zone_name"`
	ZoneSubscriptionType      string `json:"zone_subscription_type"`
	AccountID                 string `json:"account_id"`
	AccountName               string `json:"account_name"`
	AdministratorEmailAddress string `json:"administrator_email_address"`
}

// CreateRecordRequest represents a request to create an zone
type CreateZoneRequest struct {
	AccountID        string `json:"account_id"`
	Name             string `json:"name"`
	SubscriptionType string `json:"subscription_type"`
}
