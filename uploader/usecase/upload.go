package usecase

import (
	"context"
	"math/rand"
	"time"

	"github.com/interrupted-network/fake-uploader/uploader/domain/uploader"
)

func (uc *useCase) Upload(ctx context.Context,
	request *uploader.Request) (*uploader.Result, error) {
	var client *client
	for {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		client = uc.clients[rand.Intn(len(uc.clients))]
		if client.isConnected {
			break
		}
	}
	request.Client = client
	return uc.upload(ctx, request)
}

func (uc *useCase) upload(ctx context.Context,
	request *uploader.Request) (*uploader.Result, error) {
	bytes := make([]byte, request.Size)
	result := new(uploader.Result)
	if request.Deadline == 0 {
		request.Deadline = time.Second * 10
	}
	err := request.Client.SetWriteDeadline(time.Now().Add(request.Deadline))
	if err != nil {
		return nil, err
	}
	l, err := request.Client.Write(bytes)
	if err != nil {
		return result, err
	}
	result.SentLen = int64(l)
	// if l > 0 {
	// 	uc.logger.Debugf("sent %s", byteCountIEC(result.SentLen))
	// }
	return result, nil
}
