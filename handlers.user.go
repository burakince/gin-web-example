package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func generateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

func showRegistrationPage(c *gin.Context) {
	data := gin.H{
		"title": "Register",
	}
	render(c, data, "register.html")
}

func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if _, err := registerNewUser(username, password); err == nil {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		data := gin.H{
			"title": "Successful registration & Login",
		}
		render(c, data, "login-successful.html")
	} else {
		data := gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error(),
		}
		c.HTML(http.StatusBadRequest, "register.html", data)
	}
}
