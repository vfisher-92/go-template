package paginator

const defaultLimit = 20
const defaultOffset = 0

type Paginator struct {
	Limit int
	Page  int
}

func (p *Paginator) GetLimit() int {
	limit := defaultLimit

	if p.Limit > 0 {
		limit = p.Limit
	}

	return limit
}

func (p *Paginator) GetOffset() int {
	offset := defaultOffset
	limit := defaultLimit

	if p.Page > 0 {
		offset = (p.Page - 1) * limit
	}

	return offset
}
