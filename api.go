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

func History(c *fgin.Context) {

	records := Records()
	collection := []interface{}{}

	for _, r := range records {
		collection = append(collection, r)
	}

	c.Collection(collection, NewRecordTransformer())
}

func InitRouter(base string, r *gin.Engine) *gin.Engine {
	r.Use(Cors())

	v1 := r.Group(base + "/api/v1")
	v1.GET("/system", fgin.H(GetSystemInfo))
	v1.GET("/history", fgin.H(History))

	return r
}
