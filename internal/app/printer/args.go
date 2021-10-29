package printer

import (
	"flag"
)

const (
	ALL  = "ALL"
	CPU  = "CPU"
	MEM  = "MEM"
	DISK = "DISK"
)

func Setup() {
	all := flag.Bool("a", false, " : All Information")
	cpu := flag.Bool("c", false, " : CPU Information")
	mem := flag.Bool("m", false, " : Memory Information")
	disk := flag.Bool("d", false, " : Disk Information")

	flag.Parse()

	argsCheckup(all, cpu, mem, disk)
}

func argsCheckup(all, cpu, mem, disk *bool) {
	switch {
	case *all == true:
		Printer().Print(ALL)
	case *cpu == true:
		Printer().Print(CPU)
	case *mem == true:
		Printer().Print(MEM)
	case *disk == true:
		Printer().Print(DISK)
	}
}
