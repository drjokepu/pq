package pq

type PQTx struct {
	pcn *PQConn
}

func newPQTx(pcn *PQConn) *PQTx {
	return &PQTx{pcn: pcn}
}

func (ptx *PQTx) Commit() error {
	return ptx.pcn.cn.Commit()
}

func (ptx *PQTx) Rollback() error {
	return ptx.pcn.cn.Rollback()
}
