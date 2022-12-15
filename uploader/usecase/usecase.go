package usecase

import (
	"github.com/interrupted-network/fake-uploader/log"
	"github.com/interrupted-network/fake-uploader/uploader/domain/uploader"
	"github.com/spf13/viper"
)

type useCase struct {
	logger log.Logger
	config config
}

func New(logger log.Logger, registry *viper.Viper) uploader.UseCase {
	uc := &useCase{
		logger: logger,
	}

	if err := registry.Unmarshal(&uc.config); err != nil {
		panic(err)
	}
	uc.config.initialize()
	return uc
}
