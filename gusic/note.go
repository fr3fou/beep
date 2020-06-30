package gusic

import "math"

// NewNote is a constructor for a Note
func NewNote(step int, duration NoteDuration, volume float64) Note {
	return Note{
		Frequency: 440 * math.Pow(twelfthrootof2, float64(step)),
		Duration:  duration,
		Volume:    volume,
	}
}

// Rest is a handy wrapper around NewNote with the given duration and volume and frequency 0
func Rest(duration NoteDuration) Note {
	return NewNote(0, duration, 0)
}
