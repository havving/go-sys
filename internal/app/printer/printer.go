package printer

import (
	"fmt"
	"os"
)

func Printer(mode string) {
	switch mode {
	case ALL:
		printAll()
	case CPU:
		printCpu()
	case MEM:
		printMem()
	}
}

func printAll() {
	printCpu()
	printMem()
}

func printCpu() {

}

func printMem() {
	fmt.Fprintf(os.Stdout, "%18s %10s %10s\n",
		"total", "used", "free")
}
