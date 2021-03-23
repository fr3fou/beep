package gusic

import (
	"math"
)

// NewNote is a constructor for a Note
func NewNote(step int, duration NoteDuration, volume float64) SingleNote {
	return SingleNote{
		Frequency: 440 * math.Pow(twelfthrootof2, float64(step)),
		Duration:  duration,
		Volume:    volume,
	}
}

// Samples returns the samples
func (note SingleNote) Samples(sampleRate float64, generator Generator, envelope ADSR) []float64 {
	noteSamples := []float64{}

	duration := math.Ceil(sampleRate * note.Duration.Seconds())
	for j := 1.0; j <= duration; j++ {
		t := 2 * math.Pi * j * (note.Frequency / sampleRate)
		noteSamples = append(noteSamples, note.SampleAt(t, sampleRate, generator))
	}
	ApplyADSR(noteSamples, envelope)
	return noteSamples
}

// SampleAt returns the sample at a given point of `t`
// `t` is how far a wave is in its cycle (phase)
func (note SingleNote) SampleAt(t float64, sampleRate float64, generator Generator) float64 {
	return generator(t) * note.Volume
}

// Rest is a handy wrapper around NewNote with the given duration and volume and frequency 0
func Rest(duration NoteDuration) SingleNote {
	return NewNote(0, duration, 0)
}
