package service

import (
	sigar "github.com/cloudfoundry/gosigar"
	"go-sys/internal/pkg/util"
	"os"
)

func (s *SysModel) GetHost() {
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	if host != "" {
		s.Host.HostName = host
	} else {
		s.Host.HostName = "UnKnown"
	}
}

func (s *SysModel) GetCpu() {
	cpu := sigar.Cpu{}
	cpu.Get()

	s.Cpu.Idle = cpu.Idle
	s.Cpu.Sys = cpu.Sys
	s.Cpu.User = cpu.User
	s.Cpu.Wait = cpu.Wait
	s.Cpu.Total = cpu.Total()
}

func (s *SysModel) GetMem() {
	mem := sigar.Mem{}
	swap := sigar.Swap{}

	mem.Get()
	swap.Get()

	s.Mem.Used = util.MemFormat(mem.Used)
	s.Mem.Free = util.MemFormat(mem.Free)
	s.Mem.ActualUsed = util.MemFormat(mem.ActualUsed)
	s.Mem.ActualFree = util.MemFormat(mem.ActualFree)
	s.Mem.Total = util.MemFormat(mem.Total)

	s.Swap.Used = util.MemFormat(swap.Used)
	s.Swap.Free = util.MemFormat(swap.Free)
	s.Swap.Total = util.MemFormat(swap.Total)
}

func (s *SysModel) GetDisk() {
	fsList := sigar.FileSystemList{}
	fsList.Get()

	for _, fs := range fsList.List {
		dirName := fs.DirName

		usage := sigar.FileSystemUsage{}
		usage.Get(dirName)

		s.Disk.DevName = fs.DevName
		s.Disk.Total = usage.Total
		s.Disk.Used = usage.Used
		s.Disk.Avail = usage.Avail
		s.Disk.DirName = dirName
		s.Disk.UsePercent = sigar.FormatPercent(usage.UsePercent())
	}
}
