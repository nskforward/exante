package exante

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
	"unicode/utf8"
)

const (
	accessTokenRefreshInterval = time.Hour
)

type jwt struct {
	ClientID      string   `json:"iss"`
	ApplicationID string   `json:"sub"`
	IssuedAt      int64    `json:"iat"`
	Expiration    int64    `json:"exp"`
	Rights        []string `json:"aud"`
}

func (client *Client) refreshAccessToken() {
	if client.accessTokenExpiration < time.Now().Add(time.Minute).Unix() {
		client.accessTokenExpiration = time.Now().Add(accessTokenRefreshInterval).Unix()
		client.accessToken = generateAccessToken(client.clientID, client.appID, client.sharedKey, client.accessTokenExpiration)
	}
}

func generateAccessToken(clientID, appID, sharedKey string, expiration int64) string {
	data, err := json.Marshal(jwt{
		ClientID:      clientID,
		ApplicationID: appID,
		Rights:        []string{"symbols", "ohlc", "feed", "change", "crossrates", "orders", "summary", "accounts", "transactions"},
		IssuedAt:      time.Now().Unix(),
		Expiration:    expiration,
	})
	if err != nil {
		panic(fmt.Errorf("cannot marshal jwt token: %w", err))
	}

	payload := base64.RawStdEncoding.EncodeToString(toUTF8(string(data)))
	var buffer bytes.Buffer
	buffer.WriteString("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.")
	buffer.WriteString(payload)
	sha := hmac256(sharedKey, buffer.Bytes())
	buffer.WriteRune('.')
	buffer.WriteString(sha)
	return string(buffer.Bytes())
}

func toUTF8(source string) []byte {
	bs := make([]byte, len(source)*utf8.UTFMax)
	count := 0
	for _, r := range source {
		count += utf8.EncodeRune(bs[count:], r)
	}
	bs = bs[:count]
	return bs
}

func hmac256(secret string, source []byte) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(source)
	return base64.RawStdEncoding.EncodeToString(h.Sum(nil))
}
