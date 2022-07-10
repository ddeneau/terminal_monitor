package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)


/* Get the virual memory */
func getVirtualMemory() string {
	v, _ := mem.VirtualMemory()

	total, used, available := v.Total, v.Used, v.Available

	memorySummary := fmt.Sprintf("total: %d mb, \n Used: %d mb, \n Available: %d mb", total, used, available)

	return memorySummary
}

/* Get # cores and percentage used */
func getCPU() string {
	v, _ := cpu.Info()

	coreTime, _ := cpu.Percent(0, false)

	cpuInfo := cpu.InfoStat(v[0])

	cpuSummary := fmt.Sprintf("Cores: %d  Used: %f", cpuInfo.Cores, coreTime[0])

	return cpuSummary
}