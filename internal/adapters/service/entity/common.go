package entity

type GenericIdRequest[T any] struct {
	Id T
}

type GenericListRequest[T any] struct {
	Ids    []T
	Search string
}
