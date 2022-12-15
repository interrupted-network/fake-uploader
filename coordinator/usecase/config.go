package usecase

import "time"

type config struct {
	UploadSize struct {
		Min uint
		Max uint
	}
	RxTxMaxRatio float32
	Interval     time.Duration
	Concurrent   int
}

func (c *config) initialize() {
	if c.Interval == 0 {
		c.Interval = time.Minute
	}
	if c.Concurrent == 0 {
		c.Concurrent = 1
	}
}
