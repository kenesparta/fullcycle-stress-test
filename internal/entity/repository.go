package entity

type StatsRepo interface {
	Get(rf RequestFlag, cm ConcurrencyMgmt)
}
