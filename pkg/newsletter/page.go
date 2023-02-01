package newsletter

import "fmt"

type Page[T any] struct {
	Number   int `json:"number"`
	Elements []T `json:"elements"`
}

func (p Page[T]) New(data []T, limit int, offset int) Page[T] {
	fmt.Println("data", len(data))

	p.Number = 1
	if len(data) <= offset {
		p.Number = (len(data) - offset) / limit
	}

	if offset >= len(data) {
		p.Elements = make([]T, 0)
	}

	if offset+limit > len(data) && offset < len(data) {
		p.Elements = data[offset:]
	}
	if limit+offset <= len(data) {
		p.Elements = data[offset : offset+limit]
	}

	if len(data) == 0 {
		p.Number = (len(data) - offset) / limit
	}

	fmt.Println("limit", limit)
	fmt.Println("offset", offset)

	return p
}
