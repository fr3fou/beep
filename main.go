package main

import (
	"encoding/binary"
	"math"
	"os"
	"time"
)

type Mapper func(x float64) float64

type Offset float64

const (
	A4 Offset = 0
)

const (
	SampleRate float64 = 48000
)

// A is the 1/12th root of 2
var A = math.Pow(2, 1.0/12)

func main() {
	waves := Note(A4, time.Second, 0.05)
	file, err := os.Create("megolovania.bin")
	if err != nil {
		panic(err)
	}

	err = binary.Write(file, binary.LittleEndian, waves)
	if err != nil {
		panic(err)
	}
}

func Note(offset Offset, duration time.Duration, volume float64) []float64 {
	// https://pages.mtu.edu/~suits/NoteFreqCalcs.html
	freq := 440 * math.Pow(A, float64(offset))

	return MapVariadic(
		Range(1, SampleRate*duration.Seconds()),
		Multiply(freq*2*math.Pi/SampleRate),
		math.Sin,
		Multiply(volume),
	)
}

func Flatten(input [][]float64) []float64 {
	output := make([]float64, len(input)*len(input[0]))

	i := 0

	for _, row := range input {
		for _, col := range row {
			output[i] = col
			i++
		}
	}

	return output
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

func MapVariadic(arr []float64, mappers ...Mapper) []float64 {
	output := make([]float64, len(arr))

	copy(output, arr)

	for _, mapper := range mappers {
		output = Map(output, mapper)
	}

	return output
}

func Map(arr []float64, f Mapper) []float64 {
	output := make([]float64, len(arr))

	for i, elem := range arr {
		output[i] = f(elem)
	}

	return output
}
