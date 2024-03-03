package fkshopify

import (
	"fmt"
	"net/http"

	"github.com/fareskato/fkshopify/fkhttp"
)

type shopifyCollection struct {
	ShopifyID     int    `json:"id"`
	Title         string `json:"title"`
	ProductsCount int    `json:"products_count"`
}

type shopifyCollectionResponse struct {
	Collection shopifyCollection `json:"collection"`
}

func (s Shopify) GetShopifyCollection(id string) (*shopifyCollection, error) {
	var colRes shopifyCollectionResponse
	colUrl := fmt.Sprintf("%s/collections/%s.json", s.InitStoreUrl(), id)
	colres, res, err := fkhttp.HttpShopifyGet(colRes, colUrl)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, err
	}
	return &colres.Collection, nil
}
