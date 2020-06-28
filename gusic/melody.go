package gusic

// NewMelody is a constructor for a melody
func NewMelody(sampleRate int, bpm int, noteLength int, generator Generator, adsr ADSR, notes ...Note) *Melody {
	if notes == nil {
		notes = []Note{}
	}

	return &Melody{
		SampleRate: sampleRate,
		BPM:        bpm,
		NoteLength: noteLength,
		Envelope:   adsr,
		Generator:  generator,
		Notes:      notes,
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
	length := len(m.Notes)

	ratios := m.Envelope.GetRatios()

	attackEnd := int(ratios.AttackRatio * float64(length))
	decayEnd := int(ratios.DecayRatio*float64(length)) + attackEnd
	sustainEnd := int(ratios.SustainRatio*float64(length)) + decayEnd

	m.Envelope.Attack(m.Notes[:attackEnd])
	m.Envelope.Decay(m.Notes[attackEnd:decayEnd])
	m.Envelope.Sustain(m.Notes[decayEnd:sustainEnd])
	m.Envelope.Release(m.Notes[sustainEnd:])
}
