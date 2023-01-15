package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// init web router
	routeBase := "/monitor"

	r := InitRouter(routeBase, gin.Default())

	r.GET(routeBase, func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "/")
	})

	// run monitor task
	m := NewMonitor()
	m.Run(context.Background())

	// start http listener
	http.ListenAndServe(":8080", r)
}
