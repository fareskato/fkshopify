package fkutils

import "errors"

var (
	ErrWebHookMissed     = errors.New("missing webhook key")
	ErrHmacHeaderMissed  = errors.New("missing X-Shopify-Hmac-SHA256 header")
	ErrInvalidWebHookKey = errors.New("webhook verification error, please check if the webhook key is valid")
)
