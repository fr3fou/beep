package mego

import "math"

// NewNote is a constructor for a Note
func NewNote(step Step, duration NoteDuration, volume float64) *Note {
	freq := 440 * math.Pow(A, float64(offset))

	return &Note{
		Frequency: freq,
		Duruation: duration,
		Volume:    volume,
	}
}
