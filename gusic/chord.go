package gusic

import "math"

// NewChord is a Chord constructor
func NewChord(notes ...SingleNote) Chord {
	return Chord{
		Notes: notes,
	}
}

// Samples returns the samples
func (c Chord) Samples(sampleRate float64, generator Generator) []float64 {
	chordSamples := [][]float64{}

	longestNote := 0.0
	for _, note := range c.Notes {
		samples := note.Samples(sampleRate, generator)
		chordSamples = append(chordSamples, samples)
		longestNote = math.Max(longestNote, float64(len(samples)))
	}

	finalSamples := make([]float64, int64(longestNote))
	for i := int64(0); i < int64(longestNote); i++ {
		for _, sample := range chordSamples {
			if i < int64(len(sample)) {
				finalSamples[i] += sample[i]
			}
		}
	}

	return finalSamples
}
