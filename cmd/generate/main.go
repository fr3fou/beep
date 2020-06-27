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
		gusic.NewADSR(0.30, 0.10, 0.30, 0.30),
		// variadic optional notes to begin with
	)

	m.AddNotes(
		gusic.A(4, mego.Quaver, volume),
		gusic.B(4, mego.Quaver, volume),
		gusic.DS(4, mego.Quaver, volume),
		gusic.FS(4, mego.Quaver, volume),
		gusic.E(5, mego.Quaver, volume),
		gusic.A(4, mego.Quaver, volume),
	)

	// pcm := m.PCM()       // returns the PCM array
	// err := m.Write(file) // write PCM to an io.Writer

	// err := wav.Write(m, file) // write wav to an io.writer
}
