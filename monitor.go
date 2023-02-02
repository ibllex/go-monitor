package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Monitor struct {
	stop func()
	dur  time.Duration
}

func NewMonitor(dur time.Duration) *Monitor {
	return &Monitor{dur: dur}
}

// Run monitor
func (m *Monitor) Run(ctx context.Context) {
	// create a cancelable context
	ctx, m.stop = context.WithCancel(ctx)
	// start background task
	go func() {
		// start loop
		for {
			select {
			case <-ctx.Done():
				// end loop
				return
			case <-m.do(ctx): // panic recovery
			}
		}
	}()
}

// Stop monitor
func (m *Monitor) Stop() {
	if m.stop != nil {
		m.stop()
	}
}

func (m *Monitor) do(ctx context.Context) <-chan error {
	errCh := make(chan error)
	go func() {
		t := time.NewTicker(m.dur)

		defer func() {
			t.Stop()

			if r := recover(); r != nil {
				errCh <- fmt.Errorf("panic with error: %v", r)
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				m.next()
			}
		}
	}()

	return errCh
}

// Get & save system information
func (m *Monitor) next() {
	r := SaveRecord(&Record{
		LoadPercent:     SystemLoad().Percent,
		CpuUsagePercent: CpuUsage(),
		MemUsagePercent: MemeoryInfo().UsedPercent,
	})

	if r.Error != nil {
		log.Print(r.Error)
	}
}
