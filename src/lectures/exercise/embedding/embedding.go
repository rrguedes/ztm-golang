//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import "fmt"

type Bytes int
type Celcius float32

// Interfaces
type Measurer interface {
	GetAverage() int
}

// Structs
type BandwidthUsage struct {
	amount []Bytes
}
type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

// Methods
func (b *BandwidthUsage) GetAverage(seconds int) int {
	sum := 0
	for i := 0; i < seconds; i++ {
		current := b.amount[i]
		sum += int(current)
	}
	return sum / seconds
}

func (c *CpuTemp) GetAverage(seconds int) int {
	sum := 0
	for i := 0; i < seconds; i++ {
		current := c.temp[i]
		sum += int(current)
	}
	return sum / seconds
}

func (m *MemoryUsage) GetAverage(seconds int) int {
	sum := 0
	for i := 0; i < seconds; i++ {
		current := m.amount[i]
		sum += int(current)
	}
	return sum / seconds
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	seconds := 5
	avgBandwitch := bandwidth.GetAverage(seconds)
	avgTemperature := temp.GetAverage(seconds)
	avgMemory := memory.GetAverage(seconds)

	fmt.Printf("--- Average Bandwitch for %v seconds: %v \n", seconds, avgBandwitch)
	fmt.Printf("--- Average Temperature for %v seconds: %v \n", seconds, avgTemperature)
	fmt.Printf("--- Average Memory usage for %v seconds: %v \n", seconds, avgMemory)

}
