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

//
// SystemInfo transformer
//

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
		result["cpu_usage"] = Decimal(s.CpuUsage)
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

//
// Record transformer
//

type RecordTransformer struct {
	*fractal.BaseTransformer
}

func (t *RecordTransformer) Transform(data fractal.Any) fractal.M {
	result := fractal.M{}

	if r := t.toRecord(data); r != nil {
		result["id"] = r.ID
		result["load"] = Decimal(r.LoadPercent)
		result["cpu"] = Decimal(r.CpuUsagePercent)
		result["memory"] = Decimal(r.MemUsagePercent)
		result["date"] = r.CreatedAt.Format("01/02 15:04:05")
	}

	return result
}

func (t *RecordTransformer) toRecord(data fractal.Any) *Record {

	switch b := data.(type) {
	case *Record:
		return b
	case Record:
		return &b
	}

	return nil
}

func NewRecordTransformer() *RecordTransformer {
	return &RecordTransformer{&fractal.BaseTransformer{}}
}
