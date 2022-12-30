package http

import (
	"time"
)

const (
	accessTokenRefreshInterval = time.Hour
)

func (client *Client) refreshAccessToken() {
	if client.accessTokenExpiration < time.Now().Add(time.Minute).Unix() {
		client.accessTokenExpiration = time.Now().Add(accessTokenRefreshInterval).Unix()
		client.accessToken = generateAccessToken(client.clientID, client.appID, client.sharedKey, client.accessTokenExpiration)
	}
}

func generateAccessToken(clientID, appID, sharedKey string, expiration int64) string {
	return GenerateJWT(JWT{
		Issuer:    clientID,
		Subject:   appID,
		Audience:  []string{"symbols", "ohlc", "feed", "change", "crossrates", "orders", "summary", "accounts", "transactions"},
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: expiration,
	}, []byte(sharedKey))
}
