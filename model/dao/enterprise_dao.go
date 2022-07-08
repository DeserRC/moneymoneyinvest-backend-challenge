package dao

import (
	"database/sql"
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/model/entity"
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/model/service"
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/util"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	table    = "CREATE TABLE IF NOT EXISTS enterprise (cnpj VARCHAR(256) PRIMARY KEY, invoicing INT, address VARCHAR(256), name VARCHAR(256), cpf VARCHAR(256), phone VARCHAR(256), email VARCHAR(256));"
	contains = "SELECT 1 FROM enterprise WHERE cnpj = ?;"
	find     = "SELECT cnpj, invoicing, address, name, cpf, phone, email FROM enterprise WHERE cnpj = ?;"
	findAll  = "SELECT cnpj, invoicing, address, name, cpf, phone, email FROM enterprise;"
	save     = "REPLACE INTO enterprise VALUES (?,?,?,?,?,?,?);"
	delete   = "DELETE FROM enterprise WHERE cnpj = ?;"
)

func CreateTable() error {
	_, err := update(table)
	return err
}

func Contains(cnpj string) (bool, error) {
	rows, err := query(contains, cnpj)

	if err != nil {
		return false, err
	}

	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			field := zap.Error(err)
			util.Logger.Error("Error closing rows", field)
		}
	}(rows)

	next := rows.Next()
	return next, nil
}

func Find(cnpj string) (*entity.Enterprise, error) {
	rows, err := query(find, cnpj)

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			field := zap.Error(err)
			util.Logger.Error("Error closing rows", field)
		}
	}(rows)

	if !rows.Next() {
		return nil, errors.New("No data found")
	}

	client := entity.Client{}
	enterprise := entity.Enterprise{Requester: client}

	err = rows.Scan(&enterprise.Cnpj, &enterprise.Invoicing, &enterprise.Address, &client.Name,
		&client.Cpf, &client.Phone, &client.Email)

	if err != nil {
		return nil, err
	}

	return &enterprise, nil
}

func FindAll() (*[]entity.Enterprise, error) {
	rows, err := query(findAll)

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			field := zap.Error(err)
			util.Logger.Error("Error closing rows", field)
		}
	}(rows)

	var enterprises []entity.Enterprise

	for rows.Next() {
		client := entity.Client{}
		enterprise := entity.Enterprise{Requester: client}

		err = rows.Scan(&enterprise.Cnpj, &enterprise.Invoicing, &enterprise.Address, &client.Name,
			&client.Cpf, &client.Phone, &client.Email)

		if err != nil {
			return nil, err
		}

		enterprises = append(enterprises, enterprise)
	}

	return &enterprises, nil
}

func Save(enterprise entity.Enterprise) error {
	_, err := update(save, enterprise.Cnpj, enterprise.Invoicing, enterprise.Address,
		enterprise.Requester.Name, enterprise.Requester.Cpf, enterprise.Requester.Phone,
		enterprise.Requester.Email)

	return err
}

func Delete(cnpj string) error {
	_, err := update(delete, cnpj)
	return err
}

func prepare(query string) (*sql.Stmt, error) {
	return service.Connection.Prepare(query)
}

func update(query string, values ...interface{}) (sql.Result, error) {
	statement, err := prepare(query)

	if err != nil {
		return nil, err
	}

	defer func(statement *sql.Stmt) {
		if err := statement.Close(); err != nil {
			field := zap.Error(err)
			util.Logger.Error("Error closing statement", field)
		}
	}(statement)

	return statement.Exec(values...)
}

func query(query string, values ...interface{}) (*sql.Rows, error) {
	statement, err := prepare(query)

	if err != nil {
		return nil, err
	}

	defer func(statement *sql.Stmt) {
		if err := statement.Close(); err != nil {
			field := zap.Error(err)
			util.Logger.Error("Error closing statement", field)
		}
	}(statement)

	return statement.Query(values...)
}
