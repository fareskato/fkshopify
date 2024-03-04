package fkshopify

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"

	"github.com/fareskato/fkshopify/fkutils"
)

// verifyWebHook takes the request bosdy and the X-Shopify-Hmac-SHA256 header sended
// by the web hook so U need to extract the X-Shopify-Hmac-SHA256 header
// from the request then pass it to the function
func (s Shopify) verifyWebHook(data []byte, hmacHeader string) bool {
	secret := []byte(s.storeWebHookKey)
	computedHMAC := hmac.New(sha256.New, secret)
	computedHMAC.Write(data)
	expectedHMAC := computedHMAC.Sum(nil)
	decodedHMAC, err := base64.StdEncoding.DecodeString(hmacHeader)
	if err != nil {
		return false
	}
	return hmac.Equal(expectedHMAC, decodedHMAC)
}

// WebHookHandleActionEntity uses to habdle web hooks for sources like product, order ....etc
// this function handle create,update and delete web hooks
// in case of create or update the sended data from shopify conatins resource filels like
// for example:
// {
// "title": "Example T-Shirt",
// "updated_at": "2021-12-31T19:00:00-05:00",
// "vendor": "Acme",
// .... other fileds
// }
// in case of delete the sended data contains only the deleted item id like:
//
//	{
//	  "id": 788032119674292900
//	}
//
// for more unfo refer to: https://shopify.dev/docs/api/admin-rest/2024-01/resources/webhook#event-topics
func WebHookHandleActionEntity[T any](s Shopify, t T, hmacHeader string, reqBody []byte) (*T, error) {
	if s.storeWebHookKey == "" {
		return nil, fkutils.ErrWebHookMissed
	}
	if hmacHeader == "" {
		return nil, fkutils.ErrHmacHeaderMissed
	}

	if !s.verifyWebHook(reqBody, hmacHeader) {
		return nil, fkutils.ErrInvalidWebHookKey
	}
	err := json.Unmarshal(reqBody, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
