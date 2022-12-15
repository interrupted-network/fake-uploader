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
					uc.logger.DebugF("%s: %s sent",
						byteCountIEC(result.SentLen), result.Address)
				}
			}()
		}
		wg.Wait()
	}
	return nil
}

func (uc *useCase) upload() (*uploader.Result, error) {
	size := uint(uc.config.UploadSize.Min) +
		uint(rand.Intn(int(uc.config.UploadSize.Max)))
	return uc.uploader.UploadRandomTarget(context.Background(), uint(size))
}
