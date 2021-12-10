package http

type Pagination struct {
	Limit  uint64 `json:"limit" swag_example:"20"`
	Offset uint64 `json:"offset" swag_example:"30"`
	Total  uint64 `json:"total" swag_example:"100"`
}

var DefaultPagination = Pagination{
	Limit:  uint64(20),
	Offset: uint64(5),
	Total:  uint64(500),
}

func (p *Pagination) SetTotal(total uint64) {
	p.Total = total
}

func (p *Pagination) SetLimit() {
	if p.Limit == 0 {
		p.Limit = p.Total
	}
}

type PaginationResult struct {
	Pagination Pagination  `json:"pagination"`
	Result     interface{} `json:"results"`
}
