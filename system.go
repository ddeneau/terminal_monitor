package main

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

// Stores data about RAM.
type RAM struct {
	total    	float32
	used      	float32
	available 	float32
}

/* Get the virual memory */
func getVirtualMemory() RAM {
	v, _ := mem.VirtualMemory()
	total, used, available := float32(v.Total), float32(v.Used), float32(v.Available)
	ramModel := RAM{}
	
	ramModel.total = total / (10e8)
	ramModel.used = used / (10e8)
	ramModel.available = available / (10e8)

	return ramModel
}

/* Get # cores and percentage used
returns a int32 and float64, so uses two return values instead of a struct. */
func getCPU() (int32, float64) {
	infoIndex, _ := cpu.Info()
	coreTime, _ := cpu.Percent(0, true)
	infoStat := cpu.InfoStat(infoIndex[0])
	cores, time := infoStat.Cores, coreTime[0]

	return cores, time
}

/* Get model name of CPU */
func getCPUTitle() string {
	infoIndex, _ := cpu.Info()
	infoStat := cpu.InfoStat(infoIndex[0])

	return infoStat.ModelName
}
