package coordinator

type UseCase interface {
	GetMessageQueue() chan []byte

	Process() error
	Start()
}
