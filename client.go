package exante_http

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	ctx                   context.Context
	cancel                context.CancelFunc
	accountID             string
	clientID              string
	appID                 string
	sharedKey             string
	serverAddr            string
	accessToken           string
	accessTokenExpiration int64
}

func NewClient(ctx context.Context, accountID, serverAddr, clientID, appID, sharedKey string) (*Client, error) {

	if serverAddr == "" {
		return nil, fmt.Errorf("accountID must be defined")
	}

	if serverAddr == "" {
		return nil, fmt.Errorf("serverAddr must be defined")
	}

	if clientID == "" {
		return nil, fmt.Errorf("clientID must be defined")
	}

	if appID == "" {
		return nil, fmt.Errorf("appID must be defined")
	}

	if sharedKey == "" {
		return nil, fmt.Errorf("sharedKey must be defined")
	}

	client := &Client{
		ctx:        ctx,
		accountID:  accountID,
		clientID:   clientID,
		appID:      appID,
		sharedKey:  sharedKey,
		serverAddr: serverAddr,
	}

	accounts, err := client.GetUserAccounts()
	if err != nil {
		return nil, fmt.Errorf("ping request failed: %w", err)
	}

	found := false
	for _, acc := range accounts {
		if acc.AccountID == accountID {
			found = true
			if acc.Status != "Full" {
				fmt.Println("[warn] account", acc.AccountID, "has permissions:", acc.Status)
			}
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("accountId '%s' does not exist on server. Available accounts: %v", accountID, accounts)
	}

	scoped, cancel := context.WithCancel(ctx)
	client.ctx = scoped
	client.cancel = cancel

	return client, nil
}

func (client *Client) Close() {
	client.cancel()
}

func (client *Client) executeHTTPRequest(req *http.Request) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(client.ctx, 10*time.Second)
	defer cancel()
	req.WithContext(ctx)

	client.refreshAccessToken()
	req.Header.Add("Authorization", strings.Join([]string{"Bearer", client.accessToken}, " "))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 399 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("cannot read response body: %w", err)
		}
		return nil, fmt.Errorf("bad http response code: %s: %s", resp.Status, string(data))
	}

	return resp, err
}

func (client *Client) closeResponse(body io.Closer) {
	err := body.Close()
	if err != nil {
		fmt.Println("error: cannot correctly close response body stream")
	}
}
