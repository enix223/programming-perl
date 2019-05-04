package main

import (
	"encoding/binary"
	"flag"
	"io"
	"log"
	"os"
)

// readUInt32 read uint32 from the file reader
func readUInt32(input io.Reader) <-chan uint32 {
	res := make(chan uint32)
	go func() {
		defer close(res)
		log.Print("Reading numbers from input file...")
		for {
			var i uint32
			err := binary.Read(input, binary.BigEndian, &i)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					log.Fatalf("Failed to read input file: %v", err)
				}
			}

			res <- i
		}
	}()
	return res
}

// buildNumExistList map each number from the input source,
// and build an index for the existence of each number, and
// output to out channel sequencially.
func buildNumExistList(input <-chan uint32) <-chan uint32 {
	res := make(chan uint32)
	go func() {
		numMap := make([]bool, 10e7, 10e7)
		defer close(res)

		// mark each number from input and turn that num exist flag on
		log.Print("Buiding number existence list...")
		for i := range input {
			if numMap[i] {
				log.Fatalf("Duplicate number found: %d", i)
			}
			numMap[i] = true
		}

		log.Print("Reading sorted numbers...")
		for idx, exist := range numMap {
			if exist {
				res <- uint32(idx)
			}
		}
	}()
	return res
}

// writeSortedNums write each number in the sorted sequence into out file
func writeSortedNums(input <-chan uint32, out io.Writer) {
	for i := range input {
		binary.Write(out, binary.BigEndian, &i)
	}

	log.Print("Done")
}

// sort sort the numbers from input file, and write them to output file
func sort(infile, outfile string) {
	inf, err := os.Open(infile)
	if err != nil {
		log.Fatalf("Failed to open input file: %v", err)
	}
	defer inf.Close()

	outf, err := os.OpenFile(outfile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to open output file: %v", err)
	}
	defer outf.Close()

	num := readUInt32(inf)
	sorted := buildNumExistList(num)
	writeSortedNums(sorted, outf)
}

func sortNum(infile, outfile string) {
	inf, err := os.Open(infile)
	if err != nil {
		log.Fatalf("Failed to open input file: %v", err)
	}
	defer inf.Close()

	outf, err := os.OpenFile(outfile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to open output file: %v", err)
	}
	defer outf.Close()

	numMap := make([]byte, 10e7, 10e7)
	log.Print("Reading numbers from input file...")
	for {
		var i uint32
		err := binary.Read(inf, binary.BigEndian, &i)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("Failed to read input file: %v", err)
			}
		}

		if numMap[i] == 1 {
			log.Fatalf("Duplicate number found: %d", i)
		}

		numMap[i] = 1
	}

	log.Print("Reading sorted numbers...")
	for idx, exist := range numMap {
		if exist == 1 {
			err := binary.Write(outf, binary.BigEndian, uint32(idx))
			if err != nil {
				log.Fatalf("Failed to write number to outfile: %v", err)
			}
		}
	}

	log.Print("Done")
}

func main() {
	infile := flag.String("infile", "", "input random numbers file")
	outfile := flag.String("outfile", "", "output sorted numbers file")
	flag.Parse()

	if len(*infile) == 0 {
		log.Fatal("infile argument not specify")
	}

	if len(*outfile) == 0 {
		log.Fatal("outfile argument not specify")
	}

	// sort(*infile, *outfile)
	sortNum(*infile, *outfile)
}
