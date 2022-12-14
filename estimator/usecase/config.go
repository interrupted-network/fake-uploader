package usecase

type config struct {
	Command struct {
		Name string
		Args []string
	}
	InterfaceName string
}

func (c *config) initialize() {
	if c.InterfaceName == "" {
		panic("interface name can not be empty")
	}
}
