package gusic

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
	Envelope   *ADSR
}

// ADSR is an Envelope implementation
type ADSR struct {
	attack      float64
	decay       float64
	susain      float64
	release     float64
	currentStep float64
}

// Note is a note
type Note struct {
	Frequency float64
	Duration  NoteDuration
	Volume    float64
}