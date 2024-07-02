package domain

type WebResponse[T any] struct {
	Data  T      `json:"data"`
	Error string `json:"error,omitempty"`
}
