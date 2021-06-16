package utils

import "strings"

func IsLink(input string) bool {
	if strings.HasPrefix(input, "http://") {
		return true
	}

	if strings.HasPrefix(input, "https://") {
		return true
	}

	return false
}

func GetWebsiteName(url string) string {
	if strings.Contains(url, "imgur.com") {
		return "Imgur"
	} else if strings.Contains(url, "youtube.com") {
		return "YouTube"
	} else if strings.Contains(url, "cdn.discordapp.com") {
		return "Discord CDN"
	} else {
		return "Сторонний сайт"
	}
}
