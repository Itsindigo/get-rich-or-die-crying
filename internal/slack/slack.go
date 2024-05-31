package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Slack struct {
	HookURL    string
	HTTPClient *http.Client
}

func (s *Slack) request(method string, url string, body io.Reader, result *string) (res *http.Response, err error) {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err = s.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return res, fmt.Errorf("error reading response body: %w", err)
	}

	if result != nil {
		*result = string(bodyBytes)
	}

	if res.StatusCode != 200 {
		return res, fmt.Errorf("request failed: %d, reason: %s", res.StatusCode, *result)
	}

	return res, nil
}

func (s *Slack) SendMessage(ctx context.Context, blocks Blocks) (string, error) {
	var response string

	data, err := json.Marshal(blocks)

	if err != nil {
		return "", err
	}

	body := bytes.NewReader(data)

	_, err = s.request("POST", s.HookURL, body, &response)

	if err != nil {

		return response, err
	}

	return response, nil
}

func NewSlack(webhookID string) *Slack {
	hookUrl := fmt.Sprintf("https://hooks.slack.com/services/%s", webhookID)

	return &Slack{
		HookURL: hookUrl,
		HTTPClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}
