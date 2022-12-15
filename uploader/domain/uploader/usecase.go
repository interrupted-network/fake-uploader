package uploader

import "context"

type UseCase interface {
	Upload(ctx context.Context, request *Request) error
}
