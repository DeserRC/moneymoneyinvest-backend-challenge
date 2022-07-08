package controller

import (
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/model/repository"
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowEnterprise(context *gin.Context) {
	util.Logger.Debug("Request from Show Enterprises")
	enterprises, err := repository.FindAll()

	if err != nil {
		cause := err.Error()
		json := gin.H{"error": cause}

		context.JSON(http.StatusInternalServerError, json)
		return
	}

	context.JSON(http.StatusOK, enterprises)
}
