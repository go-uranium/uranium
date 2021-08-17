package postgres

import (
	"github.com/go-uranium/uranium/model/category"
)

var (
	SQLGetCategories = `SELECT tid, tname, name, color FROM categories;`
)

func (pg *Postgres) GetCategories() ([]*category.Category, error) {
	rows, err := pg.db.Query(SQLGetCategories)
	if err != nil {
		return nil, err
	}
	var cates []*category.Category
	for rows.Next() {
		cate := &category.Category{}
		err := rows.Scan(&cate.TID, &cate.TName, &cate.Name, &cate.Color)
		if err != nil {
			return nil, err
		}
		cates = append(cates, cate)
	}
	return cates, nil
}
