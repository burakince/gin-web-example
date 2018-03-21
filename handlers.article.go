package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	data := gin.H{
		"title":   "Home Page",
		"payload": articles,
	}
	render(c, data, "index.html")
}

func getArticle(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := getArticleByID(articleID); err == nil {
			data := gin.H{
				"title":   article.Title,
				"payload": article,
			}
			render(c, data, "article.html")
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func showArticleCreationPage(c *gin.Context) {
	data := gin.H{
		"title": "Create New Article",
	}
	render(c, data, "create-article.html")
}

func createArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := createNewArticle(title, content); err == nil {
		data := gin.H{
			"title":   "Submission Successful",
			"payload": a,
		}
		render(c, data, "submission-successful.html")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
