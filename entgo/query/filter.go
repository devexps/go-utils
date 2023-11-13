package query

import (
	"encoding/json"
	"strings"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	"github.com/devexps/go-micro/v2/encoding"

	"github.com/devexps/go-utils/stringcase"
)

type FilterOp int

const (
	FilterNot                   = "not"         // Not equal to
	FilterIn                    = "in"          // Check if value is in list
	FilterNotIn                 = "not_in"      // Not in list
	FilterGTE                   = "gte"         // Greater than or equal to the passed value
	FilterGT                    = "gt"          // Greater than passed value
	FilterLTE                   = "lte"         // Less than or equal to the passed value
	FilterLT                    = "lt"          // Less than passed value
	FilterRange                 = "range"       // Whether it is between the two given values
	FilterIsNull                = "isnull"      // Is it empty
	FilterNotIsNull             = "not_isnull"  // Whether it is not empty
	FilterContains              = "contains"    // Whether to contain the specified substring
	FilterInsensitiveContains   = "icontains"   // Case-insensitive, whether to contain the specified substring
	FilterStartsWith            = "startswith"  // Start with value
	FilterInsensitiveStartsWith = "istartswith" // Case-insensitive, starts with value
	FilterEndsWith              = "endswith"    // End with value
	FilterInsensitiveEndsWith   = "iendswith"   // Case-insensitive, ends with value
	FilterExact                 = "exact"       // Exact match
	FilterInsensitiveExact      = "iexact"      // Case-insensitive, exact match
	FilterRegex                 = "regex"       // Regular expression
	FilterInsensitiveRegex      = "iregex"      // Case-insensitive, regular expression
	FilterSearch                = "search"      // Research all
)

type DatePart int

const (
	DatePartDate        DatePart = iota // Date
	DatePartYear                        // Year
	DatePartISOYear                     // ISO 8601 - Number of weeks in year
	DatePartQuarter                     // Quarter
	DatePartMonth                       // Month
	DatePartWeek                        // ISO 8601 - week number, number of week in year
	DatePartWeekDay                     // Day of the week
	DatePartISOWeekDay                  // Day of the week (ISO)
	DatePartDay                         // Day
	DatePartTime                        // Hours: Minutes: Seconds
	DatePartHour                        // Hour
	DatePartMinute                      // Minute
	DatePartSecond                      // Second
	DatePartMicrosecond                 // Microseconds
)

var dateParts = [...]string{
	DatePartDate:        "date",
	DatePartYear:        "year",
	DatePartISOYear:     "iso_year",
	DatePartQuarter:     "quarter",
	DatePartMonth:       "month",
	DatePartWeek:        "week",
	DatePartWeekDay:     "week_day",
	DatePartISOWeekDay:  "iso_week_day",
	DatePartDay:         "day",
	DatePartTime:        "time",
	DatePartHour:        "hour",
	DatePartMinute:      "minute",
	DatePartSecond:      "second",
	DatePartMicrosecond: "microsecond",
}

func hasDatePart(str string) bool {
	for _, item := range dateParts {
		if str == item {
			return true
		}
	}
	return false
}

// QueryCommandToWhereConditions converts query commands into selection conditions
func QueryCommandToWhereConditions(strJson string, isOr bool) (error, func(s *sql.Selector)) {
	if len(strJson) == 0 {
		return nil, nil
	}

	codec := encoding.GetCodec("json")

	queryMap := make(map[string]string)
	var queryMapArray []map[string]string
	if err1 := codec.Unmarshal([]byte(strJson), &queryMap); err1 != nil {
		if err2 := codec.Unmarshal([]byte(strJson), &queryMapArray); err2 != nil {
			return err2, nil
		}
	}
	return nil, func(s *sql.Selector) {
		var ps []*sql.Predicate
		ps = append(ps, processQueryMap(s, queryMap)...)
		for _, v := range queryMapArray {
			ps = append(ps, processQueryMap(s, v)...)
		}
		if isOr {
			s.Where(sql.Or(ps...))
		} else {
			s.Where(sql.And(ps...))
		}
	}
}

func processQueryMap(s *sql.Selector, queryMap map[string]string) []*sql.Predicate {
	var ps []*sql.Predicate
	for k, v := range queryMap {
		key := stringcase.ToSnakeCase(k)

		keys := strings.Split(key, "__")

		if cond := oneFieldFilter(s, keys, v); cond != nil {
			ps = append(ps, cond)
		}
	}
	return ps
}

func BuildFilterSelector(andFilterJsonString, orFilterJsonString string) (error, []func(s *sql.Selector)) {
	var err error
	var queryConditions []func(s *sql.Selector)

	var andSelector func(s *sql.Selector)
	err, andSelector = QueryCommandToWhereConditions(andFilterJsonString, false)
	if err != nil {
		return err, nil
	}
	if andSelector != nil {
		queryConditions = append(queryConditions, andSelector)
	}
	var orSelector func(s *sql.Selector)
	err, orSelector = QueryCommandToWhereConditions(orFilterJsonString, true)
	if err != nil {
		return err, nil
	}
	if orSelector != nil {
		queryConditions = append(queryConditions, orSelector)
	}
	return nil, queryConditions
}

func oneFieldFilter(s *sql.Selector, keys []string, value string) *sql.Predicate {
	var cond *sql.Predicate

	if len(keys) == 1 {
		field := keys[0]
		cond = filterEqual(s, field, value)
	} else if len(keys) == 2 {
		if len(keys[0]) == 0 {
			return nil
		}
		field := keys[0]
		op := strings.ToLower(keys[1])
		switch op {
		case FilterNot:
			cond = filterNot(s, field, value)
		case FilterIn:
			cond = filterIn(s, field, value)
		case FilterNotIn:
			cond = filterNotIn(s, field, value)
		case FilterGTE:
			cond = filterGTE(s, field, value)
		case FilterGT:
			cond = filterGT(s, field, value)
		case FilterLTE:
			cond = filterLTE(s, field, value)
		case FilterLT:
			cond = filterLT(s, field, value)
		case FilterRange:
			cond = filterRange(s, field, value)
		case FilterIsNull:
			cond = filterIsNull(s, field, value)
		case FilterNotIsNull:
			cond = filterIsNotNull(s, field, value)
		case FilterContains:
			cond = filterContains(s, field, value)
		case FilterInsensitiveContains:
			cond = filterInsensitiveContains(s, field, value)
		case FilterStartsWith:
			cond = filterStartsWith(s, field, value)
		case FilterInsensitiveStartsWith:
			cond = filterInsensitiveStartsWith(s, field, value)
		case FilterEndsWith:
			cond = filterEndsWith(s, field, value)
		case FilterInsensitiveEndsWith:
			cond = filterInsensitiveEndsWith(s, field, value)
		case FilterExact:
			cond = filterExact(s, field, value)
		case FilterInsensitiveExact:
			cond = filterInsensitiveExact(s, field, value)
		case FilterRegex:
			cond = filterRegex(s, field, value)
		case FilterInsensitiveRegex:
			cond = filterInsensitiveRegex(s, field, value)
		case FilterSearch:
			cond = filterSearch(s, field, value)
		default:
			cond = filterDatePart(s, op, field, value)
		}
	}
	return cond
}

// filterEqual = equality operation
// SQL: WHERE "name" = "tom"
func filterEqual(s *sql.Selector, field, value string) *sql.Predicate {
	return sql.EQ(s.C(field), value)
}

// filterNot NOT Inequality operation
// SQL: WHERE NOT ("name" = "tom")
// or： WHERE "name" <> "tom"
// You can use NOT to filter out NULL, but you can't use <> and !=.
func filterNot(s *sql.Selector, field, value string) *sql.Predicate {
	return sql.Not(sql.EQ(s.C(field), value))
}

// filterIn IN operation
// SQL: WHERE name IN ("tom", "jimmy")
func filterIn(s *sql.Selector, field, value string) *sql.Predicate {
	var values []any
	if err := json.Unmarshal([]byte(value), &values); err == nil {
		return sql.In(s.C(field), values...)
	}
	return nil
}

// filterNotIn NOT IN operation
// SQL: WHERE name NOT IN ("tom", "jimmy")`
func filterNotIn(s *sql.Selector, field, value string) *sql.Predicate {
	var values []any
	if err := json.Unmarshal([]byte(value), &values); err == nil {
		return sql.NotIn(s.C(field), values...)
	}
	return nil
}

// filterGTE GTE (Greater Than or Equal) Greater than or equal to >= operation
// SQL: WHERE "create_time" >= "2023-10-25"
func filterGTE(s *sql.Selector, field, value string) *sql.Predicate {
	return sql.GTE(s.C(field), value)
}

// filterGT GT (Greater than) greater than >operation
// SQL: WHERE "create_time" > "2023-10-25"
func filterGT(s *sql.Selector, field, value string) *sql.Predicate {
	return sql.GT(s.C(field), value)
}

// filterLTE LTE (Less Than or Equal) less than or equal to <= operation
// SQL: WHERE "create_time" <= "2023-10-25"
func filterLTE(s *sql.Selector, field, value string) *sql.Predicate {
	return sql.LTE(s.C(field), value)
}

// filterLT LT (Less than) less than < operation
// SQL: WHERE "create_time" < "2023-10-25"
func filterLT(s *sql.Selector, field, value string) *sql.Predicate {
	return sql.LT(s.C(field), value)
}

// filterRange BETWEEN operation in the range of values
// SQL: WHERE "create_time" BETWEEN "2023-10-25" AND "2024-10-25"
// or： WHERE "create_time" >= "2023-10-25" AND "create_time" <= "2024-10-25"
func filterRange(s *sql.Selector, field, value string) *sql.Predicate {
	var values []any
	if err := json.Unmarshal([]byte(value), &values); err == nil {
		if len(values) != 2 {
			return nil
		}
		return sql.And(
			sql.GTE(s.C(field), values[0]),
			sql.LTE(s.C(field), values[1]),
		)
	}
	return nil
}

// filterIsNull IS NULL operation
// SQL: WHERE name IS NULL
func filterIsNull(s *sql.Selector, field, _ string) *sql.Predicate {
	return sql.IsNull(s.C(field))
}

// filterIsNotNull IS NOT NULL operation
// SQL: WHERE name IS NOT NULL
func filterIsNotNull(s *sql.Selector, field, _ string) *sql.Predicate {
	return sql.Not(sql.IsNull(s.C(field)))
}

// filterContains LIKE fuzzy query before and after
// SQL: WHERE name LIKE '%L%';
func filterContains(s *sql.Selector, field, value string) *sql.Predicate {
	return sql.Contains(s.C(field), value)
}

// filterInsensitiveContains ILIKE fuzzy query before and after
// SQL: WHERE name ILIKE '%L%';
func filterInsensitiveContains(s *sql.Selector, field, value string) *sql.Predicate {
	return sql.ContainsFold(s.C(field), value)
}

// filterStartsWith LIKE Prefix + fuzzy query
// SQL: WHERE name LIKE 'La%';
func filterStartsWith(s *sql.Selector, field, value string) *sql.Predicate {
	return sql.HasPrefix(s.C(field), value)
}

// filterInsensitiveStartsWith ILIKE prefix + fuzzy query
// SQL: WHERE name ILIKE 'La%';
func filterInsensitiveStartsWith(s *sql.Selector, field, value string) *sql.Predicate {
	return sql.EqualFold(s.C(field), value+"%")
}

// filterEndsWith LIKE suffix + fuzzy query
// SQL: WHERE name LIKE '%a';
func filterEndsWith(s *sql.Selector, field, value string) *sql.Predicate {
	return sql.HasSuffix(s.C(field), value)
}

// filterInsensitiveEndsWith ILIKE suffix + fuzzy query
// SQL: WHERE name ILIKE '%a';
func filterInsensitiveEndsWith(s *sql.Selector, field, value string) *sql.Predicate {
	return sql.EqualFold(s.C(field), "%"+value)
}

// filterExact LIKE operation precise comparison
// SQL: WHERE name LIKE 'a';
func filterExact(s *sql.Selector, field, value string) *sql.Predicate {
	return sql.Like(s.C(field), value)
}

// filterInsensitiveExact ILIKE operation is case-insensitive and exact comparison
// SQL: WHERE name ILIKE 'a';
func filterInsensitiveExact(s *sql.Selector, field, value string) *sql.Predicate {
	return sql.EqualFold(s.C(field), value)
}

// filterRegex regular search
// MySQL: WHERE title REGEXP BINARY '^(An?|The) +'
// Oracle: WHERE REGEXP_LIKE(title, '^(An?|The) +', 'c');
// PostgreSQL: WHERE title ~ '^(An?|The) +';
// SQLite: WHERE title REGEXP '^(An?|The) +';
func filterRegex(s *sql.Selector, field, value string) *sql.Predicate {
	p := sql.P()
	p.Append(func(b *sql.Builder) {
		switch s.Builder.Dialect() {
		case dialect.Postgres:
			b.Ident(s.C(field)).WriteString(" ~ ")
			b.Arg(value)
			break
		case dialect.MySQL:
			b.Ident(s.C(field)).WriteString(" REGEXP BINARY ")
			b.Arg(value)
			break
		case dialect.SQLite:
			b.Ident(s.C(field)).WriteString(" REGEXP ")
			b.Arg(value)
			break
		case dialect.Gremlin:
			break
		}
	})
	return p
}

// filterInsensitiveRegex regular search is not case sensitive
// MySQL: WHERE title REGEXP '^(an?|the) +'
// Oracle: WHERE REGEXP_LIKE(title, '^(an?|the) +', 'i');
// PostgreSQL: WHERE title ~* '^(an?|the) +';
// SQLite: WHERE title REGEXP '(?i)^(an?|the) +';
func filterInsensitiveRegex(s *sql.Selector, field, value string) *sql.Predicate {
	p := sql.P()
	p.Append(func(b *sql.Builder) {
		switch s.Builder.Dialect() {
		case dialect.Postgres:
			b.Ident(s.C(field)).WriteString(" ~* ")
			b.Arg(strings.ToLower(value))
			break
		case dialect.MySQL:
			b.Ident(s.C(field)).WriteString(" REGEXP ")
			b.Arg(strings.ToLower(value))
			break
		case dialect.SQLite:
			b.Ident(s.C(field)).WriteString(" REGEXP ")
			if !strings.HasPrefix(value, "(?i)") {
				value = "(?i)" + value
			}
			b.Arg(strings.ToLower(value))
			break
		case dialect.Gremlin:
			break
		}
	})
	return p
}

// filterSearch research all
// SQL:
func filterSearch(s *sql.Selector, _, _ string) *sql.Predicate {
	p := sql.P()
	p.Append(func(b *sql.Builder) {
		switch s.Builder.Dialect() {

		}
	})
	return nil
}

// filterDatePart timestamp extraction date
// SQL: select extract(quarter from timestamp '2018-08-15 12:10:10');
func filterDatePart(s *sql.Selector, datePart, field, value string) *sql.Predicate {
	return nil
}
