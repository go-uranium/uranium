package category

import (
	"github.com/lib/pq"

	"github.com/go-ushio/ushio/common/scan"
)

type Category struct {
	TID   int
	TName string
	Name  string
	Color string
	Admin []int
}

func ScanCategory(scanner scan.Scanner) (*Category, error) {
	cate := &Category{}
	err := scanner.Scan(&cate.TID, &cate.TName, &cate.Name,
		&cate.Color, pq.Array(&cate.Admin))
	if err != nil {
		return &Category{}, err
	}
	return cate, nil
}
