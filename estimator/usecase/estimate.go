package usecase

import (
	"os/exec"

	"github.com/interrupted-network/fake-uploader/estimator/domain/estimator"
)

func (uc *useCase) Estimate() (*estimator.Result, error) {
	out, err := exec.Command(
		uc.config.Command.CMD, uc.config.Command.Args...).Output()
	if err != nil {
		uc.logger.ErrorF("error on execute command, err: %v", err)
	}

	uc.logger.DebugF(string(out))
	return nil, nil
}
