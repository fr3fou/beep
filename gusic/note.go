package gusic

import "math"

// NewNote is a constructor for a Note
func NewNote(step int, duration NoteDuration, volume float64) Note {
	return Note{
		Frequency: 440 * math.Pow(twelfthrootof2, float64(step)),
		Duruation: duration,
		Volume:    volume,
	}
}

// Pause is a handy wrapper around NewNote with duration 0 and volume 0
func Pause(duration NoteDuration) Note {
	return NewNote(0, duration, 0)
}
