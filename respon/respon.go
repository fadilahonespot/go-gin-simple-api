package respon

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type Respon struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    interface{}
}


//-----------------------
func HandleSucces(c *gin.Context, data interface{}) {
	var returnData = Respon{
		Status:  "0000",
		Message: "Success",
		Data:    data,
	}
	c.JSON(http.StatusOK, returnData)
}
//-----------------------

//-----------------------
func HandleError(c *gin.Context, message string) {
	var returnData = Respon{
		Status:  "501",
		Message: message,
	}
	c.JSON(http.StatusBadRequest, returnData)
}
//-----------------------
