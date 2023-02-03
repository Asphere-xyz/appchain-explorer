package main

import (
	"github.com/Ankr-network/appchain-explorer/shared/container"
	"time"
)

func main() {
	di := container.NewContainer()
	container.MustStartDefault(di)
	for {
		time.Sleep(1 * time.Second)
	}
}
