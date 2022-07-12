package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

/* Get the virual memory */
func getVirtualMemory() string {
	v, _ := mem.VirtualMemory()

	total, used, available := v.Total/1000000, v.Used/1000000, v.Available/1000000

	memorySummary := fmt.Sprintf(" total: %d mb, \n Used: %d mb, \n Available: %d mb", total, used, available)

	return memorySummary
}

/* Get # cores and percentage used */
func getCPU() string {
	infoIndex, _ := cpu.Info()

	coreTime, _ := cpu.Percent(0, true)

	infoStat := cpu.InfoStat(infoIndex[0])

	cpuSummary := fmt.Sprintf(`	Model: %s 
	# of Cores: %d 
	Usage: %f %%`,
		infoStat.ModelName,
		infoStat.Cores, coreTime[0])

	return cpuSummary
}
