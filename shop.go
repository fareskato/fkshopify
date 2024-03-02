package fkshopify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Shop struct {
	ShopifyID     int    `json:"id"`
	Name          string `json:"name"`
	Phone         string `json:"phone,omitempty"`
	Email         string `json:"email,omitempty"`
	AddressOne    string `json:"address1,omitempty"`
	Country       string `json:"country,omitempty"`
	CountryName   string `json:"country_name,omitempty"`
	City          string `json:"city,omitempty"`
	Currency      string `json:"currency"`
	CustomerEmail string `json:"customer_email,omitempty"`
}

type ShopResponse struct {
	Shop Shop `json:"shop"`
}

func (s Shopify) GetShopData() (*Shop, error) {
	res, err := http.Get(fmt.Sprintf("%s/shop.json", s.InitStoreUrl()))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var shopRes ShopResponse
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &shopRes); err != nil {
		return nil, err
	}
	return &shopRes.Shop, nil
}
