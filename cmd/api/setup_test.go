package main

import (
	"authentification/cmd/api/controllers"
	"os"
	"testing"
)

var testApp App

func TestMain(m *testing.M) {
	testApp.Authentication = controllers.NewTestAuthentication()
	testApp.DB = nil
	os.Exit(m.Run())
}
