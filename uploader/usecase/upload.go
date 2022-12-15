package usecase

import (
	"context"
	"math/rand"
	"net"

	"github.com/interrupted-network/fake-uploader/uploader/domain/uploader"
)

func (uc *useCase) UploadRandomTarget(ctx context.Context, size uint) error {
	target := uc.config.Targets[rand.Intn(len(uc.config.Targets))]
	request := &uploader.Request{
		Network: target.Network,
		Address: target.Address,
		Size:    size,
	}
	return uc.Upload(ctx, request)
}

func (uc *useCase) Upload(ctx context.Context, request *uploader.Request) error {
	conn, err := net.Dial("tcp", request.Address)
	if err != nil {
		return err
	}

	bytes := make([]byte, 10000)
	for i := 0; i < int(request.Size); i += len(bytes) {
		_, err = conn.Write(bytes)
		if err != nil {
			return err
		}
	}
	return nil
}
