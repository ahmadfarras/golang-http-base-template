package helper

import (
	"ahmadfarras/golang-http-base-template/app/configuration/logger"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func JsonEncode(ctx context.Context, code int, w http.ResponseWriter, resp interface{}) {
	log := logger.FromCtx(ctx)

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	encoder := json.NewEncoder(w)

	err := encoder.Encode(resp)
	if err != nil {
		log.Panic(err.Error())
	}
}

func JsonDecode(ctx context.Context, r io.Reader, req interface{}) error {
	log := logger.FromCtx(ctx)

	decoder := json.NewDecoder(r)
	err := decoder.Decode(&req)
	if err != nil {
		log.Panic(err.Error())
	}

	return nil
}
