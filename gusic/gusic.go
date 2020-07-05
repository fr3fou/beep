package gusic

import "time"

// Generator is a function for producing samples
type Generator = func(x float64) float64

// Easer is an easing function meant to be used in conjunction with an EasedADSR
// Both its range and domain should be in [0, 1]
type Easer = func(t float64) float64

// NoteDuration is the duration of a single note
type NoteDuration = time.Duration

// Run represents a single run (stave) in sheet music, containing notes
type Run struct {
	Notes []Note
}

// Melody is a melody
type Melody struct {
	Runs       []Run
	SampleRate float64
	NoteLength int
	BPM        int
	Generator  Generator
	Envelope   ADSR

	breve          NoteDuration
	semibreve      NoteDuration
	minim          NoteDuration
	crotchet       NoteDuration
	quaver         NoteDuration
	semiquaver     NoteDuration
	demisemiquaver NoteDuration
}

// EasedADSR is an eased Envelope implementation
type EasedADSR struct {
	easing            Easer
	ratios            ADSRRatios
	attackMultiplier  float64
	decayMultiplier   float64
	sustainMultiplier float64
}

// ADSRRatios is a struct for determining what duration should each state in an ADSR take.
type ADSRRatios struct {
	AttackRatio  float64
	DecayRatio   float64
	SustainRatio float64
	ReleaseRatio float64
}

// SingleNote is a note
type SingleNote struct {
	Frequency float64
	Duration  NoteDuration
	Volume    float64
}

// Chord is a group of notes that are to be played at the same time
type Chord []SingleNote

// ADSR defines what an Envelope should behave like
type ADSR interface {
	Attack(samples []float64)
	Decay(samples []float64)
	Sustain(samples []float64)
	Release(samples []float64)
	Ratios() ADSRRatios
}

// Note is an interface for returning samples
type Note interface {
	Samples(sampleRate float64, generator Generator, adsr ADSR) []float64
}

// ApplyADSR applies all the stages of an ADSR to an array of notes
func ApplyADSR(noteSamples []float64, envelope ADSR) {
	length := len(noteSamples)

	ratios := envelope.Ratios()

	attackEnd := int(ratios.AttackRatio * float64(length))
	decayEnd := int(ratios.DecayRatio*float64(length)) + attackEnd
	sustainEnd := int(ratios.SustainRatio*float64(length)) + decayEnd

	envelope.Attack(noteSamples[:attackEnd])
	envelope.Decay(noteSamples[attackEnd:decayEnd])
	envelope.Sustain(noteSamples[decayEnd:sustainEnd])
	envelope.Release(noteSamples[sustainEnd:])
}
