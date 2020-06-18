package mego

// NewMelody is a constructor for a melody
func NewMelody(sampleRate int, bpm int, noteLength int, generator Generator, envelope Envelope, notes ...Note) *Melody {
	if notes == nil {
		notes = []Note{}
	}

	return &Melody{
		SampleRate: sampleRate,
		BPM:        bpm,
		NoteLength: noteLength,
		Envelope:   &envelope,
		Generator:  &generator,
		Note:       notes,
	}
}

// AddNote adds a note
func (m *Melody) AddNote(note Note) {
	m.Note = append(m.Note, note)
}

// AddNotes adds notes
func (m *Melody) AddNotes(notes ...Note) {
	m.Note = append(m.Note, notes...)
}
