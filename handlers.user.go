package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-contrib/location"
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

func showLoginPage(c *gin.Context) {
	data := gin.H{
		"title": "Login",
	}
	render(c, data, "login.html")
}

func performLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if isUserValid(username, password) {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		url := location.Get(c)

		data := gin.H{
			"title": "Successful Login",
			"url":   url.String(),
		}
		render(c, data, "login-successful.html")

	} else {
		data := gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided",
		}
		c.HTML(http.StatusBadRequest, "login.html", data)
	}
}

func logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
