package data

import "github.com/go-ushio/ushio/core/category"

func (data *Data) GetCategories() ([]*category.Category, error) {
	rows, err := data.db.Query(data.sentence.SQLGetCategories)
	if err != nil {
		return nil, err
	}
	var cates []*category.Category
	for rows.Next() {
		cate, err := category.ScanCategory(rows)
		if err != nil {
			return nil, err
		}
		cates = append(cates, cate)
	}
	return cates, nil
}
