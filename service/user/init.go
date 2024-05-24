package userService

import (
	// "olin_psef/domain/model/general"
	// "olin_psef/infra"
	// "olin_psef/repository"

	"github.com/lukaslinardi/xyz_multifinance_api/domain/model/general"
	"github.com/lukaslinardi/xyz_multifinance_api/infra"
	repository "github.com/lukaslinardi/xyz_multifinance_api/repositories"
	"github.com/sirupsen/logrus"
)

type UserData struct {
	User UserService
}

func NewUser(repo repository.Repo, conf general.AppService, dbList *infra.DatabaseList, logger *logrus.Logger) UserData {
	return UserData{
		User: newUserService(repo.Database, conf, dbList, logger),
	}
}
