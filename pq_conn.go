package pq

import (
	"database/sql/driver"
)

type PQConn struct {
	cn *conn
}

func newPQConn(cn *conn) *PQConn {
	return &PQConn{cn: cn}
}

func OpenPQConn(name string) (*PQConn, error) {
	cn, err := Open(name)
	if err == nil {
		return newPQConn(cn.(*conn)), nil
	} else {
		return nil, err
	}
}

func (pcn *PQConn) Close() error {
	return pcn.Close()
}

func (pcn *PQConn) Begin() (*PQTx, error) {
	_, err := pcn.cn.Begin()
	if err == nil {
		return newPQTx(pcn), nil
	} else {
		return nil, err
	}
}

func (pcn *PQConn) Prepare(query string) (*PQStmt, error) {
	st, err := pcn.cn.Prepare(query)
	if err == nil {
		return newPQStmt(st.(*stmt)), nil
	} else {
		return nil, err
	}
}

func (pcn *PQConn) Exec(query string, args ...interface{}) (PQResult, error) {
	values := make([]driver.Value, len(args))
	for i, arg := range args {
		values[i] = arg.(driver.Value)	
	}
	
	res, err := pcn.cn.Exec(query, values)
	if err != nil {
		return res.(PQResult), nil
	} else {
		return PQResult(0), err
	}
}