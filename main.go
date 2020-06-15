package main

import (
	"encoding/binary"
	"math"
	"os"
	"time"
)

type (
	NoteDuration = time.Duration
	Offset       float64
	Mapper       func(x float64) float64
)

const (
	C4  Offset = -9
	CS4 Offset = -8
	D4  Offset = -7
	DS4 Offset = -6
	E4  Offset = -5
	F4  Offset = -4
	FS4 Offset = -3
	G4  Offset = -2
	GS4 Offset = -1
	A4  Offset = 0
	AS4 Offset = 1
	B4  Offset = 2
	//
	C5 Offset = 3
)

// We define the length of a beat by the time signature (4/4 in the case of Angry Birds and Megalovania)
// We do that by determining which note length will be = 1 beat, e.g. 4/4 -> (4) means Crotchet so
// 1 Crotchet = 1 Beat => 1 Semibreve = 4 beats
const (
	Semibreve      NoteDuration = NoteDuration(time.Minute / BPM * BeatNoteLength)
	Minim          NoteDuration = Semibreve / 2
	Crotchet       NoteDuration = Semibreve / 4
	Quaver         NoteDuration = Semibreve / 8
	Semiquaver     NoteDuration = Semibreve / 16
	Demisemiquaver NoteDuration = Semibreve / 32
)

const (
	SampleRate = 48000.0
	BPM        = 60.0
	// BeatNoteLength is derived from the time signature (in our case 4/4)
	// e.g. for the 7/8 time signature, our note length will be 8 (i.e. an eigth note)
	// so since we're basing our definitions by a semibreve we multiple by the reverse -> 1 / notelength.
	BeatNoteLength = 4
)

// A is the 1/12th root of 2
var A = math.Pow(2, 1.0/12)

func main() {
	waves := Flatten(
		Note(E4, Semiquaver, 0.05),
		Note(FS4, Semiquaver, 0.05),
		Note(G4, Quaver, 0.05),
		Note(E4, Quaver, 0.05),
		Note(B4, Quaver, 0.05),
		Note(E4, Semiquaver, 0.05),
		Note(FS4, Semiquaver, 0.05),
		Note(G4, Quaver, 0.05),
		Note(B4, Quaver, 0.05),
		Note(B4, Crotchet, 0.05),
		//
		Note(B4, Semiquaver, 0.05),
		Note(C5, Semiquaver, 0.05),
		Note(B4, Semiquaver, 0.05),
		Note(A4, Semiquaver, 0.05),
		Note(G4, Quaver, 0.05),
		Note(G4, Semiquaver, 0.05),
		Note(FS4, Semiquaver, 0.05),
		Note(E4, Crotchet, 0.05),
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

func Pause(duration NoteDuration) []float64 {
	return Note(0, duration, 0)
}

func Note(offset Offset, duration NoteDuration, volume float64) []float64 {
	// https://pages.mtu.edu/~suits/NoteFreqCalcs.html
	freq := 440 * math.Pow(A, float64(offset))

	return MapVariadic(
		Range(1, SampleRate*duration.Seconds()),
		Multiply(freq*2*math.Pi/SampleRate),
		math.Sin,
		Multiply(volume),
	)
}

func Envelope(wave []float64, attack, decay, sustain float64) []float64 {
	attackEnd := int(attack * float64(len(wave)))            // index of attackEnd
	decayEnd := int(decay*float64(len(wave))) + attackEnd    // index of decayEnd
	sustainEnd := int(sustain*float64(len(wave))) + decayEnd // index of sustainEnd

	releaseLength := len(wave) - sustainEnd

	i := 0.0

	attackStep := 1.35 / float64(attackEnd)
	attackPart := Map(wave[:attackEnd], func(x float64) float64 {
		val := x * i
		i += attackStep
		return val
	})

	decayStep := 0.35 / float64(decayEnd-attackEnd)
	decayPart := Map(wave[attackEnd:decayEnd], func(x float64) float64 {
		val := x * i
		i -= decayStep
		return val
	})

	sustainPart := Map(wave[decayEnd:sustainEnd], func(x float64) float64 {
		return x
	})

	releaseStep := 1.0 / float64(releaseLength)
	releasePart := Map(wave[sustainEnd:], func(x float64) float64 {
		val := x * i
		i -= releaseStep
		return val
	})

	return Flatten(
		attackPart,
		decayPart,
		sustainPart,
		releasePart,
	)
}

func Flatten(input ...[]float64) []float64 {
	output := []float64{}

	for _, row := range input {
		output = append(output, row...)
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
