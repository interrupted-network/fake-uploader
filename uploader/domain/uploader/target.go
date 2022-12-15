package uploader

import "time"

type Target struct {
	Network     string
	Address     string
	DialTimeout time.Duration
	RWTimeout   time.Duration
}
