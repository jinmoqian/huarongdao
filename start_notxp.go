//go:build !windowsxp

package main

import "C"
import "time"

//export start
func start() {
	err := work()
	if err != nil {
		errorMessage(err.Error())
	}
}

const keepAliveTimeout = 5 * time.Second
