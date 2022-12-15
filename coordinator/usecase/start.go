package usecase

import "time"

func (uc *useCase) Start() {
	if uc.started {
		return
	}
	uc.started = true
	go uc.start()
	uc.logger.DebugF("started")
}

func (uc *useCase) start() {
	for uc.started {
		if err := uc.Process(); err != nil {
			uc.logger.ErrorF("process. error: %v", err)
			time.Sleep(time.Second * 5)
			continue
		}
		time.Sleep(uc.config.Interval)
	}
}
