package gusic

// NewMelody is a constructor for a melody
func NewMelody(sampleRate int, bpm int, noteLength int, generator Generator, adsr ADSR) *Melody {
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
func (m *Melody) ApplyADSR() {
	ApplyADSR(m.Envelope, m.Notes)
}
