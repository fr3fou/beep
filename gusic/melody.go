package gusic

// NewMelody is a constructor for a melody
func NewMelody(sampleRate int, bpm int, noteLength int, generator Generator, adsr ADSR) *Melody {
	return &Melody{
		SampleRate: sampleRate,
		BPM:        bpm,
		NoteLength: noteLength,
		Envelope:   adsr,
		Generator:  generator,
		Notes:      []Note{},
	}
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
func (m *Melody) ApplyADSR() {
	ApplyADSR(m.Envelope, m.Notes)
}
