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
	"go.uber.org/zap"
)

func StartHttpApp() {
	log := logger.InitLogger()
	defer log.Sync()

	if err := godotenv.Load(".env"); err != nil {
		log.Panic(err.Error())
	}

	db := config.NewSqlDb(log)
	router := routes.InitRoute(db)

	StartServer(router, log)
}

func StartServer(router *httprouter.Router, log *zap.Logger) {
	server := http.Server{
		Addr:    fmt.Sprintf(":%v", os.Getenv(constant.APP_PORT)),
		Handler: router,
	}

	log.Info("Server listening on" + server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
