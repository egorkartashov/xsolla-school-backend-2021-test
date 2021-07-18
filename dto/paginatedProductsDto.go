package dto

type PaginatedProductsDto struct {
	Pagination Pagination   `json:"pagination"`
	Data       []ProductDto `json:"data"`
}
