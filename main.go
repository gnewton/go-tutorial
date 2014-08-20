package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Worker interface {
	Work(id string, fastaChannel chan Fasta, closeChannel chan bool)
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fastaChannel := make(chan Fasta, 100)
	closeChannel := make(chan bool)

	go extract(reader, fastaChannel)
	numWorkers := 12

	for i := 0; i < numWorkers; i++ {
		worker := new(PrintWorker)
		go worker.Work(strconv.Itoa(i), fastaChannel, closeChannel)
	}

	for i := 0; i < numWorkers; i++ {
		_ = <-closeChannel
	}
}

func fastaPrint(fa Fasta) {
	fmt.Println("         ID = " + fa.GetId())
	fmt.Println("Description = " + fa.GetDescription())
	fmt.Println("   Sequence = " + fa.GetSequence())
}
