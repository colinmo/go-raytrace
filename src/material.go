package main

type Material struct {
	Color     Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func NewMaterial() Material {
	return Material{
		Color:     NewColor(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
}

func (m Material) Equals(m2 Material) bool {
	return m.Color.Equals(m2.Color) &&
		epsilonEquals(m.Ambient, m2.Ambient) &&
		epsilonEquals(m.Diffuse, m2.Diffuse) &&
		epsilonEquals(m.Specular, m2.Specular) &&
		epsilonEquals(m.Shininess, m2.Shininess)
}
