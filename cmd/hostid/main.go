package main

import (
	"fmt"
)

func main() {
	if id, err := gohostid.GetID(); err == nil {
		fmt.Println("Host ID:", id)
		return
	}
	fmt.Println("ERROR: Cannot get unique Host ID")
}
