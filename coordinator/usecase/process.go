package usecase

import (
	"context"
	"math/rand"
	"sync"

	"github.com/interrupted-network/fake-uploader/uploader/domain/uploader"
)

func (uc *useCase) Process() error {
	estimated, err := uc.estimator.Estimate()
	if err != nil {
		return err
	}

	uc.logger.DebugF("estimated. Rx: %s, Tx: %s, Ratio: %f",
		byteCountIEC(estimated.Rx), byteCountIEC(estimated.Tx), estimated.RxTxRatio)

	var totalSent int64 = 0
	if estimated.RxTxRatio > uc.config.RxTxMaxRatio {
		var wg sync.WaitGroup
		for i := 0; i < int(uc.config.Concurrent); i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				var err error

				var result *uploader.Result
				result, err = uc.upload()
				if err != nil {
					uc.logger.DebugF("error while uploading. error: %v", err)
				}
				if result != nil {
					totalSent += result.SentLen
				}
			}()
		}
		wg.Wait()
		uc.logger.DebugF("%s sent", byteCountIEC(totalSent))
	}
	return nil
}

func (uc *useCase) upload() (*uploader.Result, error) {
	size := uint(uc.config.UploadSize.Min) +
		uint(rand.Intn(int(uc.config.UploadSize.Max)))
	return uc.uploader.UploadRandomTarget(context.Background(), uint(size))
}
