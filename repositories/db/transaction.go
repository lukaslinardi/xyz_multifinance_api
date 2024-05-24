package db

import (
	"github.com/lukaslinardi/xyz_multifinance_api/infra"
	"github.com/sirupsen/logrus"
)

type TransactionConfig struct {
	db  *infra.DatabaseList
	log *logrus.Logger
}

func newTransaction(db *infra.DatabaseList, logger *logrus.Logger) TransactionConfig {
	return TransactionConfig{
		db:  db,
		log: logger,
	}
}

type Transaction interface {}
