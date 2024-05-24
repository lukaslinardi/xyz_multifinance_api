package transactionService

import (
	"github.com/lukaslinardi/xyz_multifinance_api/domain/model/general"
	"github.com/lukaslinardi/xyz_multifinance_api/infra"
	"github.com/lukaslinardi/xyz_multifinance_api/repositories/db"
	"github.com/sirupsen/logrus"
)

type TransactionService struct {
	db     db.Database
	conf   general.AppService
	dbConn *infra.DatabaseList
	log    *logrus.Logger
}

func newTransactionService(db db.Database, conf general.AppService, dbConn *infra.DatabaseList, logger *logrus.Logger) TransactionService {
	return TransactionService{
		db:     db,
		conf:   conf,
		dbConn: dbConn,
		log:    logger,
	}
}

type Transaction interface {}
