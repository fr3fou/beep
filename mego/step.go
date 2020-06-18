package mego

// Step is a step from the pitch standard (A440)
type Step float64

const a float64 = 1.059463094359

// Some predefined note Steps in the 4th and 5th octave with relation to A4
const (
	C3  Step = -21
	CS3 Step = -20
	D3  Step = -19
	DS3 Step = -18
	E3  Step = -17
	F3  Step = -16
	FS3 Step = -15
	G3  Step = -14
	GS3 Step = -13
	A3  Step = -12
	AS3 Step = -11
	B3  Step = -10

	C4  Step = -9
	CS4 Step = -8
	D4  Step = -7
	DS4 Step = -6
	E4  Step = -5
	F4  Step = -4
	FS4 Step = -3
	G4  Step = -2
	GS4 Step = -1
	A4  Step = 0
	AS4 Step = 1
	B4  Step = 2

	C5  Step = 3
	CS5 Step = 4
	D5  Step = 5
	DS5 Step = 6
	E5  Step = 7
	F5  Step = 8
	FS5 Step = 9
	G5  Step = 10
	GS5 Step = 11
	A5  Step = 12
	AS5 Step = 13
	B5  Step = 14

	C6  Step = 15
	CS6 Step = 16
)
