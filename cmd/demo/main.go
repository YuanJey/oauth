package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	ssoGroup := r.Group("/sso")
	{
		ssoGroup.GET("/oauth", Oauth)
		ssoGroup.GET("/code", Code)
		ssoGroup.GET("/token", Token)
		ssoGroup.GET("/userInfo", UserInfo)
	}
	address := "0.0.0.0:8000"
	err := r.Run(address)
	if err != nil {
		panic("api start failed " + err.Error())
	}
}
func Test() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	}
}
