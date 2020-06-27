package gusic

func NewADSR(a, d, s, r float64) *ADSR {
	return &ADSR{
		a, d, s, r, 0,
	}
}

func (a *ADSR) Attack(notes []Note) []Note {
	output := make([]Note, len(notes))
	copy(output, notes)

	length := len(notes)
	index := int(a.attack * float64(length))

	multiplier := 0.0
	for i := range notes[:index] {
		output[i].Volume *= multiplier
		multiplier += float64(i+1) * a.attack * 2 // add 1 to offset by 1
	}

	return output
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
