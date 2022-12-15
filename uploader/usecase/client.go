package usecase

import (
	"net"
	"time"

	"github.com/interrupted-network/fake-uploader/log"
	"github.com/interrupted-network/fake-uploader/uploader/domain/uploader"
)

type client struct {
	net.Conn
	logger   log.Logger
	msgQueue <-chan []byte
	target   *uploader.Target

	sleepChan     <-chan bool
	sleepDuration time.Duration
	started       bool
}

func newClient(logger log.Logger,
	target *uploader.Target, msgQueue <-chan []byte) *client {
	c := &client{
		logger:        logger,
		msgQueue:      msgQueue,
		target:        target,
		sleepDuration: time.Second,
	}
	return c
}

func (c *client) Start() {
	if c.started {
		return
	}
	c.started = true
	go c.checkLiveness()
	go c.start()
}

func (c *client) checkLiveness() {
	for c.started {
		if c.Conn == nil {
			var err error
			c.Conn, err = net.DialTimeout(
				c.target.Network,
				c.target.Address,
				c.target.DialTimeout,
			)
			if err != nil {
				c.errored(err)
				time.Sleep(c.sleepDuration)
				continue
			}
		}
		time.Sleep(time.Second)
	}
}

func (c *client) errored(err error) {
	c.Conn = nil
	if c.sleepDuration < time.Minute*10 {
		c.sleepDuration *= 2
	}
	time.Sleep(c.sleepDuration)
}

func (c *client) start() {
	for c.started {
		if c.Conn == nil {
			time.Sleep(c.sleepDuration)
			continue
		}
		msg := <-c.msgQueue
		_, err := c.Write(msg)
		if err != nil {
			c.errored(err)
			continue
		}
		c.logger.DebugF("%s sent", byteCountIEC(int64(len(msg))))
		if c.sleepDuration > time.Second {
			c.sleepDuration /= 2
		}
	}
}
