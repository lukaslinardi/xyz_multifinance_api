package transactionHandler

import (

	// cg "olin_psef/domain/constants/general"
	// mc "olin_psef/domain/model/auth"
	// "olin_psef/domain/model/general"
	// "olin_psef/domain/utils"

	transactionService "github.com/lukaslinardi/xyz_multifinance_api/service/transaction"

	"github.com/lukaslinardi/xyz_multifinance_api/domain/model/general"
	"github.com/sirupsen/logrus"
)

type TransactionHandler struct {
	Transaction transactionService.Transaction
	conf        general.AppService
	log         *logrus.Logger
}

func NewTransactionHandler(transaction transactionService.Transaction, conf general.AppService, logger *logrus.Logger) TransactionHandler {
	return TransactionHandler{
		Transaction: transaction,
		conf:        conf,
		log:         logger,
	}
}
