package genricresponses

import "net/http"

type GenericResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

type UserAuthFailedLogin struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

var GenericInternalServerErrorResponse = GenericResponse{
	StatusCode: http.StatusInternalServerError,
	Message:    "It's not you. It's us",
}

var GenericBadRequestResponse = GenericResponse{
	StatusCode: http.StatusBadRequest,
	Message:    "Bad/Invalid Request",
}
