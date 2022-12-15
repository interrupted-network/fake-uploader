package usecase

import (
	"time"

	"github.com/interrupted-network/fake-uploader/uploader/domain/uploader"
)

type config struct {
	Targets []*uploader.Target
}

func (c *config) initialize() {
	for _, t := range c.Targets {
		if t.DialTimeout == 0 {
			t.DialTimeout = time.Second
		}
		if t.RWTimeout == 0 {
			t.RWTimeout = time.Second * 3
		}
	}
}
