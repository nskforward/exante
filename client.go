package exante

import (
	"context"
	"fmt"
)

type Client struct {
	ctx                   context.Context
	accountID             string
	clientID              string
	appID                 string
	sharedKey             string
	serverAddr            string
	accessToken           string
	accessTokenExpiration int64
}

func NewClient(ctx context.Context, accountID, serverAddr, clientID, appID, sharedKey string) (*Client, error) {
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

	return client, nil
}
