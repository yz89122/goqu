// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import exp "github.com/yz89122/goqu/v10/exp"

import mock "github.com/stretchr/testify/mock"
import sb "github.com/yz89122/goqu/v10/internal/sb"

// SQLDialect is an autogenerated mock type for the SQLDialect type
type SQLDialect struct {
	mock.Mock
}

// Dialect provides a mock function with given fields:
func (_m *SQLDialect) Dialect() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ToDeleteSQL provides a mock function with given fields: b, clauses
func (_m *SQLDialect) ToDeleteSQL(b sb.SQLBuilder, clauses exp.DeleteClauses) {
	_m.Called(b, clauses)
}

// ToInsertSQL provides a mock function with given fields: b, clauses
func (_m *SQLDialect) ToInsertSQL(b sb.SQLBuilder, clauses exp.InsertClauses) {
	_m.Called(b, clauses)
}

// ToSelectSQL provides a mock function with given fields: b, clauses
func (_m *SQLDialect) ToSelectSQL(b sb.SQLBuilder, clauses exp.SelectClauses) {
	_m.Called(b, clauses)
}

// ToTruncateSQL provides a mock function with given fields: b, clauses
func (_m *SQLDialect) ToTruncateSQL(b sb.SQLBuilder, clauses exp.TruncateClauses) {
	_m.Called(b, clauses)
}

// ToUpdateSQL provides a mock function with given fields: b, clauses
func (_m *SQLDialect) ToUpdateSQL(b sb.SQLBuilder, clauses exp.UpdateClauses) {
	_m.Called(b, clauses)
}
