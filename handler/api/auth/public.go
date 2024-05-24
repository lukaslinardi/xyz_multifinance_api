package authHandler

import (
	"context"
	//	"crypto/sha256"
	"strings"

	//"crypto/rand"
	"fmt"
	"time"

	//	"time"
	//	"crypto/sha256"
	//	"fmt"
	"net/http"

	//	"strings"
	//"time"

	"github.com/lukaslinardi/xyz_multifinance_api/domain/utils"

	cg "github.com/lukaslinardi/xyz_multifinance_api/domain/constants/general"
	dg "github.com/lukaslinardi/xyz_multifinance_api/domain/model/general"

	"github.com/sirupsen/logrus"
)

type PublicHandler struct {
	log  *logrus.Logger
	Conf dg.AppService
}

func NewPublicHandler(conf dg.AppService, logger *logrus.Logger) PublicHandler {
	return PublicHandler{
		log:  logger,
		Conf: conf,
	}
}

type Session struct{}

func (ph PublicHandler) AuthValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		respData := utils.ResponseData{
			Status: cg.Fail,
		}

		authorization := req.Header.Get("Authorization")
		authorizationID := req.Header.Get("Authorization-ID")

		if authorization == "" {
			respData.Message = "Token Not Valid"
			utils.WriteResponse(res, respData, http.StatusUnauthorized)
			return
		}

		if authorizationID == "" {
			respData.Message = "Token Not Valid"
			utils.WriteResponse(res, respData, http.StatusUnauthorized)
			return
		}

		authUnix, err := utils.StrToInt64(authorizationID)
		if err != nil {
			fmt.Println("error3")
			respData.Message = "Token Not Valid"
			utils.WriteResponse(res, respData, http.StatusUnauthorized)
			return
		}

		authTime := time.Unix(authUnix, 0)
		if time.Now().UTC().Unix() > (authTime.UTC().Add(cg.Time1Min)).Unix() {
			fmt.Println("error4")
			respData.Message = "Token Not Valid"
			utils.WriteResponse(res, respData, http.StatusUnauthorized)
			return
		}

		// authCompareByte := sha256.Sum256([]byte(fmt.Sprintf("%s%s", ph.Conf.Authorization.Public.SecretKey, authorizationID)))
		// authCompare := fmt.Sprintf("%x", authCompareByte)
		//
		// if authCompare != authorization {
		// 	fmt.Println("error5")
		// 	respData.Message = "Token Not Valid"
		// 	utils.WriteResponse(res, respData, http.StatusUnauthorized)
		// 	return
		// }

		authorizationKeyHeader := req.Header.Get("Authorization-Key")
		if authorizationKeyHeader != "" {
			fmt.Println("error6")
			if !strings.Contains(authorizationKeyHeader, "Bearer") {
				fmt.Println("error7")
				ph.log.Error(fmt.Errorf("invalid Token Format"))
				respData.Message = "Invalid Token Format"
				utils.WriteResponse(res, respData, http.StatusBadRequest)
				return
			}
			accessToken := strings.Replace(authorizationKeyHeader, "Bearer ", "", -1)

			claims, err := utils.CheckAccessToken(accessToken)
			if err != nil {
				fmt.Println("error8")
				respData.Message = "Token expired"
				utils.WriteResponse(res, respData, http.StatusBadRequest)
				return
			}

			ctx := context.WithValue(req.Context(), "session", claims["session"])
			req = req.WithContext(ctx)

			next.ServeHTTP(res, req)
			return
		}

		ctx := context.WithValue(req.Context(), Session{}, authorization)
		req = req.WithContext(ctx)

		next.ServeHTTP(res, req)
	})
}
