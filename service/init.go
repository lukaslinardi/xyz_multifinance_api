package service

import (
	"github.com/lukaslinardi/fullstack_engineer_sprint_asia/domain/model/general"
	"github.com/lukaslinardi/fullstack_engineer_sprint_asia/infra"
	repository "github.com/lukaslinardi/fullstack_engineer_sprint_asia/repositories"
	authService "github.com/lukaslinardi/fullstack_engineer_sprint_asia/service/auth"
	userService "github.com/lukaslinardi/fullstack_engineer_sprint_asia/service/user"
	"github.com/sirupsen/logrus"
)

type Service struct {
	User userService.UserData
	Auth authService.AuthData
}

func NewService(repo repository.Repo, conf general.AppService, dbList *infra.DatabaseList, logger *logrus.Logger) Service {
	return Service{
		User: userService.NewUser(repo, conf, dbList, logger),
		Auth: authService.NewAuth(repo, conf, dbList, logger),
	}
}
