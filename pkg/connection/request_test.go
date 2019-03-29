package connection

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewValidationError_PopulatesMessage(t *testing.T) {
	err := NewValidationError("test error 1")

	assert.Equal(t, "test error 1", err.Message)
}

func TestAPIRequestBodyDefaultValidator(t *testing.T) {
	type teststruct struct {
		TestProperty string `validate:"required"`
	}

	t.Run("Valid_NoError", func(t *testing.T) {
		ts := teststruct{
			TestProperty: "testvalue1",
		}
		v := APIRequestBodyDefaultValidator{}

		err := v.Validate(ts)

		assert.Nil(t, err)
	})

	t.Run("Invalid_ReturnsValidationError", func(t *testing.T) {
		ts := teststruct{}
		v := APIRequestBodyDefaultValidator{}

		err := v.Validate(ts)

		assert.NotNil(t, err)
		assert.IsType(t, &ValidationError{}, err)
	})
}

func TestAPIRequestParameters_WithPagination_SetsPagination(t *testing.T) {
	p := APIRequestPagination{
		Page: 5,
	}
	params := APIRequestParameters{}

	params.WithPagination(p)

	assert.Equal(t, p, params.Pagination)
}

func TestAPIRequestParameters_WithSorting_SetsSorting(t *testing.T) {
	s := APIRequestSorting{
		Property: "testproperty1",
	}
	params := APIRequestParameters{}

	params.WithSorting(s)

	assert.Equal(t, s, params.Sorting)
}

func TestAPIRequestParameters_WithFilter_AddsFilter(t *testing.T) {
	f := APIRequestFiltering{
		Property: "testproperty1",
		Operator: EQOperator,
		Value:    []string{"testvalue1"},
	}
	params := APIRequestParameters{}

	params.WithFilter(f)

	assert.Len(t, params.Filtering, 1)
	assert.Equal(t, f, params.Filtering[0])
}

func TestAPIRequestParameters_WithFilters_AddsFilters(t *testing.T) {
	f1 := APIRequestFiltering{
		Property: "testproperty1",
		Operator: EQOperator,
		Value:    []string{"testvalue1"},
	}
	f2 := APIRequestFiltering{
		Property: "testproperty2",
		Operator: EQOperator,
		Value:    []string{"testvalue2"},
	}

	f := []APIRequestFiltering{f1, f2}
	params := APIRequestParameters{}

	params.WithFilters(f)

	assert.Len(t, params.Filtering, 2)
	assert.Equal(t, f1, params.Filtering[0])
	assert.Equal(t, f2, params.Filtering[1])
}

func TestParseOperator(t *testing.T) {
	type testoperator struct {
		Operator         string
		ExpectedOperator APIRequestFilteringOperator
	}
	operators := []testoperator{
		testoperator{
			Operator:         "EQ",
			ExpectedOperator: EQOperator,
		},
		testoperator{
			Operator:         "LK",
			ExpectedOperator: LKOperator,
		},
		testoperator{
			Operator:         "GT",
			ExpectedOperator: GTOperator,
		},
		testoperator{
			Operator:         "LT",
			ExpectedOperator: LTOperator,
		},
		testoperator{
			Operator:         "IN",
			ExpectedOperator: INOperator,
		},
		testoperator{
			Operator:         "NEQ",
			ExpectedOperator: NEQOperator,
		},
		testoperator{
			Operator:         "NIN",
			ExpectedOperator: NINOperator,
		},
		testoperator{
			Operator:         "NLK",
			ExpectedOperator: NLKOperator,
		},
	}

	for _, operator := range operators {
		t.Run(fmt.Sprintf("%s_Exact_Parses", operator.Operator), func(t *testing.T) {
			parsedOperator, err := ParseOperator(operator.Operator)

			assert.Nil(t, err)
			assert.Equal(t, operator.ExpectedOperator, parsedOperator)
		})
	}

	t.Run("MixedCase_Parses", func(t *testing.T) {
		operator, err := ParseOperator("gT")

		assert.Nil(t, err)
		assert.Equal(t, GTOperator, operator)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		_, err := ParseOperator("invalidoperator")

		assert.NotNil(t, err)
	})
}

func TestAPIRequestFilteringOperator_String(t *testing.T) {
	t.Run("Valid_String", func(t *testing.T) {
		o := GTOperator

		s := o.String()

		assert.Equal(t, "gt", s)
	})

	t.Run("Invalid_Unknown", func(t *testing.T) {
		var o APIRequestFilteringOperator = 999

		s := o.String()

		assert.Equal(t, "Unknown", s)
	})
}

func TestRequestAll_Invoke(t *testing.T) {
	t.Run("MultiplePages_ExpectedCalls", func(t *testing.T) {
		totalCalls := 0
		f := func(parameters APIRequestParameters) (ResponseBody, error) {
			totalCalls++

			return &APIResponseBody{
				Metadata: APIResponseMetadata{
					Pagination: APIResponseMetadataPagination{
						TotalPages: 10,
					},
				},
			}, nil
		}

		r := RequestAll{
			GetNext: f,
		}

		err := r.Invoke(APIRequestParameters{})

		assert.Nil(t, err)
		assert.Equal(t, 10, totalCalls)
	})

	t.Run("Error", func(t *testing.T) {
		totalCalls := 0
		f := func(parameters APIRequestParameters) (ResponseBody, error) {
			totalCalls++

			return &APIResponseBody{
				Metadata: APIResponseMetadata{
					Pagination: APIResponseMetadataPagination{
						TotalPages: 10,
					},
				},
			}, errors.New("test error 1")
		}

		r := RequestAll{
			GetNext: f,
		}

		err := r.Invoke(APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, "test error 1", err.Error())
		assert.Equal(t, 1, totalCalls)
	})
}
