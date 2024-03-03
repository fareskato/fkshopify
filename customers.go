package fkshopify

import (
	"fmt"
	"strings"

	"github.com/fareskato/fkshopify/fkhttp"
)

type shopifyCustomerAddress struct {
	ShopifyID  int    `json:"id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	AddressOne string `json:"address1,omitempty"`
	Country    string `json:"country,omitempty"`
	City       string `json:"city,omitempty"`
}

type shopifyCustomer struct {
	ShopifyID   int                      `json:"id"`
	FirstName   string                   `json:"first_name,omitempty"`
	LastName    string                   `json:"last_name,omitempty"`
	Email       string                   `json:"email,omitempty"`
	Currency    string                   `json:"currency,omitempty"`
	OrdersCount int                      `json:"orders_count"`
	TotalSpent  string                   `json:"total_spent,omitempty"`
	Addresses   []shopifyCustomerAddress `json:"addresses,omitempty"`
}

type shopifyCustomersResponse struct {
	Customers []shopifyCustomer `json:"customers"`
}

func (s Shopify) GetAllShopifyCustomers() ([]shopifyCustomer, error) {
	var customers []shopifyCustomer
	var customersRes shopifyCustomersResponse
	customersUrl := fmt.Sprintf("%s/customers.json?limit=250", s.InitStoreUrl())
	for {
		customersres, res, err := fkhttp.HttpShopifyGet(customersRes, customersUrl)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customersres.Customers...)
		linkHeader := res.Header.Get("Link")
		if linkHeader == "" {
			break
		}
		if strings.Contains(linkHeader, `rel="next"`) {
			pageInfo := fkhttp.ExtractPageInfoFromLinkHeader(linkHeader)
			if pageInfo == "" {
				break
			}
			customersUrl = fmt.Sprintf("%s&page_info=%s", customersUrl, pageInfo)
		} else {
			break
		}
	}
	return customers, nil
}
