// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package stormpath

import (
	"github.com/jmcvetta/restclient"
	"log"
)

// An Href is a container for URL links
type Href struct {
	Href string `json:"href"`
}

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
	Groups     Href   `json:"groups,omitempty"`     // A link to the groups that the account belongs to.
	Directory  Href   `json:"directory,omitempty"`  // A link to the owning directory.
	Tenant     Href   `json:"tenant,omitempty"`     // A link to the tenant owning the directory the group belongs to.
	app        *Application
}

// Delete removes an account from Stormpath.
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

// Update saves the account to Stormpath.
func (a *Account) Update() error {
	e := new(StormpathError)
	rr := restclient.RequestResponse{
		Userinfo: a.app.userinfo(),
		Url:      a.Href,
		Method:   "POST",
		Data: &a,
		Error:    e,
	}
	status, err := restclient.Do(&rr)
	if err != nil {
		return err
	}
	if status != 200 {
		log.Println(status)
		log.Println(e)
		return BadResponse
	}
	return nil // Successful update
}
