package gusic

func NewADSR(a, d, s, r float64) Envelope {
	return &ADSR{
		a, d, s, r,
	}
}

func (a *ADSR) Attack(notes []Note) []Note {
	panic("not implemented") // TODO: Implement
}

func (a *ADSR) Decay(notes []Note) []Note {
	panic("not implemented") // TODO: Implement
}

func (a *ADSR) Sustain(notes []Note) []Note {
	panic("not implemented") // TODO: Implement
}

func (a *ADSR) Release(notes []Note) []Note {
	panic("not implemented") // TODO: Implement
}
