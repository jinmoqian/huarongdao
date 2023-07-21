//go:build windowsxp

package main

import "time"

func start() {
	err := work()
	if err != nil {
		errorMessage(err.Error())
	}
}

const keepAliveTimeout = 20 * time.Second
