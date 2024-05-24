package userHandler

import (
	"github.com/lukaslinardi/xyz_multifinance_api/domain/model/general"
	"github.com/lukaslinardi/xyz_multifinance_api/service"
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
