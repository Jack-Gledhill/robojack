package debug

import (
	"math"
	"runtime"
	"time"

	"github.com/Jack-Gledhill/robojack/config"
	"github.com/Jack-Gledhill/robojack/log"
)

// Runtime is populated at runtime with information about the program's resource usage
var Runtime = RuntimeInfo{
	LogLevel:  log.Level.String(),
	Mode:      config.Mode(),
	StartTime: time.Now(),
}

// RuntimeSnapshot is a snapshot of the results of RuntimeInfo's functions
// A snapshot can be taken by calling Runtime.Snapshot()
type RuntimeSnapshot struct {
	ActiveGoroutines int           `json:"active_goroutines"`
	GCCycles         uint32        `json:"gc_cycles"`
	GCUsage          int           `json:"gc_usage"`
	LogLevel         string        `json:"log_level"`
	MemoryAllocated  uint64        `json:"memory_allocated"`
	Mode             string        `json:"mode"`
	StartTime        time.Time     `json:"start_time"`
	Uptime           time.Duration `json:"uptime"`
}

// RuntimeInfo holds info about the program's resource usage, uptime, log level and running environment
type RuntimeInfo struct {
	LogLevel  string    `json:"log_level"`
	Mode      string    `json:"mode"`
	StartTime time.Time `json:"start_time"`
}

// Snapshot takes a snapshot of the program's resource usage when called
func (r *RuntimeInfo) Snapshot() RuntimeSnapshot {
	return RuntimeSnapshot{
		ActiveGoroutines: runtime.NumGoroutine(),
		GCCycles:         r.GCCycles(),
		GCUsage:          r.GCUsage(),
		LogLevel:         r.LogLevel,
		MemoryAllocated:  r.MemoryAllocated(),
		Mode:             r.Mode,
		StartTime:        r.StartTime,
		Uptime:           r.Uptime(),
	}
}

// Uptime calculates the duration of time the program has been running for
func (r *RuntimeInfo) Uptime() time.Duration {
	return time.Since(r.StartTime)
}

// ActiveGoroutines returns the number of goroutines that are currently alive
func (r *RuntimeInfo) ActiveGoroutines() int {
	return runtime.NumGoroutine()
}

// GCUsage returns the % of CPU time used by the garbage collector since the program started
func (r *RuntimeInfo) GCUsage() int {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return int(math.Round(m.GCCPUFraction * 100))
}

// GCCycles returns the number of completed garbage collector cycles since the program started
func (r *RuntimeInfo) GCCycles() uint32 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return m.NumGC
}

// MemoryAllocated returns the current size of the program's heap in bytes
func (r *RuntimeInfo) MemoryAllocated() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return m.Sys
}
