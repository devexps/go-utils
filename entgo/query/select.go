package query

import (
	"entgo.io/ent/dialect/sql"
	"github.com/devexps/go-utils/stringcase"
)

func BuildFieldSelect(s *sql.Selector, fields []string) {
	if len(fields) > 0 {
		for i, field := range fields {
			switch {
			case field == "id_" || field == "_id":
				field = "id"
			}
			fields[i] = stringcase.ToSnakeCase(field)
		}
		s.Select(fields...)
	}
}

func BuildFieldSelector(fields []string) (error, func(s *sql.Selector)) {
	if len(fields) > 0 {
		return nil, func(s *sql.Selector) {
			BuildFieldSelect(s, fields)
		}
	} else {
		return nil, nil
	}
}
