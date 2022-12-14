package estimator

type UseCase interface {
	Estimate() (*Result, error)
}
