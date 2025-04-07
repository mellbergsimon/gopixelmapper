package main

import (
	"fmt"

	"github.com/mellbergsimon/gopixelmapper/findartnet"
)

func main() {

	artnetAddresses := findartnet.FindArtnet()
	for _, addr := range artnetAddresses {
		fmt.Printf("ArtNet at: %v\n", addr)
	}

}
