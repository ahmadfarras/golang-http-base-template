package helper

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func JsonEncode(code int, w http.ResponseWriter, resp interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		logrus.Panic(err)
	}
	w.Write(jsonResp)
}
