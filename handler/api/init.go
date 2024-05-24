package api

import (
	"github.com/lukaslinardi/fullstack_engineer_sprint_asia/domain/model/general"
	userHandler "github.com/lukaslinardi/fullstack_engineer_sprint_asia/handler/api/user"
	"github.com/lukaslinardi/fullstack_engineer_sprint_asia/service"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	User userHandler.UserDataHandler
}

func NewHandler(sv service.Service, conf general.AppService, logger *logrus.Logger) Handler {
	return Handler{
		User: userHandler.NewUserDataHandler(sv, conf, logger),
	}
}
