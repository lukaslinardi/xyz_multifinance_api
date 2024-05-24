package repository

import (
	"github.com/lukaslinardi/fullstack_engineer_sprint_asia/domain/model/general"
	"github.com/lukaslinardi/fullstack_engineer_sprint_asia/infra"

	database "github.com/lukaslinardi/fullstack_engineer_sprint_asia/repositories/db"
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
