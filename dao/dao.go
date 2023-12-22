package dao

import (
	"LanshanTeamwork8/model"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

var redisClient *redis.Client
var db *sqlx.DB

func InitDatabase() {

	//初始化redis连接
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	//检查redis连接是否正常
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("redis连接失败:%v", err)
		return
	}
	fmt.Println("redis连接成功")

	//初始化mysql数据库连接
	dsn := "root:123456@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("mysql连接失败:%v", err)
		return
	}
	fmt.Println("mysql连接成功")
}

func GetUserFromCache(userID int) (*model.User, error) {
	//构建redis键名
	key := fmt.Sprintf("user:%d", userID)

	//根据key从redis中读出value
	ctx := context.Background()
	val, err := redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("user not found in cache")
	} else if err != nil {
		return nil, fmt.Errorf("failed to get user from cache:%v", err)
	}

	//将读出的缓存数据反序列化为User结构体
	user := &model.User{}
	err = json.Unmarshal([]byte(val), user)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user data:%v", err)
	}

	return user, nil
}

func GetUserFromDatabase(userID int) (*model.User, error) {
	//执行mysql查询语句获取用户信息
	query := "select id,username,password from user where id=?"
	user := &model.User{}
	err := db.Get(user, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			//用户不存在
			return nil, fmt.Errorf("user not found in database")
		}
		return nil, fmt.Errorf("failed to get user from database:%v", err)
	}
	return user, nil
}

func CacheUser(user *model.User) error {
	//做一个JSON序列化
	data, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user data:%v", err)
	}

	//构建redis键名
	key := fmt.Sprintf("user:%d", user.ID)

	//将用户信息存储到redis储存中，设置过期时间为1小时
	var ctx = context.Background()
	redisClient.Set(ctx, key, string(data), time.Hour)
	return nil
}
