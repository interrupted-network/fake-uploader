package usecase

import "time"

type config struct {
	UploadSize struct {
		Min uint
		Max uint
	}
	TotalBalancerEnabled    bool
	TxRxMinRatio            float32
	TxRxMaxRatio            float32
	Interval                time.Duration
	Concurrent              int
	RealtimeBalancerEnabled bool
	RealtimeTxRxRatio       float32
}

func (c *config) initialize() {
	if c.Interval == 0 {
		c.Interval = time.Minute
	}
	if c.Concurrent == 0 {
		c.Concurrent = 1
	}
	if c.RealtimeTxRxRatio == 0 {
		c.Concurrent = 10
	}
}
