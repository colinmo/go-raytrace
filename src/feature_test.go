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
type rays map[string]Ray
type slices map[string]map[int]float64
type shapes map[string]Shaper
type intersections map[string]Intersection
type arrayintersections map[string]map[int]Intersection

type tupletest struct {
	Tuples             tuples
	Colors             colors
	Canvases           canvases
	PPMs               ppms
	Matrices           matrices
	Rays               rays
	Slices             slices
	Shapes             shapes
	Intersections      intersections
	ArrayIntersections arrayintersections
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
				tt.Rays = rays{}
				tt.Slices = slices{}
				tt.Shapes = shapes{}
				tt.Intersections = intersections{}
				tt.ArrayIntersections = arrayintersections{}
				return ctx, nil
			})

			// Tuple
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) ← tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.makeATuple)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) is a point$`, tt.aIsAPoint)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) is a vector$`, tt.aIsAVector)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) is not a point$`, tt.aIsNotAPoint)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) is not a vector$`, tt.aIsNotAVector)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+)\.([xyzw]) = (-?\d+\.\d+)`, tt.aFloatValueEqual)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) = tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.tupleEqualsTuple)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) ← point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.tuplepPoint)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) ← vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.tuplevVector)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) \+ tuple\.([a-zA-Z0-9_]+) = tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.tupleAPlusBEqualsC)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) - tuple\.([a-zA-Z0-9_]+) = vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.pointMinusPointEqualsVector)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) - tuple\.([a-zA-Z0-9_]+) = point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.pointMinusVectorEqualsPoint)
			ctx.Step(`^-tuple\.([a-zA-Z0-9_]+) = tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.negativeTupleEquals)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) \* (-?\d+(?:\.\d+)?) = tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.tupleMultipliedScalarEquals)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) \/ (-?\d+(?:\.\d+)?) = tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.tupleDividedScalarEquals)
			ctx.Step(`^magnitude\(tuple\.([a-zA-Z0-9_]+)\) = (-?\d+(?:\.\d+)?)$`, tt.magnitudeTupleEquals)
			ctx.Step(`^magnitude\(tuple\.([a-zA-Z0-9_]+)\) = √(\d+(?:\.\d+)?)$`, tt.magnitudeTupleEqualsSqrt)
			ctx.Step(`^normalize\(tuple\.([a-zA-Z0-9_]+)\) = vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.normalizeTupleEqualsVector)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) ← normalize\(tuple\.([a-zA-Z0-9_]+)\)$`, tt.tupleNormalizeEqualsTuple)
			ctx.Step(`^dot\(tuple\.([a-zA-Z0-9_]+), tuple\.([a-zA-Z0-9_]+)\) = (-?\d+(?:\.\d+)?)$`, tt.dotTupleaTupleb)
			ctx.Step(`^cross\(tuple\.([a-zA-Z0-9_]+), tuple\.([a-zA-Z0-9_]+)\) = vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.crossTuplebTupleaVector)

			// Color
			ctx.Step(`^colors\.([a-zA-Z0-9_]+)\.([a-zA-Z0-9_]+) = (-?\d+(?:\.\d+)?)$`, tt.checkColor)
			ctx.Step(`^colors\.([a-zA-Z0-9_]+) ← color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.makeColor)
			ctx.Step(`^colors\.([a-zA-Z0-9_]+) \* (\d+) = color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.multiplyColourScalar)
			ctx.Step(`^colors\.([a-zA-Z0-9_]+) \+ colors\.([a-zA-Z0-9_]+) = color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.AddColour)
			ctx.Step(`^colors\.([a-zA-Z0-9_]+) - colors\.([a-zA-Z0-9_]+) = color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.SubtractColour)
			ctx.Step(`^colors\.([a-zA-Z0-9_]+) \* colors\.([a-zA-Z0-9_]+) = color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.multiplyColourTuple)

			// Canvas
			ctx.Step(`^canvas\.([a-zA-Z0-9_]+) ← canvas\((\d+), (\d+)\)$`, tt.makeCanvas)
			ctx.Step(`^canvas\.([a-zA-Z0-9_]+)\.height = (\d+)$`, tt.canvasHeightEquals)
			ctx.Step(`^canvas\.([a-zA-Z0-9_]+)\.width = (\d+)$`, tt.canvasWidthEquals)
			ctx.Step(`^every pixel of canvas\.([a-zA-Z0-9_]+) is color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.everyPixelOfCanvascIsColor)
			ctx.Step(`^pixel_at\(canvas\.([a-zA-Z0-9_]+), (\d+), (\d+)\) = colors\.([a-zA-Z0-9_]+)$`, tt.pixelAt)
			ctx.Step(`^write_pixel\(canvas\.([a-zA-Z0-9_]+), (\d+), (\d+), colors\.([a-zA-Z0-9_]+)\)$`, tt.writePixelAt)
			ctx.Step(`^every pixel of canvas\.([a-zA-Z0-9_]+) is set to color\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.everyPixelOfCanvascIsSetToColor)
			ctx.Step(`^lines (\d+)-(\d+) of ppm\.([a-zA-Z0-9_]+) are$`, tt.linesOfPpmppmAre)
			ctx.Step(`^ppm\.([a-zA-Z0-9_]+) ← canvas_to_ppm\(canvas\.([a-zA-Z0-9_]+)\)$`, tt.saveCanvasAsPPM)
			ctx.Step(`^ppm\.([a-zA-Z0-9_]+) ends with a newline character$`, tt.ppmEndsWithANewlineCharacter)

			// Matrices
			ctx.Step(`^matrix.([a-zA-Z0-9_]+)\[(\d+),(\d+)\] = (-?\d+(?:\.\d+)?)$`, tt.checkMatrixValueAtXY)
			ctx.Step(`^the following (\d+)x(\d+) matrix matrix.([a-zA-Z0-9_]+):$`, tt.makeXbyXMatrixFromTable)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) = matrix\.([a-zA-Z0-9_]+)$`, tt.matrixAEqualsMatrixB)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) != matrix\.([a-zA-Z0-9_]+)$`, tt.matrixANotEqualsMatrixB)
			ctx.Step(`^the following matrix matrix\.([a-zA-Z0-9_]+):$`, tt.makeMatrixFromTable)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) \* matrix\.([a-zA-Z0-9_]+) is the following matrix:$`, tt.matrixATimesMatrixB)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) \* tuple\.([a-zA-Z0-9_]+) = tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.matrixATimesTupleB)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) \* tuple\.([a-zA-Z0-9_]+) = tuple\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.matrixATimesTupleBPoint)
			ctx.Step(`^IdentityMatrix \* tuple\.([a-zA-Z0-9_]+) = tuple\.([a-zA-Z0-9_]+)$`, tt.identityTimesTuple)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) \* IdentityMatrix = matrix\.([a-zA-Z0-9_]+)$`, tt.matrixTimesIdentity)
			ctx.Step(`^transpose\(matrix\.([a-zA-Z0-9_]+)\) is the following matrix:$`, tt.transposematrixAIsTheFollowingMatrix)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) = identity_matrix$`, tt.matrixAIdentity_matrix)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) ← transpose\(identity_matrix\)$`, tt.matrixATransposeidentity_matrix)
			ctx.Step(`^determinant\(matrix.([a-zA-Z0-9_]+)\) = (-?\d+(?:\.\d+)?)$`, tt.determinantMatrix)
			ctx.Step(`^the following 2x2 matrix ([a-zA-Z0-9_]+):$`, tt.theFollowing2X2MatrixA)
			ctx.Step(`^submatrix\(matrix\.([a-zA-Z0-9_]+), (\d+), (\d+)\) is the following (\d+)x(\d+) matrix:$`, tt.submatrixmatrixAIsTheFollowingXMatrix)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) ← submatrix\(matrix\.([a-zA-Z0-9_]+), (\d+), (\d+)\)$`, tt.matrixBSubmatrixmatrixA)
			ctx.Step(`^minor\(matrix\.([a-zA-Z0-9_]+), (\d+), (\d+)\) = (-?\d+)$`, tt.minormatrixA)
			ctx.Step(`^cofactor\(matrix\.([a-zA-Z0-9_]+), (\d+), (\d+)\) = (-?\d+)$`, tt.cofactormatrixA)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) is invertible$`, tt.matrixAIsInvertible)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) is not invertible$`, tt.matrixAIsNotInvertible)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+)\[(\d+),(\d+)\] = (-?\d+)\/(\d+)$`, tt.matrixBRowColEquals)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) ← inverse\(matrix\.([a-zA-Z0-9_]+)\)$`, tt.matrixBInversematrixA)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) is the following (\d+)x(\d+) matrix:$`, tt.matrixBIsTheFollowingXMatrix)
			ctx.Step(`^inverse\(matrix\.([a-zA-Z0-9_]+)\) is the following (\d+)x(\d+) matrix:$`, tt.inversematrixAIsTheFollowingXMatrix)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) \* inverse\(matrix\.([a-zA-Z0-9_]+)\) = matrix\.([a-zA-Z0-9_]+)$`, tt.matrixCInversematrixBMatrixA)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) ← matrix\.([a-zA-Z0-9_]+) \* matrix\.([a-zA-Z0-9_]+)$`, tt.matrixCMatrixAMatrixB)

			// TRANSFORMS
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) \* tuple\.([a-zA-Z0-9_]+) = point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.matrixinvTuplepPoint)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) ← translation\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.matrixtransformTranslation)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) \* tuple\.([a-zA-Z0-9_]+) = point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.matrixtransformTuplepPoint)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) \* tuple\.([a-zA-Z0-9_]+) = tuple\.([a-zA-Z0-9_]+)$`, tt.matrixtransformTuplevTuplev)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) \* tuple\.([a-zA-Z0-9_]+) = vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.matrixinvTuplevVector)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) ← scaling\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.matrixtransformScaling)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) ← rotation_x\(π \/ (-?\d+(?:\.\d+)?)\)$`, tt.transformPiFractionRotation_x)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) ← rotation_y\(π \/ (\d+)\)$`, tt.transformPiFractionRotation_y)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) ← rotation_z\(π \/ (\d+)\)$`, tt.transformPiFractionRotation_z)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) \* tuple\.([a-zA-Z0-9_]+) = point\((-?\d+(?:\.\d+)?), (-)?√(\d+)\/(\d+), (-)?√(\d+)\/(\d+)\)$`, tt.matrixYZSqrt)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) \* tuple\.([a-zA-Z0-9_]+) = point\((-)?√(\d+)\/(\d+), (-)?√(\d+)\/(\d+), (-?\d+(?:\.\d+)?)\)$`, tt.matrixXYSqrt)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) \* tuple\.([a-zA-Z0-9_]+) = point\((-)?√(\d+)\/(\d+), (-?\d+(?:\.\d+)?), (-)?√(\d+)\/(\d+)\)$`, tt.matrixXZSqrt)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) ← shearing\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.matrixtransformShearing)

			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) ← matrix\.([a-zA-Z0-9_]+) \* matrix\.([a-zA-Z0-9_]+) \* matrix\.([a-zA-Z0-9_]+)$`, tt.matrixTMatrixCMatrixBMatrixA)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) ← matrix\.([a-zA-Z0-9_]+) \* tuple\.([a-zA-Z0-9_]+)$`, tt.tuplepMatrixATuplep)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) = point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.tuplepPointEquals)

			// RAYS
			ctx.Step(`^ray\.([a-zA-Z0-9_]+) ← ray\(tuple.([a-zA-Z0-9_]+), tuple.([a-zA-Z0-9_]+)\)$`, tt.rayrRayoriginDirection)
			ctx.Step(`^ray\.([a-zA-Z0-9_]+)\.direction = tuple.([a-zA-Z0-9_]+)$`, tt.rayrdirectionDirection)
			ctx.Step(`^ray\.([a-zA-Z0-9_]+)\.origin = tuple.([a-zA-Z0-9_]+)$`, tt.rayroriginOrigin)
			ctx.Step(`^position\(ray\.([a-zA-Z0-9_]+), (-?\d+(?:\.\d+)?)\) = point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)$`, tt.positionrayrPoint)
			ctx.Step(`^ray\.([a-zA-Z0-9_]+) ← ray\(point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\), vector\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)\)$`, tt.rayrRaypointVector)
			ctx.Step(`^ray\.([a-zA-Z0-9_]+)\.direction = vector\((-?\d+(?:\.\d+)?),\s*(-?\d+(?:\.\d+)?),\s*(-?\d+(?:\.\d+)?)\)$`, tt.rayrDirectionVector)
			ctx.Step(`^ray\.([a-zA-Z0-9_]+)\.origin = point\((-?\d+(?:\.\d+)?),\s*(-?\d+(?:\.\d+)?),\s*(-?\d+(?:\.\d+)?)\)$`, tt.rayrOriginPoint)
			ctx.Step(`^ray\.([a-zA-Z0-9_]+) ← transform\(ray\.([a-zA-Z0-9_]+), matrix.([a-zA-Z0-9_]+)\)$`, tt.rayrTransformrayrM)

			// SHAPES AND INTERSECTIONS
			ctx.Step(`^slice\.([a-zA-Z0-9_]+)\[(\d+)\] = (-?\d+(?:\.\d+)?)$`, tt.sliceIndexEquals)
			ctx.Step(`^sphere\.([a-zA-Z0-9_]+) ← sphere\(\)$`, tt.defaultSphere)
			ctx.Step(`^slice\.([a-zA-Z0-9_]+)\.count = (\d+)$`, tt.sliceCount)
			ctx.Step(`^slice\.([a-zA-Z0-9_]+) ← intersect\(sphere\.([a-zA-Z0-9_]+), ray\.([a-zA-Z0-9_]+)\)$`, tt.slicexsIntersectspheresRayr)
			ctx.Step(`^intersection\.([a-zA-Z0-9_]+) ← intersection\((-?\d+(?:\.\d+)?), sphere\.([a-zA-Z0-9_]+)\)$`, tt.intersectioniIntersectionSpheres)
			ctx.Step(`^intersection\.([a-zA-Z0-9_]+)\.object = sphere\.([a-zA-Z0-9_]+)$`, tt.intersectioniobjectSpheres)
			ctx.Step(`^intersection\.([a-zA-Z0-9_]+)\.t = (-?\d+(?:\.\d+)?)$`, tt.intersectionit)
			ctx.Step(`^set_transform\(sphere\.([a-zA-Z0-9_]+), matrix\.([a-zA-Z0-9_]+)\)$`, tt.set_transformspheresMatrixt)
			ctx.Step(`^sphere\.([a-zA-Z0-9_]+)\.transform = identity_matrix$`, tt.spherestransformIdentity_matrix)
			ctx.Step(`^sphere\.([a-zA-Z0-9_]+)\.transform = matrix\.([a-zA-Z0-9_]+)$`, tt.spherestransformMatrixt)
			ctx.Step(`^set_transform\(sphere\.([a-zA-Z0-9_]+), scaling\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)\)$`, tt.setTransformspheresScaling)
			ctx.Step(`^set_transform\(sphere\.([a-zA-Z0-9_]+), translation\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)\)$`, tt.setTransformspheresTranslation)

			//ctx.Step(`^intersection\.([a-zA-Z0-9_]+) ← hit\(slice\.([a-zA-Z0-9_]+)\)$`, tt.intersectioniHitslicexs)

			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+) ← intersections\(intersection\.([a-zA-Z0-9_]+), intersection\.([a-zA-Z0-9_]+)\)$`, tt.slicexsIntersectionsintersectioniIntersectioni)
			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+)\[(\d+)\]\.t = (\d+)$`, tt.slicexsDT)
			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+)\.count = (\d+)$`, tt.arrayintersectionsxscount)
			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+) ← intersect\(sphere\.([a-zA-Z0-9_]+), ray\.([a-zA-Z0-9_]+)\)$`, tt.arrayintersectionsxsIntersectspheresRayr)
			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+)\[(\d+)\]\.object = sphere\.([a-zA-Z0-9_]+)$`, tt.arrayintersectionsxsObjectSpheres)
			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+)\[(\d+)\] = (-?\d+(?:\.\d+)?)$`, tt.arrayintersectionsxs)
			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+) ← intersections\(intersection\.([a-zA-Z0-9_]+), intersection\.([a-zA-Z0-9_]+), intersection\.([a-zA-Z0-9_]+), intersection\.([a-zA-Z0-9_]+)\)$`, tt.arrayintersectionsxsIntersectionsintersectioniIntersectioniIntersectioniIntersectioni)
			ctx.Step(`^intersection\.([a-zA-Z0-9_]+) ← hit\(arrayintersections\.([a-zA-Z0-9_]+)\)$`, tt.intersectioniHitarrayintersectionsxs)
			ctx.Step(`^intersection\.([a-zA-Z0-9_]+) = intersection\.([a-zA-Z0-9_]+)$`, tt.intersectioniIntersectioni)
			ctx.Step(`^intersection\.([a-zA-Z0-9_]+) is nothing$`, tt.intersectioniIsNothing)
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
func (tt *tupletest) matrixATimesTupleBPoint(varName1, varName2 string, x, y, z float64) error {
	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("matrix %s not available", varName1)
	}
	b, ok := tt.Tuples[varName2]
	if !ok {
		return fmt.Errorf("matrix %s not available", varName2)
	}
	d := a.MultiplyTuple(b)
	if d.EqualsTuple(NewPoint(x, y, z)) {
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
	d := IdentityMatrix()

	if d.MultiplyTuple(a).EqualsTuple(a) {
		return nil
	}
	return fmt.Errorf("2matrix multiplication by tuple fail\n%s\n%s", d.ToString(), a.ToString())

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

	d := a.MultiplyMatrix(IdentityMatrix())
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

func (tt *tupletest) matrixAIdentity_matrix(varName string) error {
	a, ok := tt.Matrices[varName]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName)
	}
	if a.EqualsMatrix(IdentityMatrix()) {
		return nil
	}
	return fmt.Errorf("Transposed identity wasn't identity")
}

func (tt *tupletest) matrixATransposeidentity_matrix(varName string) error {
	d := IdentityMatrix()
	tt.Matrices[varName] = d.Transpose()
	return nil
}

func (tt *tupletest) determinantMatrix(varName string, expect float64) error {
	a, ok := tt.Matrices[varName]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName)
	}
	if a.Determinant() == expect {
		return nil
	}
	return fmt.Errorf("determinant was %f not %f", a.Determinant(), expect)
}
func (tt *tupletest) theFollowing2X2MatrixA(varName string, values *godog.Table) error {
	tt.Matrices[varName] = tableToMatrix(values)
	return nil
}
func (tt *tupletest) submatrixmatrixAIsTheFollowingXMatrix(varName string, subRow, subCol, newRow, newCol int, expect *godog.Table) error {
	a, ok := tt.Matrices[varName]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName)
	}
	target := tableToMatrix(expect)
	check := a.Submatrix(subRow, subCol)
	if check.EqualsMatrix(target) {
		return nil
	}
	return fmt.Errorf("Submatrix fail %v vs %v\n", a.Submatrix(subRow, subCol), target)
}

func (tt *tupletest) matrixBSubmatrixmatrixA(varName1, varName2 string, row, col int) error {
	a, ok := tt.Matrices[varName2]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName2)
	}

	tt.Matrices[varName1] = a.Submatrix(row, col)
	return nil
}

func (tt *tupletest) minormatrixA(varName1 string, row, col int, target float64) error {

	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName1)
	}

	if a.Minor(row, col) == target {
		return nil
	}
	return fmt.Errorf("minor failed %f != %f", a.Minor(row, col), target)
}

func (tt *tupletest) cofactormatrixA(varName1 string, row, col int, target float64) error {
	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName1)
	}

	if a.Cofactor(row, col) == target {
		return nil
	}
	return fmt.Errorf("cofactor failed %f != %f", a.Cofactor(row, col), target)
}

func (tt *tupletest) matrixAIsInvertible(varName1 string) error {
	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName1)
	}
	if a.Invertable() {
		return nil
	}
	return fmt.Errorf("Matrix is not invertable")
}

func (tt *tupletest) matrixAIsNotInvertible(varName1 string) error {
	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName1)
	}
	if !a.Invertable() {
		return nil
	}
	return fmt.Errorf("Matrix is invertable")
}

func (tt *tupletest) matrixBRowColEquals(varName1 string, x, y int, nom, denom float64) error {
	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName1)
	}

	if epsilonEquals(a.Cells[x][y], nom/denom) {
		return nil
	}
	return fmt.Errorf("Cell %d,%d is %f not %f", x, y, a.Cells[x][y], nom/denom)
}

func (tt *tupletest) matrixBInversematrixA(varName1, varName2 string) error {
	a, ok := tt.Matrices[varName2]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName2)
	}
	tt.Matrices[varName1] = a.Inverse()
	return nil
}

func (tt *tupletest) matrixBIsTheFollowingXMatrix(varName1 string, x, y int, expect *godog.Table) error {
	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName1)
	}
	bob := tableToMatrix(expect)

	if a.EqualsMatrix(bob) {
		return nil
	}
	return fmt.Errorf("Matrices are not equal \n%v\n%v", a, bob)
}

func (tt *tupletest) inversematrixAIsTheFollowingXMatrix(varName1 string, arg1, arg2 int, expect *godog.Table) error {
	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName1)
	}
	bob := tableToMatrix(expect)
	b := a.Inverse()

	if b.EqualsMatrix(bob) {
		return nil
	}
	return fmt.Errorf("Matrices are not equal \n%v\n%v", b, bob)
}

func (tt *tupletest) matrixCInversematrixBMatrixA(varName1, varName2, varName3 string) error {
	c, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName1)
	}
	b, ok := tt.Matrices[varName2]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName2)
	}
	a, ok := tt.Matrices[varName3]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName3)
	}

	d := c.MultiplyMatrix(b.Inverse())
	if d.EqualsMatrix(a) {
		return nil
	}
	return fmt.Errorf("matrix inverse reversal fail")
}

func (tt *tupletest) matrixCMatrixAMatrixB(varName1, varName2, varName3 string) error {
	b, ok := tt.Matrices[varName3]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName3)
	}
	a, ok := tt.Matrices[varName2]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName2)
	}

	tt.Matrices[varName1] = a.MultiplyMatrix(b)
	return nil
}

func (tt *tupletest) matrixinvTuplepPoint(varName1, varName2 string, x, y, z float64) error {
	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName1)
	}
	b, ok := tt.Tuples[varName2]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName2)
	}

	q := a.MultiplyTuple(b)
	if q.EqualsTuple(NewPoint(x, y, z)) {
		return nil
	}
	return fmt.Errorf("Matrix inversion by tuple fail %v", q)
}

func (tt *tupletest) matrixtransformTranslation(varName1 string, x, y, z float64) error {
	tt.Matrices[varName1] = NewTranslation(x, y, z)
	return nil
}

func (tt *tupletest) matrixtransformTuplepPoint(varName1, varName2 string, x, y, z float64) error {
	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName1)
	}
	b, ok := tt.Tuples[varName2]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName2)
	}

	d := a.MultiplyTuple(b)
	if d.EqualsTuple(NewPoint(x, y, z)) {
		return nil
	}
	return fmt.Errorf("transform by point fail")
}

func (tt *tupletest) matrixtransformTuplevTuplev(varName1, varName2 string) error {
	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName1)
	}
	b, ok := tt.Tuples[varName2]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName2)
	}

	d := a.MultiplyTuple(b)
	if d.EqualsTuple(b) {
		return nil
	}
	return fmt.Errorf("transform by vector fail")
}

func (tt *tupletest) matrixinvTuplevVector(varName1, varName2 string, x, y, z float64) error {
	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName1)
	}
	b, ok := tt.Tuples[varName2]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName2)
	}

	d := a.MultiplyTuple(b)

	if d.EqualsTuple(NewVector(x, y, z)) {
		return nil
	}
	return fmt.Errorf("transform by vector fail")
}

func (tt *tupletest) matrixtransformScaling(varName1 string, x, y, z float64) error {
	tt.Matrices[varName1] = NewScaling(x, y, z)
	return nil
}

func (tt *tupletest) transformPiFractionRotation_x(varName1 string, denom float64) error {
	tt.Matrices[varName1] = NewRotationX(math.Pi / denom)
	return nil
}

func (tt *tupletest) transformPiFractionRotation_y(varName1 string, denom float64) error {
	tt.Matrices[varName1] = NewRotationY(math.Pi / denom)
	return nil
}
func (tt *tupletest) transformPiFractionRotation_z(varName1 string, denom float64) error {
	tt.Matrices[varName1] = NewRotationZ(math.Pi / denom)
	return nil
}

func (tt *tupletest) matrixYZSqrt(varName1, varName2 string, x float64, ys string, y1, y2 float64, zs string, z1, z2 float64) error {
	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName1)
	}
	b, ok := tt.Tuples[varName2]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName2)
	}

	y := math.Sqrt(y1) / y2
	if ys == "-" {
		y *= -1
	}
	z := math.Sqrt(z1) / z2
	if zs == "-" {
		z *= -1
	}

	p := NewPoint(x, y, z)

	if a.MultiplyTuple(b).EqualsTuple(p) {
		return nil
	}
	return fmt.Errorf("YZ Multiplytuple sqrt")
}

func (tt *tupletest) matrixXYSqrt(varName1, varName2 string, xs string, x1, x2 float64, ys string, y1, y2 float64, z float64) error {
	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName1)
	}
	b, ok := tt.Tuples[varName2]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName2)
	}

	x := math.Sqrt(x1) / x2
	if xs == "-" {
		x *= -1
	}
	y := math.Sqrt(y1) / y2
	if ys == "-" {
		y *= -1
	}

	p := NewPoint(x, y, z)

	if a.MultiplyTuple(b).EqualsTuple(p) {
		return nil
	}
	return fmt.Errorf("XY Multiplytuple sqrt")
}

func (tt *tupletest) matrixXZSqrt(varName1, varName2 string, xs string, x1, x2 float64, y float64, zs string, z1, z2 float64) error {
	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName1)
	}
	b, ok := tt.Tuples[varName2]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName2)
	}

	x := math.Sqrt(x1) / x2
	if xs == "-" {
		x *= -1
	}
	z := math.Sqrt(z1) / z2
	if zs == "-" {
		z *= -1
	}

	p := NewPoint(x, y, z)

	if a.MultiplyTuple(b).EqualsTuple(p) {
		return nil
	}
	return fmt.Errorf("ZX Multiplytuple sqrt \n%v\n%v", a.MultiplyTuple(b), p)
}

func (tt *tupletest) matrixtransformShearing(varName1 string, xy, xz, yx, yz, zx, zy float64) error {
	tt.Matrices[varName1] = NewShearing(xy, xz, yx, yz, zx, zy)
	return nil

}
func (tt *tupletest) matrixTMatrixCMatrixBMatrixA(varName1, varName2, varName3, varName4 string) error {
	c, ok := tt.Matrices[varName2]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName2)
	}
	b, ok := tt.Matrices[varName3]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName3)
	}
	a, ok := tt.Matrices[varName4]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName4)
	}

	d := c.MultiplyMatrix(b)
	d = d.MultiplyMatrix(a)
	tt.Matrices[varName1] = d
	return nil
}

func (tt *tupletest) tuplepMatrixATuplep(varName1, varName2, varName3 string) error {

	b, ok := tt.Matrices[varName2]
	if !ok {
		return fmt.Errorf("!matrix %s not available", varName2)
	}
	a, ok := tt.Tuples[varName3]
	if !ok {
		return fmt.Errorf("tuple %s not available", varName3)
	}

	tt.Tuples[varName1] = b.MultiplyTuple(a)

	return nil
}

func (tt *tupletest) tuplepPointEquals(varName string, x, y, z float64) error {
	a, ok := tt.Tuples[varName]
	if !ok {
		return fmt.Errorf("tuple %s not available", varName)
	}
	if a.EqualsTuple(NewPoint(x, y, z)) {
		return nil
	}
	return fmt.Errorf("Point mismatch %s <-> %s", a.ToString(), NewPoint(x, y, z).ToString())
}

func (tt *tupletest) rayrRayoriginDirection(varName1, varName2, varName3 string) error {
	o, ok := tt.Tuples[varName2]
	if !ok {
		return fmt.Errorf("tuple %s not available", varName2)
	}
	d, ok := tt.Tuples[varName3]
	if !ok {
		return fmt.Errorf("tuple %s not available", varName3)
	}
	tt.Rays[varName1] = NewRay(o, d)
	return nil
}

func (tt *tupletest) rayrdirectionDirection(varName1, varName2 string) error {
	a, ok := tt.Rays[varName1]
	if !ok {
		return fmt.Errorf("Ray %s not available", varName1)
	}
	b, ok := tt.Tuples[varName2]
	if !ok {
		return fmt.Errorf("Tuple %s not available", varName2)
	}

	if a.Direction.EqualsTuple(b) {
		return nil
	}
	return fmt.Errorf("Bad direction")

}

func (tt *tupletest) rayroriginOrigin(varName1, varName2 string) error {

	a, ok := tt.Rays[varName1]
	if !ok {
		return fmt.Errorf("Ray %s not available", varName1)
	}
	b, ok := tt.Tuples[varName2]
	if !ok {
		return fmt.Errorf("Tuple %s not available", varName2)
	}

	if a.Origin.EqualsTuple(b) {
		return nil
	}
	return fmt.Errorf("Bad Origin")
}

func (tt *tupletest) positionrayrPoint(varName1 string, t, x, y, z float64) error {
	a, ok := tt.Rays[varName1]
	if !ok {
		return fmt.Errorf("Ray %s not available", varName1)
	}
	if a.Position(t).EqualsTuple(NewPoint(x, y, z)) {
		return nil
	}
	return fmt.Errorf("Position on ray fail")
}

func (tt *tupletest) rayrRaypointVector(varName1 string, px, py, pz, vx, vy, vz float64) error {
	tt.Rays[varName1] = NewRay(NewPoint(px, py, pz), NewVector(vx, vy, vz))
	return nil
}

func (tt *tupletest) sliceIndexEquals(varName1 string, index int, value float64) error {
	a, ok := tt.Slices[varName1]
	if !ok {
		return fmt.Errorf("Slice %s not available", varName1)
	}

	if a[index] == value {
		return nil
	}
	return fmt.Errorf("Slice %s[%d] wrong %f", varName1, index, value)
}

func (tt *tupletest) defaultSphere(varName1 string) error {
	tt.Shapes[varName1] = NewSphere()
	return nil
}

func (tt *tupletest) sliceCount(varName1 string, count int) error {
	a, ok := tt.Slices[varName1]
	if !ok {
		return fmt.Errorf("Slice %s not available", varName1)
	}

	if len(a) == count {
		return nil
	}
	return fmt.Errorf("Wrong size %s - %d not %d", varName1, len(a), count)

}

func (tt *tupletest) slicexsIntersectspheresRayr(varName1, varName2, varName3 string) error {
	a, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("Sphere %s not available", varName2)
	}
	b, ok := tt.Rays[varName3]
	if !ok {
		return fmt.Errorf("Ray %s not available", varName3)
	}

	tt.ArrayIntersections[varName1] = a.Intersects(b)
	return nil
}

// func (tt *tupletest) intersectioniHitslicexs(varName1, varName2 string) error {
// 	b, ok := tt.Slices[varName2]
// 	if !ok {
// 		return fmt.Errorf("Slice %s not available", varName2)
// 	}
//
// 	tt.Intersections[varName1] = b.Hit()
// 	return godog.ErrPending
//
// }
//

func (tt *tupletest) intersectioniIntersectionSpheres(varName1 string, t float64, varName2 string) error {
	a, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("Sphere %s not available", varName2)
	}

	tt.Intersections[varName1] = NewIntersection(t, a)
	return nil
}

func (tt *tupletest) intersectioniobjectSpheres(varName1, varName2 string) error {
	a, ok := tt.Intersections[varName1]
	if !ok {
		return fmt.Errorf("Sphere %s not available", varName1)
	}

	b, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("Sphere %s not available", varName2)
	}

	if a.ObjectEquals(b) {
		return nil
	}
	return fmt.Errorf("Object mismatch for intersection %v|%v", a.Object, b)
}

func (tt *tupletest) intersectionit(varName1 string, value float64) error {
	a, ok := tt.Intersections[varName1]
	if !ok {
		return fmt.Errorf("Sphere %s not available", varName1)
	}

	if a.T == value {
		return nil
	}
	return fmt.Errorf("Big intersect fail")
}
func (tt *tupletest) slicexsIntersectionsintersectioniIntersectioni(varName1, varName2, varName3 string) error {
	b, ok := tt.Intersections[varName2]
	if !ok {
		return fmt.Errorf("Sphere %s not available", varName2)
	}
	c, ok := tt.Intersections[varName3]
	if !ok {
		return fmt.Errorf("Sphere %s not available", varName3)
	}
	tt.ArrayIntersections[varName1] = Intersections(b, c)
	return nil
}

func (tt *tupletest) slicexsDT(varName1 string, index int, expect float64) error {
	a, ok := tt.ArrayIntersections[varName1]
	if !ok {
		return fmt.Errorf("IntersectionArray %s not available", varName1)
	}
	if a[index].T == expect {
		return nil
	}
	return fmt.Errorf("IntersectionArray index %d fail, is %f not %f", index, a[index].T, expect)
}

func (tt *tupletest) arrayintersectionsxscount(varName1 string, count int) error {
	a, ok := tt.ArrayIntersections[varName1]
	if !ok {
		return fmt.Errorf("IntersectionArray %s not available", varName1)
	}
	if len(a) == count {
		return nil
	}
	return fmt.Errorf("Bad count of intersectionarray %s [%d not %d]", varName1, len(a), count)
}

func (tt *tupletest) arrayintersectionsxsIntersectspheresRayr(varName1, varName2, varName3 string) error {
	b, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("Sphere %s not available", varName2)
	}
	c, ok := tt.Rays[varName3]
	if !ok {
		return fmt.Errorf("Sphere %s not available", varName3)
	}
	tt.ArrayIntersections[varName1] = b.Intersects(c)
	return nil
}

// ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+)\[(\d+)\]\.object = sphere\.([a-zA-Z0-9_]+)$`, tt.arrayintersectionsxsObjectSpheres)
func (tt *tupletest) arrayintersectionsxsObjectSpheres(varName1 string, index int, varName2 string) error {
	a, ok := tt.ArrayIntersections[varName1]
	if !ok {
		return fmt.Errorf("Sphere2 %s not available", varName1)
	}
	b, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("Sphere3 %s not available", varName2)
	}

	if a[index].Object.Equals(b) {
		return nil
	}
	return fmt.Errorf("Array Intersections %s[%d].Object isn't sphere %s", varName1, index, varName2)

}

// ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+)\[(\d+)\] = (-?\d+(?:\.\d+)?)$`, tt.arrayintersectionsxs)
func (tt *tupletest) arrayintersectionsxs(varName1 string, index int, value float64) error {
	a, ok := tt.ArrayIntersections[varName1]
	if !ok {
		return fmt.Errorf("AyInters %s not available", varName1)
	}
	if a[index].T == value {
		return nil
	}
	return fmt.Errorf("Array intersect fail for %s[%d] - %f rather than %f", varName1, index, a[index].T, value)
}

func (tt *tupletest) arrayintersectionsxsIntersectionsintersectioniIntersectioniIntersectioniIntersectioni(varName1, varName2, varName3, varName4, varName5 string) error {
	b, ok := tt.Intersections[varName2]
	if !ok {
		return fmt.Errorf("Inters %s not available", varName2)
	}
	c, ok := tt.Intersections[varName3]
	if !ok {
		return fmt.Errorf("Inters %s not available", varName3)
	}
	d, ok := tt.Intersections[varName4]
	if !ok {
		return fmt.Errorf("Inters %s not available", varName4)
	}
	e, ok := tt.Intersections[varName5]
	if !ok {
		return fmt.Errorf("Inters %s not available", varName5)
	}

	tt.ArrayIntersections[varName1] = Intersections(b, c, d, e)
	return nil
}

// ctx.Step(`^intersection\.([a-zA-Z0-9_]+) ← hit\(arrayintersections\.([a-zA-Z0-9_]+)\)$`, tt.intersectioniHitarrayintersectionsxs)
func (tt *tupletest) intersectioniHitarrayintersectionsxs(varName1, varName2 string) error {
	b, ok := tt.ArrayIntersections[varName2]
	if !ok {
		return fmt.Errorf("AyInters %s not available", varName2)
	}
	ok, inter := Hit(b)
	if ok {
		tt.Intersections[varName1] = inter
	}
	return nil
}

// ctx.Step(`^intersection\.([a-zA-Z0-9_]+) = intersection\.([a-zA-Z0-9_]+)$`, tt.intersectioniIntersectioni)
func (tt *tupletest) intersectioniIntersectioni(varName1, varName2 string) error {
	a, ok := tt.Intersections[varName1]
	if !ok {
		return fmt.Errorf("Inters %s not available", varName1)
	}
	b, ok := tt.Intersections[varName2]
	if !ok {
		return fmt.Errorf("Inters %s not available", varName2)
	}
	if a.Equals(b) {
		return nil
	}
	return fmt.Errorf("Inters don't match %v,%v", a, b)
}

// ctx.Step(`^intersection\.([a-zA-Z0-9_]+) is nothing$`, tt.intersectioniIsNothing)
func (tt *tupletest) intersectioniIsNothing(varName1 string) error {
	_, ok := tt.Intersections[varName1]
	if ok {
		return fmt.Errorf("Inters %s should not be available", varName1)
	}
	return nil
}

func (tt *tupletest) rayrDirectionVector(varName1 string, x, y, z float64) error {
	a, ok := tt.Rays[varName1]
	if !ok {
		return fmt.Errorf("Ray %s not available", varName1)
	}
	if a.Direction.EqualsTuple(NewVector(x, y, z)) {
		return nil
	}
	return fmt.Errorf("ray direction %s is not %s", a.Direction.ToString(), NewVector(x, y, z).ToString())
}

func (tt *tupletest) rayrOriginPoint(varName1 string, x, y, z float64) error {
	a, ok := tt.Rays[varName1]
	if !ok {
		return fmt.Errorf("Ray %s not available", varName1)
	}
	if a.Origin.EqualsTuple(NewPoint(x, y, z)) {
		return nil
	}
	return fmt.Errorf("ray direction %s is not %s", a.Direction.ToString(), NewPoint(x, y, z).ToString())
}

func (tt *tupletest) rayrTransformrayrM(varName1, varName2, varName3 string) error {
	b, ok := tt.Rays[varName2]
	if !ok {
		return fmt.Errorf("Ray %s not available", varName2)
	}
	c, ok := tt.Matrices[varName3]
	if !ok {
		return fmt.Errorf("Matrix %s not available", varName3)
	}

	tt.Rays[varName1] = b.Transform(c)
	return nil
}

func (tt *tupletest) set_transformspheresMatrixt(varName1, varName2 string) error {
	a, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("Shape %s not available", varName1)
	}
	b, ok := tt.Matrices[varName2]
	if !ok {
		return fmt.Errorf("Matrix %s not available", varName2)
	}

	a.SetTransform(b)
	return nil
}

func (tt *tupletest) spherestransformIdentity_matrix(varName1 string) error {
	a, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("Shape %s not available", varName1)
	}
	jeff := a.GetTransform()
	if jeff.EqualsMatrix(IdentityMatrix()) {
		return nil
	}
	return fmt.Errorf("Sphere transform not identity")
}

func (tt *tupletest) spherestransformMatrixt(varName1, varName2 string) error {
	a, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("Shape %s not available", varName1)
	}
	b, ok := tt.Matrices[varName2]
	if !ok {
		return fmt.Errorf("Matrix %s not available", varName2)
	}
	jeff := a.GetTransform()
	if jeff.EqualsMatrix(b) {
		return nil
	}
	return fmt.Errorf("Sphere transform not %v", b)
}

func (tt *tupletest) setTransformspheresScaling(varName1 string, x, y, z float64) error {
	_, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("Shape %s not available", varName1)
	}
	tt.Shapes[varName1].SetTransform(NewScaling(x, y, z))
	return nil
}

func (tt *tupletest) setTransformspheresTranslation(varName1 string, x, y, z float64) error {
	_, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("Shape %s not available", varName1)
	}
	tt.Shapes[varName1].SetTransform(NewTranslation(x, y, z))
	return nil
}
