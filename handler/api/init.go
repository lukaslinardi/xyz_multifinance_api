package api

import (
	"github.com/lukaslinardi/xyz_multifinance_api/domain/model/general"
	userHandler "github.com/lukaslinardi/xyz_multifinance_api/handler/api/user"
	"github.com/lukaslinardi/xyz_multifinance_api/service"
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
