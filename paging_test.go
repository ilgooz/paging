package paging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = map[Paging][]int{
	Paging{Page: 0, Limit: 0, Count: 0}:   []int{10, 0, 1, 0},
	Paging{Page: 1, Limit: 10, Count: 19}: []int{10, 0, 1, 2},
	Paging{Page: 3, Limit: 5, Count: 248}: []int{5, 10, 3, 50},
	Paging{Page: 2, Limit: 1, Count: 100}: []int{1, 1, 2, 100},
}

func TestAll(t *testing.T) {
	ass := assert.New(t)

	for c, val := range cases {
		p := c.Calc()
		ass.Equal(val, []int{p.Limit, p.Offset, p.Page, p.TotalPages})
	}
}
