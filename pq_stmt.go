package pq

import (
	"database/sql/driver"
)

type PQStmt struct {
	st *stmt
}

type PQResult driver.RowsAffected

func newPQStmt(st *stmt) *PQStmt {
	return &PQStmt{st: st}
}

func (pqr PQResult) LastInsertId() (int64, error) {
	return driver.RowsAffected(pqr).LastInsertId()
}

func (pqr PQResult) RowsAffected() (int64, error) {
	return driver.RowsAffected(pqr).RowsAffected()
}
