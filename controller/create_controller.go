package controller

import (
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/model/entity"
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/model/repository"
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func CreateEnterprise(context *gin.Context) {
	util.Logger.Debug("Request from Create Enterprise")
	var enterprise entity.Enterprise

	if err := context.ShouldBindJSON(&enterprise); err != nil {
		cause := err.Error()
		json := gin.H{"error": cause}

		context.JSON(http.StatusBadRequest, json)
		return
	}

	if err := entity.ValidateEnterprise(&enterprise); err != nil {
		cause := err.Error()
		json := gin.H{"error": cause}

		context.JSON(http.StatusBadRequest, json)
		return
	}

	err := repository.Save(enterprise)

	if err != nil {
		field := zap.Error(err)
		util.Logger.Error("There was an error creating the enterprise", field)

		cause := err.Error()
		json := gin.H{"error": cause}

		context.JSON(http.StatusInternalServerError, json)
		return
	}

	context.JSON(http.StatusOK, enterprise)
}
