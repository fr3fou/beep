package gusic

// NewChord is a Chord constructor
func NewChord(notes ...SingleNote) Chord {
	return Chord{
		Notes: notes,
	}
}

func (c *Chord) Samples(sampleRate float64, generator Generator) []float64 {

}
