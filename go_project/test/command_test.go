package main

import (
	"fmt"
	"log"
	"os/exec"
	"testing"
)

func TestCommand(t *testing.T) {

	cmd := exec.Command("ls", "-l", "/Users/louis")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}
