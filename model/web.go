package model

type WebResponse struct {
	Code   int    `json:"code"`
	Data   any    `json:"data,omitempty"`
	Errors string `json:"errors,omitempty"`
	Token  string `json:"token,omitempty"`
}

func ToWebResponse[T any](code int, toModelResponse func() T) WebResponse {
	data := toModelResponse()

	return WebResponse{
		Code: code,
		Data: data,
	}
}
