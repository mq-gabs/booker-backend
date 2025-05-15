package utils

type ListResponse[T any] struct {
	List  []T `json:"list"`
	Count int `json:"count"`
}
