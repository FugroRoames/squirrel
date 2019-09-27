package squirrel

type unionBuilder struct {
	SelectBuilders []SelectBuilder
}

// Union - Return the SQL UNION of a series of select queries.
func Union(b ...SelectBuilder) unionBuilder {
	return unionBuilder{SelectBuilders: b}
}

func (u unionBuilder) ToSql() (sqlStr string, args []interface{}, err error) {
	n := len(u.SelectBuilders)
	argsOffset := 0
	for i, b := range u.SelectBuilders {
		s, a, e := b.ToSqlWithArgsOffset(argsOffset)

		sqlStr += s
		if i != n-1 {
			sqlStr += " UNION "
		}
		args = append(args, a...)
		argsOffset += len(args)

		if e != nil {
			return "", []interface{}{}, e
		}
	}

	return sqlStr, args, err
}
