package db

import (
	"github.com/lukaslinardi/fullstack_engineer_sprint_asia/infra"
	"github.com/sirupsen/logrus"
)

type UserConfig struct {
	db  *infra.DatabaseList
	log *logrus.Logger
}

func newUser(db *infra.DatabaseList, logger *logrus.Logger) UserConfig {
	return UserConfig{
		db:  db,
		log: logger,
	}
}

type User interface {}
