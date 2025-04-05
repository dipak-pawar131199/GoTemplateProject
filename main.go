package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gotemplateproject/internal/db"
	"gotemplateproject/internal/handler"

)

func main() {
	fmt.Println("Golang template project")
	db.InitDB()
	r := gin.Default()

	r.GET("/users", handler.GetUsers)
	r.Run(":8080")

	defer db.DB.Close()
}
