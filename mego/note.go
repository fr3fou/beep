package mego

import "math"

// NewNote is a constructor for a Note
func NewNote(step Step, duration NoteDuration, volume float64) Note {
	freq := 440 * math.Pow(a, float64(step))

	return Note{
		Frequency: freq,
		Duruation: duration,
		Volume:    volume,
	}
}

// Pause is a handy wrapper around NewNote with duration 0 and volume 0
func Pause(duration NoteDuration) Note {
	return NewNote(0, duration, 0)
}
