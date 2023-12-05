package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

const readmePath = "README.md"

func updateReadme() error {
	timestamp := time.Now().Format("02/01/2006 15:04:05")
	content := fmt.Sprintf("# Hello from me %s\n", timestamp)

	return ioutil.WriteFile(readmePath, []byte(content), 0644)
}

func main() {
	if err := updateReadme(); err != nil {
		fmt.Println("Error updating README.md:", err)
	}
}
