package usecase

import (
	"github.com/interrupted-network/fake-uploader/estimator/domain/estimator"
	"github.com/interrupted-network/fake-uploader/log"
	"github.com/spf13/viper"
)

type useCase struct {
	logger log.Logger
	config config
}

func New(logger log.Logger, registry *viper.Viper) estimator.UseCase {
	uc := &useCase{
		logger: logger,
	}

	if err := registry.Unmarshal(&uc.config); err != nil {
		panic(err)
	}
	uc.config.initialize()
	return uc
}
