package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fastaChannel := make(chan Fasta, 100)
	closeChannel := make(chan bool)

	go extract(reader, fastaChannel)
	numWorkers := 12

	for i := 0; i < numWorkers; i++ {
		go worker(i, fastaChannel, closeChannel)
	}

	for i := 0; i < numWorkers; i++ {
		_ = <-closeChannel
	}
}

func worker(id int, fastaChannel chan Fasta, closeChannel chan bool) {
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

func fastaPrint(fa Fasta) {
	fmt.Println("         ID = " + fa.GetId())
	fmt.Println("Description = " + fa.GetDescription())
	fmt.Println("   Sequence = " + fa.GetSequence())
}
