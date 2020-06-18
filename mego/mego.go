package mego

import "time"

// Generator is a function for producing samples
type Generator = func(x float64) float64

// NoteDuration is the duration of a single note
type NoteDuration = time.Duration

// Melody is a melody
type Melody struct {
	Note       []Note
	SampleRate int
	NoteLength int
	BPM        int
	Generator  *Generator
	Envelope   *Envelope
}

// Envelope is an ADSR configuration
type Envelope struct {
	Attack  float64
	Decay   float64
	Sustain float64
	Release float64
}

// Note is a note
type Note struct {
	Frequency int
	Duruation NoteDuration
}
