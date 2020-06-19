package gusic

type OctaveContext struct {
	Octave int
}

func OctaveScope(octave int, f func(*OctaveContext) []Note) []Note {
	return f(&OctaveContext{Octave: octave})
}

func (o *OctaveContext) VolumeScope(volume float64, f func(ctx *VolumeContext) []Note) []Note {
	// now what?
}

func (o *OctaveContext) A(duration NoteDuration, volume float64) Note {
	return A(o.Octave, duration, volume)
}
