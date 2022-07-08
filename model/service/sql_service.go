package service

import (
	"database/sql"
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/util"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var Connection *sql.DB

func InitConnection(folder string) error {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		if err := os.MkdirAll(folder, 0777); err != nil {
			return err
		}
	}

	path := folder + "/enterprises.db"
	connection, err := sql.Open("sqlite3", path)

	if err != nil {
		return err
	}

	Connection = connection

	util.Logger.Info("Database have been started!")
	return nil
}

func CloseConnection() error {
	err := Connection.Close()

	if err != nil {
		return err
	}

	util.Logger.Info("Database have been closed!")
	return nil
}
