package main

import (
	"github.com/YuanJey/oauth/demo"
	"github.com/YuanJey/oauth/oauth"
	"github.com/gin-gonic/gin"
	"net/http"
)

var auths = make(map[string]oauth.AuthProvider)

const (
	qw  = "qw"
	idm = "idm"
)

func init() {
	auths[qw] = &demo.QWProvider{
		ClientID:     "",
		ClientSecret: "",
		SSOURL:       "",
	}
	auths[idm] = &demo.IDMProvider{}
}

func Oauth(c *gin.Context) {
	redirectURI := c.Query("redirect_uri")
	state := c.Query("state")
	c.SetCookie("redirect_uri", redirectURI, 3600, "/", "", false, true)
	c.SetCookie("state", state, 3600, "/", "", false, true)
	authorizationURL := auths[qw].GetAuthorizationURL(redirectURI, state)
	c.Redirect(http.StatusFound, authorizationURL)
}
func Code(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")
	redirectURI, err := c.Cookie("redirect_uri")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "redirect_uri cookie not found"})
		return
	}

	redirectURIModule := redirectURI + "&code=" + code + "&state=" + state
	c.Redirect(http.StatusFound, redirectURIModule)
}
func Token(c *gin.Context) {
	accessToken, err := auths[qw].ExchangeCodeForToken(c.Query("code"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to exchange code for token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
func UserInfo(c *gin.Context) {
	accessToken := c.Query("access_token")
	userInfo, err := auths[qw].GetUserInfo(accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user info"})
		return
	}
	c.JSON(http.StatusOK, userInfo)
}
