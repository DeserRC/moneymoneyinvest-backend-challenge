package router

import (
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/controller"
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/util"
	"github.com/gin-gonic/gin"
)

func HandleRequest() error {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	server.GET("/enterprises/", controller.ShowEnterprise)
	server.GET("/enterprises/:cnpj", controller.FindEnterprise)

	server.POST("/enterprises", controller.CreateEnterprise)
	server.PATCH("/enterprises/:cnpj", controller.EditEnterprise)

	server.DELETE("/enterprises/:cnpj", controller.DeleteEnterprise)
	err := server.Run(":8080")

	if err != nil {
		return err
	}

	util.Logger.Info("Server routes have been started!")
	return nil
}
