package chainviewer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/g3co/twtest/pkg/ethrpc"
	"net/http"
)

const nodeHost = "https://cloudflare-eth.com"

func sendRequest(rq ethrpc.Request, resp interface{}) error {
	b := new(bytes.Buffer)

	err := json.NewEncoder(b).Encode(rq)
	if err != nil {
		return fmt.Errorf("client: could not encode request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, nodeHost, b)
	if err != nil {
		return fmt.Errorf("client: could not create request: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("client: error making http request: %w", err)
	}

	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return fmt.Errorf("client: could not decode response body: %w", err)
	}

	return nil
}
