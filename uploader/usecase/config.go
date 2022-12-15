package usecase

type config struct {
	UploadSize struct {
		Min uint
		Max uint
	}
	Targets []struct {
		Network string
		Address string
	}
}

func (c *config) initialize() {
}
