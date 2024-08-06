package main

import (
	"fmt"
	"sky-take-out-gin/internal/utils/initialize"
)

func main() {
	err := initialize.Initialize()
	if err != nil {
		fmt.Println(err)
	}
}
