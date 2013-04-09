// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package stormpath

import "errors"

var (
	BadResponse = errors.New("Bad response from server.")
)

type StormpathError struct {
	Status           int
	Message          string
	Code             int
	MoreInfo         string
	DeveloperMessage string
}
