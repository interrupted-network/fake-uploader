package usecase

func (uc *useCase) Start() {
	if uc.started {
		return
	}
	uc.started = true
	go uc.beginProcessTotalBalancer()
	go uc.beginProcessRealtimeBalancer()
	uc.logger.Debugf("started")
}
