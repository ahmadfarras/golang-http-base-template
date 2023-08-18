package usecase

import (
	"ahmadfarras/golang-http-base-template/app/domain/model/request"
	res "ahmadfarras/golang-http-base-template/app/domain/model/response"
	"context"
)

type CategoryUsecase interface {
	Create(ctx context.Context, request request.CategoryCreateRequest) error
	Update(ctx context.Context, id int, request request.CategoryUpdateRequest) error
	GetById(ctx context.Context, id int) (res.CategoryDetailResponse, error)
	GetAll(ctx context.Context) ([]res.CategoryDetailResponse, error)
	Delete(ctx context.Context, id int) error
}
