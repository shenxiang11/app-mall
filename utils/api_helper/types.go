package api_helper

import "errors"

type Response struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

var (
	ErrInvalidBody = errors.New("请检查你的请求踢")
)
