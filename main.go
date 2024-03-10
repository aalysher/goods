package main

import (
	"fmt"

	"github.com/aalysher/goods/internal/service"
)

func main() {
	_, err := service.New()
	if err != nil {
		fmt.Println(err)
	}
}
