package usecase

import (
	"context"
	"math/rand"
	"net"
	"time"

	"github.com/interrupted-network/fake-uploader/uploader/domain/uploader"
)

func (uc *useCase) Upload(ctx context.Context,
	size uint) (*uploader.Result, error) {
	target := uc.config.Targets[rand.Intn(len(uc.config.Targets))]
	request := &uploader.Request{
		Target: target,
		Size:   size,
	}
	return uc.upload(ctx, request)
}

func (uc *useCase) upload(ctx context.Context,
	request *uploader.Request) (*uploader.Result, error) {
	uc.logger.DebugF("connecting %s(%s)...",
		request.Target.Address, request.Target.Network)
	conn, err := net.DialTimeout(
		request.Target.Network,
		request.Target.Address,
		request.Target.DialTimeout,
	)
	if err != nil {
		return nil, err
	}
	uc.logger.DebugF("%s: connected", request.Target.Address)

	conn.SetDeadline(time.Now().Add(request.Target.RWTimeout))

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
