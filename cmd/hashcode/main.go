package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/julianvilas/hashcode"
)

func main() {
	os.Exit(mainWithExitCode())
}

func mainWithExitCode() int {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <dataset_file>\n", os.Args[0])
		return 1
	}
	datasetFile := os.Args[1]

	log.Printf("loading file %s", datasetFile)
	problem, err := parseDataset(datasetFile)
	if err != nil {
		log.Println(err)
		return 1
	}

	log.Printf("simulating problem: %s", problem)
	result, err := problem.Solve()
	if err != nil {
		log.Println(err)
		return 1
	}

	log.Printf("writing results")
	fmt.Println(result)

	return 0
}

/*
	Assumption: dataset files are composed by a header and a line for every entry:

		HEADER
		ENTRY1
		ENTRY2
		...
		ENTRYN
*/
func parseDataset(datasetFile string) (hashcode.Problem, error) {
	f, err := os.Open(datasetFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	p := hashcode.NewProblem()

	scanner := bufio.NewScanner(f)

	// Read header.
	if ok := scanner.Scan(); !ok {
		return nil, fmt.Errorf("can not parse header: %v", scanner.Err())
	}

	h := scanner.Text()

	log.Printf("adding header: %s", h)
	if err := p.AddHeader(h); err != nil {
		return nil, err
	}

	log.Printf("adding entries")
	// Read entries.
	for scanner.Scan() {
		if err := p.AddEntry(scanner.Text()); err != nil {
			return nil, err
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("can not parse entries: %v", err)
	}

	return p, nil
}
