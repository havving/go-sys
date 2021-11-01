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
	all := flag.Bool("a", false, " : All Information(default: false)")
	cpu := flag.Bool("c", false, " : CPU Information(default: false)")
	mem := flag.Bool("m", false, " : Memory Information(default: false)")

	flag.Parse()

	argsCheckup(all, cpu, mem)
}

func argsCheckup(all, cpu, mem *bool) {
	switch {
	case *all == true:
		Printer().Print(ALL)
	case *cpu == true:
		Printer().Print(CPU)
	case *mem == true:
		Printer().Print(MEM)
	default:
		break
	}
}
