package utils

type Paging struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func (p *Paging) Offset() int {
	return (p.Page - 1) * p.PageSize
}

func (p *Paging) Limit() int {
	return p.PageSize
}
