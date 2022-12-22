package usecase

func (uc *useCase) Start() {
	if uc.started {
		return
	}
	uc.started = true

	go uc.beginCheckRealtime()
	go uc.beginFillChanRealtime()
}
