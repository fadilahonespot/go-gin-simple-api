package routers

import (
	"go-gin-simpe-api/controller"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type PersonController struct {
	DB *gorm.DB
}

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