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

// Envelope is an envelope
type Envelope interface {
	Attack([]Note) []Note
	Decay([]Note) []Note
	Sustain([]Note) []Note
	Release([]Note) []Note
}

// ADSR is an Envelope implementation
type ADSR struct {
	Attack  float64
	Decay   float64
	Sustain float64
	Release float64
}

// Note is a note
type Note struct {
	Frequency float64
	Duruation NoteDuration
	Volume    float64
}