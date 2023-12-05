package input

import "regexp"

// Checks that the passed url is a valid
// youtube url to pass to yt-dlt
func IsValidYoutubeUrl(url string) bool {
	match, _ := regexp.MatchString(`^https:\/\/www\.youtube\.com\/watch\?v=(?:\w|-|=|&)*$`, url)
	return match
}
