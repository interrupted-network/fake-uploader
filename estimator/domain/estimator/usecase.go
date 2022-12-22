package estimator

type UseCase interface {
	GetRealtimeSpeedChan() <-chan *Speed
	Start()
	Estimate() (*Result, error)
}
