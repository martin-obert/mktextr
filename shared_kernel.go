package mktextr

type PagedResult[T any] struct {
	Items      []T `json:"items"`
	TotalCount int `json:"totalCount"`
}
