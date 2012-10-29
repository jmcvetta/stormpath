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

type RequestConfig struct {
	Method      string
	Href        string
	QueryString map[string]string // Do we have to call this QueryString even tho it's a map?
	Headers     map[string]string
	Body        string
}

type Request struct {
	RequestConfig
}

func NewRequest(c RequestConfig) (r *Request, err error) {
	split := strings.Split(c.Href, "?")
	if len(split) > 1 {
		r.Href = split[0]
		qsStr := split[1]
		for _, pair := range strings.Split(qsStr, "&") {
			kv := strings.Split(pair, "=")
			if len(kv) == 2 {
				r.QueryString[kv[0]] = kv[1]
			} else {
				// We'll return the error, but first we will go ahead and 
				// construct as much of the Request object as we can.
				err = errors.New("Invalid query string: " + qsStr)
			}
		}
	} else {
		r.Href = c.Href
	}
	r.Method = strings.ToUpper(c.Method)
	r.Headers = c.Headers
	r.Headers["Content-Length"] = strconv.Itoa((len(c.Body)))
	return r, err

}

// ResourceUri returns the Request's href as a URL object.
func (r *Request) ResourceUri() (*url.URL, error) {
	// Is url.ParseRequestURI() more appropriate than url.Parse()?
	return url.ParseRequestURI(r.Href)
}

// ToSQueryString returns a string representation of this Request's queryString
// map.
func (r *Request) ToSQueryString(canonical bool) string {
	var result string
	for k, v := range r.QueryString {
		encK := util.EncodeUrl(k, false, canonical)
		encV := util.EncodeUrl(v, false, canonical)
		if result != "" {
			result += "&"
		}
		result += encK + "=" + encV
	}
	return result
}
