package main

import (
	"fmt"
	"os/exec"
)

func openUI(keepAlive bool, url string) (func(), error) {
	err := exec.Command("xdg-open", url).Start()
	return nil, err
}
func errorMessage(message string) {
	fmt.Println(message)
}
func main() {
	start()
}
