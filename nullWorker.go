package main

import (
	"math/rand"
	"time"
)

type NullWorker struct {
}

var r *rand.Rand = rand.New(rand.NewSource(99))

func (pw *NullWorker) Work(id string, fastaChannel chan Fasta, resultsChannel chan string) {
	// do some work
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	resultsChannel <- id
}
