package routes

import (
	"ahmadfarras/golang-http-base-template/app/domain/usecase"
	"ahmadfarras/golang-http-base-template/app/infrastructure/persistence"
	"ahmadfarras/golang-http-base-template/app/interface/controller"
	"database/sql"

	"github.com/julienschmidt/httprouter"
)

func CategoryRoute(db *sql.DB, router *httprouter.Router) {
	categoryRepository := persistence.NewCategoryRepositoryImpl(db)
	categoryUsecase := usecase.NewCategoryUsecaseImpl(categoryRepository)
	categoryController := controller.NewCategoryController(categoryUsecase)

	router.POST("/category", categoryController.Create)
	router.GET("/category", categoryController.GetAll)
	router.GET("/category/:categoryId", categoryController.GetById)
	router.PUT("/category/:categoryId", categoryController.Update)
	router.DELETE("/category/:categoryId", categoryController.Delete)
}
