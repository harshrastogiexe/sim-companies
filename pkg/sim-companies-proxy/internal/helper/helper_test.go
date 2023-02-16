package helper_test

import (
	"net/http"
	"testing"

	"github.com/harshrastogiexe/pkg/sim-companies-proxy/internal/helper"
)

func Test_IsStatusCodeOk(t *testing.T) {
	type TestCase struct {
		code     int
		expected bool
	}

	tcs := []TestCase{
		{code: http.StatusNotFound, expected: false},
		{code: http.StatusInternalServerError, expected: false},
		{code: http.StatusOK, expected: true},
		{code: http.StatusCreated, expected: true},
	}

	for _, tc := range tcs {
		res := http.Response{StatusCode: tc.code}
		if isOk := helper.IsResponseStatusOk(&res); tc.expected != isOk {
			t.Errorf("expected %v when status code is %d, got %v when", tc.expected, tc.code, isOk)
		}
	}
}
