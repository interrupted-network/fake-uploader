package app

func (a *app) Start() {
	a.estimator.UseCase.Estimate()
}
