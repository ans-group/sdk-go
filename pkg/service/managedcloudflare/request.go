package managedcloudflare

// CreateRecordRequest represents a request to create an account
type CreateAccountRequest struct {
	Name string `json:"name"`
}

// CreateRecordRequest represents a request to create an account member
type CreateAccountMemberRequest struct {
	EmailAddress string `json:"email_address"`
}

type CreateOrchestrationRequest struct {
	ZoneName                  string `json:"zone_name"`
	ZoneSubscriptionType      string `json:"zone_subscription_type"`
	AccountID                 string `json:"account_id"`
	AccountName               string `json:"account_name"`
	AdministratorEmailAddress string `json:"administrator_email_address"`
}
