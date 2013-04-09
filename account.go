// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package stormpath

// Users are referred to as user account objects or accounts. The username and
// email fields for accounts are unique within a directory and are used to log
// into applications.
type Account struct {
	Href       string `json:"href"`       // The resource fully qualified location URI
	Username   string `json:"username"`   // The username for the account. Must be unique across the owning directory. If not specified, the username will default to the email field.	1 < N <= 255 characters
	Email      string `json:"email"`      // The email address for the account. Must be unique across the owning directory.	1 < N <= 255 characters
	Password   string `json:"password"`   // The password for the account. Only include this property if setting or changing the account password.	1 < N <= 255 characters
	GivenName  string `json:"givenName"`  // The given (first) name for the account holder.	1 < N <= 255 characters
	MiddleName string `json:"middleName"` // The middle (second) name for the account holder.	1 < N <= 255 characters
	Surname    string `json:"surname"`    // The surname (last name) for the account holder.	1 < N <= 255 characters
	Status     string `json:"status"`     // Enabled accounts are able to login to their assigned applications.	Enum	enabled,disabled
	Groups     string `json:"groups"`     // A link to the groups that the account belongs to.
	Directory  string `json:"directory"`  // A link to the owning directory.
	Tenant     string `json:"tenant"`     // A link to the tenant owning the directory the group belongs to.
}

func (a *Application) CreateAccount(new Account) (Account, error) {
	url := a.Href + "/accounts"
	acct := Account{}
	e := new(interface{})
	rr := restclient.RequestResponse{
Url: url,
Method: "POST",
	Data: &new,
	Result: &acct,
	Error: &e,
}
status, err := restclient.Do(&rr)
if err != nil {
return acct, err
}
if status != 409 {
return acct, BadResponse
}
return acct, nil
}
