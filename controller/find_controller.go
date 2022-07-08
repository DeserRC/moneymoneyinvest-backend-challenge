package controller

import (
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/model/repository"
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func FindEnterprise(context *gin.Context) {
	util.Logger.Debug("Request from Find Enterprise")

	cnpj := context.Params.ByName("cnpj")
	contains, err := repository.Exist(cnpj)

	if err != nil {
		field := zap.Error(err)
		util.Logger.Error("There was an error verifying the enterprise", field)

		cause := err.Error()
		json := gin.H{"error": cause}

		context.JSON(http.StatusInternalServerError, json)
		return
	}

	if !contains {
		json := gin.H{"error": "No enterprise was found with this CNPJ"}
		context.JSON(http.StatusBadRequest, json)

		return
	}

	enterprise, err := repository.Find(cnpj)

	if err != nil {
		field := zap.Error(err)
		util.Logger.Error("There was an error getting the enterprise", field)

		cause := err.Error()
		json := gin.H{"error": cause}

		context.JSON(http.StatusInternalServerError, json)
		return
	}

	context.JSON(http.StatusOK, enterprise)
}
