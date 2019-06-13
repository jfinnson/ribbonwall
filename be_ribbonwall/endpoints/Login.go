package endpoints

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"

	logger "github.com/ribbonwall/common/logging"
)

func LoginHandler(c *gin.Context) {
	authConfig := c.Keys["services"].(*Endpoints).Services.Config.Auth
	domain := authConfig.Domain
	aud := authConfig.Audience

	conf := &oauth2.Config{
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
		RedirectURL:  authConfig.RedirectURL,
		Scopes:       []string{"openid"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://" + domain + "/authorize",
			TokenURL: "https://" + domain + "/oauth/token",
		},
	}

	audience := oauth2.SetAuthURLParam("audience", aud)
	url := conf.AuthCodeURL("", audience)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func CallbackHandler(c *gin.Context) {
	authConfig := c.Keys["services"].(*Endpoints).Services.Config.Auth
	domain := authConfig.Domain

	conf := &oauth2.Config{
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
		RedirectURL:  authConfig.RedirectURL,
		Scopes:       []string{"openid"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://" + domain + "/authorize",
			TokenURL: "https://" + domain + "/oauth/token",
		},
	}

	code := c.Query("code")

	token, err := conf.Exchange(context.TODO(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	session.Set("id_token", token.Extra("id_token"))
	session.Set("access_token", token.AccessToken)
	session.Save()

	logger.Infof("access_token: %s", token.AccessToken)

	// Redirect to logged in page
	c.Redirect(http.StatusSeeOther, "/competition_results")

}
