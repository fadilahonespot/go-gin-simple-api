package main

import (
	"go-gin-simpe-api/config"
	"go-gin-simpe-api/routers"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	port := "7861"

	db := config.DbConnect()
	defer db.Close()

	// r := router.SetupRouter()

	r := gin.Default()
	routers.SetupRouter(db, r)
	
	err := r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
