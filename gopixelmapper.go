package main

import (

	// "github.com/mellbergsimon/gopixelmapper/findartnet"

	"github.com/mellbergsimon/gopixelmapper/ndi"
)

func main() {

	// artnetAddresses := findartnet.FindArtnet()
	// for _, addr := range artnetAddresses {
	// 	fmt.Printf("ArtNet at: %v\n", addr)
	// }

	ndi.GetNDI()
}
