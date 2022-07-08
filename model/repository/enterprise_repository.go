package repository

import (
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/model/dao"
	"github.com/DeserRC/moneymoneyinvest-backend-challenge/model/entity"
)

func Exist(cnpj string) (bool, error) {
	return dao.Contains(cnpj)
}

func Find(cnpj string) (*entity.Enterprise, error) {
	return dao.Find(cnpj)
}

func FindAll() (*[]entity.Enterprise, error) {
	return dao.FindAll()
}

func Save(enterprise entity.Enterprise) error {
	return dao.Save(enterprise)
}

func Delete(cnpj string) error {
	return dao.Delete(cnpj)
}
