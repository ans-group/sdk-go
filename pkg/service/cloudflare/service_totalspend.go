package cloudflare

import (
	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetTotalSpendMonthToDate retrieves a total spend for current month
func (s *Service) GetTotalSpendMonthToDate() (TotalSpend, error) {
	body, err := connection.Get[TotalSpend](s.connection, "/cloudflare/v1/total-spend/month-to-date", connection.APIRequestParameters{})
	return body.Data, err
}
