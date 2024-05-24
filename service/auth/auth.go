package authService

import (
	"context"
	"errors"

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
