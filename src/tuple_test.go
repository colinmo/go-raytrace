package main

import (
	"context"
	"fmt"
	"github.com/cucumber/godog"
	"strconv"
	"testing"
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

func (tt *tupletest) tupleEqualsTuple(varName, x, y, z, w string) error {
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
	a, ok := tt.Tuples[varName]
	if !ok {
		return fmt.Errorf("tuple %snot available", varName)
	}
	tempTuple := Tuple{X: xF, Y: yF, Z: zF, W: wF}
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
		if !EpsilonEquals(a.X, valF) {
			return fmt.Errorf("X value %f didn't equal %f", a.X, valF)
		}
	case "y":
		if !EpsilonEquals(a.Y, valF) {
			return fmt.Errorf("Y value %f didn't equal %f", a.Y, valF)
		}
	case "z":
		if !EpsilonEquals(a.Z, valF) {
			return fmt.Errorf("Z value %f didn't equal %f", a.Z, valF)
		}
	case "w":
		if !EpsilonEquals(a.W, valF) {
			return fmt.Errorf("W value %f didn't equal %f", a.W, valF)
		}
	}
	return nil
}

func (tt *tupletest) makeATuple(varName string, x, y, z, w string) error {
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
	tt.Tuples[varName] = Tuple{
		X: xF,
		Y: yF,
		Z: zF,
		W: wF,
	}
	return nil
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
