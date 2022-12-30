package http

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type JWT struct {
	ID        string   `json:"jti,omitempty"`
	Issuer    string   `json:"iss,omitempty"`
	Subject   string   `json:"sub,omitempty"`
	Audience  []string `json:"aud,omitempty"`
	ExpiresAt int64    `json:"exp,omitempty"`
	NotBefore int64    `json:"nbf,omitempty"`
	IssuedAt  int64    `json:"iat,omitempty"`
}

func GenerateJWT(claims JWT, secret []byte) string {
	data, err := json.Marshal(claims)
	if err != nil {
		panic(err)
	}
	var buffer bytes.Buffer
	buffer.WriteString("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.")
	buffer.WriteString(base64.RawURLEncoding.EncodeToString(data))
	h := hmac.New(sha256.New, secret)
	h.Write(buffer.Bytes())
	buffer.WriteRune('.')
	buffer.WriteString(base64.RawURLEncoding.EncodeToString(h.Sum(nil)))
	return buffer.String()
}
