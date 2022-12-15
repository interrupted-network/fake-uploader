package app

import (
	"github.com/interrupted-network/fake-uploader/coordinator"
	"github.com/interrupted-network/fake-uploader/estimator"
	"github.com/interrupted-network/fake-uploader/uploader"
)

func (a *app) initializeModules() {
	a.initializeEstimator()
	a.initializeUploader()
	a.initializeCoordinator()
}

func (a *app) initializeEstimator() {
	a.estimator = estimator.Initialize(
		a.logger.WithPrefix("estimator"), a.viper.Sub("estimator"))
}

func (a *app) initializeUploader() {
	a.uploader = uploader.Initialize(
		a.logger.WithPrefix("uploader"), a.viper.Sub("uploader"))
}

func (a *app) initializeCoordinator() {
	a.coordinator = coordinator.Initialize(
		a.logger.WithPrefix("coordinator"), a.viper.Sub("coordinator"),
		a.estimator.UseCase, a.uploader.UseCase)
}
