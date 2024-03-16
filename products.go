package fkshopify

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/fareskato/fkshopify/fkhttp"
)

type ShopifyProductImage struct {
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
	Images      []ShopifyProductImage   `json:"images"`
}

type shopifyProductsResponse struct {
	Products []shopifyProduct
}

// ShopifyProductsFetchOptions defines options when U fetch products from shopify store
type ShopifyProductsFetchOptions struct {
	Fields       string
	CollectionID string
}

// GetAllShopifyProducts takes options(shopifyProductsFetchOptions)
// so U can defined specific product data to fetch like title, vendor for example:
// id,title,tags,variants,vendor,published_at,body_html,status,images
// also if U want U can fetch products of specific collection so U can pass
// the collection id
func (s Shopify) GetAllShopifyProducts(options ShopifyProductsFetchOptions) ([]shopifyProduct, error) {
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
		prsRes, res, err := fkhttp.HttpShopifyGet(productsRes, productsUrl)
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
		if strings.Contains(linkHeader, `rel="previous"`) {
			break
		}
	}
	return products, nil
}

// PushProductImagesToShopify delets the old images and create the passed images slice
func (s Shopify) PushProductImagesToShopify(ctx context.Context, id int, images []ShopifyProductImage) error {
	// prepare data
	type Data map[string]any
	data := Data{
		"product": map[string]any{
			"id":     id,
			"images": images,
		},
	}
	url := fmt.Sprintf("%s/products/%d.json", s.InitStoreUrl(), id)
	resp, err := fkhttp.HttpShopifyRequestWithHeaders(ctx, http.MethodPut, url, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// PushProductCostToshopify will update the cost per item value.
func (s Shopify) PushProductCost(ctx context.Context, iiID int, cost string) error {
	// prepare data
	type Data map[string]any
	data := Data{
		"inventory_item": map[string]any{
			"inventory_item_id": iiID,
			"cost":              cost,
		},
	}
	url := fmt.Sprintf("%s/inventory_items/%d.json", s.InitStoreUrl(), iiID)
	resp, err := fkhttp.HttpShopifyRequestWithHeaders(ctx, http.MethodPut, url, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// all good
	return nil
}

// PushProductQty updated the product qty which needed the location id.
func (s Shopify) PushProductQty(ctx context.Context, locID, iiID, qty int) error {
	// prepare data
	type Data map[string]any
	data := Data{
		"location_id":       locID,
		"inventory_item_id": iiID,
		"available":         qty,
	}
	url := fmt.Sprintf("%s/inventory_levels/set.json", s.InitStoreUrl())
	resp, err := fkhttp.HttpShopifyRequestWithHeaders(ctx, http.MethodPost, url, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// all good
	return nil
}
