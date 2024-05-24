package transactionHandler

import (
	"github.com/lukaslinardi/xyz_multifinance_api/domain/model/general"
	"github.com/lukaslinardi/xyz_multifinance_api/service"
	"github.com/sirupsen/logrus"
)

type TransactionDataHandler struct {
	Transaction TransactionHandler
}

func NewTransactionDataHandler(sv service.Service, conf general.AppService, logger *logrus.Logger) TransactionDataHandler {
	return TransactionDataHandler{
		Transaction: NewTransactionHandler(sv.Auth.Auth, conf, logger),
	}
}
