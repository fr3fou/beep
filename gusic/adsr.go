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
		output[i].Volume *= multiplier
		// We add + 1 to compensate for the fact that arrays start at 0
		multiplier += float64(i+1) * a.attack * 1.25
	}
	a.currentStep = multiplier

	return output
}

func (a *ADSR) Decay(notes []Note) []Note {
	// ok
}

func (a *ADSR) Sustain(notes []Note) []Note {
	panic("not implemented") // TODO: Implement
}

func (a *ADSR) Release(notes []Note) []Note {
	panic("not implemented") // TODO: Implement
}
