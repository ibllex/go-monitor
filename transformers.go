package main

import "github.com/ibllex/go-fractal"

type SystemInfo struct {
	Disks    []*DiskInfo
	MemTotal uint64
	MemUsed  uint64
	CpuCount int
	CpuUsage float64
	Load     float64
}

type SystemInfoTransformer struct {
	*fractal.BaseTransformer
}

func (t *SystemInfoTransformer) Transform(data fractal.Any) fractal.M {
	result := fractal.M{}

	if s := t.toSystemInfo(data); s != nil {
		result["disks"] = s.Disks
		result["mem_total"] = s.MemTotal
		result["mem_used"] = s.MemUsed
		result["cpu_count"] = s.CpuCount
		result["cpu_usage"] = s.CpuUsage
		result["load"] = s.Load
	}

	return result
}

func (t *SystemInfoTransformer) toSystemInfo(data fractal.Any) *SystemInfo {

	switch b := data.(type) {
	case *SystemInfo:
		return b
	case SystemInfo:
		return &b
	}

	return nil
}

func NewSystemInfoTransformer() *SystemInfoTransformer {
	return &SystemInfoTransformer{&fractal.BaseTransformer{}}
}
