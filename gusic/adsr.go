package gusic

func NewADSR(a, d, s, r float64) *ADSR {
	return &ADSR{
		a, d, s, r, 0.0,
	}
}

func (a *ADSR) Attack(notes []Note) []Note {
	output := make([]Note, len(notes))
	copy(output, notes)

	length := len(notes)
	index := int(a.attack * float64(length))

	multiplier := a.currentStep
	for i := range notes[:index] {

	}
	a.currentStep = multiplier

	return output
}

func (a *ADSR) Decay(notes []Note) []Note {
	//
}

func (a *ADSR) Sustain(notes []Note) []Note {
	panic("not implemented") // TODO: Implement
}

func (a *ADSR) Release(notes []Note) []Note {
	panic("not implemented") // TODO: Implement
}

func (a *ADSR) Apply(notes []Note) []Note {
	var v []Note = notes

	v = a.Attack(v)
	v = a.Decay(v)
	v = a.Sustain(v)
	v = a.Release(v)

	return v
}
