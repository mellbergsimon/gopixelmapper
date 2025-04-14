package ndi

import (
	"fmt"

	"github.com/mellbergsimon/gondi"
)

func GetNDI() {

	fmt.Println("Initializing NDI")
	gondi.InitLibrary("")

	version := gondi.GetVersion()
	fmt.Printf("NDI version: %s\n", version)

}
