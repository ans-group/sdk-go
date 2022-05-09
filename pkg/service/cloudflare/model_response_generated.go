package cloudflare

import "github.com/ukfast/sdk-go/pkg/connection"

// GetAccountSliceResponseBody represents an API response body containing []Account data
type GetAccountSliceResponseBody struct {
	connection.APIResponseBody
	Data []Account `json:"data"`
}

// GetAccountResponseBody represents an API response body containing Account data
type GetAccountResponseBody struct {
	connection.APIResponseBody
	Data Account `json:"data"`
}

// GetAccountMemberSliceResponseBody represents an API response body containing []AccountMember data
type GetAccountMemberSliceResponseBody struct {
	connection.APIResponseBody
	Data []AccountMember `json:"data"`
}

// GetAccountMemberResponseBody represents an API response body containing AccountMember data
type GetAccountMemberResponseBody struct {
	connection.APIResponseBody
	Data AccountMember `json:"data"`
}

// GetSpendPlanSliceResponseBody represents an API response body containing []SpendPlan data
type GetSpendPlanSliceResponseBody struct {
	connection.APIResponseBody
	Data []SpendPlan `json:"data"`
}

// GetSpendPlanResponseBody represents an API response body containing SpendPlan data
type GetSpendPlanResponseBody struct {
	connection.APIResponseBody
	Data SpendPlan `json:"data"`
}

// GetSubscriptionSliceResponseBody represents an API response body containing []Subscription data
type GetSubscriptionSliceResponseBody struct {
	connection.APIResponseBody
	Data []Subscription `json:"data"`
}

// GetSubscriptionResponseBody represents an API response body containing Subscription data
type GetSubscriptionResponseBody struct {
	connection.APIResponseBody
	Data Subscription `json:"data"`
}

// GetZoneSliceResponseBody represents an API response body containing []Zone data
type GetZoneSliceResponseBody struct {
	connection.APIResponseBody
	Data []Zone `json:"data"`
}

// GetZoneResponseBody represents an API response body containing Zone data
type GetZoneResponseBody struct {
	connection.APIResponseBody
	Data Zone `json:"data"`
}

// GetTotalSpendSliceResponseBody represents an API response body containing []TotalSpend data
type GetTotalSpendSliceResponseBody struct {
	connection.APIResponseBody
	Data []TotalSpend `json:"data"`
}

// GetTotalSpendResponseBody represents an API response body containing TotalSpend data
type GetTotalSpendResponseBody struct {
	connection.APIResponseBody
	Data TotalSpend `json:"data"`
}
