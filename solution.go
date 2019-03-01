package hashcode

// Here is where you just implement the problem solution.

type ProblemImplementation struct{}

func NewProblem() Problem {
	return &ProblemImplementation{}
}

func (p *ProblemImplementation) AddHeader(header string) error {
	return nil
}

func (p *ProblemImplementation) AddEntry(entry string) error {
	return nil
}

func (p *ProblemImplementation) String() string {
	return ""
}

func (p *ProblemImplementation) Solve() (Result, error) {
	return nil, nil
}
