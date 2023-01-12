package main

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

type DiskInfo struct {
	Device     string
	MountPoint string
	Total      uint64
	Used       uint64
}

// Disks info
func DisksInfo() []*DiskInfo {
	disks := []*DiskInfo{}

	partitions, _ := disk.Partitions(false)
	for _, p := range partitions {
		if strings.Contains(filepath.Base(p.Device), "loop") {
			continue
		}

		usage, _ := disk.Usage(p.Mountpoint)
		disks = append(disks, &DiskInfo{
			Device:     p.Device,
			MountPoint: p.Mountpoint,
			Total:      usage.Total,
			Used:       usage.Used,
		})
	}

	return disks
}

// Get memory info
func MemeoryInfo() *mem.VirtualMemoryStat {
	// memory info
	v, _ := mem.VirtualMemory()
	return v
}

// Get cpu count
func CpuCount() int {
	cpuCount, _ := cpu.Counts(true)
	return cpuCount
}

// Get cpu usage
func CpuUsage() float64 {
	cpuUsage, _ := cpu.Percent(time.Second, false)
	if len(cpuUsage) <= 0 {
		return 0
	}

	return cpuUsage[0]
}

type LoadInfo struct {
	Load1      float64
	Load5      float64
	Load15     float64
	Percentage float64
}

// Get system load
func SystemLoad() *LoadInfo {
	stat, _ := load.Avg()
	return &LoadInfo{
		Load1:      stat.Load1,
		Load5:      stat.Load5,
		Load15:     stat.Load15,
		Percentage: stat.Load1 / (float64(CpuCount()*2) * 0.75) * 100,
	}
}
