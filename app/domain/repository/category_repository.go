package repository

import (
	"ahmadfarras/golang-http-base-template/app/domain/model/aggregate"
	"context"
)

type CategoryRepository interface {
	Save(ctx context.Context, category aggregate.Category) error
	Update(ctx context.Context, category *aggregate.Category) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]aggregate.Category, error)
	GetById(ctx context.Context, id int) (*aggregate.Category, error)
}
