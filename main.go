package main

import (
	"go-worker-template/cmd"
	"go-worker-template/config"
)

func main() {
	config.SetEnvironment()

	w := new(cmd.Worker)
	w.StartWorker()
}
