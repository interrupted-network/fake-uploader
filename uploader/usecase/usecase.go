package usecase

import (
	"time"

	"github.com/interrupted-network/fake-uploader/log"
	"github.com/interrupted-network/fake-uploader/uploader/domain/uploader"
	"github.com/spf13/viper"
)

type useCase struct {
	logger log.Logger
	config config

	clients []*client
}

func New(logger log.Logger, registry *viper.Viper) uploader.UseCase {
	uc := &useCase{
		logger:  logger,
		clients: make([]*client, 0),
	}

	if err := registry.Unmarshal(&uc.config); err != nil {
		panic(err)
	}
	uc.config.initialize()
	return uc
}

func (uc *useCase) Initialize(msgQueue <-chan []byte) {
	for _, target := range uc.config.Targets {
		uc.clients = append(uc.clients, newClient(uc.logger, target, msgQueue))
	}
}

func (uc *useCase) Start() {
	for _, c := range uc.clients {
		c.Start()
	}
	go uc.start()
}

func (uc *useCase) start() {
	for {
		numAlive := 0
		for _, c := range uc.clients {
			if c.isConnected {
				numAlive++
			}
		}
		uc.logger.DebugF("alive connections: %d", numAlive)
		time.Sleep(time.Second * 10)
	}
}
