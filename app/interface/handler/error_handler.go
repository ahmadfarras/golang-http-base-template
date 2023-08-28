package handler

import (
	errDomain "ahmadfarras/golang-http-base-template/app/domain/error"
	"ahmadfarras/golang-http-base-template/app/interface/response"
	"errors"
)

func HandleError(err error) response.HttpResponse {

	if errors.Is(err, errDomain.CategoryNotFoundError) {
		return response.NotFound()
	}

	return response.InternalError()
}
