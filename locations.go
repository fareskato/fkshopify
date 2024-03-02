package fkshopify

import (
	"fmt"
	"strings"

	"github.com/fareskato/fkshopify/fkhttp"
)

type Location struct {
	ShopifyID   int    `json:"id"`
	Name        string `json:"name"`
	Active      bool   `json:"active"`
	AddressOne  string `json:"address1,omitempty"`
	AddressTwo  string `json:"address2,omitempty"`
	City        string `json:"city,omitempty"`
	Zip         string `json:"zip,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	CountryName string `json:"country_name,omitempty"`
}

type LocationResponse struct {
	Location Location `json:"location"`
}

type LocationsResponse struct {
	Locations []Location `json:"locations"`
}

func (s Shopify) GetAllShopifyLocations() ([]Location, error) {
	var locations []Location
	var locsRes LocationsResponse
	locationsUrl := fmt.Sprintf("%s/locations.json?limit=250", s.InitStoreUrl())
	for {
		locres, res, err := fkhttp.HttpGet(locsRes, locationsUrl)
		if err != nil {
			return nil, err
		}
		locations = append(locations, locres.Locations...)
		linkHeader := res.Header.Get("Link")
		if linkHeader == "" {
			break
		}
		if strings.Contains(linkHeader, `rel="next"`) {
			pageInfo := fkhttp.ExtractPageInfoFromLinkHeader(linkHeader)
			if pageInfo == "" {
				break
			}
			locationsUrl = fmt.Sprintf("%s&page_info=%s", locationsUrl, pageInfo)
		} else {
			break
		}
	}
	return locations, nil
}
