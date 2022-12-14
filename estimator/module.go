package estimator

import (
	"github.com/interrupted-network/fake-uploader/estimator/domain/estimator"
	"github.com/interrupted-network/fake-uploader/estimator/usecase"
	"github.com/interrupted-network/fake-uploader/log"
	"github.com/spf13/viper"
)

type Module struct {
	UseCase estimator.UseCase
}

func Initialize(logger log.Logger, registry *viper.Viper) *Module {
	uc := usecase.New(logger.WithPrefix(">uc"), registry)

	mod := &Module{
		UseCase: uc,
	}
	return mod
}
