package util

import "strings"

type SqlCond struct {
	Cond *strings.Builder
	Args []any
}

func SqlCondNew(bufferSize int) *SqlCond {
	cond := new(strings.Builder)
	cond.Grow(bufferSize)

	return &SqlCond{
		Cond: cond,
		Args: nil,
	}
}

func (s *SqlCond) Get() (string, []any) {
	return s.Cond.String(), s.Args
}

func (s *SqlCond) Eq(key string, val any) *SqlCond {
	s.Cond.WriteString(key)
	s.Cond.WriteString(" = ? ")
	s.Args = append(s.Args, val)
	return s
}

func (s *SqlCond) NotEq(key string, val any) *SqlCond {
	s.Cond.WriteString(key)
	s.Cond.WriteString(" != ? ")
	s.Args = append(s.Args, val)
	return s
}

func (s *SqlCond) inWrp(logic, key string, arr []any) *SqlCond {
	if len(arr) <= 0 {
		return s
	}

	var tmp strings.Builder
	tmp.Grow(len(logic) + len(key) + (len(arr) * 3) + 2)

	tmp.WriteString(key)
	tmp.WriteString(logic)
	tmp.WriteString("(")
	for range arr {
		tmp.WriteString("?, ")
	}

	tmpStr := tmp.String()
	tmpStr = tmpStr[0 : len(tmpStr)-2]

	s.Cond.WriteString(tmpStr)
	s.Cond.WriteString(") ")
	s.Args = append(s.Args, arr...)
	return s
}

func (s *SqlCond) In(key string, arr []any) *SqlCond {
	return s.inWrp(" IN ", key, arr)
}

func (s *SqlCond) NotIn(key string, arr []any) *SqlCond {
	return s.inWrp(" NOT IN ", key, arr)
}

func (s *SqlCond) And() *SqlCond {
	s.Cond.WriteString(" AND ")
	return s
}

func (s *SqlCond) Or() *SqlCond {
	s.Cond.WriteString(" OR ")
	return s
}

func (s *SqlCond) Begin() *SqlCond {
	s.Cond.WriteString(" (")
	return s
}

func (s *SqlCond) End() *SqlCond {
	s.Cond.WriteString(") ")
	return s
}
