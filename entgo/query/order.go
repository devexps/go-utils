package query

import (
	"strings"

	"entgo.io/ent/dialect/sql"
)

// QueryCommandToOrderConditions query commands converted into sorting conditions
func QueryCommandToOrderConditions(orderBys []string) (error, func(s *sql.Selector)) {
	if len(orderBys) == 0 {
		return nil, nil
	}
	return nil, func(s *sql.Selector) {
		for _, v := range orderBys {
			if strings.HasPrefix(v, "-") {
				key := v[1:]
				if len(key) == 0 {
					continue
				}
				BuildOrderSelect(s, key, true)
			} else {
				if len(v) == 0 {
					continue
				}
				BuildOrderSelect(s, v, false)
			}
		}
	}
}

// BuildOrderSelect .
func BuildOrderSelect(s *sql.Selector, field string, desc bool) {
	if desc {
		s.OrderBy(sql.Desc(s.C(field)))
	} else {
		s.OrderBy(sql.Asc(s.C(field)))
	}
}

// BuildOrderSelector .
func BuildOrderSelector(orderBys []string, defaultOrderField string) (error, func(s *sql.Selector)) {
	if len(orderBys) == 0 {
		return nil, func(s *sql.Selector) {
			BuildOrderSelect(s, defaultOrderField, true)
		}
	} else {
		return QueryCommandToOrderConditions(orderBys)
	}
}
