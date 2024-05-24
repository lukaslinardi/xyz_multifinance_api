package authService

import (
	"github.com/lukaslinardi/xyz_multifinance_api/domain/model/general"
	"github.com/lukaslinardi/xyz_multifinance_api/infra"
	repository "github.com/lukaslinardi/xyz_multifinance_api/repositories"
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
