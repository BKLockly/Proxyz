/*
 * @Author: Lockly
 * @Date: 2025-02-16 17:35:06
 * @LastEditors: Lockly
 * @LastEditTime: 2025-02-18 18:10:44
 */

package info

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type ProxyMetrics struct {
	mu   sync.RWMutex
	stop chan struct{}

	cpuUsage   float64
	memUsageMB float64
}

func NewProxyMetrics(stopChan chan struct{}) *ProxyMetrics {
	return &ProxyMetrics{
		stop: stopChan,
	}
}

func (pm *ProxyMetrics) StopMonitoring() {
	close(pm.stop)
}

func (pm *ProxyMetrics) StartMonitoring(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			percent, _ := cpu.Percent(time.Second, false)
			if len(percent) > 0 {
				pm.mu.Lock()
				pm.cpuUsage = percent[0]
				pm.mu.Unlock()
			}

			var m runtime.MemStats
			runtime.ReadMemStats(&m)

			pm.mu.Lock()
			pm.memUsageMB = float64(m.Alloc) / 1024 / 1024
			pm.mu.Unlock()
		case <-pm.stop:
			return
		}
	}
}

func (pm *ProxyMetrics) GetMetrics() map[string]string {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	v, _ := mem.VirtualMemory()

	return map[string]string{
		"cpu_usage":       fmt.Sprintf("%.2f%%", pm.cpuUsage),
		"mem_usage_self":  fmt.Sprintf("%.2f MB", pm.memUsageMB),
		"mem_usage_total": fmt.Sprintf("%.2f%%", v.UsedPercent),
	}
}
