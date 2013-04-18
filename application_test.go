// Copyright (C) 2013 Jason McVetta, all rights reserved.

package stormpath

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	app := setupApplication(t)
	tmpl := createAccountTemplate(t)
	acct0, err := app.CreateAccount(tmpl)
	if err != nil {
		t.Error(err)
	}
	defer acct0.Delete()
	acct1, err := app.Authenticate(tmpl.Email, tmpl.Password)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, acct0.Email, acct1.Email)
	assert.Equal(t, acct0.Password, acct1.Password)
}
