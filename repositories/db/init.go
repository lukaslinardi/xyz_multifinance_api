package db

import (
	"github.com/lukaslinardi/xyz_multifinance_api/infra"
	"github.com/sirupsen/logrus"
)

type Database struct {
	User User
	Auth Auth
}

func NewDatabase(db *infra.DatabaseList, logger *logrus.Logger) Database {
	return Database{
		User: newUser(db, logger),
		Auth: newAuth(db, logger),
	}
}
