package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

const (
	clientID     = "dc3e12fdda79a72348ae"
	clientSecret = "0db4b33e1d29bd4e02a5d95cc2f0ad9ff56df742"
	oauthState   = "normal"
)

var oauthConfig = &oauth2.Config{
	ClientID:     clientID,
	ClientSecret: clientSecret,
	RedirectURL:  "http://localhost:8080/GithubCallback",
	Scopes:       []string{"", ""},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "http://localhost:8888/oauth2/auth",
		TokenURL: "http://localhost:8888/oauth2/token",
	},
}

func main() {
	engine := gin.Default()

	engine.LoadHTMLGlob("templates/**/**")
	engine.GET("/", handleMain)
	engine.GET("/login", handleLogin)
	engine.GET("/GithubCallback", handleGithubCallback)
	engine.Run()
}

func handleMain(c *gin.Context) {
	c.HTML(200, "posts/index.html", nil)
}

func handleLogin(c *gin.Context) {
	//AuthCodeURL将网址返回到OAuth 2.0提供程序的同意页面，该页面会明确要求所需范围的权限。
	url := oauthConfig.AuthCodeURL(oauthState)
	log.Println("请求授权界面的地址url：", url)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func handleGithubCallback(c *gin.Context) {
	state := c.Request.ParseForm.GET("state")
	if state != oauthState {
		log.Fatalf("invalid oauth state, expected '%s', got '%s'\n", oauthState, state)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	log.Printf("callback parm info======>%#v\n", c.Request.Form)

	code := c.Request.ParseForm.GET("code")
	log.Println("code is ======>", code)
	token, err := oauthConfig.Exchange(context.Background(), code)
	log.Println("token is ========>", token)
	if err != nil {
		log.Fatal("code exchange failed,err:%v\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

}
