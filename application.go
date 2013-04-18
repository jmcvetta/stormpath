// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package stormpath

import (
	"encoding/base64"
	"github.com/jmcvetta/restclient"
	"log"
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

// CreateAccount creates a new account accessible to the application.
func (app *Application) CreateAccount(template Account) (Account, error) {
	/*
		data := &map[string]string{
			"username":  username,
			"password":  password,
			"email":     email,
			"surname":   surname,
			"givenName": givenName,
		}
	*/
	url := app.Href + "/accounts"
	acct := Account{}
	e := new(StormpathError)
	rr := restclient.RequestResponse{
		Userinfo: app.userinfo(),
		Url:      url,
		Method:   "POST",
		Data:     &template,
		Result:   &acct,
		Error:    e,
	}
	status, err := restclient.Do(&rr)
	if err != nil {
		return acct, err
	}
	acct.app = app
	if status != 201 {
		log.Println(status)
		log.Println(e)
		return acct, BadResponse
	}
	return acct, nil
}

// Authenticate with Stormpath using supplied credentials.  Username may be
// either a username or the user's email.
func (app *Application) Authenticate(username, password string) (Account, error) {
	acct := Account{}
	s := username + ":" + password
	value := base64.URLEncoding.EncodeToString([]byte(s))
	m := map[string]string{
		"type":  "basic",
		"value": value,
	}
	loginUrl := app.Href + "/loginAttempts"
	var res struct {
		Account struct {
			Href string `json:"href"`
		} `json:"account"`
	}
	e := new(StormpathError)
	rr := restclient.RequestResponse{
		Userinfo: app.userinfo(),
		Url:      loginUrl,
		Method:   "POST",
		Data:     &m,
		Result:   &res,
		Error:    &e,
	}
	status, err := restclient.Do(&rr)
	if err != nil {
		return acct, err
	}
	if status != 200 {
		log.Println(status)
		log.Println(res)
		log.Println(e)
		return acct, InvalidUsernamePassword
	}
	return app.GetAccount(res.Account.Href)
}

func (app *Application) GetAccount(href string) (Account, error) {
	acct := Account{}
	e := new(StormpathError)
	rr := restclient.RequestResponse{
		Userinfo: app.userinfo(),
		Url:      href,
		Method:   "GET",
		Result:   &acct,
		Error:    e,
	}
	status, err := restclient.Do(&rr)
	if err != nil {
		return acct, err
	}
	acct.app = app
	if status != 200 {
		log.Println(status)
		log.Println(e)
		return acct, BadResponse
	}
	return acct, nil
}
