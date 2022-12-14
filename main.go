package main

import (
	"github.com/interrupted-network/fake-uploader/app"
)

func main() {
	application := app.New()
	application.Initialize()

	application.Start()
}
