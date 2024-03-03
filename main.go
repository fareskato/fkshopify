package fkshopify

import (
	"fmt"
)

type Shopify struct {
	storeUser       string
	storePassword   string
	storeName       string
	storeApiVersion string
	storeWebHookKey string
	hmacHeader      string
}

// init the store with store credentials so U can interact
// with shopify admin rest APIs
func New(su, sp, sn, apiV, swhk string) Shopify {
	return Shopify{
		storeUser:       su,
		storePassword:   sp,
		storeName:       sn,
		storeApiVersion: apiV,
		storeWebHookKey: swhk,
		hmacHeader:      "X-Shopify-Hmac-SHA256",
	}
}

// InitStoreUrl returns the base url of shopify admin api
// so U can append the requested resource and enjoy)
func (s Shopify) InitStoreUrl() string {
	return fmt.Sprintf("https://%s:%s@%s.myshopify.com/admin/api/%s", s.storeUser, s.storePassword, s.storeName, s.storeApiVersion)
}

// InitStoreHeaders return shopify needed headers(in case U need it)
func (s Shopify) InitStoreHeaders() map[string]string {
	return map[string]string{
		"Content-Type":           "application/json",
		"X-Shopify-Access-Token": s.storePassword,
	}
}

// func (s Shopify) InitWebHook(hm, k string) fkwebhooks.WebHook {
// 	return fkwebhooks.WebHook{
// 		HmacHeader: hm,
// 		Key:        k,
// 	}
// }
