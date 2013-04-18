// Copyright (C) 2013 Jason McVetta, all rights reserved.

package stormpath

import (
	"github.com/bmizerany/assert"
	"github.com/darkhelmet/env"
	"github.com/jmcvetta/randutil"
	"log"
	"testing"
)

func setupApplication(t *testing.T) *Application {
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)
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
		t.Fatal(err)
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

func TestGetAccount(t *testing.T) {
	app := setupApplication(t)
	tmpl := createAccountTemplate(t)
	acct0, _ := app.CreateAccount(tmpl)
	defer acct0.Delete()
	acct1, err := app.GetAccount(acct0.Href)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, acct0, acct1)
}

func TestUpdateAccount(t *testing.T) {
	app := setupApplication(t)
	tmpl := createAccountTemplate(t)
	acct0, _ := app.CreateAccount(tmpl)
	defer acct0.Delete()
	acct0.GivenName = "Mister"
	acct0.Surname = "Spock"
	err := acct0.Update()
	if err != nil {
		t.Fatal(err)
	}
	acct1, _ := app.GetAccount(acct0.Href)
	assert.Equal(t, "Mister", acct1.GivenName)
	assert.Equal(t, "Spock", acct1.Surname)
}
