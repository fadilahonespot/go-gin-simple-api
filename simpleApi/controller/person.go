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
	ID        int  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Respon struct {
	gorm.Model
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    interface{}
}

func (e *Person) TableName() string {
	return "person"
}

func CreatePersonController(DB *gorm.DB, r *gin.Engine) {
	personController := PersonController{DB}

	r.GET("/person", personController.viewAll)
	r.POST("/person", personController.createPerson)
	r.PUT("/person", personController.update)
	r.GET("/person/:id", personController.viewById)
	r.DELETE("/person/:id", personController.delete)

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

func (e *PersonController) viewAll(c *gin.Context) {
	var person []Person
	err := e.DB.Table("person").Find(&person).Error
	if err != nil {
		e.handleError(c, "Ooppss server somting wrong")
		return
	}
	e.handleSucces(c, person)
}

func (e *PersonController) createPerson(c *gin.Context) {
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

func (e *PersonController) update(c *gin.Context) {
	var person = Person{}
	err := c.Bind(&person)
	if err != nil {
		e.handleError(c, "internal server error")
		return
	}

	id := person.ID
	checkPerson := Person{}
	err = e.DB.Table("person").Where("id = ?", id).First(&checkPerson).Error
	if err != nil {
		e.handleError(c, "id is not exis")
		return
	}

	err = e.DB.Table("person").Where("id = ?", id).Update(&person).Error
	if err != nil {
		e.handleError(c, "failed to update data")
		return
	}
	e.handleSucces(c, person)
}

func (e *PersonController) viewById(c *gin.Context) {
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

func (e *PersonController) delete(c *gin.Context) {
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
