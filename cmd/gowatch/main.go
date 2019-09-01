package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func exec_process(p string, a string) {
	cmd := exec.Command(p, a)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
}
func main() {
	exec_process("ls", "-lah")
}
