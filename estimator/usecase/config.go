package usecase

type config struct {
	InterfaceName string
}

func (c *config) initialize() {
	if c.InterfaceName == "" {
		panic("interface name can not be empty")
	}
}
