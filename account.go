// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package stormpath

import (
	"github.com/jmcvetta/restclient"
	"log"
)

// Users are referred to as user account objects or accounts. The username and
// email fields for accounts are unique within a directory and are used to log
// into applications.
type Account struct {
	Href       string `json:"href,omitempty"`       // The resource fully qualified location URI
	Username   string `json:"username,omitempty"`   // The username for the account. Must be unique across the owning directory. If not specified, the username will default to the email field.	1 < N <= 255 characters
	Email      string `json:"email,omitempty"`      // The email address for the account. Must be unique across the owning directory.	1 < N <= 255 characters
	Password   string `json:"password,omitempty"`   // The password for the account. Only include this property if setting or changing the account password.	1 < N <= 255 characters
	GivenName  string `json:"givenName,omitempty"`  // The given (first) name for the account holder.	1 < N <= 255 characters
	MiddleName string `json:"middleName,omitempty"` // The middle (second) name for the account holder.	1 < N <= 255 characters
	Surname    string `json:"surname,omitempty"`    // The surname (last name) for the account holder.	1 < N <= 255 characters
	Status     string `json:"status,omitempty"`     // Enabled accounts are able to login to their assigned applications.	Enum	enabled,disabled
	Groups     string `json:"groups,omitempty"`     // A link to the groups that the account belongs to.
	Directory  string `json:"directory,omitempty"`  // A link to the owning directory.
	Tenant     string `json:"tenant,omitempty"`     // A link to the tenant owning the directory the group belongs to.
	app        *Application
}

// func (a *Application) CreateAccount(username, password, email, surname, givenName string) (*Account, error) {
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

func (a *Account) Delete() error {
	e := new(StormpathError)
	rr := restclient.RequestResponse{
		Userinfo: a.app.userinfo(),
		Url:      a.Href,
		Method:   "DELETE",
		Error:    e,
	}
	status, err := restclient.Do(&rr)
	if err != nil {
		return err
	}
	if status != 204 {
		log.Println(status)
		log.Println(e)
		return BadResponse
	}
	return nil // Successful deletion
}
