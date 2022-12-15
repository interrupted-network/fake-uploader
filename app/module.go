package app

import (
	"github.com/interrupted-network/fake-uploader/estimator"
	"github.com/interrupted-network/fake-uploader/uploader"
)

func (a *app) initializeModules() {
	a.initializeEstimator()
	a.initializeUploader()
}

func (a *app) initializeEstimator() {
	a.estimator = estimator.Initialize(
		a.logger.WithPrefix("estimator"), a.viper.Sub("estimator"))
}

func (a *app) initializeUploader() {
	a.uploader = uploader.Initialize(
		a.logger.WithPrefix("uploader"), a.viper.Sub("uploader"))
}
