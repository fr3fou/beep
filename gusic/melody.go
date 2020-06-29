package gusic

import "math"

// NewMelody is a constructor for a melody
func NewMelody(sampleRate float64, bpm int, noteLength int, generator Generator, adsr ADSR) *Melody {
	m := &Melody{
		SampleRate: sampleRate,
		BPM:        bpm,
		NoteLength: noteLength,
		Envelope:   adsr,
		Generator:  generator,
		Notes:      []Note{},
	}

	m.calculateDurations()

	return m
}

func (m *Melody) calculateDurations() {
	// m.breve =
	// m.semibreve =
	// m.crotchet =
	// m.quaver =
	// m.semiquaver =
	// m.demisemiquaver =
}

// AddNote adds a note
func (m *Melody) AddNote(note Note) {
	m.Notes = append(m.Notes, note)
}

// AddNotes adds notes
func (m *Melody) AddNotes(notes ...Note) {
	m.Notes = append(m.Notes, notes...)
}

// ApplyADSR applies all the stages of an ADSR to an array of notes
func (m *Melody) applyADSR() {
	applyADSR(m.Envelope, m.Notes)
}

func (m *Melody) compute() []float64 {
	melody := []float64{}

	m.applyADSR()

	for _, note := range m.Notes {
		samples := []float64{}

		for j := 1.0; j < m.SampleRate*note.Duration.Seconds(); j++ {
			val := math.Sin(note.Frequency*2*math.Pi/m.SampleRate) * note.Volume
			samples = append(samples, val)
		}

		melody = append(melody, samples...)
	}

	return melody
}

// func (m *Melody) PCM() []byte {

// }
