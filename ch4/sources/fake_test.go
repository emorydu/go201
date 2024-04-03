package main

import "testing"

type Result struct {
	Count int
}

func (r Result) Int() int { return r.Count }

type Rows []struct{}

type Stmt interface {
	Close() error
	NumInput() int
	Exec(stmt string, args ...any) (Result, error)
	Query(args []string) (Rows, error)
}

func MaleCount(s Stmt) (int, error) {
	result, err := s.Exec("SELECT COUNT(*) FROM employee_tab WHERE gender=?", 5)
	if err != nil {
		return 0, err
	}
	return result.Int(), nil
}

type fakeStmtForMaleCount struct {
	Stmt
}

func TestEmployeeMaleCount(t *testing.T) {
	f := fakeStmtForMaleCount{}
	c, _ := MaleCount(f)
	if c != 5 {
		t.Errorf("want: %d, actual: %d", 5, c)
		return
	}
}
