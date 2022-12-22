package app

import (
	"math/rand"
	"time"
)

func (a *app) Initialize() {
	rand.Seed(time.Now().Unix())

	a.loadConfig()
	a.initializeModules()
	a.logger.Debugf("initialized")
}
