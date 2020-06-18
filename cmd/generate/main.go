package main

import (
	"math"

	"github.com/fr3fou/mego/mego"
)

func main() {
	sampleRate := 48000
	bpm := 120
	noteLength := 4 // 4/4, therefore 4
	volume := 0.05

	m := mego.NewMelody(
		sampleRate,
		bpm,
		noteLength,
		math.Sin,
		mego.NewADSR(0.30, 0.10, 0.30, 0.30),
		// variadic optional notes to begin with
	)

	m.AddNotes(
		mego.A(4, mego.Quaver, volume),
		mego.B(4, mego.Quaver, volume),
		mego.DS(4, mego.Quaver, volume),
		mego.FS(4, mego.Quaver, volume),
		mego.E(5, mego.Quaver, volume),
		mego.A(4, mego.Quaver, volume),
	)

	pcm := m.PCM()       // returns the PCM array
	err := m.Write(file) // write PCM to an io.Writer

	err := wav.Write(m, file) // write wav to an io.writer
}
