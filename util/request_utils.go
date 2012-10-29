package util

import (
	"net/url"
	"strings"
)

/*
In the Ruby and Java SDKs, these functions are class methods on RequestUtils. 
That doesn't fit well with Go's more functional idiom, so in this implimentation
they are plain old functions.
*/

// EncodeUrl encodes a URL (durrrrh)
func EncodeUrl(value string, path, canonical bool) string {
	// Note that url.QueryEscape() does not work exactly like Ruby's URI.escape()
	encoded := url.QueryEscape(value)
	if canonical {
		strmap := map[string]string{
			"+":   "%20",
			"*":   "%2A",
			"%7E": "~",
		}
		for key, value := range strmap {
			encoded = strings.Replace(encoded, key, value, -1)
		}
	}
	if path {
		s := "%2F"
		if strings.Contains(encoded, s) {
			encoded = strings.Replace(encoded, s, "/", -1)
		}
	}
	return encoded
}

// DefaultPort returns true if the specified URI uses a standard port (i.e. 
// http == 80 or https == 443), false otherwise.
func DefaultPort(u *url.URL) bool {
	scheme := u.Scheme
	parts := strings.Split(u.Host, ":")
	var port string
	switch len(parts) {
	case 1:
		port = ""
	case 2:
		port = parts[1]
	default:
		// Would it be better to return an error here?
		return false
	}
	result := port == "" || port == "0" || (port == "80" && scheme == "http") || (port == "443" && scheme == "https")
	return result // Success!
}
