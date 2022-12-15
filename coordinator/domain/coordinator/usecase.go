package coordinator

type UseCase interface {
	Process() error
	Start()
}
