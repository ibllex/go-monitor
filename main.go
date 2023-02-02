package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

var (
	cfgFile string
	dbFile  string
	cfg     Config
)

func main() {
	// Parse cmd flags
	flag.StringVar(&cfgFile, "config", "data/monitor.ini", "config file path")
	flag.StringVar(&dbFile, "db", "data/monitor.db", "database file")

	flag.Parse()

	cfg = Config{
		Port: 8080,
		Slug: "monitor",
	}

	// Load config
	ini.MapTo(&cfg, cfgFile)

	cfgHanfler := ini.Empty()
	if err := cfgHanfler.ReflectFrom(&cfg); err != nil {
		log.Fatal(err)
	} else {
		cfgHanfler.SaveTo(cfgFile)
	}

	// Clear history regularly, as we only keep the last 30 days of records
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for range ticker.C {
			CleanRecords(10 * 24 * time.Hour)
		}
	}()

	// Init the web router
	routeBase := "/" + cfg.Slug

	r := InitRouter(routeBase, gin.Default())

	r.GET(routeBase, func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "/")
	})

	// Run monitor task
	m := NewMonitor(1 * time.Minute)
	m.Run(context.Background())

	// Start the http listener
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r); err != nil {
		log.Fatal(err)
	}
}
