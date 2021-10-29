package util

const (
	DiskFormat    = "%-15s %4d %4d %5d %4s %-15s\n"
	DiskFormatStr = "%-15s %4s %4s %5s %4s %-15s\n"
)

func MemFormat(val uint64) uint64 {
	return val / 1024
}
