package query

import (
	"entgo.io/ent/dialect/sql"

	_ "github.com/devexps/go-micro/v2/encoding/json"
)

// BuildQuerySelector builds a paging filter query
func BuildQuerySelector(
	andFilterJsonString, orFilterJsonString string,
	page, pageSize int32, noPaging bool,
	orderBys []string, defaultOrderField string,
	selectFields []string,
) (err error, whereSelectors []func(s *sql.Selector), querySelectors []func(s *sql.Selector)) {
	err, whereSelectors = BuildFilterSelector(andFilterJsonString, orFilterJsonString)
	if err != nil {
		return err, nil, nil
	}
	var orderSelector func(s *sql.Selector)
	err, orderSelector = BuildOrderSelector(orderBys, defaultOrderField)
	if err != nil {
		return err, nil, nil
	}
	pageSelector := BuildPaginationSelector(page, pageSize, noPaging)

	var fieldSelector func(s *sql.Selector)
	err, fieldSelector = BuildFieldSelector(selectFields)

	if len(whereSelectors) > 0 {
		querySelectors = append(querySelectors, whereSelectors...)
	}
	if orderSelector != nil {
		querySelectors = append(querySelectors, orderSelector)
	}
	if pageSelector != nil {
		querySelectors = append(querySelectors, pageSelector)
	}
	if fieldSelector != nil {
		querySelectors = append(querySelectors, fieldSelector)
	}
	return
}
