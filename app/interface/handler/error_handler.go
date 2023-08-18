package handler

import (
	errDomain "ahmadfarras/golang-http-base-template/app/domain/error"
	"ahmadfarras/golang-http-base-template/app/interface/response"
	"errors"

	"github.com/sirupsen/logrus"
)

func HandleError(err error) response.HttpResponse {

	if errors.Is(err, errDomain.CategoryNotFoundError) {
		logrus.Warn(err)
		return response.NotFound()
	}

	logrus.Error(err)
	return response.InternalError()
}
