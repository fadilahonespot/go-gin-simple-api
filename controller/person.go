package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	models "go-gin-simpe-api/Models"
	"go-gin-simpe-api/respon"
	"strconv"
)
type PersonController struct {
	DB *gorm.DB
}

func (e *PersonController) ViewAll(c *gin.Context) {
	var person []models.Person
	err := e.DB.Table("persons").Find(&person).Error
	if err != nil {
		respon.HandleError(c, "Ooppss server somting wrong")
		return
	}
	respon.HandleSucces(c, person)
}

func (e *PersonController) CreatePerson(c *gin.Context) {
	var person = models.Person{}
	err := c.Bind(&person)
	if err != nil {
		respon.HandleError(c, "failed to insert data")
		return
	}

	err = e.DB.Table("persons").Save(&person).Error
	if err != nil {
		respon.HandleError(c, "Oppsss server somting wrong")
		return
	}
	respon.HandleSucces(c, person)
}

func (e *PersonController) Update(c *gin.Context) {
	id := c.Param("id")
	var person = models.Person{}
	err := c.Bind(&person)
	if err != nil {
		respon.HandleError(c, "internal server error")
		return
	}

	checkPerson := models.Person{}
	err = e.DB.Table("persons").Where("id = ?", id).First(&checkPerson).Error
	if err != nil {
		respon.HandleError(c, "id is not exis")
		return
	}

	err = e.DB.Table("persons").Where("id = ?", id).Update(&person).Find(&person).Error
	if err != nil {
		respon.HandleError(c, "failed to update data")
		return
	}
	respon.HandleSucces(c, person)
}

func (e *PersonController) ViewById(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		respon.HandleError(c, "id has be number")
		return
	}
	var person = models.Person{}
	err = e.DB.Table("persons").Where("id = ?", id).First(&person).Error
	if err != nil {
		respon.HandleError(c, "id not exsis")
		return
	}
	respon.HandleSucces(c, person)
}

func (e *PersonController) Delete(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		respon.HandleError(c, "id has be number")
		return
	}
	var person = models.Person{}
	err = e.DB.Table("persons").Where("id = ?", id).First(&person).Error
	if err != nil {
		respon.HandleError(c, "id not exsis")
		return
	}

	err = e.DB.Table("persons").Where("id = ?", id).Delete(&person).Error
	if err != nil {
		respon.HandleError(c, "failed delete data")
		return
	}
	respon.HandleSucces(c, "Delete data success")
}
