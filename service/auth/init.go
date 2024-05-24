package authService

import (
	"github.com/lukaslinardi/fullstack_engineer_sprint_asia/domain/model/general"
	"github.com/lukaslinardi/fullstack_engineer_sprint_asia/infra"
	repository "github.com/lukaslinardi/fullstack_engineer_sprint_asia/repositories"
	"github.com/sirupsen/logrus"
)

type AuthData struct {
	Auth AuthService
}

func NewAuth(repo repository.Repo, conf general.AppService, dbList *infra.DatabaseList, logger *logrus.Logger) AuthData {
	return AuthData{
		Auth: newAuthService(repo.Database, conf, dbList, logger),
	}
}
