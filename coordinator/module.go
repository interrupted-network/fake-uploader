package coordinator

import (
	"github.com/interrupted-network/fake-uploader/coordinator/domain/coordinator"
	"github.com/interrupted-network/fake-uploader/coordinator/usecase"
	"github.com/interrupted-network/fake-uploader/estimator/domain/estimator"
	"github.com/interrupted-network/fake-uploader/log"
	"github.com/interrupted-network/fake-uploader/uploader/domain/uploader"
	"github.com/spf13/viper"
)

type Module struct {
	UseCase coordinator.UseCase
}

func Initialize(logger log.Logger, registry *viper.Viper,
	estimator estimator.UseCase, uploader uploader.UseCase) *Module {
	uc := usecase.New(logger.WithPrefix(">uc"), registry, estimator, uploader)

	mod := &Module{
		UseCase: uc,
	}
	return mod
}
