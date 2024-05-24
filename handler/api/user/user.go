package userHandler

import (
	// "encoding/json"
	// "io/ioutil"
	// "net/http"

	"github.com/lukaslinardi/xyz_multifinance_api/domain/model/general"
	"github.com/sirupsen/logrus"

	userService "github.com/lukaslinardi/xyz_multifinance_api/service/user"
)

type UserHandler struct {
	User userService.User
	conf general.AppService
	log  *logrus.Logger
}

func NewUserHandler(user userService.User, conf general.AppService, logger *logrus.Logger) UserHandler {
	return UserHandler{
		User: user,
		conf: conf,
		log:  logger,
	}
}

