package uploader

import (
	"github.com/interrupted-network/fake-uploader/log"
	"github.com/interrupted-network/fake-uploader/uploader/domain/uploader"
	"github.com/interrupted-network/fake-uploader/uploader/usecase"
	"github.com/spf13/viper"
)

type Module struct {
	UseCase uploader.UseCase
}

func Initialize(logger log.Logger, registry *viper.Viper) *Module {
	uc := usecase.New(logger.WithPrefix(">uc"), registry)

	mod := &Module{
		UseCase: uc,
	}
	return mod
}
