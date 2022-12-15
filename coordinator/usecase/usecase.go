package usecase

import (
	"github.com/interrupted-network/fake-uploader/coordinator/domain/coordinator"
	"github.com/interrupted-network/fake-uploader/estimator/domain/estimator"
	"github.com/interrupted-network/fake-uploader/log"
	"github.com/interrupted-network/fake-uploader/uploader/domain/uploader"
	"github.com/spf13/viper"
)

type useCase struct {
	logger log.Logger
	config config

	estimator estimator.UseCase
	uploader  uploader.UseCase

	started bool
}

func New(logger log.Logger, registry *viper.Viper,
	estimator estimator.UseCase, uploader uploader.UseCase) coordinator.UseCase {
	uc := &useCase{
		logger:    logger,
		estimator: estimator,
		uploader:  uploader,
	}

	if err := registry.Unmarshal(&uc.config); err != nil {
		panic(err)
	}
	uc.config.initialize()
	return uc
}
