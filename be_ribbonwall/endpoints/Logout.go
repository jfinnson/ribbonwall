package endpoints

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func LogoutHandler(c *gin.Context) {
	authConfig := c.Keys["services"].(*Endpoints).Services.Config.Auth
	domain := authConfig.Domain

	var Url *url.URL
	Url, err := url.Parse("https://" + domain)
	if err != nil {
		panic("boom")
	}

	Url.Path += "/v2/logout"
	parameters := url.Values{}
	parameters.Add("returnTo", authConfig.HomeURL)
	parameters.Add("client_id", authConfig.ClientID)
	Url.RawQuery = parameters.Encode()

	// Clear session
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusTemporaryRedirect, Url.String())
}
