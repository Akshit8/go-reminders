package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// reminderBody represents reminder request body
type reminderBody struct {
	ID       string        `json:"id"`
	Title    string        `json:"title"`
	Message  string        `json:"message"`
	Duration time.Duration `json:"duration"`
}

// HTTPClient represents the HTTP client which communicates with reminders background API
type HTTPClient struct {
	client     *http.Client
	BackendURI string
}

// NewHTTPClient creates a new instance of HTTPClient
func NewHTTPClient(backendURI string) *HTTPClient {
	return &HTTPClient{
		client:     &http.Client{},
		BackendURI: backendURI,
	}
}

// Create calls the create API endpoint
func (c HTTPClient) Create(title, message string, duration time.Duration) ([]byte, error) {
	requestBody := reminderBody{
		Title:    title,
		Message:  message,
		Duration: duration,
	}

	return c.apiCall(
		http.MethodPost,
		"/reminders",
		&requestBody,
		http.StatusCreated,
	)
}

// Edit calls the edit API endpoint
func (c HTTPClient) Edit(id string, title, message string, duration time.Duration) ([]byte, error) {
	requestBody := reminderBody{
		ID:       id,
		Title:    title,
		Message:  message,
		Duration: duration,
	}

	return c.apiCall(
		http.MethodPatch,
		"/reminders/"+id,
		&requestBody,
		http.StatusOK,
	)
}

// Fetch calls the fetch API endpoint
func (c HTTPClient) Fetch(ids []string) ([]byte, error) {
	idsSet := strings.Join(ids, ",")
	return c.apiCall(
		http.MethodGet,
		"/reminders/"+idsSet,
		nil,
		http.StatusOK,
	)

}

// Delete calls the fetch API endpoint
func (c HTTPClient) Delete(ids []string) error {
	idsSet := strings.Join(ids, ",")
	_, err := c.apiCall(
		http.MethodDelete,
		"/reminders/"+idsSet,
		nil,
		http.StatusNoContent,
	)
	return err
}

// Healthy checks whether a given host is up and running
func (c HTTPClient) Healthy(host string) bool {
	res, err := http.Get(host + "/health")
	if err != nil || res.StatusCode != http.StatusOK {
		return false
	}
	return true
}

func (c HTTPClient) apiCall(method, path string, body interface{}, resCode int) ([]byte, error) {
	bs, err := json.Marshal(body)
	if err != nil {
		e := wrapError("could not marshal request body", err)
		return nil, e
	}

	req, err := http.NewRequest(method, c.BackendURI+path, bytes.NewReader(bs))
	if err != nil {
		e := wrapError("could not create request", err)
		return []byte{}, e
	}

	res, err := c.client.Do(req)
	if err != nil {
		e := wrapError("error making http call", err)
		return []byte{}, e
	}

	resBody, err := c.readResBody(res.Body)
	if err != nil {
		return []byte{}, err
	}

	if res.StatusCode != resCode {
		if len(resBody) > 0 {
			fmt.Printf("response body:\n%s\n", resBody)
		}
		return []byte{}, fmt.Errorf("expected response code: %d, got %d", resCode, res.StatusCode)
	}

	return []byte(resBody), err
}

// readResBody reads response body
func (c HTTPClient) readResBody(b io.Reader) (string, error) {
	bs, err := ioutil.ReadAll(b)
	if err != nil {
		e := wrapError("could not read response body", err)
		return "", e
	}

	if len(bs) == 0 {
		return "", nil
	}

	var buff bytes.Buffer
	if err := json.Indent(&buff, bs, "", "/t"); err != nil {
		e := wrapError("could not indent json", err)
		return "", e
	}

	return buff.String(), nil
}
