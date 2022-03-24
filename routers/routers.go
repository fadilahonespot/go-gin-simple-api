package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-gin-simpe-api/controller"
)

type PersonController struct {
	DB *gorm.DB
}

// func SetupRouter() *gin.Engine {
// 	r := gin.Default()

// 	v1 := r.Group("/v1")
// 	{
// 		v1.GET("book", controller.ListBook)
// 	v1.POST("book", Controllers.AddNewBook)
// 	v1.GET("book/:id", Controllers.GetOneBook)
// 	v1.PUT("book/:id", Controllers.PutOneBook)
// 	v1.DELETE("book/:id", Controllers.DeleteBook)
// 	}

// 	return r
// }

func SetupRouter(DB *gorm.DB, r *gin.Engine) {
	personController := controller.PersonController{DB}

	v1 := r.Group("/v1")
	{
		v1.GET("/person", personController.ViewAll)
		v1.POST("/person", personController.CreatePerson)
		v1.PUT("/person/:id", personController.Update)
		v1.GET("/person/:id", personController.ViewById)
		v1.DELETE("/person/:id", personController.Delete)
	}

}