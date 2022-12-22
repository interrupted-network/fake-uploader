package uploader

import "context"

type UseCase interface {
	Initialize(msgQueue <-chan []byte)
	Start()

	Upload(ctx context.Context, request *Request) (*Result, error)
}
