package usecase

import (
	"context"
	"math/rand"

	"github.com/interrupted-network/fake-uploader/uploader/domain/uploader"
)

func (uc *useCase) Process() error {
	estimated, err := uc.estimator.Estimate()
	if err != nil {
		return err
	}

	uc.logger.DebugF("estimated. Tx: %s, Rx: %s, Ratio: %f",
		byteCountIEC(estimated.Tx), byteCountIEC(estimated.Rx), estimated.TxRxRatio)

	if (!uc.isUploadStarted && estimated.TxRxRatio < uc.config.TxRxMinRatio) ||
		(uc.isUploadStarted && estimated.TxRxRatio < uc.config.TxRxMaxRatio) {
		uc.isUploadStarted = true
		for i := 0; i < int(uc.config.Concurrent); i++ {
			uc.beginUpload()
		}
	} else {
		uc.isUploadStarted = false
	}
	return nil
}

func (uc *useCase) upload() (*uploader.Result, error) {
	size := uint(uc.config.UploadSize.Min) +
		uint(rand.Intn(int(uc.config.UploadSize.Max)))
	return uc.uploader.Upload(context.Background(), uint(size))
}

func (uc *useCase) beginUpload() {
	size := int64(uc.config.UploadSize.Min) +
		int64(rand.Intn(int(uc.config.UploadSize.Max)))
	uc.logger.DebugF("putting %s in queue", byteCountIEC(size))
	bytes := make([]byte, size)
	uc.msgQueue <- bytes
}
