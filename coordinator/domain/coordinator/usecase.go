package coordinator

type UseCase interface {
	GetMessageQueue() chan []byte

	Start()
}
