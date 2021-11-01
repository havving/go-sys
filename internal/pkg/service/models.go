package service

type SysModel struct {
	Host HostModel
	Cpu  CpuSysModel
	Mem  MemSysModel
	Swap SwapSysModel
}

type HostModel struct {
	Host     string
	HostName string
}

type CpuSysModel struct {
	Idle  uint64
	Sys   uint64
	User  uint64
	Wait  uint64
	Total uint64
}

type MemSysModel struct {
	Used       uint64
	Free       uint64
	ActualUsed uint64
	ActualFree uint64
	Total      uint64
}

type SwapSysModel struct {
	Used  uint64
	Free  uint64
	Total uint64
}
