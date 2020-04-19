package main

import (
	"context"
	//"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

const (
	clientID     = "dc3e12fdda79a72348ae"
	clientSecret = "0db4b33e1d29bd4e02a5d95cc2f0ad9ff56df742"
	oauthState   = "random"
)

var oauthConfig = &oauth2.Config{
	ClientID:     clientID,
	ClientSecret: clientSecret,
	RedirectURL:  "http://localhost:8080/GithubCallback",
	Scopes:       []string{"user", "gist"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "http://github.com/login/oauth/authorize",
		TokenURL: "http://github.com/login/oauth/access_token",
	},
}

func main() {
	engine := gin.Default()

	engine.LoadHTMLGlob("templates/**/*")
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
	//通过请求回跳url地址获取state
	c.Request.ParseForm()
	log.Printf("callback parm info======>%#v\n", c.Request.Form)
	state:=c.Request.URL.Query().Get("state")
	if state != oauthState {
		log.Fatalf("invalid oauth state, expected '%s', got '%s'\n", oauthState, state)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	
    //获取code
	code := c.Request.Form.Get("code")
	log.Println("code is ======>", code)
	//通过code获取token
	token, err := oauthConfig.Exchange(context.Background(), code)
	log.Println("token is ========>", token)
	if err != nil {
		log.Fatalf("code exchange failed,err:%v\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"token":token,
		"msg":"获取token成功",
	})
	// //通过token获取客户信息
	// userInfoURL:="https://api.github.com/user"
	// response,err:=http.Get(userInfoURL)
	// if err!=nil{
	// 	log.Fatalf("Get token failed,err:%#v\n",err)
	// 	return
	// }
	// defer response.Body.Close()
	// contents,err:=ioutil.ReadAll(response.Body)
	// c.JSON(http.StatusOK,gin.H{
	// 	"code":200,
	// 	"msg":string(contents),
	// })

}
