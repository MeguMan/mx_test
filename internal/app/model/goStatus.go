package model

type GoroutineStatus struct {
	Id string
	Finished bool
	RowsStats *RowsStats
}
