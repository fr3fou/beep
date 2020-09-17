package gusic

// Dotted makes a note dotted.
func (m *Melody) Dotted(n NoteDuration) NoteDuration {
	return n + 1/2*n
}

// Breve is a breve
func (m *Melody) Breve() NoteDuration {
	return m.breve
}

// Semibreve is a semibreve
func (m *Melody) Semibreve() NoteDuration {
	return m.semibreve
}

// Minim is a minim
func (m *Melody) Minim() NoteDuration {
	return m.minim
}

// Crotchet is a crotchet
func (m *Melody) Crotchet() NoteDuration {
	return m.crotchet
}

// Quaver is a quaver
func (m *Melody) Quaver() NoteDuration {
	return m.quaver
}

// Semiquaver is a semiquaver
func (m *Melody) Semiquaver() NoteDuration {
	return m.semiquaver
}

// Demisemiquaver is a demisemiquaver
func (m *Melody) Demisemiquaver() NoteDuration {
	return m.demisemiquaver
}

// DoubleWhole is a double whole
func (m *Melody) DoubleWhole() NoteDuration {
	return m.breve
}

// Whole is a whole
func (m *Melody) Whole() NoteDuration {
	return m.semibreve
}

// Half is a half
func (m *Melody) Half() NoteDuration {
	return m.minim
}

// Quarter is a quarter
func (m *Melody) Quarter() NoteDuration {
	return m.crotchet
}

// Eighth is an eighth
func (m *Melody) Eighth() NoteDuration {
	return m.quaver
}

// Sixteenth is a sixteenth
func (m *Melody) Sixteenth() NoteDuration {
	return m.semiquaver
}

// ThirtySecond is a thirty second
func (m *Melody) ThirtySecond() NoteDuration {
	return m.demisemiquaver
}
