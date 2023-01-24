package main

import (
	"fmt"
	"math"
	"testing"

	"gonum.org/v1/gonum/mat"
)

/* TESTS */
func TestATupleWithW10IsAPoint(t *testing.T) {
	a := Tuple{4.3, -4.2, 3.1, 1.0}
	if a.X != 4.3 {
		t.Errorf("a.X != 4.3 -> %f", a.X)
	}
	if a.Y != -4.2 {
		t.Errorf("a.Y != -4.2 -> %f", a.Y)
	}
	if a.Z != 3.1 {
		t.Errorf("a.Z != 3.1 -> %f", a.Z)
	}
	if a.W != 1.0 {
		t.Errorf("a.W != 1.0 -> %f", a.W)
	}
	if !a.isPoint() {
		t.Errorf("a is not a point")
	}
	if a.isVector() {
		t.Errorf("a is a vector")
	}
}

func TestATupleWithW00IsAVector(t *testing.T) {
	a := Tuple{4.3, -4.2, 3.1, 0.0}
	if a.X != 4.3 {
		t.Errorf("a.X != 4.3 -> %f", a.X)
	}
	if a.Y != -4.2 {
		t.Errorf("a.Y != -4.2 -> %f", a.Y)
	}
	if a.Z != 3.1 {
		t.Errorf("a.Z != 3.1 -> %f", a.Z)
	}
	if a.W != 0.0 {
		t.Errorf("a.W != 0.0 -> %f", a.W)
	}
	if a.isPoint() {
		t.Errorf("a is a point")
	}
	if !a.isVector() {
		t.Errorf("a is not a vector")
	}
}

func TestPointCreatesTuplesWithW1(t *testing.T) {
	p := point(4, -4, 3)
	comp := Tuple{4, -4, 3, 1}
	if !p.equals(comp) {
		t.Errorf("p not a point")
	}
}

func TestVectorCreatesTuplesWithW0(t *testing.T) {
	v := vector(4, -4, 3)
	comp := Tuple{4, -4, 3, 0}
	if !v.equals(comp) {
		t.Errorf("v is not a vector")
	}
}

func TestAddingTwoTuples(t *testing.T) {
	a1 := Tuple{3, -2, 5, 1}
	a2 := Tuple{-2, 3, 1, 0}
	a3 := Tuple{1, 1, 6, 1}
	if a1.add(a2) != a3 {
		t.Errorf("Failed adding two tuples")
	}
}

func TestSubtractingTwoPoints(t *testing.T) {
	a1 := point(3, 2, 1)
	a2 := point(5, 6, 7)
	a3 := vector(-2, -4, -6)
	if !a1.sub(a2).equals(a3) {
		t.Errorf("Failed subtracting two points")
		t.Errorf("%f,%f,%f,%f vs %f,%f,%f,%f = %f,%f,%f,%f", a1.X, a1.Y, a1.Z, a1.W, a2.X, a2.Y, a2.Z, a2.W, a1.sub(a2).X, a1.sub(a2).Y, a1.sub(a2).Z, a1.sub(a2).W)
	}
}

func TestSubtractVectorFromPoint(t *testing.T) {
	p := point(3, 2, 1)
	v := vector(5, 6, 7)
	p2 := point(-2, -4, -6)
	if p.sub(v) != p2 {
		t.Errorf("Failed substracting vector from point")
	}
}

func TestSubtractingTwoVectors(t *testing.T) {
	v1 := vector(3, 2, 1)
	v2 := vector(5, 6, 7)
	v3 := vector(-2, -4, -6)
	if v1.sub(v2) != v3 {
		t.Error("Failed subtracting two vectors")
	}
}

func TestSubtractingAVectorFromZeroVector(t *testing.T) {
	zero := vector(0, 0, 0)
	v := vector(1, -2, 3)
	v2 := vector(-1, 2, -3)
	if zero.sub(v) != v2 {
		t.Errorf("Failed subtracting from zero vector")
	}
}

func TestNegatingATuple(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	a2 := Tuple{-1, 2, -3, 4}
	if !a.neg().equals(a2) {
		t.Error("FAiled negating")
	}
}

func TestMultiplyingTupleByAScalar(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	if !a.mult(3.5).equals(Tuple{3.5, -7, 10.5, -14}) {
		t.Error("Multiply by scalar fails")
	}
}
func TestMultiplyingTupleByAFraction(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	if !a.mult(0.5).equals(Tuple{0.5, -1, 1.5, -2}) {
		t.Error("Multiply by scalar fails")
	}
}

func TestDividingTupleByAScalar(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	if !a.div(2).equals(Tuple{0.5, -1, 1.5, -2}) {
		t.Error("Dividing by scalar fails")
	}
}

func TestMagnitudeVectors(t *testing.T) {
	v := vector(1, 0, 0)
	if v.magnitude() != 1 {
		t.Error("Magnitude fail")
	}
	v = vector(0, 1, 0)
	if v.magnitude() != 1 {
		t.Error("Magnitude fail")
	}
	v = vector(0, 0, 1)
	if v.magnitude() != 1 {
		t.Error("Magnitude fail")
	}
	v = vector(1, 2, 3)
	if !epsilonEquals(v.magnitude(), math.Sqrt(14)) {
		t.Error("Magnitude fail")
	}
	v = vector(-1, -2, 3)
	if !epsilonEquals(v.magnitude(), math.Sqrt(14)) {
		t.Error("Magnitude fail")
	}

}

func TestNormaliseVector(t *testing.T) {
	v := vector(4, 0, 0)
	if !v.normalize().equals(vector(1, 0, 0)) {
		t.Error("Normalize fail")
	}
	v = vector(1, 2, 3)
	if !v.normalize().equals(vector(0.26726, 0.53452, 0.80178)) {
		t.Error("Normalize fail")
	}
}

func TestDotProductTwoTuples(t *testing.T) {
	a := vector(1, 2, 3)
	b := vector(2, 3, 4)
	if a.dot(b) != 20 {
		t.Error("Dot product fail")
	}
}

func TestCrossProductTwoVectors(t *testing.T) {
	a := vector(1, 2, 3)
	b := vector(2, 3, 4)
	if !a.cross(b).equals(vector(-1, 2, -1)) {
		t.Error("Cross fail")
	}
	if !b.cross(a).equals(vector(1, -2, 1)) {
		t.Error("Cross fail")
	}

}

// COLORS
func TestColorsAreRGB(t *testing.T) {
	c := Color{-0.5, 0.4, 1.7}
	if !epsilonEquals(c.Red, -0.5) || !epsilonEquals(c.Green, 0.4) || !epsilonEquals(c.Blue, 1.7) {
		t.Error("Color fail")
	}
}

func TestAddingColors(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}
	if !c1.add(c2).equals(Color{1.6, 0.7, 1.0}) {
		t.Error("Color add fail")
	}
}

func TestSubtractingColors(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}
	if !c1.sub(c2).equals(Color{0.2, 0.5, 0.5}) {
		t.Error("Color sub fail")
	}
}

func TestMultiplyingColorByScalar(t *testing.T) {
	c := Color{0.2, 0.3, 0.4}
	if !c.multS(2).equals(Color{0.4, 0.6, 0.8}) {
		t.Error("Color mult fail")
	}
}
func TestMultiplyingColors(t *testing.T) {
	c1 := Color{1, 0.2, 0.4}
	c2 := Color{0.9, 1, 0.1}
	if !c1.multC(c2).equals(Color{0.9, 0.2, 0.04}) {
		t.Error("Color mult fail")
	}
}

// CANVAS
func TestCreatingACanvas(t *testing.T) {
	c := NewCanvas(10, 20)
	if c.W != 10 {
		t.Error("Width fail")
	}
	if c.H != 20 {
		t.Error("Height fail")
	}
	black := Color{0, 0, 0}
	for i := 0; i < c.W; i++ {
		for j := 0; j < c.H; j++ {
			if c.pixelAt(i, j) != black {
				t.Error("Default black image fail")
			}
		}
	}
}

func TestWritePixelToCanvas(t *testing.T) {
	c := NewCanvas(10, 20)
	red := Color{1, 0, 0}

	c.writePixel(2, 3, red)
	if !c.pixelAt(2, 3).equals(red) {
		t.Error("Pixel at 2,3 not red")
	}
}

func TestConstructingPPMHeader(t *testing.T) {
	c := NewCanvas(5, 3)
	ppm := c.toPPM()
	header := ppm[:11]
	if header != "P3\n5 3\n255\n" {
		t.Errorf("Header fail %s", header)
	}
}

func TestConstructingThePPMPixelData(t *testing.T) {
	c := NewCanvas(5, 3)
	c1 := Color{1.5, 0, 0}
	c2 := Color{0, 0.5, 0}
	c3 := Color{-0.5, 0, 1}
	c.writePixel(0, 0, c1)
	c.writePixel(2, 1, c2)
	c.writePixel(4, 2, c3)
	ppm := c.toPPM()
	guts := ppm[11:]
	if guts != "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0\n0 0 0 0 0 0 0 128 0 0 0 0 0 0 0\n0 0 0 0 0 0 0 0 0 0 0 0 0 0 255\n" {
		t.Errorf("Guts fail \n%s vs \n%s", guts, "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0\n0 0 0 0 0 0 0 128 0 0 0 0 0 0 0\n0 0 0 0 0 0 0 0 0 0 0 0 0 0 255\n")
	}
}

func TestSplittingLongLinesInPPM(t *testing.T) {
	c := NewCanvas(10, 2)
	c1 := Color{1, 0.8, 0.6}
	for j := 0; j < 2; j++ {
		for i := 0; i < 10; i++ {
			c.writePixel(i, j, c1)
		}
	}
	ppm := c.toPPM()
	guts := ppm[12:]
	if guts != "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204\n153 255 204 153 255 204 153 255 204 153 255 204 153\n255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204\n153 255 204 153 255 204 153 255 204 153 255 204 153\n" {
		t.Errorf("Guts fail \n[%s] vs \n[%s]", guts, "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204\n153 255 204 153 255 204 153 255 204 153 255 204 153\n255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204\n153 255 204 153 255 204 153 255 204 153 255 204 153\n")
	}
}

// MATRIX
func TestConstructInspect4x4Matrix(t *testing.T) {
	var values = []float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5}
	m := mat.NewDense(4, 4, values)
	if m.At(0, 0) != 1 {
		t.Errorf("M[0,0] failed")
	}
	if m.At(0, 3) != 4 {
		t.Errorf("M[0,4] failed")
	}
	if m.At(3, 2) != 15.5 {
		t.Errorf("M[3,2] failed")
	}
}

func TestConstruct2x2Matrix(t *testing.T) {
	var values = []float64{-3, 5, 1, -2}
	m := mat.NewDense(2, 2, values)
	if m.At(0, 0) == -3 && m.At(0, 1) == 5 && m.At(1, 0) == 1 && m.At(1, 1) == -2 {
		//
	} else {
		t.Errorf("2x2=fail")
	}
}

func TestConstruct3x3Matrix(t *testing.T) {
	var values = []float64{-3, 5, 0, 1, -2, -7, 0, 1, 1}
	m := mat.NewDense(3, 3, values)
	if m.At(0, 0) == -3 && m.At(1, 1) == -2 && m.At(2, 2) == 1 {
		//
	} else {
		t.Errorf("3x3=fail")
	}
}

func TestMatrixEqualityEquals(t *testing.T) {
	var values = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}
	m := mat.NewDense(4, 4, values)
	values = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}
	n := mat.NewDense(4, 4, values)
	if !mat.EqualApprox(m, n, epsilon) {
		t.Errorf("Matrix equal fail")
	}
}
func TestMatrixEqualityNotEquals(t *testing.T) {
	var values = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}
	m := mat.NewDense(4, 4, values)
	values = []float64{2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	n := mat.NewDense(4, 4, values)
	if mat.EqualApprox(m, n, epsilon) {
		t.Errorf("Matrix non-equal fail")
	}
}

func TestMatrixMult(t *testing.T) {
	var values = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}
	a := mat.NewDense(4, 4, values)
	values = []float64{-2, 1, 2, 3, 3, 2, 1, -1, 4, 3, 6, 5, 1, 2, 7, 8}
	b := mat.NewDense(4, 4, values)
	values = []float64{20, 22, 50, 48, 44, 54, 114, 108, 40, 58, 110, 102, 16, 26, 46, 42}
	c := mat.NewDense(4, 4, values)
	d := mat.NewDense(4, 4, nil)
	d.Mul(a, b)
	fmt.Println(d)
	if !mat.EqualApprox(d, c, epsilon) {
		t.Errorf("Matrix Mul fail")
	}
}

func TestMatrixTupleMult(t *testing.T) {
	var values = []float64{1, 2, 3, 4, 2, 4, 4, 2, 8, 6, 4, 1, 0, 0, 0, 1}
	a := mat.NewDense(4, 4, values)
	b := Tuple{1, 2, 3, 1}
	if !b.matrixMult(a).equals(Tuple{18, 24, 33, 1}) {
		t.Errorf("Matrix by Tuple fail")
	}
}

func TestMatrixIdentityMult(t *testing.T) {
	var values = []float64{0, 1, 2, 4, 1, 2, 4, 8, 2, 4, 8, 16, 4, 8, 16, 32}
	a := mat.NewDense(4, 4, values)
	b := identityMatrix()
	c := mat.NewDense(4, 4, nil)
	c.Mul(a, b)
	if !mat.EqualApprox(c, a, epsilon) {
		t.Errorf("Matrix identity mult fail")
	}
}

func TestTupleIdentityMult(t *testing.T) {
	a := Tuple{1, 2, 3, 4}
	b := identityMatrix()
	if !a.matrixMult(b).equals(a) {
		t.Errorf("Tuple identity mult fail")
	}
}

func TestTransposeMatrix(t *testing.T) {
	var values = []float64{0, 9, 3, 0, 9, 8, 0, 8, 1, 8, 5, 3, 0, 0, 5, 8}
	a := mat.NewDense(4, 4, values)
	b := a.T()
	values = []float64{0, 9, 1, 0, 9, 8, 8, 0, 3, 0, 5, 5, 0, 8, 3, 8}
	c := mat.NewDense(4, 4, values)
	if !mat.EqualApprox(b, c, epsilon) {
		t.Errorf("Matrix transpose fail")
	}
}

func TestDeterminant2x2Matrix(t *testing.T) {
	var values = []float64{1, 5, -3, 2}
	a := mat.NewDense(2, 2, values)
	if mat.Det(a) != 17 {
		t.Errorf("Determinent fail")
	}
}

func TestSubmatrixOf3x3Matrix(t *testing.T) {
	var values = []float64{1, 5, 0, -3, 2, 7, 0, 6, -3}
	a := mat.NewDense(3, 3, values)
	b := subMatrix(a, 0, 2)
	if !mat.EqualApprox(b, mat.NewDense(2, 2, []float64{-3, 2, 0, 6}), epsilon) {
		t.Errorf("3x3Submatrix fail")
	}
}

func TestSubmatrixOf4x4Matrix(t *testing.T) {
	var values = []float64{-6, 1, 1, 6, -8, 5, 8, 6, -1, 0, 8, 2, -7, 1, -1, 1}
	a := mat.NewDense(4, 4, values)
	b := subMatrix(a, 2, 1)
	if !mat.EqualApprox(b, mat.NewDense(3, 3, []float64{-6, 1, 6, -8, 8, 6, -7, -1, 1}), epsilon) {
		t.Errorf("4x4Submatrix fail")
	}
}

func TestCalculateMinor3x3Matrix(t *testing.T) {
	var values = []float64{3, 5, 0, 2, -1, -7, 6, -1, 5}
	a := mat.NewDense(3, 3, values)
	b := subMatrix(a, 1, 0)
	if !epsilonEquals(mat.Det(b), 25) {
		t.Errorf("Minor3x3 det fail")
	}
	if !epsilonEquals(minor(a, 1, 0), 25) {
		t.Errorf("Minor3x3 minor fail")
	}
}

func TestCofactor3x3MAtrix(t *testing.T) {
	var values = []float64{3, 5, 0, 2, -1, -7, 6, -1, 5}
	a := mat.NewDense(3, 3, values)
	if !epsilonEquals(minor(a, 0, 0), -12) {
		t.Errorf("First cofactor fail")
	}
	if !epsilonEquals(cofactor(a, 0, 0), -12) {
		t.Errorf("Second cofactor fail")
	}
	if !epsilonEquals(minor(a, 1, 0), 25) {
		t.Errorf("Third cofactor fail")
	}
	if !epsilonEquals(cofactor(a, 1, 0), -25) {
		t.Errorf("Fourth cofactor fail")
	}
}

func TestCalcDeterm3x3(t *testing.T) {
	var values = []float64{1, 2, 6, -5, 8, -4, 2, 6, 4}
	a := mat.NewDense(3, 3, values)
	if !epsilonEquals(cofactor(a, 0, 0), 56) {
		fmt.Println(cofactor(a, 0, 0))
		t.Errorf("First test det 3x3 fail")
	}
	if !epsilonEquals(cofactor(a, 0, 1), 12) {
		fmt.Println(cofactor(a, 0, 0))
		t.Errorf("First test det 3x3 fail")
	}
	if !epsilonEquals(cofactor(a, 0, 2), -46) {
		fmt.Println(cofactor(a, 0, 0))
		t.Errorf("First test det 3x3 fail")
	}
	if !epsilonEquals(mat.Det(a), -196) {
		fmt.Println(cofactor(a, 0, 0))
		t.Errorf("First test det 3x3 fail")
	}
}

func TestCalcDeterm4x4(t *testing.T) {
	var values = []float64{-2, -8, 3, 5, -3, 1, 7, 3, 1, 2, -9, 6, -6, 7, 7, -9}
	a := mat.NewDense(4, 4, values)
	if !epsilonEquals(cofactor(a, 0, 0), 690) {
		fmt.Println(cofactor(a, 0, 0))
		t.Errorf("First test det 4x43 fail")
	}
	if !epsilonEquals(cofactor(a, 0, 1), 447) {
		fmt.Println(cofactor(a, 0, 0))
		t.Errorf("First test det 4x4 fail")
	}
	if !epsilonEquals(cofactor(a, 0, 2), 210) {
		fmt.Println(cofactor(a, 0, 0))
		t.Errorf("First test det 4x4 fail")
	}
	if !epsilonEquals(cofactor(a, 0, 3), 51) {
		fmt.Println(cofactor(a, 0, 0))
		t.Errorf("First test det 4x4 fail")
	}
	if !epsilonEquals(mat.Det(a), -4071) {
		fmt.Println(cofactor(a, 0, 0))
		t.Errorf("First test det 4xx4 fail")
	}

}

func TestInvertibility(t *testing.T) {
	var values = []float64{6, 4, 4, 4, 5, 5, 7, 6, 4, -9, 3, -7, 9, 1, 7, -6}
	a := mat.NewDense(4, 4, values)
	if !epsilonEquals(mat.Det(a), -2120) {
		t.Errorf("X")
	}
	if !isInvertable(a) {
		t.Errorf("X")
	}
	a = mat.NewDense(4, 4, []float64{-4, 2, -2, -3, 9, 6, 2, 6, 0, -5, 1, -5, 0, 0, 0, 0})
	if !epsilonEquals(mat.Det(a), 0) {
		t.Errorf("X")
	}
	if isInvertable(a) {
		t.Errorf("X")
	}

}

func TestInverse(t *testing.T) {
	var values = []float64{-5, 2, 6, -8, 1, -5, 1, 8, 7, 7, -6, -7, 1, -3, 7, 4}
	a := mat.NewDense(4, 4, values)
	var b mat.Dense
	b.Inverse(a)
	if !epsilonEquals(mat.Det(a), 532) {
		t.Errorf("X")
	}
	if !epsilonEquals(cofactor(a, 2, 3), -160) {
		t.Errorf("x")
	}
	if !epsilonEquals(b.At(3, 2), -160.0/532.0) {
		fmt.Printf("%f, %f", b.At(3, 2), -160.0/532.0)
		t.Errorf("z")
	}
	if !epsilonEquals(cofactor(a, 3, 2), 105) {
		t.Errorf("a")
	}
	if !epsilonEquals(b.At(2, 3), 105.0/532.0) {
		t.Errorf("b")
	}
	if !mat.EqualApprox(&b, mat.NewDense(4, 4, []float64{0.21805, 0.45113, 0.24060, -0.04511, -0.80827, -1.45677, -0.44361, 0.52068, -0.07895, -0.22368, -0.05263, 0.19737, -0.52256, -0.81391, -0.30075, 0.30639}), epsilon) {
		t.Errorf("c")
	}
}

func TestInverse2(t *testing.T) {
	a := mat.NewDense(4, 4, []float64{8, -5, 9, 2, 7, 5, 6, 1, -6, 0, 9, 6, -3, 0, -9, -4})
	var b mat.Dense
	b.Inverse(a)
	if !mat.EqualApprox(&b, mat.NewDense(4, 4, []float64{-0.15385, -0.15385, -0.28205, -0.53846, -0.07692, 0.12308, 0.02564, 0.03077, 0.35897, 0.35897, 0.43590, 0.92308, -0.69231, -0.69231, -0.76923, -1.92308}), epsilon) {
		t.Errorf("1")
	}
}

func TestInverse3(t *testing.T) {
	a := mat.NewDense(4, 4, []float64{9, 3, 0, 9, -5, -2, -6, -3, -4, 9, 6, 4, -7, 6, 6, 2})
	var b mat.Dense
	b.Inverse(a)
	if !mat.EqualApprox(&b, mat.NewDense(4, 4, []float64{-0.04074, -0.07778, 0.14444, -0.22222, -0.07778, 0.03333, 0.36667, -0.33333, -0.02901, -0.14630, -0.10926, 0.12963, 0.17778, 0.06667, -0.26667, 0.33333}), epsilon) {
		t.Errorf("1")
	}

}

func TestMultiplyProductByInverse(t *testing.T) {
	a := mat.NewDense(4, 4, []float64{3, -9, 7, 3, 3, -8, 2, -9, -4, 4, 4, 1, -6, 5, -1, 1})
	b := mat.NewDense(4, 4, []float64{8, 2, 2, 2, 3, -1, 7, 0, 7, 0, 5, 4, 6, -2, 0, 5})
	var c, d, e mat.Dense
	c.Mul(a, b)
	d.Inverse(b)
	e.Mul(&c, &d)

	if !mat.EqualApprox(&e, a, epsilon) {
		t.Errorf("wa")
	}
}
