package userService

import (
	"github.com/lukaslinardi/fullstack_engineer_sprint_asia/domain/model/general"
	"github.com/lukaslinardi/fullstack_engineer_sprint_asia/infra"
	"github.com/lukaslinardi/fullstack_engineer_sprint_asia/repositories/db"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	db     db.Database
	conf   general.AppService
	dbConn *infra.DatabaseList
	log    *logrus.Logger
}

func newUserService(db db.Database, conf general.AppService, dbConn *infra.DatabaseList, logger *logrus.Logger) UserService {
	return UserService{
		db:     db,
		conf:   conf,
		dbConn: dbConn,
		log:    logger,
	}
}

type User interface {}
