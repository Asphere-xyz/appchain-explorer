package main

import (
	"github.com/Ankr-network/ankr-protocol/shared/container"
	"time"
)

func main() {
	di := container.NewContainer()
	container.MustStartDefault(di)
	for {
		time.Sleep(1 * time.Second)
	}
}
