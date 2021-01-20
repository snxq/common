package gorm

import (
	"context"

	"gorm.io/gorm"
)

// Filter 数据库查询条件 Where
type Filter struct {
	Query string
	Args  []interface{}
}

// Query 数据库查询请求
type Query struct {
	Ctx context.Context

	DB            *gorm.DB
	Filters       []Filter
	Count         *int64
	Order         string
	Limit, Offset int
}

// QueryDB 数据库查询
// Limit == -1 时不分页
func (qr *Query) QueryDB(result interface{}) error {
	q := qr.DB.WithContext(qr.Ctx).Model(result)
	for _, query := range qr.Filters {
		q = q.Where(query.Query, query.Args...)
	}
	if qr.Order != "" {
		q = q.Order(qr.Order)
	}
	if qr.Count != nil {
		q = q.Count(qr.Count)
	}
	if qr.Limit == 0 {
		qr.Limit = -1
	}

	return q.Offset(qr.Offset).Limit(qr.Limit).Find(result).Error
}
