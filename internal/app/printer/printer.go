package printer

import (
	"fmt"
	"go-sys/internal/pkg/service"
	"go-sys/internal/pkg/util"
	"os"
)

type GeneralSysModel struct {
	service.SysModel
}

// 초기화
//func (g *GeneralSysModel) initialize() {
//	//err := g.GetMem()
//	err := service.GetMem()
//	if err != nil {
//		panic(err)
//	}
//}

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
		printAll(g)
	case CPU:
		g.printCpu()
	case MEM:
		g.printMem()
	}
}

func printAll(g GeneralSysModel) {
	g.printCpu()
	g.printMem()
}

func (g GeneralSysModel) printCpu() {
	fmt.Fprintf(os.Stdout, "%18s %10s %10s %10s %10s\n",
		"Idle", "Sys", "User", "Wait", "Total")

	fmt.Fprintf(os.Stdout, "Cpu:    %10d %10d %10d %10d %10d\n",
		util.Format(g.Cpu.Idle), util.Format(g.Cpu.Sys), util.Format(g.Cpu.User), util.Format(g.Cpu.Wait), util.Format(g.Cpu.Total))
}

func (g GeneralSysModel) printMem() {
	fmt.Fprintf(os.Stdout, "\n%18s %10s %10s\n",
		"total", "used", "free")

	fmt.Fprintf(os.Stdout, "Mem:    %10d %10d %10d\n",
		util.Format(g.Mem.Total), util.Format(g.Mem.Used), util.Format(g.Mem.Free))

	fmt.Fprintf(os.Stdout, "-/+ buffers/cache: %10d %10d\n",
		util.Format(g.Mem.ActualUsed), util.Format(g.Mem.ActualFree))

	fmt.Fprintf(os.Stdout, "Swap:   %10d %10d %10d\n",
		util.Format(g.Swap.Total), util.Format(g.Swap.Used), util.Format(g.Swap.Free))

}
