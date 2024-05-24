package authService

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/lukaslinardi/xyz_multifinance_api/domain/model/auth"
	"github.com/lukaslinardi/xyz_multifinance_api/domain/model/general"
	"github.com/lukaslinardi/xyz_multifinance_api/domain/utils"
	"github.com/lukaslinardi/xyz_multifinance_api/infra"
	"github.com/lukaslinardi/xyz_multifinance_api/repositories/db"
	"github.com/sirupsen/logrus"
)

type AuthService struct {
	db     db.Database
	conf   general.AppService
	dbConn *infra.DatabaseList
	log    *logrus.Logger
}

func newAuthService(db db.Database, conf general.AppService, dbConn *infra.DatabaseList, logger *logrus.Logger) AuthService {
	return AuthService{
		db:     db,
		conf:   conf,
		dbConn: dbConn,
		log:    logger,
	}
}

type Auth interface {
	InsertUser(ctx context.Context, data auth.SignUp) (map[string]string, error)
	Login(ctx context.Context, data auth.Login) (*auth.LoginResponse, map[string]string, error)
}

func (as AuthService) Login(ctx context.Context, data auth.Login) (*auth.LoginResponse, map[string]string, error) {
	internalServerError := func(err error) (*auth.LoginResponse, map[string]string, error) {
		return nil, map[string]string{
			"en": "failed to login, try again another time",
			"id": "gagal login, coba lagi dilain waktu",
		}, err
	}

	valid := data.Validate()
	if valid != nil {
		return nil, map[string]string{
			"en": "failed to login, try again another time",
			"id": "gagal login, coba lagi dilain waktu",
		}, errors.New("data not valid")
	}

	dataUser, err := as.db.Auth.GetDataUser(ctx, data.Nik)
	if err != nil {
		as.log.WithField("request", utils.StructToString(data.Nik)).WithError(err).Errorf("Login | High | fail to get user")
		return internalServerError(err)
	}

	isValid, err := utils.ComparePassword(dataUser.Password, data.Password)
	if err != nil {
		fmt.Println(err)
		as.log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("Login | Medium | fail to compare password")
		return internalServerError(err)
	}

	if !isValid {
		return nil, map[string]string{
			"en": "password incorrect",
			"id": "password salah",
		}, errors.New("password incorrect")
	}

	session, err := utils.GetEncrypt([]byte(as.conf.KeyData.User), utils.StructToString(auth.CredentialData{
		ID:       dataUser.ID,
		Fullname: dataUser.Fullname,
		Nik:      dataUser.Nik,
	}))

	generateTime := time.Now().UTC()

	accessToken, renewToken, err := utils.GenerateJWT(session)
	if err != nil {
		as.log.WithField("request", utils.StructToString(session)).Errorf("Login | High | fail generate jwt token")
		return internalServerError(err)
	}

	result := auth.LoginResponse{
		Token: auth.TokenRes{
			Access:        accessToken,
			AccessExpired: generateTime.Add(time.Duration(as.conf.Authorization.JWT.AccessTokenDuration) * time.Minute).Format(time.RFC3339),
			Renew:         renewToken,
			RenewExpired:  generateTime.Add(time.Duration(as.conf.Authorization.JWT.RefreshTokenDuration) * time.Minute).Format(time.RFC3339),
		},
		Fullname: dataUser.Fullname,
		Nik:      dataUser.Nik,
	}

	return &result, map[string]string{
		"en": "Login Successfully",
		"id": "Login Sukses",
	}, nil
}

func (as AuthService) InsertUser(ctx context.Context, data auth.SignUp) (map[string]string, error) {

	internalServerError := func(err error) (map[string]string, error) {
		return map[string]string{
			"en": "failed to register an account, try again another time",
			"id": "gagal mendaftar akun, coba lagi dilain waktu",
		}, err
	}

	valid := data.Validate()
	if valid != nil {
		return map[string]string{
			"en": "failed to register an account, try again another time",
			"id": "gagal mendaftar akun, coba lagi dilain waktu",
		}, errors.New("data tidak valid")
	}

	isExist, err := as.db.Auth.IsExist(ctx, data.Nik)
	if err != nil {
		as.log.WithField("request", utils.StructToString(isExist)).WithError(err).Errorf("Sign Up | Low | fail to begin transaction")
		return internalServerError(err)
	}

	if isExist {
		return map[string]string{
			"en": "account is already in use",
			"id": "akun sudah digunakan",
		}, errors.New("account sudah exist")
	}

	hashedPassword, err := utils.GeneratePassword(data.Password)
	if err != nil {
		return internalServerError(err)
	}

	data.Password = hashedPassword

	tx, err := as.dbConn.Backend.Write.Begin()
	if err != nil {
		as.log.WithField("request", utils.StructToString(tx)).WithError(err).Errorf("Sign Up | Low | fail to begin transaction")
		return internalServerError(err)
	}

	err = as.db.Auth.InsertUser(ctx, tx, data)
	if err != nil {
		tx.Rollback()
		as.log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("Sign Up | High | fail to insert user")
		return internalServerError(err)
	}

	err = tx.Commit()
	if err != nil {
		as.log.WithField("request", utils.StructToString(tx)).WithError(err).Errorf("Invoice | High | fail to commit transaction")
		tx.Rollback()
		return internalServerError(err)
	}

	return map[string]string{
		"en": "Successfully",
		"id": "Sukses",
	}, nil
}
