package uploader

import "context"

type UseCase interface {
	UploadRandomTarget(ctx context.Context, size uint) (*Result, error)
	Upload(ctx context.Context, request *Request) (*Result, error)
}
