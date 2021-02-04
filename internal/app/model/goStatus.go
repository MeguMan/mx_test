package model

type GoroutineStatus struct {
	Key string
	Finished bool
	Error string
	RowsStats *RowsStats
}
