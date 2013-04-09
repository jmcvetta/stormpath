// Copyright (C) 2013 Jason McVetta, all rights reserved.

package stormpath

import (
	"github.com/darkhelmet/env"
	"github.com/jmcvetta/randutil"
	"testing"
)

func setupApplication(t *testing.T) *Application {
	spApp := env.String("STORMPATH_APP")
	apiId := env.String("STORMPATH_API_ID")
	apiSecret := env.String("STORMPATH_API_SECRET")
	s := Application{
		Href:      spApp,
		ApiId:     apiId,
		ApiSecret: apiSecret,
	}
	return &s
}

func createAccountTemplate(t *testing.T) Account {
	rnd, err := randutil.AlphaString(8)
	if err != nil {
		t.Error(err)
	}
	email := "jason.mcvetta+" + rnd + "@gmail.com"
	password := rnd + "Xy123" // Ensure we meet password requirements
	tmpl := Account{
		Username:   rnd,
		Email:      email,
		Password:   password,
		GivenName:  "James",
		MiddleName: "T",
		Surname:    "Kirk",
	}
	return tmpl
}

func TestCreateAccount(t *testing.T) {
	app := setupApplication(t)
	tmpl := createAccountTemplate(t)
	acct, err := app.CreateAccount(tmpl)
	if err != nil {
		t.Error(err)
	}
	//
	// Cleanup
	//
	acct.Delete()
}

func TestDeleteAccount(t *testing.T) {
	app := setupApplication(t)
	tmpl := createAccountTemplate(t)
	acct, _ := app.CreateAccount(tmpl)
	err := acct.Delete()
	if err != nil {
		t.Error(err)
	}
}
