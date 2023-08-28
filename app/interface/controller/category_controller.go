package controller

import (
	"ahmadfarras/golang-http-base-template/app/domain/usecase"
	"ahmadfarras/golang-http-base-template/app/interface/handler"
	"ahmadfarras/golang-http-base-template/app/interface/helper"
	"ahmadfarras/golang-http-base-template/app/interface/response"
	"net/http"
	"strconv"

	"ahmadfarras/golang-http-base-template/app/domain/model/request"

	"github.com/julienschmidt/httprouter"
)

type CategoryController struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCategoryController(categoryUsecase usecase.CategoryUsecase) *CategoryController {
	return &CategoryController{
		categoryUsecase: categoryUsecase,
	}
}

func (c *CategoryController) Create(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {

	createRequest := request.CategoryCreateRequest{}
	err := helper.JsonDecode(req.Context(), req.Body, &createRequest)
	if err != nil {
		resp := handler.HandleError(err)
		helper.JsonEncode(req.Context(), resp.StatusCode, writer, resp)
		return
	}

	err = c.categoryUsecase.Create(req.Context(), createRequest)
	if err != nil {
		resp := handler.HandleError(err)
		helper.JsonEncode(req.Context(), resp.StatusCode, writer, resp)
		return
	}

	resp := response.Created()
	helper.JsonEncode(req.Context(), resp.StatusCode, writer, resp)
}

func (c *CategoryController) GetById(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {

	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	if err != nil {
		resp := handler.HandleError(err)
		helper.JsonEncode(req.Context(), resp.StatusCode, writer, resp)
		return
	}

	res, err := c.categoryUsecase.GetById(req.Context(), categoryId)
	if err != nil {
		resp := handler.HandleError(err)
		helper.JsonEncode(req.Context(), resp.StatusCode, writer, resp)
		return
	}

	resp := response.SuccessWithData(res)
	helper.JsonEncode(req.Context(), resp.StatusCode, writer, resp)
}

func (c *CategoryController) GetAll(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	res, err := c.categoryUsecase.GetAll(req.Context())
	if err != nil {
		resp := handler.HandleError(err)
		helper.JsonEncode(req.Context(), resp.StatusCode, writer, resp)
		return
	}

	resp := response.SuccessWithData(res)
	helper.JsonEncode(req.Context(), resp.StatusCode, writer, resp)
}

func (c *CategoryController) Update(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {

	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	if err != nil {
		resp := handler.HandleError(err)
		helper.JsonEncode(req.Context(), resp.StatusCode, writer, resp)
		return
	}

	updateRequest := request.CategoryUpdateRequest{}
	err = helper.JsonDecode(req.Context(), req.Body, &updateRequest)
	if err != nil {
		resp := handler.HandleError(err)
		helper.JsonEncode(req.Context(), resp.StatusCode, writer, resp)
		return
	}

	err = c.categoryUsecase.Update(req.Context(), categoryId, updateRequest)
	if err != nil {
		resp := handler.HandleError(err)
		helper.JsonEncode(req.Context(), resp.StatusCode, writer, resp)
		return
	}

	resp := response.Success()
	helper.JsonEncode(req.Context(), resp.StatusCode, writer, resp)
}

func (c *CategoryController) Delete(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	if err != nil {
		resp := handler.HandleError(err)
		helper.JsonEncode(req.Context(), resp.StatusCode, writer, resp)
		return
	}

	err = c.categoryUsecase.Delete(req.Context(), categoryId)
	if err != nil {
		resp := handler.HandleError(err)
		helper.JsonEncode(req.Context(), resp.StatusCode, writer, resp)
		return
	}

	resp := response.Success()
	helper.JsonEncode(req.Context(), resp.StatusCode, writer, resp)
}
