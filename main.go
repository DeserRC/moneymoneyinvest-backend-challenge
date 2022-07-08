package main

import (
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/model/dao"
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/model/service"
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/router"
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/util"
	"go.uber.org/zap"
)

func main() {
	if err := util.InitLogger("environment/logs"); err != nil {
		panic(err)
		return
	}

	util.Logger.Debug("Initializing the database...")

	if err := service.InitConnection("environment/database"); err != nil {
		field := zap.Error(err)
		util.Logger.Panic("There was an error starting the database", field)

		return
	}

	defer func() {
		if err := service.CloseConnection(); err != nil {
			field := zap.Error(err)
			util.Logger.Panic("There was an error closing the database", field)
		}
	}()

	util.Logger.Debug("Creating the table...")

	if err := dao.CreateTable(); err != nil {
		field := zap.Error(err)
		util.Logger.Panic("There was an error creating the table", field)

		return
	}

	util.Logger.Debug("Initializing the routes...")

	if err := router.HandleRequest(); err != nil {
		field := zap.Error(err)
		util.Logger.Panic("There was an error starting the server routes", field)

		return
	}
}
