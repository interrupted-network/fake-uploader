package usecase

import (
	"encoding/json"
	"errors"
	"os/exec"

	"github.com/interrupted-network/fake-uploader/estimator/domain/estimator"
)

type trafficBucket struct {
	ID   int `json:"id"`
	Date struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Day   int `json:"day"`
	} `json:"date"`
	Time struct {
		Hour   int `json:"hour"`
		Minute int `json:"minute"`
	} `json:"time"`
	Rx int64 `json:"rx"`
	Tx int64 `json:"tx"`
}

type data struct {
	Vnstatversion string `json:"vnstatversion"`
	Jsonversion   string `json:"jsonversion"`
	Interfaces    []struct {
		Name    string `json:"name"`
		Alias   string `json:"alias"`
		Created struct {
			Date struct {
				Year  int `json:"year"`
				Month int `json:"month"`
				Day   int `json:"day"`
			} `json:"date"`
		} `json:"created"`
		Updated struct {
			Date struct {
				Year  int `json:"year"`
				Month int `json:"month"`
				Day   int `json:"day"`
			} `json:"date"`
			Time struct {
				Hour   int `json:"hour"`
				Minute int `json:"minute"`
			} `json:"time"`
		} `json:"updated"`
		Traffic struct {
			Total struct {
				Rx int64 `json:"rx"`
				Tx int64 `json:"tx"`
			} `json:"total"`
			Fiveminute []trafficBucket `json:"fiveminute"`
			Hour       []trafficBucket `json:"hour"`
			Day        []trafficBucket `json:"day"`
			Month      []trafficBucket `json:"month"`
			Year       []trafficBucket `json:"year"`
			Top        []trafficBucket `json:"top"`
		} `json:"traffic"`
	} `json:"interfaces"`
}

func (uc *useCase) Estimate() (*estimator.Result, error) {
	out, err := exec.Command(
		"vnstat", "-i", uc.config.InterfaceName, "--json").Output()
	if err != nil {
		uc.logger.
			WithPrefix("Estimate").
			Errorf("error on command, err: %v", err)
	}

	data := new(data)
	if err = json.Unmarshal(out, data); err != nil {
		return nil, err
	}

	if len(data.Interfaces) == 0 {
		return nil, errors.New("no interface where found")
	}
	interface0 := data.Interfaces[0]
	result := &estimator.Result{
		Rx: interface0.Traffic.Total.Rx,
		Tx: interface0.Traffic.Total.Tx,
	}
	result.TxRxRatio = float32(float64(result.Tx) / float64(result.Rx))
	return result, nil
}
