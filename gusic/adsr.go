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

func (a *ADSR) Attack() {
	length := len(a.notes)

	end := int(a.attack * float64(length))

	multiplier := a.attackMultiplier / float64(end)
	for i := range a.notes[:end] {
		a.notes[i].Volume *= multiplier * float64(i)
	}
}

func (a *ADSR) Decay() {
	panic("not implemented") // TODO: Implement
}

func (a *ADSR) Sustain() {
	panic("not implemented") // TODO: Implement
}

func (a *ADSR) Release() {
	panic("not implemented") // TODO: Implement
}

func (a *ADSR) Apply(notes []Note) {
	a.notes = notes

	a.Attack()
	a.Decay()
	a.Sustain()
	a.Release()
}
