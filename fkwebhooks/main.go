package fkwebhooks

// import (
// 	"crypto/hmac"
// 	"crypto/sha256"
// 	"encoding/base64"
// 	"encoding/json"
// 	"io"

// 	"github.com/fareskato/fkshopify/fkutils"
// )

// type WebHook struct {
// 	HmacHeader string
// 	Key        string
// }

// func verifyWebHook(data []byte) bool {
// 	secret := []byte(wh.Key)
// 	computedHMAC := hmac.New(sha256.New, secret)
// 	computedHMAC.Write(data)
// 	expectedHMAC := computedHMAC.Sum(nil)

// 	decodedHMAC, err := base64.StdEncoding.DecodeString(wh.HmacHeader)
// 	if err != nil {
// 		return false
// 	}

// 	return hmac.Equal(expectedHMAC, decodedHMAC)
// }

// func (wh WebHook) WebHookCreateEntity(reqBody io.ReadCloser) (any, error) {
// 	var whPayload any
// 	if wh.Key == "" {
// 		return nil, fkutils.ErrWebHookMissed
// 	}
// 	if wh.HmacHeader == "" {
// 		return nil, fkutils.ErrHmacHeaderMissed
// 	}
// 	data, err := io.ReadAll(reqBody)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if !wh.verifyWebHook(data) {
// 		return nil, fkutils.ErrInvalidWebHookKey
// 	}
// 	err = json.Unmarshal(data, &whPayload)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &whPayload, nil
// }
