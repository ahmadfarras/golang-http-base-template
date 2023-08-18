package response

import "net/http"

type HttpResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func Success() HttpResponse {
	return HttpResponse{
		StatusCode: http.StatusOK,
		Message:    "Success",
	}
}

func SuccessWithData(data interface{}) HttpResponse {
	return HttpResponse{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       data,
	}
}

func Created() HttpResponse {
	return HttpResponse{
		StatusCode: http.StatusCreated,
		Message:    "Created",
	}
}

func NotFound() HttpResponse {
	return HttpResponse{
		StatusCode: http.StatusNotFound,
		Message:    "Not Found",
	}
}

func InternalError() HttpResponse {
	return HttpResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Error",
	}
}
