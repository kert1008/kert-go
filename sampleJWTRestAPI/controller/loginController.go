package controller

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"sampleJWTRestAPI/models"
	"sampleJWTRestAPI/services"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

var identityKey = "id"

func newJWTAuth() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.Login); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.Login{
				ID: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals models.Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if services.GetLoginUserByIDPassword(loginVals.ID, loginVals.Password) != nil {
				return &models.Login{
					ID: loginVals.ID,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},

		Authorizator: func(data interface{}, c *gin.Context) bool {
			// if v, ok := data.(*models.Login); ok && v.ID == "admin" {
			// 	return true
			// }
			// return false
			return true
		},

		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return authMiddleware
}

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := newJWTAuth()
		auth.LoginHandler(c)
	}
}

func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		tmpBody := string(buf[0:n])

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(tmpBody)))
		login := models.Login{}
		c.BindJSON(&login)

		if err := services.CreateLoginUser(&login); err != nil {
			c.JSON(http.StatusBadRequest, "error")
		} else {
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(tmpBody)))
			auth := newJWTAuth()
			auth.LoginHandler(c)
		}
	}
}

func AuthUser() gin.HandlerFunc {
	auth := newJWTAuth()
	return auth.MiddlewareFunc()
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.Get(identityKey)
		if user := services.GetUserByID(data.(*models.Login).ID); user != nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusBadRequest, "error")
		}
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.Get(identityKey)

		user := models.User{}
		c.BindJSON(&user)

		user.ID = data.(*models.Login).ID
		if err := services.UpdateUser(&user); err != nil {
			c.JSON(http.StatusBadRequest, "error")
		} else {
			c.JSON(http.StatusOK, "ok")
		}
	}
}

func RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := newJWTAuth()
		auth.RefreshHandler(c)
	}
}
