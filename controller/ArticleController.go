package controller

import (
	models "go-gin-simpe-api/Models"
	"go-gin-simpe-api/respon"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ArticleController struct {
	DB *gorm.DB
}



//return all of articles
func (e *ArticleController) ViewAll(c *gin.Context) {
	var article []models.Article
	err := e.DB.Table("articles").Find(&article).Error
	if err != nil {
		respon.HandleError(c, "Ooppss server somting wrong")
		return
	}
	respon.HandleSucces(c, article)
}
//-----------------------

//create an Article
func (e *ArticleController) CreateArticle(c *gin.Context) {
	var article = models.Article{}
	err := c.Bind(&article)
	if err != nil {
		respon.HandleError(c, "failed to insert data")
		return
	}

	err = e.DB.Table("articles").Save(&article).Error
	if err != nil {
		respon.HandleError(c, "Oppsss server somting wrong")
		return
	}
	respon.HandleSucces(c, article)
}
//-----------------------

//update an Article
func (e *ArticleController) Update(c *gin.Context) {
	id := c.Param("id")
	var article = models.Article{}
	err := c.Bind(&article)
	if err != nil {
		respon.HandleError(c, "internal server error")
		return
	}

	checkArticle := models.Article{}
	err = e.DB.Table("articles").Where("id = ?", id).First(&checkArticle).Error
	if err != nil {
		respon.HandleError(c, "id is not exis")
		return
	}

	err = e.DB.Table("articles").Where("id = ?", id).Update(&article).Find(&article).Error
	if err != nil {
		respon.HandleError(c, "failed to update data")
		return
	}
	respon.HandleSucces(c, article)
}
//-----------------------

//return one of Articles by id
func (e *ArticleController) ViewById(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		respon.HandleError(c, "id has be number")
		return
	}
	var article = models.Article{}
	err = e.DB.Table("articles").Where("id = ?", id).First(&article).Error
	if err != nil {
		respon.HandleError(c, "id not exsis")
		return
	}
	respon.HandleSucces(c, article)
}
//-------------------

//delete an Article
func (e *ArticleController) Delete(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		respon.HandleError(c, "id has be number")
		return
	}
	var article = models.Article{}
	err = e.DB.Table("articles").Where("id = ?", id).First(&article).Error
	if err != nil {
		respon.HandleError(c, "id not exsis")
		return
	}

	err = e.DB.Table("articles").Where("id = ?", id).Delete(&article).Error
	if err != nil {
		respon.HandleError(c, "failed delete data")
		return
	}
	respon.HandleSucces(c, "Delete data success")
}
//-------------------------