package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main01() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("mongodb连接成功")
	collection := client.Database("test").Collection("people")

	//插入单个文档
	user := bson.D{{"name", "zhangsan"}, {"age", 30}}
	res, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.InsertedID)

	//插入多个文档
	users := []interface{}{
		bson.D{{"name", "lisi"}, {"age", 25}},
		bson.D{{"name", "wangwu"}, {"age", 20}},
		bson.D{{"name", "zhaoliu"}, {"age", 28}},
	}
	result, err := collection.InsertMany(context.Background(), users)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result.InsertedIDs)

	//建立查询 (无查询条件)
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println(err)
		return
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}

	// 设置查询条件
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{
					{"age", bson.D{{"$gt", 25}}},
				},
			},
		},
	}

	// 执行查询
	cursor, err = collection.Find(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		return
	}

	// 遍历查询结果
	var filteredResults []bson.M
	if err = cursor.All(context.TODO(), &filteredResults); err != nil {
		panic(err)
	}

	for _, result := range filteredResults {
		fmt.Println(result)
	}
}
