package main

import (
	"fmt"

	"github.com/mahmoud-emad/hosty/internal/host"
)

func main() {
	fmt.Println("Welcome to Hosty!")

	hosty, err := host.NewHosty()
	if err != nil {
		fmt.Println(err)
		return
	}

	h, err := host.NewHost("router", "192.168.1.9", "", 8000)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := hosty.Set(h); err != nil {
		fmt.Println(err)
	}

	h, err = hosty.Get("router")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(h)
}
