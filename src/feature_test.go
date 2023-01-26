package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v16"
)

type tuples map[string]Tuple
type colors map[string]Color
type canvases map[string]Canvas
type ppms map[string]string
type matrices map[string]Matrix

type tupletest struct {
	Tuples   tuples
	Colors   colors
	Canvases canvases
	PPMs     ppms
	Matrices matrices
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			tt := &tupletest{}

			ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
				tt.Canvases = canvases{}
				tt.Tuples = tuples{}
				tt.Colors = colors{}
				tt.PPMs = ppms{}
				tt.Matrices = matrices{}
				return ctx, nil
			})

			// Tuple
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

			// Color
			ctx.Step(`^colors\.([a-zA-Z0-9]+)\.([a-zA-Z0-9]+) = (-?\d+(?:\.\d+)?)$`, tt.checkColor)
			ctx.Step(`^colors\.([a-zA-Z0-9]+) ← color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.makeColor)
			ctx.Step(`^colors\.([a-zA-Z0-9]+) \* (\d+) = color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.multiplyColourScalar)
			ctx.Step(`^colors\.([a-zA-Z0-9]+) \+ colors\.([a-zA-Z0-9]+) = color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.AddColour)
			ctx.Step(`^colors\.([a-zA-Z0-9]+) - colors\.([a-zA-Z0-9]+) = color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.SubtractColour)
			ctx.Step(`^colors\.([a-zA-Z0-9]+) \* colors\.([a-zA-Z0-9]+) = color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.multiplyColourTuple)

			// Canvas
			ctx.Step(`^canvas\.([a-zA-Z0-9]+) ← canvas\((\d+), (\d+)\)$`, tt.makeCanvas)
			ctx.Step(`^canvas\.([a-zA-Z0-9]+)\.height = (\d+)$`, tt.canvasHeightEquals)
			ctx.Step(`^canvas\.([a-zA-Z0-9]+)\.width = (\d+)$`, tt.canvasWidthEquals)
			ctx.Step(`^every pixel of canvas\.([a-zA-Z0-9]+) is color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.everyPixelOfCanvascIsColor)
			ctx.Step(`^pixel_at\(canvas\.([a-zA-Z0-9]+), (\d+), (\d+)\) = colors\.([a-zA-Z0-9]+)$`, tt.pixelAt)
			ctx.Step(`^write_pixel\(canvas\.([a-zA-Z0-9]+), (\d+), (\d+), colors\.([a-zA-Z0-9]+)\)$`, tt.writePixelAt)
			ctx.Step(`^every pixel of canvas\.([a-zA-Z0-9]+) is set to color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.everyPixelOfCanvascIsSetToColor)
			ctx.Step(`^lines (\d+)-(\d+) of ppm\.([a-zA-Z0-9]+) are$`, tt.linesOfPpmppmAre)
			ctx.Step(`^ppm\.([a-zA-Z0-9]+) ← canvas_to_ppm\(canvas\.([a-zA-Z0-9]+)\)$`, tt.saveCanvasAsPPM)
			ctx.Step(`^ppm\.([a-zA-Z0-9]+) ends with a newline character$`, tt.ppmEndsWithANewlineCharacter)

			// Matrices
			ctx.Step(`^matrix.([a-zA-Z0-9]+)\[(\d+),(\d+)\] = (-?\d+(?:\.\d+)?)$`, tt.checkMatrixValueAtXY)
			ctx.Step(`^the following (\d+)x(\d+) matrix matrix.([a-zA-Z0-9]+):$`, tt.makeXbyXMatrixFromTable)
			ctx.Step(`^matrix\.([a-zA-Z0-9]+) = matrix\.([a-zA-Z0-9]+)$`, tt.matrixAEqualsMatrixB)
			ctx.Step(`^matrix\.([a-zA-Z0-9]+) != matrix\.([a-zA-Z0-9]+)$`, tt.matrixANotEqualsMatrixB)
			ctx.Step(`^the following matrix matrix\.([a-zA-Z0-9]+):$`, tt.makeMatrixFromTable)
			ctx.Step(`^matrix\.([a-zA-Z0-9]+) \* matrix\.([a-zA-Z0-9]+) is the following matrix:$`, tt.matrixATimesMatrixB)
			ctx.Step(`^matrix\.([a-zA-Z0-9]+) \* tuple\.([a-zA-Z0-9]+) = tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.matrixATimesTupleB)
			ctx.Step(`^IdentityMatrix \* tuple\.([a-zA-Z0-9]+) = tuple\.([a-zA-Z0-9]+)$`, tt.identityTimesTuple)
			ctx.Step(`^matrix\.([a-zA-Z0-9]+) \* IdentityMatrix = matrix\.([a-zA-Z0-9]+)$`, tt.matrixTimesIdentity)
			ctx.Step(`^transpose\(matrix\.([a-zA-Z0-9]+)\) is the following matrix:$`, tt.transposematrixAIsTheFollowingMatrix)

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

func (tt *tupletest) tupleAPlusBEqualsC(varNameA, varNameB string, x, y, z, w float64) error {
	a, ok := tt.Tuples[varNameA]
	if !ok {
		return fmt.Errorf("tuple %s not available", varNameA)
	}
	b, ok := tt.Tuples[varNameB]
	if !ok {
		return fmt.Errorf("tuple %s not available", varNameB)
	}
	tempTuple := Tuple{X: x, Y: y, Z: z, W: w}
	if a.Add(b).EqualsTuple(tempTuple) {
		return nil
	}
	return fmt.Errorf("Tuple %v doesn't equal %v", a, tempTuple)
}

func (tt *tupletest) pointMinusPointEqualsVector(varNameA, varNameB string, x, y, z float64) error {
	a, ok := tt.Tuples[varNameA]
	if !ok {
		return fmt.Errorf("tuple %s not available", varNameA)
	}
	b, ok := tt.Tuples[varNameB]
	if !ok {
		return fmt.Errorf("tuple %s not available", varNameB)
	}
	tempTuple := NewVector(x, y, z)
	if a.Subtract(b).EqualsTuple(tempTuple) {
		return nil
	}
	return fmt.Errorf("Point %v minus Point %v doesn't equal %v", a, b, tempTuple)
}

func (tt *tupletest) pointMinusVectorEqualsPoint(varNameA, varNameB string, x, y, z float64) error {
	a, ok := tt.Tuples[varNameA]
	if !ok {
		return fmt.Errorf("tuple %s not available", varNameA)
	}
	b, ok := tt.Tuples[varNameB]
	if !ok {
		return fmt.Errorf("tuple %s not available", varNameB)
	}
	tempTuple := NewPoint(x, y, z)
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

// COLORS

func (tt *tupletest) makeColor(varName string, r, g, b float64) error {
	tt.Colors[varName] = NewColor(r, g, b)
	return nil
}

func (tt *tupletest) checkColor(varName, colorIndex string, colorValue float64) error {
	var compValue float64
	a, ok := tt.Colors[varName]
	if !ok {
		return fmt.Errorf("color %snot available", varName)
	}
	switch colorIndex {
	case "Red":
		compValue = a.Red
	case "Blue":
		compValue = a.Blue
	case "Green":
		compValue = a.Green
	}
	if compValue == colorValue {
		return nil
	}
	return fmt.Errorf("didn't match %f to %f", compValue, colorValue)
}
func (tt *tupletest) multiplyColourScalar(varName string, scalar float64, r, g, b float64) error {
	a, ok := tt.Colors[varName]
	if !ok {
		return fmt.Errorf("color %snot available", varName)
	}
	if a.MultiplyScalar(scalar).Equals(NewColor(r, g, b)) {
		return nil
	}
	return fmt.Errorf("Color %v didn't match %v", a.MultiplyScalar(scalar), NewColor(r, g, b))
}
func (tt *tupletest) multiplyColourTuple(varName1, varName2 string, r, g, b float64) error {
	a1, ok := tt.Colors[varName1]
	if !ok {
		return fmt.Errorf("color %snot available", varName1)
	}
	b1, ok := tt.Colors[varName2]
	if !ok {
		return fmt.Errorf("color %snot available", varName1)
	}
	if a1.MultiplyColor(b1).Equals(NewColor(r, g, b)) {
		return nil
	}
	return fmt.Errorf("Color %v didn't match %v", a1.MultiplyColor(b1), NewColor(r, g, b))
}
func (tt *tupletest) AddColour(varName1 string, varName2 string, r, g, b float64) error {
	a1, ok := tt.Colors[varName1]
	if !ok {
		return fmt.Errorf("color %snot available", varName1)
	}
	b1, ok := tt.Colors[varName2]
	if !ok {
		return fmt.Errorf("color %snot available", varName1)
	}
	if a1.Add(b1).Equals(NewColor(r, g, b)) {
		return nil
	}
	return fmt.Errorf("Color %v didn't match %v", a1.Add(b1), NewColor(r, g, b))
}
func (tt *tupletest) SubtractColour(varName1 string, varName2 string, r, g, b float64) error {
	a1, ok := tt.Colors[varName1]
	if !ok {
		return fmt.Errorf("color %snot available", varName1)
	}
	b1, ok := tt.Colors[varName2]
	if !ok {
		return fmt.Errorf("color %snot available", varName1)
	}
	if a1.Subtract(b1).Equals(NewColor(r, g, b)) {
		return nil
	}
	return fmt.Errorf("Color %v didn't match %v", a1.Subtract(b1), NewColor(r, g, b))
}

///  CANVASES

func (tt *tupletest) makeCanvas(varName string, w, h int) error {
	tt.Canvases[varName] = NewCanvas(w, h)
	return nil
}

func (tt *tupletest) canvasHeightEquals(varName string, h int) error {
	a1, ok := tt.Canvases[varName]
	if !ok {
		return fmt.Errorf("canvas %s not available", varName)
	}
	if a1.Height == h {
		return nil
	}
	return fmt.Errorf("height of %s is %d not %d", varName, a1.Height, h)
}

func (tt *tupletest) canvasWidthEquals(varName string, w int) error {
	a1, ok := tt.Canvases[varName]
	if !ok {
		return fmt.Errorf("canvas %s not available", varName)
	}
	if a1.Width == w {
		return nil
	}
	return fmt.Errorf("width of %s is %d not %d", varName, a1.Width, w)
}

func (tt *tupletest) everyPixelOfCanvascIsColor(varName string, r, g, b float64) error {
	a1, ok := tt.Canvases[varName]
	if !ok {
		return fmt.Errorf("canvas %s not available", varName)
	}
	b1 := NewColor(r, g, b)

	for i := 0; i < a1.Width; i++ {
		for j := 0; j < a1.Height; j++ {
			if !a1.Pixels[i][j].Equals(b1) {
				return fmt.Errorf("color at %d,%d is %v not %v", i, j, a1.Pixels[i][j], b1)
			}
		}
	}
	return nil
}

func (tt *tupletest) everyPixelOfCanvascIsColorName(varName string, varName2 string) error {
	a1, ok := tt.Canvases[varName]
	if !ok {
		return fmt.Errorf("canvas %s not available", varName)
	}
	b1, ok := tt.Colors[varName2]
	if !ok {
		return fmt.Errorf("color %s not available", varName2)
	}

	for i := 0; i < a1.Width; i++ {
		for j := 0; j < a1.Height; j++ {
			if !a1.Pixels[i][j].Equals(b1) {
				return fmt.Errorf("color at %d,%d is %v not %v", i, j, a1.Pixels[i][j], b1)
			}
		}
	}
	return nil
}

func (tt *tupletest) pixelAt(varName string, x, y int, varName2 string) error {
	a1, ok := tt.Canvases[varName]
	if !ok {
		return fmt.Errorf("canvas %s not available", varName)
	}
	b1, ok := tt.Colors[varName2]
	if !ok {
		return fmt.Errorf("color %s not available", varName2)
	}
	if a1.Pixels[x][y].Equals(b1) {
		return nil
	}
	return fmt.Errorf("Pixel %d,%d is %v not %v", x, y, a1.Pixels[x][y], b1)

}

func (tt *tupletest) writePixelAt(varName string, x, y int, varName2 string) error {
	a1, ok := tt.Canvases[varName]
	if !ok {
		return fmt.Errorf("canvas %s not available", varName)
	}
	b1, ok := tt.Colors[varName2]
	if !ok {
		return fmt.Errorf("color %s not available", varName2)
	}
	a1.WritePixel(x, y, b1)
	return nil
}

func (tt *tupletest) everyPixelOfCanvascIsSetToColor(varName string, r, g, b float64) error {
	a1, ok := tt.Canvases[varName]
	if !ok {
		return fmt.Errorf("canvas %s not available", varName)
	}
	cc := NewColor(r, g, b)
	for i := range a1.Pixels {
		for j := range a1.Pixels[i] {
			tt.Canvases[varName].Pixels[i][j] = cc
		}
	}
	return nil
}

func (tt *tupletest) linesOfPpmppmAre(arg1, arg2 int, varName string, expected *godog.DocString) error {
	a1, ok := tt.PPMs[varName]
	if !ok {
		return fmt.Errorf("ppm %s not available", varName)
	}
	mep := strings.Split(a1, "\n")
	mep2 := strings.Join(mep[arg1-1:arg2], "\n")
	if mep2 == expected.Content {
		return nil
	}
	return fmt.Errorf("string mismatch of \n%s' and \n'%s'", mep2, expected)
}

func (tt *tupletest) saveCanvasAsPPM(varName1, varName2 string) error {
	a1, ok := tt.Canvases[varName2]
	if !ok {
		return fmt.Errorf("canvas %s not available", varName2)
	}
	tt.PPMs[varName1] = a1.ToPPM()
	return nil
}

func (tt *tupletest) ppmEndsWithANewlineCharacter(varName string) error {
	a1, ok := tt.PPMs[varName]
	if !ok {
		return fmt.Errorf("ppm %s not available", varName)
	}
	if a1[len(a1)-1:] == "\n" {
		return nil
	}
	return fmt.Errorf("last character of %s is %s, not newline", varName, a1[len(a1)-1:])
}

func (tt *tupletest) checkMatrixValueAtXY(varName string, x, y int, expected float64) error {
	a1, ok := tt.Matrices[varName]
	if !ok {
		return fmt.Errorf("matrix %s not available", varName)
	}
	if a1.Cells[x][y] == expected {
		return nil
	}
	return fmt.Errorf("matrix value at %d,%d is %f not %f", x, y, a1.Cells[x][y], expected)

}

func (tt *tupletest) makeMatrixFromString(varName string, stringvalue string) error {
	tt.Matrices[varName] = BuildMatrixFromString(stringvalue)
	return nil
}

func (tt *tupletest) makeXbyXMatrixFromString(x, y int, varName string, stringvalue *godog.DocString) error {
	log.Fatalf("X: %d, Y: %d, Name: %s, Value: %s\n", x, y, varName, stringvalue.Content)
	return nil
	// return tt.makeMatrixFromString(varName, stringvalue.Content)
}

func (tt *tupletest) makeXbyXMatrixFromTable(x, y int, varName string, tablevalue *godog.Table) error {
	tt.Matrices[varName] = Matrix{
		Rows:  len(tablevalue.Rows),
		Cols:  len(tablevalue.Rows[0].Cells),
		Cells: make(map[int]map[int]float64),
	}

	for x1 := range tablevalue.Rows {
		tt.Matrices[varName].Cells[x1] = make(map[int]float64)
		for y1, cell := range tablevalue.Rows[x1].Cells {
			tt.Matrices[varName].Cells[x1][y1], _ = strconv.ParseFloat(cell.Value, 64)
		}
	}
	return nil
	// return tt.makeMatrixFromString(varName, stringvalue.Content)
}

func (tt *tupletest) makeMatrixFromTable(varName string, tablevalue *godog.Table) error {
	tt.Matrices[varName] = tableToMatrix(tablevalue)
	return nil
}

func (tt *tupletest) matrixAEqualsMatrixB(varName1, varName2 string) error {

	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("matrix %s not available", varName1)
	}

	b, ok := tt.Matrices[varName2]
	if !ok {
		return fmt.Errorf("matrix %s not available", varName2)
	}

	if a.EqualsMatrix(b) {
		return nil
	}
	return fmt.Errorf("matrices do not match")
}

func (tt *tupletest) matrixANotEqualsMatrixB(varName1, varName2 string) error {
	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("matrix %s not available", varName1)
	}
	b, ok := tt.Matrices[varName2]
	if !ok {
		return fmt.Errorf("matrix %s not available", varName2)
	}

	if !a.EqualsMatrix(b) {
		return nil
	}
	return fmt.Errorf("matrices do match")
}
func (tt *tupletest) matrixATimesMatrixB(varName1, varName2 string, tablevalue *godog.Table) error {
	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("matrix %s not available", varName1)
	}
	b, ok := tt.Matrices[varName2]
	if !ok {
		return fmt.Errorf("matrix %s not available", varName2)
	}

	c := tableToMatrix(tablevalue)

	d := a.MultiplyMatrix(b)
	if d.EqualsMatrix(c) {
		return nil
	}
	return fmt.Errorf("matrix multiplaction fail\n%s", d.ToString())
}
func (tt *tupletest) matrixATimesTupleB(varName1, varName2 string, x, y, z, w float64) error {
	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("matrix %s not available", varName1)
	}
	b, ok := tt.Tuples[varName2]
	if !ok {
		return fmt.Errorf("matrix %s not available", varName2)
	}
	d := a.MultiplyTuple(b)
	if d.EqualsTuple(NewTuple(x, y, z, w)) {
		return nil
	}
	return fmt.Errorf("matrix multiplication by tuple fail\n%s", d.ToString())
}

func tableToMatrix(tablevalue *messages.PickleTable) Matrix {
	c := Matrix{
		Rows:  len(tablevalue.Rows),
		Cols:  len(tablevalue.Rows[0].Cells),
		Cells: make(map[int]map[int]float64),
	}

	for x1 := range tablevalue.Rows {
		c.Cells[x1] = make(map[int]float64)
		for y1, cell := range tablevalue.Rows[x1].Cells {
			c.Cells[x1][y1], _ = strconv.ParseFloat(cell.Value, 64)
		}
	}
	return c
}

func (tt *tupletest) identityTimesTuple(varName1 string, varName2 string) error {
	a, ok := tt.Tuples[varName1]
	if !ok {
		return fmt.Errorf("matrix %s not available", varName1)
	}
	b, ok := tt.Tuples[varName2]
	if !ok {
		return fmt.Errorf("matrix %s not available", varName2)
	}
	d := IdentityMatrix.MultiplyTuple(a)
	if d.EqualsTuple(b) {
		return nil
	}
	return fmt.Errorf("matrix multiplication by tuple fail\n%s", d.ToString())

}
func (tt *tupletest) matrixTimesIdentity(varName1, varName2 string) error {
	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("matrix %s not available", varName1)
	}
	b, ok := tt.Matrices[varName2]
	if !ok {
		return fmt.Errorf("matrix %s not available", varName2)
	}

	d := a.MultiplyMatrix(IdentityMatrix)
	if d.EqualsMatrix(b) {
		return nil
	}
	return fmt.Errorf("matrix identity multiplication fail\n%s", d.ToString())
}

func (tt *tupletest) transposematrixAIsTheFollowingMatrix(varName string, arg1 *godog.Table) error {
	a, ok := tt.Matrices[varName]
	if !ok {
		return fmt.Errorf("matrix %s not available", varName)
	}
	b := tableToMatrix(arg1)

	c := a.Transpose()
	if c.EqualsMatrix(b) {
		return nil
	}
	return fmt.Errorf("transpost fail")
}
