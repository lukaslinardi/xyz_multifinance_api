package service

import (
	"github.com/lukaslinardi/xyz_multifinance_api/domain/model/general"
	"github.com/lukaslinardi/xyz_multifinance_api/infra"
	repository "github.com/lukaslinardi/xyz_multifinance_api/repositories"
	authService "github.com/lukaslinardi/xyz_multifinance_api/service/auth"
	userService "github.com/lukaslinardi/xyz_multifinance_api/service/user"
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
