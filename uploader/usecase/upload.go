package usecase

import (
	"context"
	"math/rand"
	"net"

	"github.com/interrupted-network/fake-uploader/uploader/domain/uploader"
)

func (uc *useCase) UploadRandomTarget(ctx context.Context,
	size uint) (*uploader.Result, error) {
	target := uc.config.Targets[rand.Intn(len(uc.config.Targets))]
	request := &uploader.Request{
		Network: target.Network,
		Address: target.Address,
		Size:    size,
	}
	return uc.Upload(ctx, request)
}

func (uc *useCase) Upload(ctx context.Context,
	request *uploader.Request) (*uploader.Result, error) {
	uc.logger.DebugF("connecting %s(%s)...",
		request.Address, request.Network)
	conn, err := net.Dial(request.Network, request.Address)
	if err != nil {
		return nil, err
	}
	uc.logger.DebugF("%s: connected", request.Address)

	bytes := make([]byte, 1024*1024)
	result := new(uploader.Result)
	for i := 0; i < int(request.Size); i += len(bytes) {
		_, err = conn.Write(bytes)
		if err != nil {
			return result, err
		}
		result.SentLen += int64(len(bytes))
	}
	_ = conn.Close()
	return result, nil
}
