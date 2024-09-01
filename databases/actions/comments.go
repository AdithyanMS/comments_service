package actions

import (
	"comments_service/databases"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func insertOne(dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {

	db := databases.GetDB()
	collection := db.Client.Database(dataBase).Collection(col)

	result, err := collection.InsertOne(db.Ctx, doc)
	return result, err
}

func InsertComment(comment interface{}) (*mongo.InsertOneResult, error) {
	insertResult, err := insertOne("comments_service", "comments", comment)
	if err != nil {
		fmt.Errorf("comment not inserted: %s", err.Error())
	}
	return insertResult, err
}
