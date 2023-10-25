package test

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ VersionTestModelModel = (*customVersionTestModelModel)(nil)

type (
	// VersionTestModelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVersionTestModelModel.
	VersionTestModelModel interface {
		vTestModel
		withSession(session sqlx.Session) VersionTestModelModel
	}

	customVersionTestModelModel struct {
		*defaultVersionTestModelModel
	}
)

// NewVersionTestModelModel returns a model for the database table.
func NewVersionTestModelModel(conn sqlx.SqlConn) VersionTestModelModel {
	return &customVersionTestModelModel{
		defaultVersionTestModelModel: newVersionTestModelModel(conn),
	}
}

func (m *customVersionTestModelModel) withSession(session sqlx.Session) VersionTestModelModel {
	return NewVersionTestModelModel(sqlx.NewSqlConnFromSession(session))
}
