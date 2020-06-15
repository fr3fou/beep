package main

import (
	"encoding/binary"
	"math"
	"os"
	"time"
)

type (
	// NoteDuration is the type for the duration of a single note (SampleRate times samples)
	NoteDuration = time.Duration
	// Sample represents wave data at a particular point in time
	// 48000 Samples make up 1s of audio data at a sample rate of 48000
	Sample = float64
	// Offset is the offset of the notes with relation to A440 -> A4
	Offset float64
	// Mapper is a function type given to a mapping function which applies it to a slice of floats
	Mapper func(x float64) float64
)

// Some predefined note offsets in the 4th and 5th octave with relation to A4
const (
	C3  Offset = -21
	CS3 Offset = -20
	D3  Offset = -19
	DS3 Offset = -18
	E3  Offset = -17
	F3  Offset = -16
	FS3 Offset = -15
	G3  Offset = -14
	GS3 Offset = -13
	A3  Offset = -12
	AS3 Offset = -11
	B3  Offset = -10

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
	// SampleRate is the rate at which we sample the resultant binary musicfile
	SampleRate = 48000.0
	// BPM is the tempo of the notes we define to generate the file
	BPM = 120.0
	// BeatNoteLength is derived from the time signature (in our case 4/4)
	// e.g. for the 7/8 time signature, our note length will be 8 (i.e. an eigth note)
	// so since we're basing our definitions by a semibreve we multiple by the reverse -> 1 / notelength.
	BeatNoteLength = 4
)

// A is the 1/12th root of 2
var A = math.Pow(2, 1.0/12)

// AngryBirds is the first 3 measures of a simplified variation of the original Angry Birds theme
func AngryBirds() []Sample {
	return Flatten(
		// First measure without pauses in the beginning
		Note(E4, Semiquaver, 0.05),
		Note(FS4, Semiquaver, 0.05),
		// Second measure
		Note(G4, Quaver, 0.05),
		Note(E4, Quaver, 0.05),
		Note(B4, Quaver, 0.05),
		Note(E4, Semiquaver, 0.05),
		Note(FS4, Semiquaver, 0.05),
		Note(G4, Quaver, 0.05),
		Note(B4, Quaver, 0.05),
		Note(B4, Crotchet, 0.05),
		// Third measure without tailing pause and 2 notes
		Note(B4, Semiquaver, 0.05),
		Note(C5, Semiquaver, 0.05),
		Note(B4, Semiquaver, 0.05),
		Note(A4, Semiquaver, 0.05),
		Note(G4, Quaver, 0.05),
		Note(G4, Semiquaver, 0.05),
		Note(FS4, Semiquaver, 0.05),
		Note(E4, Crotchet, 0.05),
	)
}

func Megalovania() []Sample {
	return Flatten(
		Note(D3, Semiquaver, 0.05),
		Note(D3, Semiquaver, 0.05),
		Note(D4, Semiquaver, 0.05),
		Note(A3, Quaver, 0.05),
		Pause(Semiquaver),
		Note(GS3, Semiquaver, 0.05),
		Pause(Semiquaver),
		Note(G3, Semiquaver, 0.05),
		Pause(Semiquaver),
		Note(F3, Quaver, 0.05),
		Note(D3, Semiquaver, 0.05),
		Note(F3, Semiquaver, 0.05),
		Note(G3, Semiquaver, 0.05),
		Pause(Semiquaver),
		Note(B4, Semiquaver, 0.05),
		Note(C5, Semiquaver, 0.05),
		Note(B4, Semiquaver, 0.05),
		Note(A4, Semiquaver, 0.05),
		Note(G4, Quaver, 0.05),
		Note(G4, Semiquaver, 0.05),
		Note(FS4, Semiquaver, 0.05),
		Note(E4, Crotchet, 0.05),
	)
}

func main() {
	waves := Megalovania()
	file, err := os.Create("megolovania.bin")
	if err != nil {
		panic(err)
	}

	err = binary.Write(file, binary.LittleEndian, waves)
	if err != nil {
		panic(err)
	}
}

// Pause generates a silent samples array of a given duration
func Pause(duration NoteDuration) []Sample {
	return Note(0, duration, 0)
}

// Note generates an array of samples representing a sine wave oscilating at a specific note frequency
// for a given duration of time at a given volume
func Note(offset Offset, duration NoteDuration, volume float64) []Sample {
	// https://pages.mtu.edu/~suits/NoteFreqCalcs.html
	freq := 440 * math.Pow(A, float64(offset))

	return Envelope(
		MapVariadic(
			Range(1, SampleRate*duration.Seconds()),
			Multiply(freq*2*math.Pi/SampleRate),
			math.Sin,
			Multiply(volume),
		), 0.30, 0.10, 0.30)
}

// Envelope applies ADSR (Attack Decay Sustain Release) on an array of samples
func Envelope(wave []Sample, attack, decay, sustain float64) []Sample {
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

// Flatten merges a slice of slices of floats to a single "flat"
// slice of floats
func Flatten(input ...[]float64) []float64 {
	output := []float64{}

	for _, row := range input {
		output = append(output, row...)
	}

	return output
}

// Multiply returns a mapper function which multiplies by a given constant to then be used
// in Map and MapVariatic
func Multiply(c float64) Mapper {
	return func(x float64) float64 {
		return x * c
	}
}

// RangeStep generates a float slice containing the range low - high with a given step
func RangeStep(low, high, step float64) []float64 {
	output := make([]float64, int(high-low+1))
	index := 0

	for i := low; i <= high; i += step {
		output[index] = i
		index++
	}

	return output
}

// Range generates a float slice containing the range low - high
func Range(low, high float64) []float64 {
	return RangeStep(low, high, 1)
}

// MapVariadic sequentially applies all Mapper functions on all elements of the float slice arr
func MapVariadic(arr []float64, mappers ...Mapper) []float64 {
	output := make([]float64, len(arr))

	copy(output, arr)

	for _, mapper := range mappers {
		output = Map(output, mapper)
	}

	return output
}

// Map applies a mapper function on all elements of the float slice arr
func Map(arr []float64, f Mapper) []float64 {
	output := make([]float64, len(arr))

	for i, elem := range arr {
		output[i] = f(elem)
	}

	return output
}
