package gusic

import "math"

// NewNote is a constructor for a Note
func NewNote(step int, duration NoteDuration, volume float64) SingleNote {
	return SingleNote{
		Frequency: 440 * math.Pow(twelfthrootof2, float64(step)),
		Duration:  duration,
		Volume:    volume,
	}
}

// Samples returns the samples
func (note *SingleNote) Samples(sampleRate float64, generator Generator) []float64 {
	samples := []float64{}

	duration := math.Ceil(sampleRate * note.Duration.Seconds())
	for j := 1.0; j <= duration; j++ {
		val := generator(j*note.Frequency*2*math.Pi/sampleRate) * note.Volume
		samples = append(samples, val)
	}

	return samples
}

// Rest is a handy wrapper around NewNote with the given duration and volume and frequency 0
func Rest(duration NoteDuration) SingleNote {
	return NewNote(0, duration, 0)
}
