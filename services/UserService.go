package services

import (
	"SampleGoQmgo/database"
	"SampleGoQmgo/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func Insert() {
	userinfo := models.User{
		Name:     "Tom",
		Password: "123456",
		Age:      18,
		Email:    "tom@gmail.com",
	}

	result, err := database.QmgoConnection.InsertOne(context.TODO(), userinfo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", result)
}

func InsertMulti() {
	userinfo := []models.User{
		{
			Id:       2,
			Name:     "Jack",
			Password: "123456",
			Age:      18,
			Email:    "jack@gmail.com",
		},
		{
			Id:       3,
			Name:     "Amy",
			Password: "123456",
			Age:      18,
			Email:    "amy@gmail.com",
		},
		{
			Id:       4,
			Name:     "Terry",
			Password: "123456",
			Age:      18,
			Email:    "terry@gmail.com",
		},
	}

	result, err := database.QmgoConnection.InsertMany(context.TODO(), userinfo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", result)
}

func FindAll() {
	users := []models.User{}
	// database.QmgoConnection.Find(context.TODO(), bson.M{
	// 	"name": "Tom",
	// }).All(&users)

	database.QmgoConnection.Find(context.TODO(), bson.M{}).All(&users)
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}
}

func Aggregate() {
	matchStage := bson.D{{"$match", []bson.E{{"age", bson.D{{"$gt", 20}}}}}}

	groupStage := bson.D{{"$group", bson.D{{"_id", "$name"},
		{"email", bson.D{{"$push", "$email"}}},
		{"age", bson.D{{"$sum", "$age"}}}}}}

	var showWithInfo []bson.M

	err := database.QmgoConnection.Aggregate(context.TODO(),
		[]bson.D{matchStage, groupStage}).All(&showWithInfo)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", showWithInfo)
	for _, info := range showWithInfo {
		fmt.Printf("%+v\n", info)
	}
}
