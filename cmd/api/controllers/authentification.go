package controllers

type Authentication interface{}

type authenticationController struct{}

func NewAuthenticationController() Authentication {
	return &authenticationController{}
}
