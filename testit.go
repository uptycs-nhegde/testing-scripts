package main

import (
	"log"
	"os/exec"
)

func main() {
	// Run whoami and ifconfig after downloading from the internet
	whoamiCmd := exec.Command("whoami")
	whoamiOut, err := whoamiCmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	println("Whoami: ", string(whoamiOut))

	ifconfigCmd := exec.Command("ifconfig")
	ifconfigOut, err := ifconfigCmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	println("Ifconfig: ", string(ifconfigOut))
}
