package fkwebhooks

// type webHookProductVariant struct {
// 	ID              int    `json:"id"`
// 	SKU             string `json:"sku,omitempty"`
// 	Barcode         string `json:"barcode,omitempty"`
// 	Price           string `json:"price,omitempty"`
// 	CompareAtPrice  string `json:"compare_at_price,omitempty"`
// 	Qty             int    `json:"inventory_quantity"`
// 	InventoryItemID int    `json:"inventory_item_id"`
// }

// type webHookProductImage struct {
// 	Src string `json:"src,omitempty"`
// }

// type webHookProduct struct {
// 	ID          int                     `json:"id"`
// 	Title       string                  `json:"title"`
// 	BodyHTML    string                  `json:"body_html,omitempty"`
// 	PublishedAt string                  `json:"published_at"`
// 	Vendor      string                  `json:"vendor"`
// 	Status      string                  `json:"status,omitempty"`
// 	Tags        string                  `json:"tags,omitempty"`
// 	Variants    []webHookProductVariant `json:"variants"`
// 	Images      []webHookProductImage   `json:"images,omitempty"`
// }

/*
func NewProductWebHook(key, hmacHeader string, reqBody io.ReadCloser) (*webHookProduct, error) {
	var whPayload webHookProduct
	if key == "" {
		return nil, errors.New("missing webhook key")
	}
	if hmacHeader == "" {
		return nil, errors.New("missing X-Shopify-Hmac-SHA256 header")
	}
	data, err := io.ReadAll(reqBody)
	if err != nil {
		return nil, err
	}
	if !VerifyWebHook(data, hmacHeader, key) {
		return nil, errors.New("webhook verification error, please check if the webhook key is valid")
	}
	err = json.Unmarshal(data, &whPayload)
	if err != nil {
		return nil, err
	}
	return &whPayload, nil
}
*/
/*
func NewProductWebHook[T any](key, hmacHeader string, reqBody io.ReadCloser) (*T, error) {
	var whPayload T
	if key == "" {
		return nil, errors.New("missing webhook key")
	}
	if hmacHeader == "" {
		return nil, errors.New("missing X-Shopify-Hmac-SHA256 header")
	}
	data, err := io.ReadAll(reqBody)
	if err != nil {
		return nil, err
	}
	if !VerifyWebHook(data, hmacHeader, key) {
		return nil, errors.New("webhook verification error, please check if the webhook key is valid")
	}
	err = json.Unmarshal(data, &whPayload)
	if err != nil {
		return nil, err
	}
	return &whPayload, nil
}
*/
