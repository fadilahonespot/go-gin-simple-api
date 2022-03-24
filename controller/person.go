package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type PersonController struct {
	DB *gorm.DB
}

type Person struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Respon struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    interface{}
}

func (e *Person) TableName() string {
	return "person"
}


func (e *PersonController) handleSucces(c *gin.Context, data interface{}) {
	var returnData = Respon{
		Status:  "0000",
		Message: "Success",
		Data:    data,
	}
	c.JSON(http.StatusOK, returnData)
}

func (e *PersonController) handleError(c *gin.Context, message string) {
	var returnData = Respon{
		Status:  "501",
		Message: message,
	}
	c.JSON(http.StatusBadRequest, returnData)
}

func (e *PersonController) ViewAll(c *gin.Context) {
	var person []Person
	err := e.DB.Table("person").Find(&person).Error
	if err != nil {
		e.handleError(c, "Ooppss server somting wrong")
		return
	}
	e.handleSucces(c, person)
}

func (e *PersonController) CreatePerson(c *gin.Context) {
	var person = Person{}
	err := c.Bind(&person)
	if err != nil {
		e.handleError(c, "failed to insert data")
		return
	}

	err = e.DB.Table("person").Save(&person).Error
	if err != nil {
		e.handleError(c, "Oppsss server somting wrong")
		return
	}
	e.handleSucces(c, person)
}

func (e *PersonController) Update(c *gin.Context) {
	id := c.Param("id")
	var person = Person{}
	err := c.Bind(&person)
	if err != nil {
		e.handleError(c, "internal server error")
		return
	}

	checkPerson := Person{}
	err = e.DB.Table("person").Where("id = ?", id).First(&checkPerson).Error
	if err != nil {
		e.handleError(c, "id is not exis")
		return
	}

	err = e.DB.Table("person").Where("id = ?", id).Update(&person).Find(&person).Error
	if err != nil {
		e.handleError(c, "failed to update data")
		return
	}
	e.handleSucces(c, person)
}

func (e *PersonController) ViewById(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		e.handleError(c, "id has be number")
		return
	}
	var person = Person{}
	err = e.DB.Table("person").Where("id = ?", id).First(&person).Error
	if err != nil {
		e.handleError(c, "id not exsis")
		return
	}
	e.handleSucces(c, person)
}

func (e *PersonController) Delete(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		e.handleError(c, "id has be number")
		return
	}
	var person = Person{}
	err = e.DB.Table("person").Where("id = ?", id).First(&person).Error
	if err != nil {
		e.handleError(c, "id not exsis")
		return
	}

	err = e.DB.Table("person").Where("id = ?", id).Delete(&person).Error
	if err != nil {
		e.handleError(c, "failed delete data")
		return
	}
	e.handleSucces(c, "Delete data success")
}