package managedcloudflare

// CreateRecordRequest represents a request to create an account
type CreateAccountRequest struct {
	Name string `json:"name"`
}

// CreateRecordRequest represents a request to create an account member
type CreateAccountMemberRequest struct {
	EmailAddress string `json:"email_address"`
}
