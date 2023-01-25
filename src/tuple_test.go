package main

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/cucumber/godog"
)

type tuples map[string]Tuple

type tupletest struct {
	Tuples tuples
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			tt := &tupletest{}

			ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
				tt.Tuples = tuples{}
				return ctx, nil
			})

			ctx.Step(`^tuple\.([a-zA-Z0-9]+) ← tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.makeATuple)
			ctx.Step(`^tuple\.([a-zA-Z0-9]+) is a point$`, tt.aIsAPoint)
			ctx.Step(`^tuple\.([a-zA-Z0-9]+) is a vector$`, tt.aIsAVector)
			ctx.Step(`^tuple\.([a-zA-Z0-9]+) is not a point$`, tt.aIsNotAPoint)
			ctx.Step(`^tuple\.([a-zA-Z0-9]+) is not a vector$`, tt.aIsNotAVector)
			ctx.Step(`^tuple\.([a-zA-Z0-9]+)\.([xyzw]) = (-?\d+\.\d+)`, tt.aFloatValueEqual)
			ctx.Step(`^tuple\.([a-zA-Z0-9]+) = tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.tupleEqualsTuple)
			ctx.Step(`^tuple\.([a-zA-Z0-9]+) ← point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.tuplepPoint)
			ctx.Step(`^tuple\.([a-zA-Z0-9]+) ← vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.tuplevVector)
			ctx.Step(`^tuple\.([a-zA-Z0-9]+) \+ tuple\.([a-zA-Z0-9]+) = tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.tupleAPlusBEqualsC)
			ctx.Step(`^tuple\.([a-zA-Z0-9]+) - tuple\.([a-zA-Z0-9]+) = vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.pointMinusPointEqualsVector)
			ctx.Step(`^tuple\.([a-zA-Z0-9]+) - tuple\.([a-zA-Z0-9]+) = point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.pointMinusVectorEqualsPoint)
			ctx.Step(`^-tuple\.([a-zA-Z0-9]+) = tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.negativeTupleEquals)
			ctx.Step(`^tuple\.([a-zA-Z0-9]+) \* (-?\d+(?:\.\d+)?) = tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.tupleMultipliedScalarEquals)
			ctx.Step(`^tuple\.([a-zA-Z0-9]+) \/ (-?\d+(?:\.\d+)?) = tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.tupleDividedScalarEquals)
			ctx.Step(`^magnitude\(tuple\.([a-zA-Z0-9]+)\) = (-?\d+(?:\.\d+)?)$`, tt.magnitudeTupleEquals)
			ctx.Step(`^magnitude\(tuple\.([a-zA-Z0-9]+)\) = √(\d+(?:\.\d+)?)$`, tt.magnitudeTupleEqualsSqrt)
			ctx.Step(`^normalize\(tuple\.([a-zA-Z0-9]+)\) = vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.normalizeTupleEqualsVector)
			ctx.Step(`^tuple\.([a-zA-Z0-9]+) ← normalize\(tuple\.([a-zA-Z0-9]+)\)$`, tt.tupleNormalizeEqualsTuple)
			ctx.Step(`^dot\(tuple\.([a-zA-Z0-9]+), tuple\.([a-zA-Z0-9]+)\) = (-?\d+(?:\.\d+)?)$`, tt.dotTupleaTupleb)
			ctx.Step(`^cross\(tuple\.([a-zA-Z0-9]+), tuple\.([a-zA-Z0-9]+)\) = vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.crossTuplebTupleaVector)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func floatOrErr(float string) (float64, error) {
	if tempvar, err := strconv.ParseFloat(float, 64); err == nil {
		return tempvar, nil
	}
	return 0.0, fmt.Errorf(`"%s" could not be floated`, float)
}

func (tt *tupletest) tupleAPlusBEqualsC(varNameA, varNameB, x, y, z, w string) error {
	xF, err := floatOrErr(x)
	if err != nil {
		return err
	}
	yF, err := floatOrErr(y)
	if err != nil {
		return err
	}
	zF, err := floatOrErr(z)
	if err != nil {
		return err
	}
	wF, err := floatOrErr(w)
	if err != nil {
		return err
	}
	a, ok := tt.Tuples[varNameA]
	if !ok {
		return fmt.Errorf("tuple %s not available", varNameA)
	}
	b, ok := tt.Tuples[varNameB]
	if !ok {
		return fmt.Errorf("tuple %s not available", varNameB)
	}
	tempTuple := Tuple{X: xF, Y: yF, Z: zF, W: wF}
	if a.Add(b).EqualsTuple(tempTuple) {
		return nil
	}
	return fmt.Errorf("Tuple %v doesn't equal %v", a, tempTuple)
}

func (tt *tupletest) pointMinusPointEqualsVector(varNameA, varNameB, x, y, z string) error {
	xF, err := floatOrErr(x)
	if err != nil {
		return err
	}
	yF, err := floatOrErr(y)
	if err != nil {
		return err
	}
	zF, err := floatOrErr(z)
	if err != nil {
		return err
	}
	a, ok := tt.Tuples[varNameA]
	if !ok {
		return fmt.Errorf("tuple %s not available", varNameA)
	}
	b, ok := tt.Tuples[varNameB]
	if !ok {
		return fmt.Errorf("tuple %s not available", varNameB)
	}
	tempTuple := NewVector(xF, yF, zF)
	if a.Subtract(b).EqualsTuple(tempTuple) {
		return nil
	}
	return fmt.Errorf("Point %v minus Point %v doesn't equal %v", a, b, tempTuple)
}

func (tt *tupletest) pointMinusVectorEqualsPoint(varNameA, varNameB, x, y, z string) error {
	xF, err := floatOrErr(x)
	if err != nil {
		return err
	}
	yF, err := floatOrErr(y)
	if err != nil {
		return err
	}
	zF, err := floatOrErr(z)
	if err != nil {
		return err
	}
	a, ok := tt.Tuples[varNameA]
	if !ok {
		return fmt.Errorf("tuple %s not available", varNameA)
	}
	b, ok := tt.Tuples[varNameB]
	if !ok {
		return fmt.Errorf("tuple %s not available", varNameB)
	}
	tempTuple := NewPoint(xF, yF, zF)
	if a.Subtract(b).EqualsTuple(tempTuple) {
		return nil
	}
	return fmt.Errorf("Point %v minus Vector %v doesn't equal %v", a, b, tempTuple)
}

func (tt *tupletest) negativeTupleEquals(varNameA, x, y, z, w string) error {
	a, ok := tt.Tuples[varNameA]
	if !ok {
		return fmt.Errorf("tuple %s not available", varNameA)
	}
	temptuple, err := StringsToTuple(x, y, z, w)
	if err != nil {
		return err
	}
	if !a.Negative().EqualsTuple(temptuple) {
		return fmt.Errorf("Negative of %v is %v, which doesn't equal %v", a, a.Negative(), temptuple)
	}
	return nil
}

func (tt *tupletest) tupleMultipliedScalarEquals(varNameA string, scalar float64, x, y, z, w string) error {
	tempTuple, err := StringsToTuple(x, y, z, w)
	if err != nil {
		return err
	}
	a, ok := tt.Tuples[varNameA]
	if !ok {
		return fmt.Errorf("tuple %s not available", varNameA)
	}
	if a.MultiplyScalar(scalar).EqualsTuple(tempTuple) {
		return nil
	}
	return fmt.Errorf("Tuple %v times %f doesn't equal %v", a, scalar, tempTuple)
}

func (tt *tupletest) tupleDividedScalarEquals(varNameA string, scalar float64, x, y, z, w string) error {
	tempTuple, err := StringsToTuple(x, y, z, w)
	if err != nil {
		return err
	}
	a, ok := tt.Tuples[varNameA]
	if !ok {
		return fmt.Errorf("tuple %s not available", varNameA)
	}
	if a.DivideScalar(scalar).EqualsTuple(tempTuple) {
		return nil
	}
	return fmt.Errorf("Tuple %v times %f doesn't equal %v", a, scalar, tempTuple)
}

func (tt *tupletest) magnitudeTupleEquals(varName string, target float64) error {
	a, ok := tt.Tuples[varName]
	if !ok {
		return fmt.Errorf("tuple %s not available", varName)
	}
	if a.Magnitude() == target {
		return nil
	}
	return fmt.Errorf("Magnitude Tuple %v equals %f not %f", a, a.Magnitude(), target)

}

func (tt *tupletest) magnitudeTupleEqualsSqrt(varName string, target float64) error {
	return tt.magnitudeTupleEquals(varName, math.Sqrt(target))
}

func (tt *tupletest) normalizeTupleEqualsVector(varName, x, y, z string) error {
	a, ok := tt.Tuples[varName]
	if !ok {
		return fmt.Errorf("tuple %s not available", varName)
	}
	xF, err := floatOrErr(x)
	if err != nil {
		return err
	}
	yF, err := floatOrErr(y)
	if err != nil {
		return err
	}
	zF, err := floatOrErr(z)
	if err != nil {
		return err
	}
	if a.Normalize().EqualsTuple(NewVector(xF, yF, zF)) {
		return nil
	}
	return fmt.Errorf("Normalized tuple %v doesn't equal %v", a, NewVector(xF, yF, zF))
}

func (tt *tupletest) tupleNormalizeEqualsTuple(varNameA, varNameB string) error {
	b, ok := tt.Tuples[varNameB]
	if !ok {
		return fmt.Errorf("tuple %s not available", varNameB)
	}
	tt.Tuples[varNameA] = b.Normalize()
	return nil
}

func (tt *tupletest) dotTupleaTupleb(varNameA, varNameB string, expected float64) error {
	a, ok := tt.Tuples[varNameA]
	if !ok {
		return fmt.Errorf("tuple %s not available", varNameA)
	}
	b, ok := tt.Tuples[varNameB]
	if !ok {
		return fmt.Errorf("tuple %s not available", varNameB)
	}
	if a.DotProduct(b) == expected {
		return nil
	}
	return fmt.Errorf("dot product of %v and %v is %f not %f", a, b, a.DotProduct(b), expected)
}

func (tt *tupletest) crossTuplebTupleaVector(varNameA, varNameB, x, y, z string) error {
	a, ok := tt.Tuples[varNameA]
	if !ok {
		return fmt.Errorf("tuple %s not available", varNameA)
	}
	b, ok := tt.Tuples[varNameB]
	if !ok {
		return fmt.Errorf("tuple %s not available", varNameB)
	}
	expected, err := StringsToVector(x, y, z)
	if err != nil {
		return err
	}
	if a.CrossProduct(b).EqualsTuple(expected) {
		return nil
	}
	return fmt.Errorf("dot product of %v and %v is %v not %v", a, b, a.CrossProduct(b), expected)

}

func (tt *tupletest) tupleEqualsTuple(varName, x, y, z, w string) error {
	tempTuple, err := StringsToTuple(x, y, z, w)
	if err != nil {
		return err
	}
	a, ok := tt.Tuples[varName]
	if !ok {
		return fmt.Errorf("tuple %snot available", varName)
	}
	if a.EqualsTuple(tempTuple) {
		return nil
	}
	return fmt.Errorf("Tuple %v doesn't equal %v", a, tempTuple)

}

func (tt *tupletest) aFloatValueEqual(varName, field, xyzw string) error {
	valF, err := floatOrErr(xyzw)
	if err != nil {
		return err
	}

	a, ok := tt.Tuples[varName]
	if !ok {
		return fmt.Errorf("tuple %snot available", varName)
	}
	switch field {
	case "x":
		if !epsilonEquals(a.X, valF) {
			return fmt.Errorf("X value %f didn't equal %f", a.X, valF)
		}
	case "y":
		if !epsilonEquals(a.Y, valF) {
			return fmt.Errorf("Y value %f didn't equal %f", a.Y, valF)
		}
	case "z":
		if !epsilonEquals(a.Z, valF) {
			return fmt.Errorf("Z value %f didn't equal %f", a.Z, valF)
		}
	case "w":
		if !epsilonEquals(a.W, valF) {
			return fmt.Errorf("W value %f didn't equal %f", a.W, valF)
		}
	}
	return nil
}

func (tt *tupletest) makeATuple(varName string, x, y, z, w string) error {
	var err error
	tt.Tuples[varName], err = StringsToTuple(x, y, z, w)
	return err
}

func StringsToTuple(x, y, z, w string) (Tuple, error) {
	xF, err := floatOrErr(x)
	if err != nil {
		return Tuple{}, err
	}
	yF, err := floatOrErr(y)
	if err != nil {
		return Tuple{}, err
	}
	zF, err := floatOrErr(z)
	if err != nil {
		return Tuple{}, err
	}
	wF, err := floatOrErr(w)
	if err != nil {
		return Tuple{}, err
	}
	return Tuple{xF, yF, zF, wF}, nil
}
func StringsToVector(x, y, z string) (Tuple, error) {
	xF, err := floatOrErr(x)
	if err != nil {
		return Tuple{}, err
	}
	yF, err := floatOrErr(y)
	if err != nil {
		return Tuple{}, err
	}
	zF, err := floatOrErr(z)
	if err != nil {
		return Tuple{}, err
	}
	return NewVector(xF, yF, zF), nil
}

func (tt *tupletest) aIsAPoint(varName string) error {
	a, ok := tt.Tuples[varName]
	if !ok {
		return fmt.Errorf("tuple %snot available", varName)
	}
	if !a.isPoint() {
		return fmt.Errorf("tuple %s is not a point", varName)
	}
	return nil
}

func (tt *tupletest) aIsAVector(varName string) error {
	a, ok := tt.Tuples[varName]
	if !ok {
		return fmt.Errorf("tuple %snot available", varName)
	}
	if !a.isVector() {
		return fmt.Errorf("tuple %s is not a vector", varName)
	}
	return nil
}

func (tt *tupletest) aIsNotAPoint(varName string) error {
	a, ok := tt.Tuples[varName]
	if !ok {
		return fmt.Errorf("tuple %snot available", varName)
	}
	if a.isPoint() {
		return fmt.Errorf("tuple %s is a point", varName)
	}
	return nil
}

func (tt *tupletest) aIsNotAVector(varName string) error {
	a, ok := tt.Tuples[varName]
	if !ok {
		return fmt.Errorf("tuple %snot available", varName)
	}
	if a.isVector() {
		return fmt.Errorf("tuple %s is a vector", varName)
	}
	return nil
}

func (tt *tupletest) tuplepPoint(varName string, x, y, z string) error {
	xF, err := floatOrErr(x)
	if err != nil {
		return err
	}
	yF, err := floatOrErr(y)
	if err != nil {
		return err
	}
	zF, err := floatOrErr(z)
	if err != nil {
		return err
	}
	tt.Tuples[varName] = NewPoint(xF, yF, zF)
	return nil
}

func (tt *tupletest) tuplevVector(varName string, x, y, z string) error {
	xF, err := floatOrErr(x)
	if err != nil {
		return err
	}
	yF, err := floatOrErr(y)
	if err != nil {
		return err
	}
	zF, err := floatOrErr(z)
	if err != nil {
		return err
	}
	tt.Tuples[varName] = NewVector(xF, yF, zF)
	return nil
}
