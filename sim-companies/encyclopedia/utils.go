package encyclopedia

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func sendRequest[K any](req *http.Request) (*K, error) {
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer res.Body.Close()
	if !isResponseOk(res) {
		return nil, handleBadResponse(res)
	}
	b, _ := io.ReadAll(res.Body)
	value, err := unmarshal[K](b)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// generic function
func unmarshal[K any](blob []byte) (*K, error) {
	var val = new(K)
	err := json.Unmarshal(blob, val)
	if err != nil {
		return nil, err
	}
	return val, err
}

func isResponseOk(r *http.Response) bool {
	return r.StatusCode/100 == 2
}

func handleBadResponse(r *http.Response) error {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	m, err := unmarshal[errorResponse](b)
	if err != nil {
		return err
	}

	if r.StatusCode == http.StatusNotFound {
		return fmt.Errorf(m.Message)
	}

	return fmt.Errorf("something went wrong, %s", m.Message)
}
