package controllers

import (
	"crypto/rsa"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yamakenji24/shelfer-api/models"
	"github.com/yamakenji24/shelfer-api/service/userservice"
	"golang.org/x/crypto/bcrypt"
)

var (
	key *rsa.PrivateKey
)

func Create(c *gin.Context) {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, struct {
			Status string `json:"status"`
		}{Status: "fail"})
	}
	user.Password = bcryptPassword(user.Password)
	err := userservice.CreateModel(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, struct {
		Status string `json:"status"`
	}{Status: "success"})
}

func Login(c *gin.Context) error {
	loginparams := new(models.LoginParams)

	if err := c.Bind(loginparams); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}

	user, err := userservice.FindByUser(loginparams.Username)

	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return err
	}

	if !comparePassword(user.Password, loginparams.Password) {
		c.JSON(http.StatusUnauthorized, struct {
			Status string `json:"status"`
		}{
			Status: "fail",
		})
		return err
	}

	keyData, err := ioutil.ReadFile("./rsa/private-key.pem")
	if err != nil {
		panic(err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		panic(err)
	}

	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["name"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, struct {
		Status string `json:"status"`
		Token  string `json:"token"`
	}{
		Status: "success",
		Token:  t,
	})
	return nil
}
func comparePassword(hash string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err == nil {
		return true
	}
	return false
}
func bcryptPassword(pass string) string {
	convert, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(convert)
}
