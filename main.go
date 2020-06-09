package main

import (
	"encoding/binary"
	"math"
	"os"
)

type Mapper func(x float64) float64

var (
	SampleRate float64 = 48000
	Volume     float64 = 0.05
)

func main() {
	waves :=
		Map(
			Map(
				Map(
					Range(1, SampleRate),
					Multiply(0.2),
				),
				math.Sin,
			),
			Multiply(Volume),
		)

	file, err := os.Create("megolovania.bin")
	if err != nil {
		panic(err)
	}

	err = binary.Write(file, binary.LittleEndian, waves)
	if err != nil {
		panic(err)
	}
}

func Multiply(c float64) Mapper {
	return func(x float64) float64 {
		return x * c
	}
}

func RangeStep(low, high, step float64) []float64 {
	output := make([]float64, int(high-low+1))
	index := 0
	for i := low; i <= high; i += step {
		output[index] = i
		index++
	}

	return output
}

func Range(low, high float64) []float64 {
	return RangeStep(low, high, 1)
}

func Map(arr []float64, f Mapper) []float64 {
	output := make([]float64, len(arr))

	for i, elem := range arr {
		output[i] = f(elem)
	}

	return output
}
