package api

import (
	"github.com/lukaslinardi/xyz_multifinance_api/domain/model/general"
	authHandler "github.com/lukaslinardi/xyz_multifinance_api/handler/api/auth"
	userHandler "github.com/lukaslinardi/xyz_multifinance_api/handler/api/user"
	"github.com/lukaslinardi/xyz_multifinance_api/service"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Auth   authHandler.AuthDataHandler
	User   userHandler.UserDataHandler
	Token  authHandler.TokenHandler
	Public authHandler.PublicHandler
}

func NewHandler(sv service.Service, conf general.AppService, logger *logrus.Logger) Handler {
	return Handler{
		Auth:   authHandler.NewAuthDataHandler(sv, conf, logger),
		Token:  authHandler.NewTokenHandler(conf, logger),
		Public: authHandler.NewPublicHandler(conf, logger),
		User:   userHandler.NewUserDataHandler(sv, conf, logger),
	}
}
