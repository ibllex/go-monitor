package main

import (
	"github.com/gin-gonic/gin"
	fgin "github.com/ibllex/go-fractal/gin"
)

// GetSystemInfo get realtime system info
func GetSystemInfo(c *fgin.Context) {

	mem := MemeoryInfo()

	i := &SystemInfo{
		Disks:    DisksInfo(),
		MemTotal: mem.Total,
		MemUsed:  mem.Used,
		CpuCount: CpuCount(),
		CpuUsage: CpuUsage(),
		Load:     SystemLoad().Percent,
	}

	c.Item(i, NewSystemInfoTransformer())
}

func InitRouter(base string, r *gin.Engine) *gin.Engine {
	v1 := r.Group(base + "/api/v1")
	v1.GET("/system", fgin.H(GetSystemInfo))

	return r
}
