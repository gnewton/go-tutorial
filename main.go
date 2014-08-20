package main

//
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Worker interface {
	Work(id string, fastaChannel chan Fasta, resultsChannel chan string)
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fastaChannel := make(chan Fasta, 100)
	resultsChannel := make(chan string)

	go extract(reader, fastaChannel)
	numWorkers := 12

	for i := 0; i < numWorkers; i++ {
		worker := new(NullWorker)
		go worker.Work(strconv.Itoa(i), fastaChannel, resultsChannel)
	}

	sum := 0
	for i := 0; i < numWorkers; i++ {
		sum += len(<-resultsChannel)
	}

	fmt.Println("Results: ", sum)
}

func fastaPrint(fa Fasta) {
	fmt.Println("         ID = " + fa.GetId())
	fmt.Println("Description = " + fa.GetDescription())
	fmt.Println("   Sequence = " + fa.GetSequence())
}
