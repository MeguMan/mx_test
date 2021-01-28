package model

type GoroutineStatus struct {
	Key string
	Finished bool
	RowsStats *RowsStats
}
