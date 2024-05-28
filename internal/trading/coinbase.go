package trading

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CoinbaseAPI struct {
	BaseURL    string
	KeyName    string
	secret     string
	HTTPClient *http.Client
}

type CoinbaseAPIConfig struct {
	KeyName string
	Secret  string
}

func NewCoinbaseAPI(coinbaseApiConfig CoinbaseAPIConfig) *CoinbaseAPI {
	return &CoinbaseAPI{
		BaseURL: "api.coinbase.com/api/v3/brokerage",
		KeyName: coinbaseApiConfig.KeyName,
		secret:  coinbaseApiConfig.Secret,
		HTTPClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (cb *CoinbaseAPI) Request(method string, url string, params interface{}, result interface{}) (res *http.Response, err error) {
	var data []byte
	body := bytes.NewReader(make([]byte, 0))

	if params != nil {
		data, err = json.Marshal(params)
		if err != nil {
			return res, err
		}

		body = bytes.NewReader(data)
	}

	uri := fmt.Sprintf("https://%s%s", cb.BaseURL, url)
	req, err := http.NewRequest(method, uri, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	jwt, err := cb.getJWTHeader(method, url)

	if err != nil {
		return nil, fmt.Errorf("could not create JWT header: %v", err)
	}
	fmt.Println(jwt)

	req.Header.Add("Authorization", jwt)

	res, err = cb.HTTPClient.Do(req)

	if err != nil {
		return res, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		coinbaseError := Error{}
		if err := json.NewDecoder(res.Body).Decode(&coinbaseError); err != nil {
			return res, err
		}
		return res, error(coinbaseError)
	}

	if result != nil {
		if err = json.NewDecoder(res.Body).Decode(result); err != nil {
			return res, err
		}
	}

	return res, nil
}

func (cb *CoinbaseAPI) getJWTHeader(method, url string) (string, error) {
	jwt, err := cb.buildJWT(method, url)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Bearer %s", jwt), nil
}
