package category

import "github.com/go-ushio/ushio/common/scan"

type Category struct {
	TID   int
	TName string
	Name  string
	Color string
}

func ScanCategory(scanner scan.Scanner) (*Category, error) {
	cate := &Category{}
	err := scanner.Scan(&cate.TID, &cate.TName, &cate.Name, &cate.Color)
	if err != nil {
		return &Category{}, err
	}
	return cate, nil
}
