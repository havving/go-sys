package service

import (
	"fmt"
	sigar "github.com/cloudfoundry/gosigar"
	"go-sys/internal/models"
	"go-sys/internal/pkg/util"
	"os"
)

func (s *models.SysModel) Cpu() error {
	cpu := sigar.Cpu{}
	cpu.Get()

	s.cpu.Idle = util.Format(cpu.Idle)
	s.cpu.Sys = util.Format(cpu.Sys)
	s.cpu.User = util.Format(cpu.User)
	s.cpu.Wait = util.Format(cpu.Wait)
	s.cpu.Total = util.Format(cpu.Total())

	fmt.Fprintf(os.Stdout, "%18s %10s %10s %10s %10s\n",
		"Idle", "Sys", "User", "Wait", "Total")

	fmt.Fprintf(os.Stdout, "Cpu:    %10d %10d %10d %10d %10d\n",
		util.Format(cpu.Idle), util.Format(cpu.Sys), util.Format(cpu.User), util.Format(cpu.Wait), util.Format(cpu.Total()))

	return nil
}

func (s *SysModel) Mem() error {
	mem := sigar.Mem{}
	swap := sigar.Swap{}

	mem.Get()
	swap.Get()

	s.mem.Used = util.Format(mem.Used)
	s.mem.Free = util.Format(mem.Free)
	s.mem.ActualUsed = util.Format(mem.ActualUsed)
	s.mem.ActualFree = util.Format(mem.ActualFree)
	s.mem.Total = util.Format(mem.Total)

	s.swap.Used = util.Format(swap.Used)
	s.swap.Free = util.Format(swap.Free)
	s.swap.Total = util.Format(swap.Total)

	return nil
}
