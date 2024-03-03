package fkwebhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/fareskato/fkshopify/fkutils"
)

func verifyWebHook(data []byte, hmacHeader string, webhookKey string) bool {
	secret := []byte(webhookKey)
	computedHMAC := hmac.New(sha256.New, secret)
	computedHMAC.Write(data)
	expectedHMAC := computedHMAC.Sum(nil)

	decodedHMAC, err := base64.StdEncoding.DecodeString(hmacHeader)
	if err != nil {
		return false
	}

	return hmac.Equal(expectedHMAC, decodedHMAC)
}

func WebHookCreateEntity[T any](key, hmacHeader string, reqBody io.ReadCloser) (*T, error) {
	var whPayload T
	if key == "" {
		return nil, fkutils.ErrWebHookMissed
	}
	if hmacHeader == "" {
		return nil, fkutils.ErrHmacHeaderMissed
	}
	data, err := io.ReadAll(reqBody)
	if err != nil {
		return nil, err
	}
	if !verifyWebHook(data, hmacHeader, key) {
		return nil, fkutils.ErrInvalidWebHookKey
	}
	err = json.Unmarshal(data, &whPayload)
	if err != nil {
		return nil, err
	}
	return &whPayload, nil
}
