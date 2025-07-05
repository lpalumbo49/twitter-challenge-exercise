package main

import (
	"fmt"
	"twitter-challenge-exercise/internal"
)

func main() {
	container, err := internal.StartContainer()
	if err != nil {
		panic(fmt.Errorf("error initializing application container: %v", err))
	}

	err = container.ServeRouter()
	if err != nil {
		panic(fmt.Errorf("error starting http server: %v", err))
	}
}
