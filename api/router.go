package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	r.GET("/getinfo", GetUserInfoHandler)
	r.Run(":8080")
}
