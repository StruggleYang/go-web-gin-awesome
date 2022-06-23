package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JiraEvents(c *gin.Context) {
	fmt.Println(c.Params)
	fmt.Println(c.Request.PostForm)
	fmt.Println(c.Request.Form)
	fmt.Println(c.Request.Body)
	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}
