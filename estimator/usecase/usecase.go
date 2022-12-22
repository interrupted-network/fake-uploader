package usecase

import (
	"sync"

	"github.com/interrupted-network/fake-uploader/estimator/domain/estimator"
	"github.com/interrupted-network/fake-uploader/log"
	"github.com/spf13/viper"
)

type useCase struct {
	logger log.Logger
	config config

	started   bool
	speedChan chan *estimator.Speed

	lastParams    []string
	lastParamsMtx *sync.Mutex
}

func New(logger log.Logger, registry *viper.Viper) estimator.UseCase {
	uc := &useCase{
		logger:        logger,
		speedChan:     make(chan *estimator.Speed),
		lastParamsMtx: new(sync.Mutex),
	}

	if err := registry.Unmarshal(&uc.config); err != nil {
		panic(err)
	}
	uc.config.initialize()
	return uc
}
