package main

import (
	"fmt"

	"github.com/chartleyit/goplay/adventofcode/util"
)

type Bike struct {
	freq   int
	hist   []int
	offset int
}

// search recursively for duplicates
func search(input []int) int {
	if len(input) > 2 {
		for _, v := range input[1:] {
			if input[0] == v {
				return v
			}
		}
		return search(input[1:])
	}
	return 0
}

// calibrate runs list of frequency adjustments and returns frequency
func (b *Bike) calibrate(c []int) {
	// fmt.Println("Calibrations to execute", len(c))
	for _, adj := range c {
		b.hist = append(b.hist, b.freq)
		b.freq = b.freq + adj
		offset := search(b.hist)
		if offset != 0 {
			b.offset = offset
			return
		}
	}
}

func main() {
	// recalibrate until frequency is repeated
	calibrations := util.ReadData()
	// calibrations = []int{+1, -2, +3, +1, +1, -2}

	bike := &Bike{freq: 0}
	fmt.Println("Biginning bike calibration")
	fmt.Println("Calibrations to execute", len(calibrations))
	for i := 0; i < 10; i++ {
		bike.calibrate(calibrations)
		fmt.Println("Initial frequency", bike.freq)
		// fmt.Println("Frequency history", bike.hist)
		// fmt.Println("Calibration offset", bike.offset)
		if bike.offset != 0 {
			fmt.Println("Calibration offset", bike.offset)
			fmt.Println("Bike calibration completed")
			break
		}
		fmt.Println("Retrying calibration")
	}
}
