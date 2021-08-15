package postgres

import "github.com/go-ushio/ushio/model/category"

var (
	SQLGetCategories = `SELECT tid, tname, name, color, admin FROM categories;`
)

func (pg *Postgres) GetCategories() ([]*category.Category, error) {
	rows, err := pg.db.Query(SQLGetCategories)
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
