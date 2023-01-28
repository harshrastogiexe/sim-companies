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
	tests := []TestCase{
		{code: http.StatusNotFound, expected: false},
		{code: http.StatusInternalServerError, expected: false},
		{code: http.StatusOK, expected: true},
		{code: http.StatusCreated, expected: true},
	}

	for _, testCase := range tests {
		response := http.Response{StatusCode: testCase.code}
		if actual := helper.IsResponseStatusOk(&response); testCase.expected != actual {
			t.Errorf("expected %v when status code is %d, got %v when", testCase.expected, testCase.code, actual)
		}
	}
}
