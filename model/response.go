package model

type SuccessResponse[T any] struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   T      `json:"data"`
}

type ErrorResponse[T string | map[string]any] struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors T      `json:"errors"`
}
