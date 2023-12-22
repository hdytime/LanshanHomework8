package api

import (
	"LanshanTeamwork8/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetUserInfoHandler(c *gin.Context) {
	//1.先去缓存中读取数据//引入getUserFromCache函数
	//2.若缓存中有，返回数据，若未命中，去数据库读取//引入getUserFromDatabase函数
	//由于两个函数返回的是model.User结构体的指针，因此做一个JSON序列化的操作，
	//3.如果数据库中有，返回数据,写入缓存,//引入cacheUser函数，若未命中，返回错误信息

	//先通过请求查询id,返回id的值
	userID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// 先从 Redis 缓存中尝试获取用户信息
	user, err := dao.GetUserFromCache(userID)
	if err == nil {
		c.JSON(http.StatusOK, user)
		return
	}

	//如果缓存中不存在，从数据库获取用户信息
	user, err = dao.GetUserFromDatabase(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	//如果数据库有数据，写入到缓存中
	err = dao.CacheUser(user)
	if err != nil {
		log.Printf("Failed to cache user:%v\n", err)
	}
	c.JSON(http.StatusOK, user)
}
