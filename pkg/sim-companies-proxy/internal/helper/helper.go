package helper

import (
	"encoding/json"
	"net/http"
)

func ParseResponseJsonBody[T any](res *http.Response) (*T, error) {
	var value T
	if err := json.NewDecoder(res.Body).Decode(&value); err != nil {
		return nil, err
	}
	return &value, nil
}

// checks whether response code starts with '2'
func IsResponseStatusOk(response *http.Response) bool {
	startLiteral := response.StatusCode / 100
	return startLiteral == 2
}
