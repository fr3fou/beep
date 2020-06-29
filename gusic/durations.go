package gusic

func (m *Melody) Breve() NoteDuration {
	return m.breve
}

func (m *Melody) Semibreve() NoteDuration {
	return m.semibreve
}

func (m *Melody) Crotchet() NoteDuration {
	return m.crotchet
}

func (m *Melody) Quaver() NoteDuration {
	return m.crotchet
}

func (m *Melody) Semiquaver() NoteDuration {
	return m.semiquaver
}

func (m *Melody) Demisemiquaver() NoteDuration {
	return m.demisemiquaver
}

func (m *Melody) DoubleWhole() NoteDuration {
	return m.breve
}

func (m *Melody) Whole() NoteDuration {
	return m.semibreve
}

func (m *Melody) Half() NoteDuration {
	return m.crotchet
}

func (m *Melody) Quarter() NoteDuration {
	return m.crotchet
}

func (m *Melody) Eigth() NoteDuration {
	return m.semiquaver
}

func (m *Melody) ThirtySecond() NoteDuration {
	return m.demisemiquaver
}
