package authHandler

import (
	"github.com/lukaslinardi/fullstack_engineer_sprint_asia/domain/model/general"
	"github.com/lukaslinardi/fullstack_engineer_sprint_asia/service"
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
