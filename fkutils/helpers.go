package fkutils

import (
	"fmt"
	"html"
	"regexp"
	"strings"
	"time"
)

func SanitizeHTML(input string) string {
	// Remove HTML tags
	re := regexp.MustCompile(`<[^>]*>`)
	plainText := re.ReplaceAllString(input, " ")

	// Decode HTML entities
	decodedText := html.UnescapeString(plainText)

	// Remove extra whitespaces
	cleanedText := strings.Join(strings.Fields(decodedText), " ")

	return cleanedText
}

func FormatTimeAgo(diff time.Duration) string {
	if diff.Seconds() < 60 {
		return fmt.Sprintf("%.0f seconds ago", diff.Seconds())
	} else if diff.Minutes() < 60 {
		return fmt.Sprintf("%.0f minutes ago", diff.Minutes())
	} else if diff.Hours() < 24 {
		return fmt.Sprintf("%.0f hours ago", diff.Hours())
	} else {
		days := int(diff.Hours() / 24)
		return fmt.Sprintf("%d days ago", days)
	}
}
