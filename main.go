package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// run monitor task
	m := NewMonitor()
	m.Run(context.Background())

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "/")
	})

	// start http listener
	http.ListenAndServe(":8080", r)
}
