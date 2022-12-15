package app

import (
	l "log"
	"os"

	"github.com/interrupted-network/fake-uploader/coordinator"
	"github.com/interrupted-network/fake-uploader/estimator"
	"github.com/interrupted-network/fake-uploader/log"
	"github.com/interrupted-network/fake-uploader/uploader"
	"github.com/spf13/viper"
)

type Application interface {
	Initialize()
	Start()
}

type app struct {
	logger log.Logger
	viper  *viper.Viper

	estimator   *estimator.Module
	uploader    *uploader.Module
	coordinator *coordinator.Module
}

func New() Application {
	a := &app{
		logger: log.New(l.New(os.Stdout, "",
			l.LstdFlags|l.Lshortfile|l.Lmsgprefix)),
	}
	return a
}
