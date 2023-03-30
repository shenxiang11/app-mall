package pagination

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	DefaultPageSize = 10
	MaxPageSize     = 100
	PageVar         = "page"
	PageSizeVar     = "pageSize"
)

type Pagination struct {
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
	TotalCount int `json:"totalCount"`
	List       any `json:"list"`
}

func New(page, pageSize, total int) *Pagination {
	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}
	if page <= 0 {
		page = 1
	}

	return &Pagination{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: total,
	}
}

func NewFromGinRequest(g *gin.Context, count int) *Pagination {
	page := parseInt(g.Query(PageVar), 1)
	pageSize := parseInt(g.Query(PageSizeVar), DefaultPageSize)
	return New(page, pageSize, count)
}

func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}
