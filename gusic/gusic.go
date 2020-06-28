package gusic

import "time"

// Generator is a function for producing samples
type Generator = func(x float64) float64

// NoteDuration is the duration of a single note
type NoteDuration = time.Duration

// Melody is a melody
type Melody struct {
	Notes      []Note
	SampleRate int
	NoteLength int
	BPM        int
	Generator  Generator
	Envelope   ADSR
}

// LinearADSR is an Envelope implementation
type LinearADSR struct {
	ratios            ADSRRatios
	attackMultiplier  float64
	decayMultiplier   float64
	releaseMultiplier float64
	currentMultiplier float64
}

// ADSRRatios is a struct for determining what duration should each state in an ADSR take.
type ADSRRatios struct {
	AttackRatio  float64
	DecayRatio   float64
	SustainRatio float64
	ReleaseRatio float64
}

// ADSR defines what an Envelope should behave like
type ADSR interface {
	Attack(notes []Note)
	Decay(notes []Note)
	Sustain(notes []Note)
	Release(notes []Note)
	GetRatios() *ADSRRatios
}

// Note is a note
type Note struct {
	Frequency float64
	Duration  NoteDuration
	Volume    float64
}
