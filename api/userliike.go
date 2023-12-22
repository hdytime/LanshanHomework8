package api

import "github.com/gin-gonic/gin"

func UserLikeHandler(c *gin.Context) {
	//1.判断是否已经点赞,引入一个函数islike()
	//2.如果点赞，返回错误信息，如果没有点赞，执行点赞功能，引入like()函数
	//3.点赞成功，返回信息，点赞不成功，返回错误

}
