package usecase

import (
	"math/rand"
	"time"
)

func (uc *useCase) beginProcessTotalBalancer() error {
	if !uc.config.TotalBalancerEnabled {
		uc.logger.Debugf("Total balancer disabled")
		return nil
	}

	for uc.started {
		if err := uc.processTotalBalancer(); err != nil {
			return err
		}
		time.Sleep(uc.config.Interval)
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
		if s <= 512 {
			bytes := make([]byte, int(s))
			uc.msgQueue <- bytes
			continue
		}
		size := int(s) / uc.config.Concurrent
		for i := 0; i < uc.config.Concurrent; i++ {
			bytes := make([]byte, size)
			uc.msgQueue <- bytes
		}
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

func (uc *useCase) beginUpload() {
	size := int64(uc.config.UploadSize.Min) +
		int64(rand.Intn(int(uc.config.UploadSize.Max)))
	uc.logger.Debugf("putting %s in queue", byteCountIEC(size))
	bytes := make([]byte, size)
	uc.msgQueue <- bytes
}
