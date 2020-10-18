package types

import "strconv"

type PaginatorRequest struct {
	Page int64 `form:"Page" json:"Page,omitempty"`
	Size int64 `form:"Size" json:"Size,omitempty"`
}

func (p *PaginatorRequest) ToMap() map[string]string {
	return map[string]string{
		"Page": strconv.FormatInt(p.Page, 10),
		"Size": strconv.FormatInt(p.Size, 10),
	}
}

type PaginatorResponse struct {
	Page       int `json:"Page,omitempty"`
	Size       int `json:"Size,omitempty"`
	Total      int `json:"Total,omitempty"`
	TotalItems int `json:"TotalItems,omitempty"`
}
