package main

import "math"

type Light struct {
	Intensity Color
	Position  Tuple
}

func NewLight(position Tuple, intensity Color) Light {
	return Light{
		Intensity: intensity,
		Position:  position,
	}
}

func (l Light) Transform(m Matrix) Light {
	return NewLight(m.MultiplyTuple(l.Position), l.Intensity)
}

func Lighting(
	material Material,
	object Shaper,
	light Light,
	point Tuple,
	eyev Tuple,
	normalv Tuple,
	inShadow bool) Color {
	color := material.Color
	if material.HasPattern {
		color = material.Pattern.ColorAtObject(object, point)
	}
	effectiveColor := color.MultiplyColor(light.Intensity)
	lightV := light.Position.Subtract(point).Normalize()
	ambient := effectiveColor.MultiplyScalar(material.Ambient)
	black := NewColor(0, 0, 0)
	var diffuse Color
	var specular Color

	lightDotNormal := lightV.DotProduct(normalv)
	if inShadow {
		return ambient
	}
	if lightDotNormal < 0 {
		diffuse = black
		specular = black
	} else {
		diffuse = effectiveColor.MultiplyScalar(material.Diffuse).MultiplyScalar(lightDotNormal)
		reflectV := lightV.Negative().Reflect(normalv)
		reflectDotEye := reflectV.DotProduct(eyev)
		if reflectDotEye <= 0 {
			specular = black
		} else {
			factor := math.Pow(reflectDotEye, material.Shininess)
			specular = light.Intensity.MultiplyScalar(material.Specular).MultiplyScalar(factor)
		}
	}
	return ambient.Add(diffuse).Add(specular)
}
