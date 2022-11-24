package metrics

import (
	"fmt"
	"sync"
	"time"
)

type ValueWithUnit struct {
	Value uint64 `json:"value"`
	Unit  string `json:"unit"`
}

type MetricsCollection struct {
	Uptime        Uptime
	countersMap   map[string]uint64
	countersMutex sync.RWMutex
}

func (mc *MetricsCollection) RegisterCounter(counterName string) {
	mc.countersMutex.Lock()
	defer mc.countersMutex.Unlock()
	if _, ok := mc.countersMap[counterName]; ok {
		panic(fmt.Sprintf("counterName %q already exists", counterName))
	}
	mc.countersMap[counterName] = 0
}

func (mc *MetricsCollection) IncrementCounter(counterName string) {
	mc.countersMutex.Lock()
	defer mc.countersMutex.Unlock()
	mc.countersMap[counterName]++
}

func (mc *MetricsCollection) GetCountersMap() map[string]uint64 {
	mc.countersMutex.RLock()
	defer mc.countersMutex.RUnlock()
	return mc.countersMap
}

func NewMetricsCollection(startTime time.Time) *MetricsCollection {
	return &MetricsCollection{
		Uptime:      NewUpTime(startTime),
		countersMap: make(map[string]uint64),
	}
}
