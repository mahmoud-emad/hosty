package cmd

import "fmt"

func printHeader() {
	fmt.Printf("%-15s %-20s %-8s %-15s\n", "Name", "Address", "Port", "User")
}

func printHost(name, addr, port, user string) {
	fmt.Printf("%-15s %-20s %-8s %-15s\n", name, addr, port, user)
}
