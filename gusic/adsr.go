package gusic

func NewADSR(attack, attackMultiplier, decay, decayMultiplier, sustain, release, releaseMultiplier float64) *ADSR {
	return &ADSR{
		attack:            attack,
		attackMultiplier:  attackMultiplier,
		decay:             decay,
		decayMultiplier:   decayMultiplier,
		sustain:           sustain,
		release:           release,
		releaseMultiplier: releaseMultiplier,
	}
}

func (a *ADSR) Attack(notes []Note) []Note {
	length := len(notes)

	output := make([]Note, length)
	copy(output, notes)

	end := int(a.attack * float64(length))

	multiplier := a.attackMultiplier / float64(end)
	for i := range notes[:end] {
		output[i].Volume *= multiplier * float64(i)
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

func (a *ADSR) Apply(notes []Note) []Note {
	var v []Note = notes

	v = a.Attack(v)
	v = a.Decay(v)
	v = a.Sustain(v)
	v = a.Release(v)

	return v
}
