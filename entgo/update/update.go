package update

import (
	"entgo.io/ent/dialect/sql"
	"github.com/devexps/go-utils/stringcase"
)

func BuildSetNullUpdate(u *sql.UpdateBuilder, fields []string) {
	if len(fields) > 0 {
		for _, field := range fields {
			field = stringcase.ToSnakeCase(field)
			u.SetNull(field)
		}
	}
}

func BuildSetNullUpdater(fields []string) (func(u *sql.UpdateBuilder), error) {
	if len(fields) > 0 {
		return func(u *sql.UpdateBuilder) {
			BuildSetNullUpdate(u, fields)
		}, nil
	} else {
		return nil, nil
	}
}
