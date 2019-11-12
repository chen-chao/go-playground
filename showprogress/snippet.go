// -*- mode:go;mode:go-playground -*-
// snippet of code @ 2019-11-07 15:11:49

// === Go Playground ===
// Execute the snippet with Ctl-Return
// Provide custom arguments to compile with Alt-Return
// Remove the snippet completely with its dir and all files M-x `go-playground-rm`

package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func startProgress(indication string, ch <-chan error) {
	symbols := []string{"-", "/", "|", "\\"}

	for i := 0; ; i++ {
		select {
		case err := <-ch:
			var finished string
			if err != nil {
				finished = "Failed, " + err.Error()
			} else {
				finished = "Done"
			}

			log.Println(indication+" ", finished)

			return
		default:
			i %= 4

			fmt.Print(indication + " " + symbols[i] + "\r")

			time.Sleep(200 * time.Millisecond)
		}
	}
}

func work() error {
	time.Sleep(10 * time.Second)
	return errors.New("oh no, failed")
}

func main() {
	finished := make(chan error)

	go func() {
		finished <- work()
	}()

	startProgress("Working!", finished)
}
