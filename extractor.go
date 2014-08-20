package main

import (
	"bufio"
	"log"
	"strings"
)

func extract(reader *bufio.Reader, fastaChannel chan Fasta) {
	var f Fasta
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				log.Print("Problem with line", line)
				log.Fatal(err)
			}
		}
		line = line[0 : len(line)-1]
		if len(line) == 0 {
			break
		}
		if line[0] == '>' {
			if f != nil {
				fastaChannel <- f
			} else {
				f = new(FastaImp)
				id, des := splitLine(line)
				f.SetId(id)
				f.SetDescription(des)
			}
		} else {
			f.SetSequence(f.GetSequence() + line)
		}
	}
	close(fastaChannel)
}

func splitLine(line string) (string, string) {
	log.Print(line)
	stuff := strings.SplitN(line, " ", 2)
	log.Print(stuff)
	if len(stuff) == 1 {
		return stuff[0][1:], ""
	} else {
		return stuff[0][1:], stuff[1]
	}
}
