package model

type WebResponse[T any] struct {
	Code   int    `json:"code"`
	Data   T      `json:"data,omitempty"`
	Errors string `json:"errors,omitempty"`
	Token  string `json:"token,omitempty"`
}
