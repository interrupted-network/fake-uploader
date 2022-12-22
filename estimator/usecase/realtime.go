package usecase

import (
	"bufio"
	"os/exec"
	"strconv"
	"time"

	"github.com/interrupted-network/fake-uploader/estimator/domain/estimator"
)

func (uc *useCase) GetRealtimeSpeedChan() <-chan *estimator.Speed {
	return uc.speedChan
}

func (uc *useCase) beginCheckRealtime() {
	cmd := exec.Command(
		"vnstat", "-i", uc.config.InterfaceName, "-l", "--rateunit", "0")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanWords)

	var parts []string = nil
	for uc.started && scanner.Scan() {
		m := scanner.Text()
		if m == "rx:" {
			uc.lastParamsMtx.Lock()
			parts = make([]string, 0)
		}
		if parts == nil {
			continue
		}
		parts = append(parts, m)
		if len(parts) == 10 {
			uc.lastParams = parts
			uc.lastParamsMtx.Unlock()
		}
	}
}

func (uc *useCase) beginFillChanRealtime() {
	for uc.started {
		if len(uc.lastParams) != 10 {
			continue
		}
		uc.lastParamsMtx.Lock()
		speed := uc.getSpeed(uc.lastParams)
		uc.lastParams = nil
		uc.lastParamsMtx.Unlock()
		if speed != nil {
			uc.logger.Debugf("realtime speed rx: %s tx: %s",
				speed.Rx.ToString(), speed.Tx.ToString())
			uc.speedChan <- speed
		}
		time.Sleep(time.Second)
	}
}

func (uc *useCase) getSpeed(parts []string) *estimator.Speed {
	result := new(estimator.Speed)
	speed, err := strconv.ParseFloat(parts[1], 32)
	if err != nil {
		return nil
	}
	result.Rx = estimator.Rate(speed)

	ratio := uc.getUnitRatio(parts[2])
	if ratio == 0 {
		return nil
	}
	result.Rx *= estimator.Rate(ratio)
	// tx
	speed, err = strconv.ParseFloat(parts[6], 32)
	if err != nil {
		return nil
	}
	result.Tx = estimator.Rate(speed)

	ratio = uc.getUnitRatio(parts[7])
	if ratio == 0 {
		return nil
	}
	result.Tx *= estimator.Rate(ratio)

	return result
}

func (uc *useCase) getUnitRatio(unit string) int64 {
	switch unit {
	case "B/s":
		return 1
	case "KiB/s":
		return 1024
	case "MiB/s":
		return 1024 * 1024
	case "GiB/s":
		return 1024 * 1024 * 1024
	default:
		uc.logger.Errorf("unknown unit: %s", unit)
	}
	return 0
}
