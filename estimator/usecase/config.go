package usecase

import "regexp"

type config struct {
	Command struct {
		CMD  string
		Args []string
	}
	RxPattern string
	TxPattern string

	rxPattern *regexp.Regexp
	txPattern *regexp.Regexp
}

func (c *config) initialize() {
	c.rxPattern = regexp.MustCompile(c.RxPattern)
	c.txPattern = regexp.MustCompile(c.TxPattern)
}
