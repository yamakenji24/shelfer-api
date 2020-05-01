package auth

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func CheckJWTHandler(c *gin.Context) {
	token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			_, err := token.Method.(*jwt.SigningMethodRSA)
			if !err {
				return nil, fmt.Errorf("unexpected sigining method")
			} else {
				return LookPublicKey()
			}
		},
	)
	if err != nil || !token.Valid {
		UnAuthorized(c)
		return
	}
}

func LookPublicKey() (*rsa.PublicKey, error) {
	key, _ := ioutil.ReadFile("./rsa/public-key.pem")
	parsedKey, err := jwt.ParseRSAPublicKeyFromPEM(key)
	return parsedKey, err
}

func UnAuthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": http.StatusText(http.StatusUnauthorized),
	})
	c.Abort()
}
