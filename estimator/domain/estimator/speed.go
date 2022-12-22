package estimator

import "fmt"

type Rate float32

func (m Rate) ToString() string {
	if m >= 1024*1024*1024 {
		return fmt.Sprintf("%.1f GiB/s", m/(1024*1024*1024))
	}
	if m >= 1024*1024 {
		return fmt.Sprintf("%.1f MiB/s", m/(1024*1024))
	}
	if m >= 1024 {
		return fmt.Sprintf("%.1f KiB/s", m/1024)
	}
	return fmt.Sprintf("%.1f B/s", m)
}

type Speed struct {
	Rx Rate
	Tx Rate
}
