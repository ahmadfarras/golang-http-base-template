package usecase

import (
	"ahmadfarras/golang-http-base-template/app/configuration/logger"
	errDomain "ahmadfarras/golang-http-base-template/app/domain/error"
	"ahmadfarras/golang-http-base-template/app/domain/model/aggregate"
	"ahmadfarras/golang-http-base-template/app/domain/model/request"
	res "ahmadfarras/golang-http-base-template/app/domain/model/response"
	"ahmadfarras/golang-http-base-template/app/domain/repository"
	"context"
)

type CategoryUsecaseImpl struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryUsecaseImpl(categoryRepository repository.CategoryRepository) CategoryUsecase {
	return &CategoryUsecaseImpl{
		categoryRepository: categoryRepository,
	}
}

func (c *CategoryUsecaseImpl) Create(ctx context.Context, request request.CategoryCreateRequest) error {
	log := logger.FromCtx(ctx)

	newCategory := aggregate.Category{Name: request.Name}
	err := c.categoryRepository.Save(ctx, newCategory)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (c *CategoryUsecaseImpl) Update(ctx context.Context, id int, request request.CategoryUpdateRequest) error {
	log := logger.FromCtx(ctx)

	category, err := c.categoryRepository.GetById(ctx, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	if category == nil {
		return errDomain.CategoryNotFoundError
	}

	updatedCategory := category.UpdateCategory(request.Name)

	err = c.categoryRepository.Update(ctx, updatedCategory)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (c *CategoryUsecaseImpl) GetById(ctx context.Context, id int) (res.CategoryDetailResponse, error) {
	log := logger.FromCtx(ctx)

	category, err := c.categoryRepository.GetById(ctx, id)
	if err != nil {
		log.Error(err.Error())
		return res.CategoryDetailResponse{}, err
	}

	if category == nil {
		return res.CategoryDetailResponse{}, errDomain.CategoryNotFoundError
	}

	return res.CreateCategoryDetailResponse(*category), nil
}

func (c *CategoryUsecaseImpl) Delete(ctx context.Context, id int) error {
	category, err := c.categoryRepository.GetById(ctx, id)
	log := logger.FromCtx(ctx)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	if category == nil {
		return errDomain.CategoryNotFoundError
	}

	err = c.categoryRepository.Delete(ctx, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
func (c *CategoryUsecaseImpl) GetAll(ctx context.Context) ([]res.CategoryDetailResponse, error) {
	log := logger.FromCtx(ctx)
	var categoriesDetailResponse []res.CategoryDetailResponse

	categories, err := c.categoryRepository.GetAll(ctx)
	if err != nil {
		log.Error(err.Error())
		return categoriesDetailResponse, err
	}

	return res.CreateCategoriesDetailResponse(categories), nil
}
