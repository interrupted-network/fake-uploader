package usecase

import (
	"errors"
	"net"
	"sync"
	"time"

	"github.com/interrupted-network/fake-uploader/log"
	"github.com/interrupted-network/fake-uploader/uploader/domain/uploader"
)

type client struct {
	net.Conn
	isConnected bool
	connMtx     *sync.Mutex

	logger   log.Logger
	msgQueue <-chan []byte
	target   *uploader.Target

	sleepDuration time.Duration
	started       bool
}

func newClient(logger log.Logger,
	target *uploader.Target, msgQueue <-chan []byte) *client {
	c := &client{
		connMtx:       new(sync.Mutex),
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
		if !c.isConnected {
			c.connMtx.Lock()
			var err error
			c.Conn, err = net.DialTimeout(
				c.target.Network,
				c.target.Address,
				c.target.DialTimeout,
			)
			c.connMtx.Unlock()
			if err != nil {
				c.errored(err)
				// c.logger.Errorf("error on connect: %v", err)
				// time.Sleep(c.sleepDuration)
				continue
			}
			c.isConnected = true
		}
		time.Sleep(time.Second)
	}
}

func (c *client) errored(err error) {
	c.connMtx.Lock()
	defer c.connMtx.Unlock()
	if c.sleepDuration < time.Minute*10 {
		c.sleepDuration *= 2
	}
	time.Sleep(c.sleepDuration)
	if !c.isConnected {
		return
	}
	c.isConnected = false
}

func (c *client) start() {
	for c.started {
		c.connMtx.Lock()
		if !c.isConnected {
			c.connMtx.Unlock()
			time.Sleep(c.sleepDuration)
			continue
		}
		msg := <-c.msgQueue
		_, err := c.Write(msg)
		c.connMtx.Unlock()
		if err != nil {
			continue
		}
		c.logger.Debugf("%s sent", byteCountIEC(int64(len(msg))))
		if c.sleepDuration > time.Second {
			c.sleepDuration /= 2
		}
	}
}

func (c *client) Write(b []byte) (n int, err error) {
	if c.Conn == nil {
		return 0, errors.New("client is not initialized")
	}
	i, err := c.Conn.Write(b)
	if err != nil {
		c.errored(err)
		c.logger.Errorf("error on write: %v", err)
		return 0, err
	}
	return i, nil
}
