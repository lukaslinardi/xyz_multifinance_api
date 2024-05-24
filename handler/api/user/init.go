package userHandler

import (
	"github.com/lukaslinardi/fullstack_engineer_sprint_asia/domain/model/general"
	"github.com/lukaslinardi/fullstack_engineer_sprint_asia/service"
	"github.com/sirupsen/logrus"
)

type UserDataHandler struct {
	User UserHandler
}

func NewUserDataHandler(sv service.Service, conf general.AppService, logger *logrus.Logger) UserDataHandler {
	return UserDataHandler{
		User: NewUserHandler(sv.User.User, conf, logger),
	}
}
