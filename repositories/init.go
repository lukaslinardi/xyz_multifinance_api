package repository

import (
	"github.com/lukaslinardi/xyz_multifinance_api/domain/model/general"
	"github.com/lukaslinardi/xyz_multifinance_api/infra"

	database "github.com/lukaslinardi/xyz_multifinance_api/repositories/db"
	"github.com/sirupsen/logrus"
)

type Repo struct {
	Database database.Database
}

func NewRepo(db *infra.DatabaseList, conf general.AppService, logger *logrus.Logger) Repo {
	return Repo{
		Database: database.NewDatabase(db, logger),
	}
}
