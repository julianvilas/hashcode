package hashcode

import "fmt"

type Problem interface {
	fmt.Stringer

	Solve() (Result, error)
	AddHeader(header string) error
	AddEntry(entry string) error
}

type Result interface {
	fmt.Stringer
}
