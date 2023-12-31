package persistence

import (
	"ahmadfarras/golang-http-base-template/app/configuration/logger"
	"ahmadfarras/golang-http-base-template/app/domain/model/aggregate"
	"ahmadfarras/golang-http-base-template/app/domain/repository"
	"context"
	"database/sql"
)

type CategoryRepositoryImpl struct {
	sql *sql.DB
}

func NewCategoryRepositoryImpl(sql *sql.DB) repository.CategoryRepository {
	return &CategoryRepositoryImpl{
		sql: sql,
	}
}

func (c *CategoryRepositoryImpl) Save(ctx context.Context, req aggregate.Category) error {
	log := logger.FromCtx(ctx)
	tx, err := c.sql.Begin()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	sql := "insert into category(name) values (?)"

	_, err = tx.ExecContext(ctx, sql, req.Name)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	tx.Commit()
	return err
}

func (c *CategoryRepositoryImpl) Update(ctx context.Context, category *aggregate.Category) error {
	log := logger.FromCtx(ctx)
	tx, err := c.sql.Begin()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	sql := "UPDATE category SET name = ? where id = ?"

	_, err = tx.ExecContext(ctx, sql, category.Name, category.Id)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	tx.Commit()
	return nil
}

func (c *CategoryRepositoryImpl) Delete(ctx context.Context, id int) error {
	log := logger.FromCtx(ctx)
	tx, err := c.sql.Begin()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	sql := "DELETE from category where id = ?"

	_, err = tx.ExecContext(ctx, sql, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	tx.Commit()
	return nil
}

func (c *CategoryRepositoryImpl) GetAll(ctx context.Context) ([]aggregate.Category, error) {
	var categories []aggregate.Category
	log := logger.FromCtx(ctx)

	tx, err := c.sql.Begin()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	sql := "SELECT * FROM category"

	rows, err := tx.QueryContext(ctx, sql)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		category := aggregate.Category{}
		err = rows.Scan(&category.Id, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (c *CategoryRepositoryImpl) GetById(ctx context.Context, id int) (*aggregate.Category, error) {
	var category *aggregate.Category = &aggregate.Category{}
	log := logger.FromCtx(ctx)

	tx, err := c.sql.Begin()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	sql := "SELECT * From category where id = ?"

	rows, err := tx.QueryContext(ctx, sql, id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	err = rows.Scan(&category.Id, &category.Name)
	if err != nil {
		return nil, err
	}

	return category, nil
}
