package authHandler

import (

	"encoding/json"
	"io"
	"net/http"

	authService "github.com/lukaslinardi/xyz_multifinance_api/service/auth"

	cg "github.com/lukaslinardi/xyz_multifinance_api/domain/constants/general"
	"github.com/lukaslinardi/xyz_multifinance_api/domain/model/auth"
	mc "github.com/lukaslinardi/xyz_multifinance_api/domain/model/auth"

	"github.com/lukaslinardi/xyz_multifinance_api/domain/model/general"
	"github.com/lukaslinardi/xyz_multifinance_api/domain/utils"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	Auth authService.Auth
	conf general.AppService
	log  *logrus.Logger
}

func NewAuthHandler(auth authService.Auth, conf general.AppService, logger *logrus.Logger) AuthHandler {
	return AuthHandler{
		Auth: auth,
		conf: conf,
		log:  logger,
	}
}

func (ah AuthHandler) Login(res http.ResponseWriter, req *http.Request) {

	respData := &utils.ResponseDataV3{
		Status: cg.Fail,
	}

	var request auth.Login

	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		respData.Message = map[string]string{
			"en": cg.HandlerErrorRequestDataEmpty,
			"id": cg.HandlerErrorRequestDataEmptyID,
		}
		respData.ErrorDebug = err.Error()
		utils.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &request)
	if err != nil {
		respData.Message = map[string]string{
			"en": cg.HandlerErrorRequestDataNotValid,
			"id": cg.HandlerErrorRequestDataNotValidID,
		}
		respData.ErrorDebug = err.Error()
		utils.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	response, message, err := ah.Auth.Login(req.Context(), request)
	if err != nil {
		respData.Message = message
		respData.ErrorDebug = err.Error()
		respData.ResponseFormatter()
		utils.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &utils.ResponseDataV3{
		Status:  cg.Success,
		Message: message,
		Detail:  response,
	}
	utils.WriteResponse(res, respData, http.StatusOK)
	return

}

func (ah AuthHandler) SignUp(res http.ResponseWriter, req *http.Request) {
	respData := &utils.ResponseDataV3{
		Status: cg.Fail,
	}

	var request mc.SignUp

	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		respData.Message = map[string]string{
			"en": cg.HandlerErrorRequestDataEmpty,
			"id": cg.HandlerErrorRequestDataEmptyID,
		}
		respData.ErrorDebug = err.Error()
		utils.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &request)
	if err != nil {
		respData.Message = map[string]string{
			"en": cg.HandlerErrorRequestDataNotValid,
			"id": cg.HandlerErrorRequestDataNotValidID,
		}
		respData.ErrorDebug = err.Error()
		utils.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	message, err := ah.Auth.InsertUser(req.Context(), request)
	if err != nil {
		respData.Message = message
		respData.ErrorDebug = err.Error()
		respData.ResponseFormatter()
		utils.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &utils.ResponseDataV3{
		Status:  cg.Success,
		Message: message,
	}

	utils.WriteResponse(res, respData, http.StatusOK)
	return

}
