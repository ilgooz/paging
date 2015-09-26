package paging

import "math"

// Paging defines fields for a pagination
type Paging struct {
	Page                int
	Limit, DefaultLimit int
	Offset              int
	TotalPages          int
	Count               int
}

// Calc produces a Paging calculated over
// Page, Limit, Count and DefaultLimit values of given Paging
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

	if p.Count < p.Limit {
		p.TotalPages = 1
	} else {
		p.TotalPages = int(math.Ceil(float64(p.Count) / float64(p.Limit)))
	}

	if p.Page > p.TotalPages {
		p.Page = p.TotalPages
	}

	p.Offset = (p.Page - 1) * p.Limit

	return &p
}
