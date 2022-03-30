package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-gin-simpe-api/controller"
)


func SetupRouter(DB *gorm.DB, r *gin.Engine) {
	personController := controller.PersonController{DB}
	articleController := controller.ArticleController{DB}


	v1 := r.Group("/v1")
	{
		v1.GET("/person", personController.ViewAll)
		v1.POST("/person", personController.CreatePerson)
		v1.PUT("/person/:id", personController.Update)
		v1.GET("/person/:id", personController.ViewById)
		v1.DELETE("/person/:id", personController.Delete)

		v1.GET("/article", articleController.ViewAll)
		v1.POST("/article", articleController.CreateArticle)
		v1.PUT("/article/:id", articleController.Update)
		v1.GET("/article/:id", articleController.ViewById)
		v1.DELETE("/article/:id", personController.Delete)
	}

}