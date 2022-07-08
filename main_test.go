package main

import (
	"bytes"
	"encoding/json"
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/controller"
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/model/dao"
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/model/entity"
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/model/repository"
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/model/service"
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

var requester = entity.Client{"Client", "00000000000", "99 999999999", "client@enterprise.com"}
var enterprise = entity.Enterprise{"00000000000000", 1000000, "Enterprise street", requester}

func Init() *gin.Engine {
	if err := util.InitLogger("environment/logs"); err != nil {
		panic(err)
		return nil
	}

	util.Logger.Debug("Initializing the database...")

	if err := service.InitConnection("environment/database"); err != nil {
		field := zap.Error(err)
		util.Logger.Panic("There was an error starting the database", field)

		return nil
	}

	util.Logger.Debug("Creating the table...")

	if err := dao.CreateTable(); err != nil {
		field := zap.Error(err)
		util.Logger.Panic("There was an error creating the table", field)

		return nil
	}

	gin.SetMode(gin.TestMode)
	return gin.Default()
}

func Shutdown() {
	if err := service.CloseConnection(); err != nil {
		field := zap.Error(err)
		util.Logger.Panic("There was an error closing the database", field)
	}
}

func CreateEnterprise() {
	err := repository.Save(enterprise)

	if err != nil {
		field := zap.Error(err)
		util.Logger.Error("There was an error creating the enterprise", field)

		return
	}
}

func DeleteEnterprise() {
	exist, err := repository.Exist(enterprise.Cnpj)

	if err != nil {
		field := zap.Error(err)
		util.Logger.Error("There was an error verifying the enterprise", field)

		return
	}

	if !exist {
		return
	}

	if err := repository.Delete(enterprise.Cnpj); err != nil {
		field := zap.Error(err)
		util.Logger.Error("There was an error deleting the enterprise", field)

		return
	}
}

func TestShow(test *testing.T) {
	server := Init()
	defer Shutdown()

	server.GET("/enterprises", controller.ShowEnterprise)
	request, _ := http.NewRequest("GET", "/enterprises", nil)

	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	assert.Equal(test, http.StatusOK, response.Code)
}

func TestFind(test *testing.T) {
	server := Init()
	defer Shutdown()

	CreateEnterprise()
	defer DeleteEnterprise()

	server.GET("/enterprises/:cnpj", controller.FindEnterprise)
	request, _ := http.NewRequest("GET", "/enterprises/00000000000000", nil)

	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	assert.Equal(test, http.StatusOK, response.Code)
}

func TestCreate(test *testing.T) {
	server := Init()
	defer Shutdown()

	server.POST("/enterprises", controller.CreateEnterprise)
	marshal, err := json.Marshal(enterprise)

	if err != nil {
		field := zap.Error(err)
		util.Logger.Error("There was an error serialization the enterprise", field)
		return
	}

	buffer := bytes.NewBuffer(marshal)

	request, _ := http.NewRequest("POST", "/enterprises", buffer)
	defer DeleteEnterprise()

	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	assert.Equal(test, http.StatusOK, response.Code)
}

func TestDelete(test *testing.T) {
	server := Init()
	defer Shutdown()

	CreateEnterprise()
	defer DeleteEnterprise()

	server.DELETE("/enterprises/:cnpj", controller.DeleteEnterprise)
	request, _ := http.NewRequest("DELETE", "/enterprises/00000000000000", nil)

	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	assert.Equal(test, http.StatusOK, response.Code)
}

func TestEdit(test *testing.T) {
	server := Init()
	defer Shutdown()

	CreateEnterprise()
	defer DeleteEnterprise()

	server.PATCH("/enterprises/:cnpj", controller.EditEnterprise)
	marshal, err := json.Marshal(enterprise)

	if err != nil {
		field := zap.Error(err)
		util.Logger.Error("There was an error serialization the enterprise", field)
		return
	}

	buffer := bytes.NewBuffer(marshal)

	request, _ := http.NewRequest("PATCH", "/enterprises/00000000000000", buffer)
	defer DeleteEnterprise()

	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	assert.Equal(test, http.StatusOK, response.Code)
}
