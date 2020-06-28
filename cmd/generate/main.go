package main

import (
	"math"

	"github.com/fr3fou/gusic/gusic"
)

func main() {
	sampleRate := 48000
	bpm := 120
	noteLength := 4 // 4/4, therefore 4
	volume := 0.05

	m := gusic.NewMelody(
		sampleRate,
		bpm,
		noteLength,
		math.Sin,
		gusic.NewLinearADSR(gusic.NewRatios(0.25, 0.25, 0.25, 0.25), 1.35, 0.35, 1.0),
	)

	m.AddNotes(
		gusic.A(4, m.Quaver, volume),
		gusic.B(4, m.Quaver, volume),
		gusic.DS(4, m.Quaver, volume),
	)

	// pcm := m.PCM()       // returns the PCM array
	// err := m.Write(file) // write PCM to an io.Writer

	// err := wav.Write(m, file) // write wav to an io.writer
}
