package paging

import "math"

type Paging struct {
	Page                int
	Limit, DefaultLimit int
	Offset              int
	TotalPages          int
	Count               int
}

func (p Paging) Calc() *Paging {
	if p.Page < 1 {
		p.Page = 1
	}

	if p.Limit < 1 {
		if p.DefaultLimit > 0 {
			p.Limit = p.DefaultLimit
		} else {
			p.Limit = 10
		}
	}

	p.Offset = (p.Page - 1) * p.Limit

	if p.Count < p.Limit {
		p.TotalPages = 1
	} else {
		p.TotalPages = int(math.Ceil(float64(p.Count) / float64(p.Limit)))
	}

	return &p
}
