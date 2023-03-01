package main

import (
	"fmt"
	"math"
)

type Material struct {
	Color           Color
	Ambient         float64
	Diffuse         float64
	Specular        float64
	Shininess       float64
	Pattern         Pattern
	HasPattern      bool
	Reflective      float64
	Transparency    float64
	RefractiveIndex float64
}

func NewMaterial() Material {
	return Material{
		Color:           NewColor(1, 1, 1),
		Ambient:         0.1,
		Diffuse:         0.9,
		Specular:        0.9,
		Shininess:       200.0,
		HasPattern:      false,
		Reflective:      0,
		Transparency:    0,
		RefractiveIndex: 1,
	}
}

func (m Material) Equals(m2 Material) bool {
	if m.HasPattern {
		if m2.HasPattern && m.Pattern.Equals(m2.Pattern) {
		} else {
			return false
		}
	} else {
		if m2.HasPattern {
			return false
		}
	}
	return m.Color.Equals(m2.Color) &&
		epsilonEquals(m.Ambient, m2.Ambient) &&
		epsilonEquals(m.Diffuse, m2.Diffuse) &&
		epsilonEquals(m.Specular, m2.Specular) &&
		epsilonEquals(m.Shininess, m2.Shininess)
}

func (m Material) ToString() string {
	var patt Pattern
	if m.HasPattern {
		patt = m.Pattern
	} else {
		patt = NewStripePattern(NewColor(0, 0, 0), NewColor(0, 0, 0))
	}
	return fmt.Sprintf(
		"C: %v, A: %v, D: %v, Sp: %f, Sh: %f, Pa: %v",
		m.Color,
		m.Ambient,
		m.Diffuse,
		m.Specular,
		m.Shininess,
		patt,
	)
}

func (m *Material) SetPattern(p Pattern) {
	m.Pattern = p
	m.HasPattern = true
}

type Pattern interface {
	GetPatternType() string
	SetTransform(t Matrix)
	GetTransform() Matrix
	GetColorString(a string) Color
	ColorAt(point Tuple) Color
	ColorAtObject(o Shaper, wp Tuple) Color
	Equals(p2 Pattern) bool
}

type TestPattern struct {
	Pattern
	PatternType string
	Transform   Matrix
}

func NewTestPattern() *TestPattern {
	return &TestPattern{
		PatternType: "test",
		Transform:   IdentityMatrix(),
	}
}

func (p *TestPattern) GetPatternType() string {
	return "test"
}
func (p *TestPattern) ColorAt(point Tuple) Color {
	return NewColor(point.X, point.Y, point.Z)
}

func (p *TestPattern) Equals(p2 Pattern) bool {
	return p.PatternType == p2.GetPatternType() &&
		p.Transform.EqualsMatrix(p2.GetTransform())
}

func (p *TestPattern) SetTransform(t Matrix) {
	p.Transform = t
}
func (p *TestPattern) GetTransform() Matrix {
	return p.Transform
}
func (p *TestPattern) ColorAtObject(o Shaper, wp Tuple) Color {
	x := o.GetTransform()
	y := x.Inverse()
	objectPoint := y.MultiplyTuple(wp)

	x = p.GetTransform()
	y = x.Inverse()
	patternPoint := y.MultiplyTuple(objectPoint)

	return p.ColorAt(patternPoint)
}
func (p *TestPattern) GetColorString(s string) Color {
	return NewColor(0, 0, 0)
}

type StripePattern struct {
	Pattern
	A, B        Color
	PatternType string
	Transform   Matrix
}

func NewStripePattern(a, b Color) *StripePattern {
	return &StripePattern{
		A:           a,
		B:           b,
		PatternType: "stripe",
		Transform:   IdentityMatrix(),
	}
}

func (p *StripePattern) GetPatternType() string {
	return "stripe"
}

func (p *StripePattern) ColorAt(point Tuple) Color {
	if math.Mod(math.Floor(point.X), 2) == 0 {
		return p.A
	}
	return p.B
}

func (p *StripePattern) Equals(p2 Pattern) bool {
	if p.PatternType != p2.GetPatternType() {
		return false
	}
	switch p.PatternType {
	case "stripe":
		return p.A.Equals(p2.GetColorString("A")) && p.B.Equals(p2.GetColorString("B"))
	}
	return false
}

func (p *StripePattern) SetTransform(t Matrix) {
	p.Transform = t
}

func (p *StripePattern) GetTransform() Matrix {
	return p.Transform
}

func (p *StripePattern) ColorAtObject(o Shaper, wp Tuple) Color {
	x := o.GetTransform()
	y := x.Inverse()
	objectPoint := y.MultiplyTuple(wp)

	x = p.GetTransform()
	y = x.Inverse()
	patternPoint := y.MultiplyTuple(objectPoint)

	return p.ColorAt(patternPoint)
}

func (p *StripePattern) GetColorString(s string) Color {
	switch s {
	case "A":
		return p.A
	case "B":
		return p.B
	}
	return p.A
}

type GradientPattern struct {
	Pattern
	A, B        Color
	PatternType string
	Transform   Matrix
}

func NewGradientPattern(a, b Color) *GradientPattern {
	return &GradientPattern{
		A:           a,
		B:           b,
		PatternType: "gradient",
		Transform:   IdentityMatrix(),
	}
}
func (p *GradientPattern) GetPatternType() string {
	return "gradient"
}
func (p *GradientPattern) ColorAt(point Tuple) Color {
	distance := p.B.Subtract(p.A)
	fraction := point.X - math.Floor(point.X)
	col := p.A.Add(distance.MultiplyScalar(fraction))

	return col
}

func (p *GradientPattern) Equals(p2 Pattern) bool {
	return p.A.Equals(p2.GetColorString("A")) && p.B.Equals(p2.GetColorString("B"))
}

func (p *GradientPattern) SetTransform(t Matrix) {
	p.Transform = t
}

func (p *GradientPattern) GetTransform() Matrix {
	return p.Transform
}

func (p *GradientPattern) ColorAtObject(o Shaper, wp Tuple) Color {
	x := o.GetTransform()
	y := x.Inverse()
	objectPoint := y.MultiplyTuple(wp)

	x = p.GetTransform()
	y = x.Inverse()
	patternPoint := y.MultiplyTuple(objectPoint)

	return p.ColorAt(patternPoint)
}

func (p *GradientPattern) GetColorString(s string) Color {
	if s == "B" {
		return p.B
	}
	return p.A
}

type RingPattern struct {
	Pattern
	A, B        Color
	PatternType string
	Transform   Matrix
}

func NewRingPattern(a, b Color) *RingPattern {
	return &RingPattern{
		A:           a,
		B:           b,
		PatternType: "ring",
		Transform:   IdentityMatrix(),
	}
}
func (p *RingPattern) GetPatternType() string {
	return "ring"
}

func (p *RingPattern) ColorAt(point Tuple) Color {
	if math.Mod(math.Floor(point.X*point.X+point.Z*point.Z), 2) == 0 {
		return p.A
	}
	return p.B
}

func (p *RingPattern) Equals(p2 Pattern) bool {
	return p.A.Equals(p2.GetColorString("A")) && p.B.Equals(p2.GetColorString("B"))
}

func (p *RingPattern) SetTransform(t Matrix) {
	p.Transform = t
}

func (p *RingPattern) GetTransform() Matrix {
	return p.Transform
}

func (p *RingPattern) ColorAtObject(o Shaper, wp Tuple) Color {
	x := o.GetTransform()
	y := x.Inverse()
	objectPoint := y.MultiplyTuple(wp)

	x = p.GetTransform()
	y = x.Inverse()
	patternPoint := y.MultiplyTuple(objectPoint)

	return p.ColorAt(patternPoint)
}

func (p *RingPattern) GetColorString(s string) Color {
	if s == "B" {
		return p.B
	}
	return p.A
}

type CheckerPattern struct {
	Pattern
	A, B        Color
	PatternType string
	Transform   Matrix
}

func (p *CheckerPattern) GetPatternType() string {
	return "checker"
}
func NewCheckerPattern(a, b Color) *CheckerPattern {
	return &CheckerPattern{
		A:           a,
		B:           b,
		PatternType: "checker",
		Transform:   IdentityMatrix(),
	}
}

func (p *CheckerPattern) ColorAt(point Tuple) Color {
	result := math.Mod(math.Floor(point.X)+math.Floor(point.Y)+math.Floor(point.Z), 2) == 0
	if result {
		return p.A
	}
	return p.B
}

func (p *CheckerPattern) Equals(p2 Pattern) bool {
	return p.A.Equals(p2.GetColorString("A")) && p.B.Equals(p2.GetColorString("B"))
}

func (p *CheckerPattern) SetTransform(t Matrix) {
	p.Transform = t
}

func (p *CheckerPattern) GetTransform() Matrix {
	return p.Transform
}

func (p *CheckerPattern) ColorAtObject(o Shaper, wp Tuple) Color {
	wp = o.WorldToObject(wp)
	x := o.GetTransform()
	y := x.Inverse()
	objectPoint := y.MultiplyTuple(wp)

	x = p.GetTransform()
	y = x.Inverse()
	patternPoint := y.MultiplyTuple(objectPoint)

	return p.ColorAt(patternPoint)
}

func (p *CheckerPattern) GetColorString(s string) Color {
	if s == "B" {
		return p.B
	}
	return p.A
}
