package main

import (
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error

	db, err = gorm.Open(sqlite.Open("data/monitor.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Record{})
}

// Record monitor data
type Record struct {
	ID              uint `gorm:"primarykey"`
	LoadPercent     float64
	CpuUsagePercent float64
	MemUsagePercent float64
	CreatedAt       time.Time `gorm:"index"`
}

// Save monitor record
func SaveRecord(r *Record) *gorm.DB {
	return db.Create(r)
}

// Get all records
func Records() (records []*Record) {
	db.Find(&records)
	return
}

func CleanRecords(keepDays time.Duration) *gorm.DB {
	return db.Where("created_at < ?", time.Now().Add(-1*keepDays)).Delete(&Record{})
}
