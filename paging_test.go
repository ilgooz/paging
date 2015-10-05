package paging

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = map[Paging][]int{
	Paging{Page: 0, Limit: 0, Count: 0}:   []int{10, 0, 1, 1},
	Paging{Page: 1, Limit: 10, Count: 19}: []int{10, 0, 1, 2},
	Paging{Page: 3, Limit: 5, Count: 248}: []int{5, 10, 3, 50},
	Paging{Page: 2, Limit: 1, Count: 100}: []int{1, 1, 2, 100},
	Paging{Page: 6, Limit: 3, Count: 13}:  []int{3, 12, 5, 5},
	Paging{Page: 2, Limit: 3, Count: 5}:   []int{3, 3, 2, 2},
}

func TestAll(t *testing.T) {
	ass := assert.New(t)

	for c, val := range cases {
		p := c.Calc()
		ass.Equal(val, []int{p.Limit, p.Offset, p.Page, p.TotalPages})
	}
}

func Example() {
	var page, limit, itemsCount = 2, 3, 5

	p := Paging{
		Page:  page,
		Limit: limit,
		Count: itemsCount,
	}.Calc()

	fmt.Println(p.Limit, p.Offset, p.Page, p.TotalPages) // 3, 3, 2, 2
}
