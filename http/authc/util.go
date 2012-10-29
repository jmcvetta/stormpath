package authc

import (
	"net/url"
	"strings"
)


// encodeUrl encodes a URL to be signed by SAuthc1.  
//
// Based on https://github.com/stormpath/stormpath-sdk-ruby/blob/master/lib/stormpath-sdk/util/request_utils.rb#L35
func encodeUrl (value string, path, canonical bool) string {
	// Note that url.QueryEscape() does not work exactly like Ruby's URI.escape()
	encoded := url.QueryEscape(value)
	if canonical {
		strmap := map[string]string {
			"+": "%20",
			"*": "%2A",
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


// defaultPort checks whether a URL contains the default port for its scheme: 80
// for http, 443 for https, or no port specified.
func defaultPort(u *url.URL) bool {
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