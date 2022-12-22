package uploader

import (
	"net"
	"time"
)

type Request struct {
	Client   net.Conn
	Size     uint
	Deadline time.Duration
}
