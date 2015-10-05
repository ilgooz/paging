### paging [![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](http://godoc.org/github.com/ilgooz/paging)
> Simple pagination calculation util for Go

## Example
```go
func main(){
  var page, limit, itemsCount = 2, 3, 5

  p := paging.Paging{
    Page:  page,
    Limit: limit,
    Count: itemsCount,
  }.Calc()

  fmt.Println(p.Limit, p.Offset, p.Page, p.TotalPages) // 3, 3, 2, 2
}
```
