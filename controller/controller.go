package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/glgaspar/pay_checker_api/data"
	"github.com/glgaspar/pay_checker_api/models"
	"io"
	"net/http"
)

func CreateBill(c *gin.Context) {
	bodyJson, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.ResultRequest{Status: false, Message: "error parsing body data", Data: err.Error()})
		return
	}
	var body models.Bill
	if err := json.Unmarshal(bodyJson, &body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.ResultRequest{Status: false, Message: "error parsing body data", Data: err.Error()})
		return
	}

	result, err := data.CreateBill(&body)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, models.ResultRequest{Status: false, Message: "error processing data", Data: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, models.ResultRequest{Status: true, Message: "", Data: result})
}

func UpdateBill(c *gin.Context) {
	bodyJson, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.ResultRequest{Status: false, Message: "error parsing body data", Data: err.Error()})
		return
	}
	var body models.Bill
	if err := json.Unmarshal(bodyJson, &body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.ResultRequest{Status: false, Message: "error parsing body data", Data: err.Error()})
		return
	}

	result, err := data.UpdateBill(&body)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, models.ResultRequest{Status: false, Message: "error processing data", Data: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, models.ResultRequest{Status: true, Message: "", Data: result})
}

func PayBill(c *gin.Context) {
	billId, err := strconv.Atoi(c.Param("billId"))
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			models.ResultRequest{
				Status:  false,
				Message: "error identifying bill",
				Data:    err.Error(),
			},
		)
		return
	}

	bills, err := data.GetList(billId)
	if err != nil || len(*bills) == 0 {
		c.IndentedJSON(
			http.StatusBadRequest,
			models.ResultRequest{
				Status:  false,
				Message: "error recovering bill information",
				Data:    err.Error(),
			},
		)
		return
	}

	formFile, _ := c.FormFile("file")
	err = c.SaveUploadedFile(formFile, "./bills/"+(*bills)[0].Path+"/"+strconv.Itoa(time.Now().Year()))
	if err != nil {
		c.IndentedJSON(
			http.StatusUnprocessableEntity,
			models.ResultRequest{
				Status:  false,
				Message: "error saving file",
				Data:    err.Error(),
			},
		)
		return
	}

	err = data.PayBill((*bills)[0].Id)
	if err != nil {
		c.IndentedJSON(
			http.StatusUnprocessableEntity,
			models.ResultRequest{
				Status:  false,
				Message: "error processing data",
				Data:    err.Error(),
			},
		)
		return
	}

	c.IndentedJSON(
		http.StatusOK,
		models.ResultRequest{
			Status:  true,
			Message: "payment successfully saved",
		},
	)
}

func GetList(c *gin.Context) {
	result, err := data.GetList()
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, models.ResultRequest{Status: false, Message: "error processing data", Data: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, models.ResultRequest{Status: true, Message: "", Data: result})
}
