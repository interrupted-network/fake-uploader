package usecase

type config struct {
	Targets []struct {
		Network string
		Address string
	}
}

func (c *config) initialize() {
}
