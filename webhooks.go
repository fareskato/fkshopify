package fkshopify

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"

	"github.com/fareskato/fkshopify/fkutils"
)

func (s Shopify) verifyWebHook(data []byte) bool {
	secret := []byte(s.storeWebHookKey)
	computedHMAC := hmac.New(sha256.New, secret)
	computedHMAC.Write(data)
	expectedHMAC := computedHMAC.Sum(nil)
	decodedHMAC, err := base64.StdEncoding.DecodeString(s.hmacHeader)
	if err != nil {
		return false
	}
	return hmac.Equal(expectedHMAC, decodedHMAC)
}
func WebHookCreateEntity[T any](s Shopify, t T, reqBody []byte) (*T, error) {
	if s.storeWebHookKey == "" {
		return nil, fkutils.ErrWebHookMissed
	}
	if s.hmacHeader == "" {
		return nil, fkutils.ErrHmacHeaderMissed
	}

	if !s.verifyWebHook(reqBody) {
		return nil, fkutils.ErrInvalidWebHookKey
	}
	err := json.Unmarshal(reqBody, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
