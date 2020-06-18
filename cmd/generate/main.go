package main

import "math"

func main() {
	sampleRate := 48000
	bpm := 120.0
	noteLength := 4 // 4/4, therefore 4

	m := mego.NewMelody(
		mego.MelodyConfig{
			SampleRate: sampleRate,
			BPM:        bpm,
			NoteLength: noteLength,
			Generator:  math.Sin,
			Envelope:   mego.NewEnvelope(0.30, 0.10, 0.30, 0.30),
		},
		// variadic optional notes to begin with
	)

	m.AddNote(
	// add one note
	)

	m.AddNotes(
	// add variadic notes
	)

	pcm := m.PCM()       // returns the PCM array
	err := m.Write(file) // write PCM to an io.Writer

	err := wav.Write(m, file) // write wav to an io.writer
}
