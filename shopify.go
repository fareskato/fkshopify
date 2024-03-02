package fkshopify

import "fmt"

type Shopify struct {
	storeUser       string
	storePassword   string
	storeName       string
	storeApiVersion string
	storeWebHookKey string
}

func New(su, sp, sn, apiV, swhk string) Shopify {
	return Shopify{
		storeUser:       su,
		storePassword:   sp,
		storeName:       sn,
		storeApiVersion: apiV,
		storeWebHookKey: swhk,
	}
}

func (s Shopify) InitStoreUrl() string {
	return fmt.Sprintf("https://%s:%s@%s.myshopify.com/admin/api/%s", s.storeUser, s.storePassword, s.storeName, s.storeApiVersion)
}

func (s Shopify) InitStoreHeaders() map[string]string {
	return map[string]string{
		"Content-Type":           "application/json",
		"X-Shopify-Access-Token": s.storePassword,
	}
}
