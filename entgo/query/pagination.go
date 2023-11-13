package query

import (
	"entgo.io/ent/dialect/sql"

	paging "github.com/devexps/go-utils/pagination"
)

// BuildPaginationSelector .
func BuildPaginationSelector(page, pageSize int32, noPaging bool) func(*sql.Selector) {
	if noPaging {
		return nil
	} else {
		return func(s *sql.Selector) {
			BuildPaginationSelect(s, page, pageSize)
		}
	}
}

// BuildPaginationSelect .
func BuildPaginationSelect(s *sql.Selector, page, pageSize int32) {
	if page < 1 {
		page = paging.DefaultPage
	}
	if pageSize < 1 {
		pageSize = paging.DefaultPageSize
	}
	offset := paging.GetPageOffset(page, pageSize)
	s.Offset(offset).Limit(int(pageSize))
}
