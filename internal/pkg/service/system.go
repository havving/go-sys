package service

import (
	sigar "github.com/cloudfoundry/gosigar"
	"go-sys/internal/pkg/util"
)

func (s *SysModel) GetCpu() {
	cpu := sigar.Cpu{}
	cpu.Get()

	s.Cpu.Idle = util.Format(cpu.Idle)
	s.Cpu.Sys = util.Format(cpu.Sys)
	s.Cpu.User = util.Format(cpu.User)
	s.Cpu.Wait = util.Format(cpu.Wait)
	s.Cpu.Total = util.Format(cpu.Total())
}

func (s *SysModel) GetMem() {
	mem := sigar.Mem{}
	swap := sigar.Swap{}

	mem.Get()
	swap.Get()

	s.Mem.Used = util.Format(mem.Used)
	s.Mem.Free = util.Format(mem.Free)
	s.Mem.ActualUsed = util.Format(mem.ActualUsed)
	s.Mem.ActualFree = util.Format(mem.ActualFree)
	s.Mem.Total = util.Format(mem.Total)

	s.Swap.Used = util.Format(swap.Used)
	s.Swap.Free = util.Format(swap.Free)
	s.Swap.Total = util.Format(swap.Total)
}
