package domain

type PaginacaoFiltro struct {
	Page int
	PageSize int
}



type PaginacaoResultado[T any] struct {
	Items []T
	TotalCount int
	Page int
	PageSize int
}