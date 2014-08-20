package main

import (
	"fmt"
)

type PrintWorker struct {
}

func (pw *PrintWorker) Work(id string, fastaChannel chan Fasta, closeChannel chan bool) {
	for {
		fa := <-fastaChannel
		if fa == nil {
			break
		}
		fmt.Println("")
		fastaPrint(fa)
	}
	closeChannel <- true
}
