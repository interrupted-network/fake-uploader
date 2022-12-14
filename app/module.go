package app

import "github.com/interrupted-network/fake-uploader/estimator"

func (a *app) initializeModules() {
	a.initializeEstimator()
}

func (a *app) initializeEstimator() {
	a.estimator = estimator.Initialize(
		a.logger.WithPrefix("estimator"), a.viper.Sub("estimator"))
}
