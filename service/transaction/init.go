package transactionService

import (
	"github.com/lukaslinardi/xyz_multifinance_api/domain/model/general"
	"github.com/lukaslinardi/xyz_multifinance_api/infra"
	repository "github.com/lukaslinardi/xyz_multifinance_api/repositories"
	"github.com/sirupsen/logrus"
)

type TransactionData struct {
	Transaction TransactionService
}

func NewTransaction(repo repository.Repo, conf general.AppService, dbList *infra.DatabaseList, logger *logrus.Logger) TransactionData {
	return TransactionData{
		Transaction: newTransactionService(repo.Database, conf, dbList, logger),
	}
}
