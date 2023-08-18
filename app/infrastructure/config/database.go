package config

import (
	"ahmadfarras/golang-http-base-template/app/configuration/constant"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func NewSqlDb() *sql.DB {
	db, err := sql.Open(
		os.Getenv(constant.DATABASE_DRIVER),
		fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
			os.Getenv(constant.DATABASE_USER),
			os.Getenv(constant.DATABASE_PASSWORD),
			os.Getenv(constant.DATABASE_HOST),
			os.Getenv(constant.DATABASE_PORT),
			os.Getenv(constant.DATABASE_NAME),
		),
	)
	if err != nil {
		logrus.Fatal(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
