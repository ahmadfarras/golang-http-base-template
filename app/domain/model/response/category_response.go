package response

import "ahmadfarras/golang-http-base-template/app/domain/model/aggregate"

type CategoryDetailResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func CreateCategoryDetailResponse(category aggregate.Category) CategoryDetailResponse {
	return CategoryDetailResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func CreateCategoriesDetailResponse(categories []aggregate.Category) []CategoryDetailResponse {
	var categoriesDetailResponse []CategoryDetailResponse
	for _, c := range categories {
		category := c
		categoryDetailResponse := CategoryDetailResponse{
			Id:   category.Id,
			Name: category.Name,
		}

		categoriesDetailResponse = append(categoriesDetailResponse, categoryDetailResponse)
	}

	return categoriesDetailResponse
}
