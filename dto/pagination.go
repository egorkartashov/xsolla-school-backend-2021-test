package dto

import "fmt"

type Pagination struct {
	Offset   int     `json:"offset"`
	Limit    int     `json:"limit"`
	Size     int     `json:"size"`
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
}

func CreatePaginationWithLinks(path string, offset, limit, size int) Pagination {
	pagination := Pagination{
		Offset: offset,
		Limit:  limit,
		Size:   size,
	}

	if size == limit {
		nextLink := fmt.Sprintf("%s?offset=%d&limit=%d", path, offset+limit, limit)
		pagination.Next = &nextLink
	}

	if offset > 0 {
		prevOffset := offset - limit
		prevLimit := limit
		if prevOffset < 0 {
			prevLimit = offset
			prevOffset = 0
		}
		previousLink := fmt.Sprintf("%s?offset=%d&limit=%d", path, prevOffset, prevLimit)
		pagination.Previous = &previousLink
	}

	return pagination
}
