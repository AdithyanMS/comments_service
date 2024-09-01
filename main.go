package main

import (
	"comments_service/controllers"
	"comments_service/databases"

	"github.com/gin-gonic/gin"
)

func main() {
	databases.Connect()
	db := databases.GetDB()
	if db.Err != nil {
		panic(db.Err)
	}
	databases.Ping(db.Client, db.Ctx)

	router := gin.Default()
	router.GET("/get", controllers.GetComments)
	router.POST("/post", controllers.PostComment)
	router.Run()

	defer databases.Close(db.Client, db.Ctx, db.Cancel)
}
