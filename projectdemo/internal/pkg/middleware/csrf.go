package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var jwtSecret []byte

const (
	//XCsrfTokenVarifyFail ...
	XCsrfTokenVarifyFail = "csrf认证失败,请重新授权获取"
	//XCsrfTokenExpried ...
	XCsrfTokenExpried   = "当前token已过期，请重新请求"
	//csrfExpiredTime 失效时间
	csrfExpiredTime =  24 * time.Hour
)

//CsrfClaims csrf
type CsrfClaims struct {
	IP string `json:"ip"`
	jwt.StandardClaims
}

// GenerateCsrfToken 生产用于csrf的token
func GenerateCsrfToken(ip string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(csrfExpiredTime)

	claims := CsrfClaims{
		ip,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "token-csrf",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

//ParseCsrfToken token解析
func ParseCsrfToken(token string) (*CsrfClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &CsrfClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*CsrfClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// CsrfMiddleware 中间件
func CsrfMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
        //默认状态码
		code := http.StatusOK
        //从请求头获取token
		token := c.GetHeader("X-XSRF-TOKEN")
        //默认返回错误提示信息
		message := XCsrfTokenVarifyFail
		//如果没有传入token
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status_code": code,
				"message": message,
			})
			c.Abort()
			return
		}
		//开始解析token
		claims, err := ParseCsrfToken(token)
		if err != nil {
			code = http.StatusUnauthorized
			//token过期
			if err.(*jwt.ValidationError).Errors  == jwt.ValidationErrorExpired {
				message = XCsrfTokenExpried
			}
		}
		//状态码不是正常有效
		if code != http.StatusOK{
			c.JSON(http.StatusUnauthorized, gin.H{
				"status_code": code,
				"message": message,
			})
			c.Abort()
			return
		}
		//设置信息到上下文
		c.Set("client_ip",claims.IP)
		c.Next()
	}
}