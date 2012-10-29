// Copyright (c) 2012 Jason McVetta.  This is Free Software, released under the 
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.

package http

import (
	"errors"
	"github.com/jmcvetta/stormpath/util"
	"net/url"
	"strconv"
	"strings"
)

type Request struct {
	method      string
	href        string
	queryString map[string]string
	headers     map[string]string
	body        string
}

type RequestConfig struct {
	HttpMethod  string
	Href        string
	QueryString map[string]string
	HttpHeaders map[string]string
	Body        string
}

func NewRequest(c RequestConfig) (r *Request, err error) {
	split := strings.Split(c.Href, "?")
	if len(split) > 1 {
		r.href = split[0]
		qsStr := split[1]
		for _, pair := range strings.Split(qsStr, "&") {
			kv := strings.Split(pair, "=")
			if len(kv) == 2 {
				r.queryString[kv[0]] = kv[1]
			} else {
				// We'll return the error, but first we will go ahead and 
				// construct as much of the Request object as we can.
				err = errors.New("Invalid query string: " + qsStr)
			}
		}
	} else {
		r.href = c.Href
	}
	r.method = strings.ToUpper(c.HttpMethod)
	r.headers = c.HttpHeaders
	r.headers["Content-Length"] = strconv.Itoa((len(c.Body)))
	return r, err

}

// ResourceUri returns the Request's href as a URL object.
func (r *Request) ResourceUri() (*url.URL, error) {
	// Is url.ParseRequestURI() more appropriate than url.Parse()?
	return url.ParseRequestURI(r.href)
}

// ToSQueryString returns a string representation of this Request's queryString
// map.
func (r *Request) ToSQueryString(canonical bool) string {
	var result string
	for k, v := range r.queryString {
		encK := util.EncodeUrl(k, false, canonical)
		encV := util.EncodeUrl(v, false, canonical)
		if result != "" {
			result += "&"
		}
		result += encK + "=" + encV
	}
	return result
}
