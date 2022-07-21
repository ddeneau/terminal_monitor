package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type systemData struct {
	cpuInfoIndex *cpu.InfoStat
}

func initializeSystemData() {
	
}

/* Get the virual memory */
func getVirtualMemory() string {
	v, _ := mem.VirtualMemory()

	total, used, available := v.Total/1000000, v.Used/1000000, v.Available/1000000

	memorySummary := fmt.Sprintf(" total: %d mb, \n Used: %d mb, \n Available: %d mb", total, used, available)

	return memorySummary
}

func getVirtualMemoryTitle() string {

	return "Virtual Memory"
}

/* Get # cores and percentage used */
func getCPU() string {
	infoIndex, _ := cpu.Info()
	coreTime, _ := cpu.Percent(0, true)
	infoStat := cpu.InfoStat(infoIndex[0])
	cores, time := infoStat.Cores, coreTime[0]

	return fmt.Sprintf(`Number of Cores: %d  Usage: %f %%`, cores, time)
}

func getCPUTitle() string {
	infoIndex, _ := cpu.Info()
	infoStat := cpu.InfoStat(infoIndex[0])

	return infoStat.ModelName
}
