// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package stormpath

import (
	"net/url"
)

// An Application in represents a real world application that can communicate
// with and be provisioned by Stormpath. After it is defined, an application is
// mapped to one or more directories or groups, whose users are then granted
// access to the application.
type Application struct {
	Href      string // Stormpath URL for this application
	ApiId     string // Stormpath API key ID
	ApiSecret string // Stormpath API key secret
}

func (a *Application) userinfo() *url.Userinfo {
	return url.UserPassword(a.ApiId, a.ApiSecret)
}
