package main

import (
	"go-gin-simpe-api/config"
	"go-gin-simpe-api/routers"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {


	db := config.DbConnect()
	defer db.Close()

	r := gin.Default()
	routers.SetupRouter(db, r)
	
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
