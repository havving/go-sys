package printer

import (
	"fmt"
	"go-sys/internal/pkg/service"
	"os"
)

type GeneralSysModel struct {
	service.SysModel
}

// Printer 생성자 함수
func Printer() *GeneralSysModel {
	g := &GeneralSysModel{}
	g.GetMem()
	g.GetCpu()

	return g
}

func (g GeneralSysModel) Print(mode string) {
	switch mode {
	case ALL:
		g.printAll()
	case CPU:
		g.printCpu()
	case MEM:
		g.printMem()
	}
	os.Exit(1)
}

func (g GeneralSysModel) printAll() {
	g.printCpu()
	g.printMem()
}

func (g GeneralSysModel) printCpu() {
	fmt.Fprintf(os.Stdout, "%18s %10s %10s %10s %10s\n",
		"Idle", "Sys", "User", "Wait", "Total")

	fmt.Fprintf(os.Stdout, "Cpu:    %10d %10d %10d %10d %10d\n",
		g.Cpu.Idle, g.Cpu.Sys, g.Cpu.User, g.Cpu.Wait, g.Cpu.Total)
}

func (g GeneralSysModel) printMem() {
	fmt.Fprintf(os.Stdout, "\n%18s %10s %10s\n",
		"total", "used", "free")

	fmt.Fprintf(os.Stdout, "Mem:    %10d %10d %10d\n",
		g.Mem.Total, g.Mem.Used, g.Mem.Free)

	fmt.Fprintf(os.Stdout, "-/+ buffers/cache: %10d %10d\n",
		g.Mem.ActualUsed, g.Mem.ActualFree)

	fmt.Fprintf(os.Stdout, "Swap:   %10d %10d %10d\n",
		g.Swap.Total, g.Swap.Used, g.Swap.Free)
}
