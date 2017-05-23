package main

import (
	"fmt"
	"os/exec"
)

func nmap(ip string) string {
	out, err := exec.Command("nmap", "-T4", "F", ip).Output()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(out))
	return string(out)
}
