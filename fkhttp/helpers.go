package fkhttp

import "regexp"

// ExtractPageInfoFromLinkHeader uses in case we have more than 250 items
// fetched from shopify (shopify limit) so we can use this function
// to gey the pagination link and get more than 250 items.
func ExtractPageInfoFromLinkHeader(linkHeader string) string {
	re := regexp.MustCompile(`page_info=([^;&>]+)`)
	matches := re.FindStringSubmatch(linkHeader)
	if len(matches) >= 2 {
		return matches[1]
	}
	return ""
}
