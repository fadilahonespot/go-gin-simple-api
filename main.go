package main

import (
	"go-gin-simpe-api/config"
	"go-gin-simpe-api/controller"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	port := "7861"

	db := config.DbConnect()
	defer db.Close()

	router := gin.Default()
	controller.CreatePersonController(db, router)
	
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
