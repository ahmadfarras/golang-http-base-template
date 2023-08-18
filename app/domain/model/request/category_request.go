package request

type CategoryCreateRequest struct {
	Name string `json:"name"`
}

type CategoryUpdateRequest struct {
	Name string `json:"name"`
}
