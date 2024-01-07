package main

import (
	"project/controllers"
	"project/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	r := gin.Default()

	r.GET("/get", controllers.GetData)
	r.GET("/get/:franchise_name", controllers.GetDataByFranchiseName)
	r.POST("/create", controllers.CreateData)
	r.PUT("/update/:id", controllers.UpdateData)

	r.Run()
}
