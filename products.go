package fkshopify

import (
	"fmt"
	"strings"

	"github.com/fareskato/fkshopify/fkhttp"
)

type shopifyProductImage struct {
	Src string `json:"src"`
}

type shopifyProductVariant struct {
	ShopifyID       int    `json:"id"`
	SKU             string `json:"sku"`
	Barcode         string `json:"barcode,omitempty"`
	Price           string `json:"price"`
	CompareAtPrice  string `json:"compare_at_price,omitempty"`
	Qty             int    `json:"inventory_quantity"`
	InventoryItemID int    `json:"inventory_item_id"`
}

type shopifyProduct struct {
	ShopifyID   int                     `json:"id"`
	Title       string                  `json:"title"`
	PublishedAt string                  `json:"published_at"`
	BodyHtml    string                  `json:"body_html"`
	Vendor      string                  `json:"vendor"`
	Status      string                  `json:"status"`
	Tags        string                  `json:"tags,omitempty"`
	Variants    []shopifyProductVariant `json:"variants,omitempty"`
	Images      []shopifyProductImage   `json:"images"`
}

type shopifyProductsResponse struct {
	Products []shopifyProduct
}

type shopifyProductsFetchOptions struct {
	Fields       string
	CollectionID string
}

// GetAllShopifyProducts takes options(shopifyProductsFetchOptions)
// so U can defined specific product data to fetch like title, vendor for example:
// id,title,tags,variants,vendor,published_at,body_html,status,images
// also if U want U can fetch products of specific collection so U can pass
// the collection id
func (s Shopify) GetAllShopifyProducts(options shopifyProductsFetchOptions) ([]shopifyProduct, error) {
	var products []shopifyProduct
	var productsRes shopifyProductsResponse
	var productsUrl string
	baseUrl := s.InitStoreUrl()
	if options.CollectionID != "" {
		productsUrl = fmt.Sprintf("%s/products.json?collection_id=%s&limit=250&fields=%s", baseUrl, options.CollectionID, options.Fields)
	} else {
		productsUrl = fmt.Sprintf("%s/products.json?&limit=250&fields=%s", baseUrl, options.Fields)
	}
	for {
		prsRes, res, err := fkhttp.HttpGet(productsRes, productsUrl)
		if err != nil {
			return nil, err
		}
		products = append(products, prsRes.Products...)
		linkHeader := res.Header.Get("Link")
		if linkHeader == "" {
			break
		}
		pageInfo := fkhttp.ExtractPageInfoFromLinkHeader(linkHeader)
		if pageInfo == "" {
			break
		}
		nextProductsUrl := fmt.Sprintf("%s/products.json?limit=250&fileds=%s", baseUrl, options.Fields)
		productsUrl = fmt.Sprintf("%s&page_info=%s", nextProductsUrl, pageInfo)
		if strings.Contains(linkHeader, `rel="next"`) {
			break
		}
	}
	return products, nil
}
