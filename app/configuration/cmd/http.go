package main

import (
	"ahmadfarras/golang-http-base-template/app/configuration/constant"
	"ahmadfarras/golang-http-base-template/app/configuration/logger"
	"ahmadfarras/golang-http-base-template/app/configuration/routes"
	"ahmadfarras/golang-http-base-template/app/infrastructure/config"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func StartHttpApp() {
	logger.InitLogrus()
	if err := godotenv.Load(".env"); err != nil {
		logrus.Panic(err)
	}

	db := config.NewSqlDb()
	router := routes.InitRoute(db)

	StartServer(router)
}

func StartServer(router *httprouter.Router) {
	server := http.Server{
		Addr:    fmt.Sprintf(":%v", os.Getenv(constant.APP_PORT)),
		Handler: router,
	}

	logrus.Info("Server listening on", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		logrus.Fatal(err)
	}
}
