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

	result, err := data.PayBill(&body)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, models.ResultRequest{Status: false, Message: "error processing data", Data: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, models.ResultRequest{Status: true, Message: "", Data: result})
}

func GetList(c *gin.Context) {
	result, err := data.GetList()
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, models.ResultRequest{Status: false, Message: "error processing data", Data: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, models.ResultRequest{Status: true, Message: "", Data: result})
}
