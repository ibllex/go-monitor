package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Clear history regularly, as we only keep the last 30 days of records
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for range ticker.C {
			CleanRecords(10 * 24 * time.Hour)
		}
	}()

	// Init the web router
	routeBase := "/monitor"

	r := InitRouter(routeBase, gin.Default())

	r.GET(routeBase, func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "/")
	})

	// Run monitor task
	m := NewMonitor(1 * time.Second)
	m.Run(context.Background())

	// Start the http listener
	http.ListenAndServe(":8080", r)
}
