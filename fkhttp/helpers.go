package fkhttp

import "regexp"

func ExtractPageInfoFromLinkHeader(linkHeader string) string {
	re := regexp.MustCompile(`page_info=([^;&>]+)`)
	matches := re.FindStringSubmatch(linkHeader)
	if len(matches) >= 2 {
		return matches[1]
	}
	return ""
}
