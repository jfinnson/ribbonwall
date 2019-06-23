package auth

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/auth0-community/go-auth0"
	"github.com/gin-gonic/gin"
	"gopkg.in/square/go-jose.v2"
	"io/ioutil"
	"net/http"

	logger "github.com/jfinnson/ribbonwall/common/logging"
)

// Wrapping a Gin endpoint with Auth0 Groups.
var validator *auth0.JWTValidator
var config Config

func InitAuth(configInit Config) {
	// Set global configs
	config = configInit

	//Creates a configuration with the Auth0 information
	data, err := ioutil.ReadFile("./domains/be_ribbonwall/config/credentials/ribbonwall.pem")
	if err != nil {
		panic("Impossible to read key form disk")
	}

	secret, err := loadPublicKey(data)
	if err != nil {
		panic("Invalid provided key")
	}
	secretProvider := auth0.NewKeyProvider(secret)
	//Aud needs to be Client ID because frontend needs to use id_token auth. Which uses ClientID as the AUD.
	configuration := auth0.NewConfiguration(secretProvider, []string{config.ClientID}, config.Issuer, jose.RS256)
	validator = auth0.NewValidator(configuration, nil)
}

func Auth0Groups(wantedGroups ...string) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {

		tok, err := validator.ValidateRequest(c.Request)
		if err != nil {
			//c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			logger.Errorf("Invalid token:", err)
			c.Redirect(http.StatusTemporaryRedirect, config.LoginURL) // Redirect to login
			c.Abort()
			return
		}

		claims := map[string]interface{}{}
		err = validator.Claims(c.Request, tok, &claims)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
			c.Abort()
			logger.Errorf("Invalid claims:", err)
			return
		}

		metadata, okMetadata := claims[fmt.Sprintf("https://%s/app_metadata", config.Audience)].(map[string]interface{})
		authorization, okAuthorization := metadata["authorization"].(map[string]interface{})
		groups, hasGroups := authorization["groups"].([]interface{})
		if !okMetadata || !okAuthorization || !hasGroups || !shouldAccess(wantedGroups, groups) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "need more privileges"})
			c.Abort()
			logger.Errorf("Need more privileges")
			return
		}
		c.Next()
	})
}

func shouldAccess(wantedGroups []string, groups []interface{}) bool {

	if len(groups) < 1 {
		return true
	}

	for _, wantedScope := range wantedGroups {

		scopeFound := false

		for _, iScope := range groups {
			scope, ok := iScope.(string)

			if !ok {
				continue
			}
			if scope == wantedScope {
				scopeFound = true
				break
			}
		}
		if !scopeFound {
			return false
		}
	}
	return true
}

// LoadPublicKey loads a public key from PEM/DER-encoded data.
func loadPublicKey(data []byte) (interface{}, error) {
	input := data

	block, _ := pem.Decode(data)
	if block != nil {
		input = block.Bytes
	}

	// Try to load SubjectPublicKeyInfo
	pub, err0 := x509.ParsePKIXPublicKey(input)
	if err0 == nil {
		return pub, nil
	}

	cert, err1 := x509.ParseCertificate(input)
	if err1 == nil {
		return cert.PublicKey, nil
	}

	return nil, fmt.Errorf("square/go-jose: parse error, got '%s' and '%s'", err0, err1)
}
