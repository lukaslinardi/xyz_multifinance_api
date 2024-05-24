package authHandler

import (
	"github.com/lukaslinardi/xyz_multifinance_api/domain/model/general"
	"github.com/lukaslinardi/xyz_multifinance_api/service"
	"github.com/sirupsen/logrus"
)

type AuthDataHandler struct {
	Auth AuthHandler
}

func NewAuthDataHandler(sv service.Service, conf general.AppService, logger *logrus.Logger) AuthDataHandler {
	return AuthDataHandler{
		Auth: NewAuthHandler(sv.Auth.Auth, conf, logger),
	}
}
