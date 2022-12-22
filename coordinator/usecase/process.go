package usecase

import (
	"context"
	"math/rand"

	"github.com/interrupted-network/fake-uploader/uploader/domain/uploader"
)

func (uc *useCase) beginProcessTotalBalancer() error {
	if !uc.config.TotalBalancerEnabled {
		uc.logger.Debugf("Total balancer disabled")
		return nil
	}
	if err := uc.processTotalBalancer(); err != nil {
		return err
	}
	return nil
}

func (uc *useCase) beginProcessRealtimeBalancer() error {
	if !uc.config.RealtimeBalancerEnabled {
		uc.logger.Debugf("Real time balancer disabled")
		return nil
	}

	ch := uc.estimator.GetRealtimeSpeedChan()
	for uc.started {
		speed := <-ch
		s := (float32(speed.Rx) * uc.config.RealtimeTxRxRatio) - float32(speed.Tx)
		if s <= 0 {
			continue
		}
		size := int(s) / uc.config.Concurrent

		// ctx, cancelFunc := context.WithTimeout(ctx, time.Second)
		// eg, _ := errgroup.WithContext(ctx)

		for i := 0; i < uc.config.Concurrent; i++ {
			bytes := make([]byte, size)
			uc.msgQueue <- bytes
			// eg.Go(func() error {
			// 	req := &uploader.Request{
			// 		// Deadline: time.Second,
			// 		Size: uint(size),
			// 	}
			// 	_, err := uc.uploader.Upload(ctx, req)
			// 	if err != nil {
			// 		// uc.logger.WithPrefix("realtime_balancer.Upload").
			// 		// 	Errorf("error on upload: %v", err)
			// 		return err
			// 	}
			// 	return nil
			// })
		}
		// eg.Wait()
		// cancelFunc()
	}

	return nil
}

func (uc *useCase) processTotalBalancer() error {
	estimated, err := uc.estimator.Estimate()
	if err != nil {
		return err
	}
	uc.logger.Debugf("total estimated. Tx: %s, Rx: %s, Ratio: %f",
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
	req := &uploader.Request{
		Size: uint(size),
	}
	return uc.uploader.Upload(context.Background(), req)
}

func (uc *useCase) beginUpload() {
	size := int64(uc.config.UploadSize.Min) +
		int64(rand.Intn(int(uc.config.UploadSize.Max)))
	uc.logger.Debugf("putting %s in queue", byteCountIEC(size))
	bytes := make([]byte, size)
	uc.msgQueue <- bytes
}
