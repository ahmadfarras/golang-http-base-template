package helper

import (
	"ahmadfarras/golang-http-base-template/app/configuration/logger"
	"context"
	"encoding/json"
	"net/http"
)

func JsonEncode(ctx context.Context, code int, w http.ResponseWriter, resp interface{}) {
	log := logger.FromCtx(ctx)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Panic(err.Error())
	}
	w.Write(jsonResp)
}
