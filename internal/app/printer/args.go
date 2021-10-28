package printer

import (
	"flag"
)

const (
	ALL = "ALL"
	CPU = "CPU"
	MEM = "MEM"
)

func Setup() {
	all := flag.Bool("a", false, " : All Information")
	cpu := flag.Bool("c", false, " : CPU Information")
	mem := flag.Bool("m", false, " : Memory Information")

	flag.Parse()

	argsCheckup(all, cpu, mem)
}

func argsCheckup(all, cpu, mem *bool) {
	switch {
	case *all == true:
		Printer(CPU)
		Printer(MEM)
	case *cpu == true:
		Printer(CPU)
	case *mem == true:
		Printer(MEM)
	}
}