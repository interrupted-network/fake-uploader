package app

func (a *app) Initialize() {
	a.loadConfig()
	a.initializeModules()
	a.logger.DebugF("initialized")
}
