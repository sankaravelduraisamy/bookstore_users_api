package controllers

import{
	"github.com/gin-ganic/gin"
	"net/http"
}

func Ping(c *gin.Context){
	c.String(http:StatusOK,format:"pong")
}