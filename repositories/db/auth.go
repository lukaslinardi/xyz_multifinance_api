package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/lukaslinardi/fullstack_engineer_sprint_asia/domain/model/auth"
	"github.com/lukaslinardi/fullstack_engineer_sprint_asia/infra"
	"github.com/sirupsen/logrus"
)

type AuthConfig struct {
	db  *infra.DatabaseList
	log *logrus.Logger
}

func newAuth(db *infra.DatabaseList, logger *logrus.Logger) AuthConfig {
	return AuthConfig{
		db:  db,
		log: logger,
	}
}

type Auth interface {
	InsertUser(ctx context.Context, tx *sql.Tx, data auth.SignUp) error
	IsExist(ctx context.Context, nik string) (bool, error)
}

func (ac AuthConfig) IsExist(ctx context.Context, nik string) (bool, error) {
    var isExist bool
	script := `SELECT exists(SELECT * FROM users where nik = ?)`

	query, args, err := ac.db.Backend.Read.In(script, nik)
	if err != nil {
		return isExist, err
	}

	query = ac.db.Backend.Read.Rebind(query)
	ac.db.Backend.Read.Get(&isExist, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return isExist, err
	}
	return isExist, nil

}

func (ac AuthConfig) InsertUser(ctx context.Context, tx *sql.Tx, data auth.SignUp) error {
	script := `INSERT INTO users
	(nik, fullname, legal_name, birth_place, birth_date, salary, ktp_picture, picture, password)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?);`

	param := make([]interface{}, 0)

	param = append(param, data.Nik)
	param = append(param, data.Fullname)
	param = append(param, data.Legalname)
	param = append(param, data.BirthPlace)
	param = append(param, time.Now())
	param = append(param, data.Salary)
	param = append(param, data.KtpPics)
	param = append(param, data.Pics)
	param = append(param, data.Password)

	query, args, err := ac.db.Backend.Read.In(script, param...)

	query = ac.db.Backend.Read.Rebind(query)

	var res *sql.Row
	if tx == nil {
		res = ac.db.Backend.Write.QueryRow(ctx, query, args...)
	} else {
		res = tx.QueryRowContext(ctx, query, args...)
	}

	if err != nil {
		return err
	}

	err = res.Err()
	if err != nil {
		return err
	}

	return nil
}
