package main

import (
	"errors"
	"fmt"

	"github.com/chartleyit/goplay/adventofcode/util"
)

type Bike struct {
	freq int
	// hist   []int
	offset *Offset
}

type Offset struct {
	key   int
	left  *Offset
	right *Offset
}

var count int

// insert will add values to offset tree
func (o *Offset) insert(k int) error {
	var err error
	count++
	if o.key < k {
		// move right
		if o.right == nil {
			o.right = &Offset{key: k}
		} else {
			err = o.right.insert(k)
		}
	} else if o.key > k {
		// move left
		if o.left == nil {
			o.left = &Offset{key: k}
		} else {
			err = o.left.insert(k)
		}
	} else {
		// fmt.Println("Found duplicate!")
		err = errors.New("Duplicate Found")
	}
	return err
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
func (b *Bike) calibrate(c []int) (int, error) {
	// fmt.Println("Calibrations to execute", len(c))
	for _, adj := range c {
		// b.hist = append(b.hist, b.freq)
		b.freq = b.freq + adj
		err := b.offset.insert(b.freq)
		if err != nil {
			return b.freq, err
		}
	}
	return 0, nil
}

func main() {
	// recalibrate until frequency is repeated
	calibrations := util.ReadData()
	// calibrations = []int{+1, -2, +3, +1, +1, -2}

	bike := &Bike{freq: 0, offset: &Offset{key: 0}}
	fmt.Println("Biginning bike calibration")
	fmt.Println("Calibrations to execute", len(calibrations))
	retry := 300
	attempt := 0
	for i := 0; i < retry; i++ {
		attempt++
		fmt.Println("Initial frequency", bike.freq)
		f, err := bike.calibrate(calibrations)
		// fmt.Println("Frequency history", bike.hist)
		// fmt.Println("Calibration offset", bike.offset)
		if err != nil {
			fmt.Println("Calibration offset", f)
			fmt.Println("Bike calibration completed")
			break
		}
		fmt.Println("Ending frequency", bike.freq)
		fmt.Println("Retrying calibration...")
	}
	fmt.Printf("\nAttempted calibrations %d\n", attempt)
}
