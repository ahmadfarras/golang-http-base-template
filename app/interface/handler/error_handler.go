package handler

import (
	errDomain "ahmadfarras/golang-http-base-template/app/domain/error"
	"ahmadfarras/golang-http-base-template/app/interface/response"
	"context"
	"errors"
)

func HandleError(ctx context.Context, err error) response.HttpResponse {

	if errors.Is(err, errDomain.CategoryNotFoundError) {
		return response.NotFound()
	}

	return response.InternalError()
}
