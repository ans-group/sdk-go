//go:generate go run ../../gen/model_paginated_gen.go -package ltaas -typename Domain,Test,Job,Threshold,Scenario -destination model_paginated.go

package ltaas

import (
	"fmt"
	"time"

	"github.com/ukfast/go-durationstring"
	"github.com/ukfast/sdk-go/pkg/connection"
)

type DomainVerificationMethod string

func (s DomainVerificationMethod) String() string {
	return string(s)
}

const (
	DomainVerificationMethodDNS        DomainVerificationMethod = "DNS"
	DomainVerificationMethodFileUpload DomainVerificationMethod = "File upload"
)

var DomainVerificationMethodEnum = []connection.Enum{
	DomainVerificationMethodDNS,
	DomainVerificationMethodFileUpload,
}

// ParseDomainVerificationMethod attempts to parse a DomainVerificationMethod from string
func ParseDomainVerificationMethod(s string) (DomainVerificationMethod, error) {
	e, err := connection.ParseEnum(s, DomainVerificationMethodEnum)
	if err != nil {
		return "", err
	}

	return e.(DomainVerificationMethod), err
}

type DomainStatus string

func (s DomainStatus) String() string {
	return string(s)
}

const (
	DomainStatusVerified    DomainStatus = "Verified"
	DomainStatusNotVerified DomainStatus = "Not verified"
)

type TestProtocol string

func (s TestProtocol) String() string {
	return string(s)
}

const (
	TestProtocolHTTP  TestProtocol = "http"
	TestProtocolHTTPS TestProtocol = "https"
)

var TestProtocolEnum = []connection.Enum{
	TestProtocolHTTP,
	TestProtocolHTTPS,
}

// ParseTestProtocol attempts to parse a TestProtocol from string
func ParseTestProtocol(s string) (TestProtocol, error) {
	e, err := connection.ParseEnum(s, TestProtocolEnum)
	if err != nil {
		return "", err
	}

	return e.(TestProtocol), err
}

type TestRecurringType string

func (s TestRecurringType) String() string {
	return string(s)
}

const (
	TestRecurringTypeDaily  TestRecurringType = "Daily"
	TestRecurringTypeOneOff TestRecurringType = "One off"
	TestRecurringTypeWeekly TestRecurringType = "Weekly"
)

type JobStatus string

func (s JobStatus) String() string {
	return string(s)
}

const (
	JobStatusPending  JobStatus = "Pending"
	JobStatusRunning  JobStatus = "Running"
	JobStatusFailed   JobStatus = "Failed"
	JobStatusStopped  JobStatus = "Stopped"
	JobStatusComplete JobStatus = "Complete"
)

type JobFailType string

func (s JobFailType) String() string {
	return string(s)
}

const (
	JobFailTypeTest           JobFailType = "Test"
	JobFailTypeInfrastructure JobFailType = "Infrastructure"
)

// TestDuration represents a load test duration
type TestDuration string

// Duration returns the test duration as time.Duration
func (d *TestDuration) Duration() time.Duration {
	if len(*d) < 8 {
		return time.Duration(0)
	}

	duration, err := time.ParseDuration(string(*d)[0:2] + "h" + string(*d)[3:5] + "m" + string(*d)[6:8] + "s")
	if err != nil {
		return time.Duration(0)
	}

	return duration
}

// ParseTestDuration parses string s and returns a pointer to an
// initialised TestDuration
func ParseTestDuration(s string) (TestDuration, error) {
	_, _, _, hours, minutes, seconds, _, _, _, err := durationstring.Parse(s)
	if err != nil {
		return TestDuration(""), err
	}

	return TestDuration(fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)), nil
}

type AgreementType string

func (s AgreementType) String() string {
	return string(s)
}

const (
	AgreementTypeSingle    AgreementType = "single"
	AgreementTypeRecurring AgreementType = "recurring"
)

var AgreementTypeEnum = []connection.Enum{
	AgreementTypeSingle,
	AgreementTypeRecurring,
}

// ParseAgreementType attempts to parse a AgreementType from string
func ParseAgreementType(s string) (AgreementType, error) {
	e, err := connection.ParseEnum(s, AgreementTypeEnum)
	if err != nil {
		return "", err
	}

	return e.(AgreementType), err
}

// Domain represents an LTaaS domain
type Domain struct {
	ID                 string                   `json:"id"`
	Name               string                   `json:"name"`
	VerificationMethod DomainVerificationMethod `json:"verification_method"`
	VerifyHash         string                   `json:"verify_hash"`
	Status             DomainStatus             `json:"status"`
	CreatedAt          connection.DateTime      `json:"created_at"`
	UpdatedAt          connection.DateTime      `json:"updated_at"`
}

// Test represents an LTaaS test
type Test struct {
	ID             string              `json:"id"`
	Name           string              `json:"name"`
	ScriptID       string              `json:"script_id"`
	ScenarioID     string              `json:"scenario_id"`
	DomainID       string              `json:"domain_id"`
	Protocol       TestProtocol        `json:"protocol"`
	Path           string              `json:"path"`
	NumberOfUsers  int                 `json:"number_of_users"`
	Duration       TestDuration        `json:"duration"`
	RecurringType  TestRecurringType   `json:"recurring_type"`
	RecurringValue int                 `json:"recurring_value"`
	NextRun        connection.DateTime `json:"next_run"`
	CreatedAt      connection.DateTime `json:"created_at"`
	UpdatedAt      connection.DateTime `json:"updated_at"`
}

// Job represents an LTaaS job
type Job struct {
	ID                 string              `json:"id"`
	TestID             string              `json:"test_id"`
	DomainID           string              `json:"domain_id"`
	ScheduledTimestamp connection.DateTime `json:"scheduled_timestamp"`
	JobStartTimestamp  connection.DateTime `json:"job_start_timestamp"`
	JobEndTimestamp    connection.DateTime `json:"job_end_timestamp"`
	Status             JobStatus           `json:"status"`
	FailType           JobFailType         `json:"fail_type"`
	CreatedAt          connection.DateTime `json:"created_at"`
	UpdatedAt          connection.DateTime `json:"updated_at"`
}

type JobResultsAxis struct {
	X connection.DateTime `json:"x"`
	Y float64             `json:"y"`
}

// JobResults represents the results of an LTaaS job
type JobResults struct {
	VirtualUsers       []JobResultsAxis `json:"virtual_users"`
	SuccessfulRequests []JobResultsAxis `json:"successful_requests"`
	FailedRequests     []JobResultsAxis `json:"failed_requests"`
	Latency            []JobResultsAxis `json:"latency"`
}

// JobSettings represents the settings of an LTaaS job
type JobSettings struct {
	Date     connection.DateTime `json:"date"`
	Type     string              `json:"type"`
	Name     string              `json:"name"`
	Duration TestDuration        `json:"duration"`
	MaxUsers int                 `json:"max_users"`
	Protocol TestProtocol        `json:"protocol"`
	Domain   string              `json:"domain"`
	Path     string              `json:"path"`
}

// Threshold represents a test threshold
type Threshold struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Query       string              `json:"query"`
	CreatedAt   connection.DateTime `json:"created_at"`
	UpdatedAt   connection.DateTime `json:"updated_at"`
}

// Scenario represents a test scenario
type Scenario struct {
	ID             string              `json:"id"`
	Name           string              `json:"name"`
	AvailableTrial bool                `json:"available_trial"`
	Formula        string              `json:"formula"`
	Description    string              `json:"description"`
	CreatedAt      connection.DateTime `json:"created_at"`
	UpdatedAt      connection.DateTime `json:"updated_at"`
}

// Agreement represents an authorisation agreement
type Agreement struct {
	Version   string `json:"version"`
	Agreement string `json:"agreement"`
}
