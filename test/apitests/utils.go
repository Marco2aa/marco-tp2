package apitests

import "bytes"
import "encoding/json"
import "io"
import "net/http"
import "time"

type CatModel struct {
	Name      string `json:"name"`
	ID        string `json:"id,omitempty"`
	BirthDate string `json:"birthDate,omitempty"`
	Color     string `json:"color,omitempty"`
}

var baseUrl = "http://localhost:8080/api"

// Global client with a proper timeout
var client = &http.Client{Timeout: 10 * time.Second}

// Wrapper to HTTP API calls, does the error handling and JSON decoding
func call(method, path string, reqBody any, code *int, result any) error {

	var body io.Reader
	if reqBody != nil {
		jsonBody, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, baseUrl+path, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// send the request
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if code != nil {
		*code = res.StatusCode
	}

	if result != nil {
		err = json.NewDecoder(res.Body).Decode(result)
	}

	return err
}
