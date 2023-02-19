package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"regexp"
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
type lights map[string]Light
type materials map[string]Material
type worlds map[string]World
type computations map[string]Computations
type cameras map[string]Camera
type patterns map[string]Pattern
type floats map[string]float64

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
	Lights             lights
	Materials          materials
	Worlds             worlds
	Computations       computations
	Cameras            cameras
	Patterns           patterns
	Floats             floats
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
				tt.Lights = lights{}
				tt.Materials = materials{}
				tt.Worlds = worlds{}
				tt.Computations = computations{}
				tt.Cameras = cameras{}
				tt.Patterns = patterns{}
				tt.Floats = floats{}
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
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) ← scaling\(([^,]+?), ([^,]+?), ([^)]+?)\)$`, tt.matrixtransformScaling)
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
			ctx.Step(`^ray\.([a-zA-Z0-9_]+) ← ray\(point\((.*), (.*), (.*)\), vector\((.*), (.*), (.*)\)\)$`, tt.rayrRaypointVector)
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
			//ctx.Step(`^set_transform\(sphere\.([a-zA-Z0-9_]+), translation\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)\)$`, tt.setTransformspheresTranslation)

			//ctx.Step(`^intersection\.([a-zA-Z0-9_]+) ← hit\(slice\.([a-zA-Z0-9_]+)\)$`, tt.intersectioniHitslicexs)

			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+) ← intersections\(intersection\.([a-zA-Z0-9_]+), intersection\.([a-zA-Z0-9_]+)\)$`, tt.slicexsIntersectionsintersectioniIntersectioni)
			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+)\[(\d+)\]\.t = (.+)$`, tt.slicexsDT)
			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+)\.count = (\d+)$`, tt.arrayintersectionsxscount)
			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+) ← intersect\(sphere\.([a-zA-Z0-9_]+), ray\.([a-zA-Z0-9_]+)\)$`, tt.arrayintersectionsxsIntersectspheresRayr)
			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+)\[(\d+)\]\.object = sphere\.([a-zA-Z0-9_]+)$`, tt.arrayintersectionsxsObjectSpheres)
			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+)\[(\d+)\] = (-?\d+(?:\.\d+)?)$`, tt.arrayintersectionsxs)
			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+) ← intersections\(intersection\.([a-zA-Z0-9_]+), intersection\.([a-zA-Z0-9_]+), intersection\.([a-zA-Z0-9_]+), intersection\.([a-zA-Z0-9_]+)\)$`, tt.arrayintersectionsxsIntersectionsintersectioniIntersectioniIntersectioniIntersectioni)
			ctx.Step(`^intersection\.([a-zA-Z0-9_]+) ← hit\(arrayintersections\.([a-zA-Z0-9_]+)\)$`, tt.intersectioniHitarrayintersectionsxs)
			ctx.Step(`^intersection\.([a-zA-Z0-9_]+) = intersection\.([a-zA-Z0-9_]+)$`, tt.intersectioniIntersectioni)
			ctx.Step(`^intersection\.([a-zA-Z0-9_]+) is nothing$`, tt.intersectioniIsNothing)

			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) ← normal_at\(sphere\.([a-zA-Z0-9_]+), point\((-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?), (-?\d+(?:\.\d+)?)\)\)$`, tt.tuplenNormal_atspheresPoint)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) ← normal_at\(sphere\.([a-zA-Z0-9_]+), point\(√(\d+)\/(\d+), √(\d+)\/(\d+), √(\d+)\/(\d+)\)\)$`, tt.tuplenNormal_atspheresPointSqrt)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) = normalize\(([a-zA-Z0-9_]+)\)$`, tt.tuplenNormalizen)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) = vector\((.+?), (.+?), (.+?)\)$`, tt.tuplenVector)
			// ctx.Step(`^tuple\.([a-zA-Z0-9_]+) = vector\(√(\d+)\/(\d+), √(\d+)\/(\d+), √(\d+)\/(\d+)\)$`, tt.tuplenVectorSqrt)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) ← scaling\((.+?), (.+?), (.+?)\) \* rotation_z\((.+?)\)$`, tt.matrixmScalingRotation_z)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) ← normal_at\(sphere\.([a-zA-Z0-9_]+), point\((.+), (.+), (.+)\)\)$`, tt.matrixnNormal_atspheresPoint)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) = vector\((.+?), (.+?), (.+?)\)$`, tt.matrixnVector)
			ctx.Step(`^set_transform\(sphere.([a-zA-Z0-9_]+), translation\((.+), (.+), (.+)\)\)$`, tt.set_transformsTranslation)
			ctx.Step(`^set_transform\(sphere\.([a-zA-Z0-9_]+), matrix.([a-zA-Z0-9_]+)\)$`, tt.set_transformspheresM)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) ← vector\((.+?), (.+?), (.+?)\)$`, tt.tuplenVectorAssign)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) ← reflect\(([a-zA-Z0-9_]+), ([a-zA-Z0-9_]+)\)$`, tt.tuplerReflectvN)

			// Chapter 6
			ctx.Step(`^light\.([a-zA-Z0-9_]+) ← point_light\(tuple\.([a-zA-Z0-9_]+), colors\.([a-zA-Z0-9_]+)\)$`, tt.lightlightPoint_lighttuplepositionColorsintensity)
			ctx.Step(`^light\.([a-zA-Z0-9_]+)\.intensity = colors\.([a-zA-Z0-9_]+)$`, tt.lightlightintensityColorsintensity)
			ctx.Step(`^light\.([a-zA-Z0-9_]+)\.position = tuple\.([a-zA-Z0-9_]+)$`, tt.lightlightpositionTupleposition)

			ctx.Step(`^material\.([a-zA-Z0-9_]+) ← material\(\)$`, tt.mMaterial)
			ctx.Step(`^material\.([a-zA-Z0-9_]+)\.ambient = (.+)$`, tt.mAmbientEquals)
			ctx.Step(`^material\.([a-zA-Z0-9_]+) = material\(\)$`, tt.materialmMaterialEquals)
			ctx.Step(`^material\.([a-zA-Z0-9_]+) ← material\(\)$`, tt.materialmMaterialAssigned)
			ctx.Step(`^material\.([a-zA-Z0-9_]+) ← sphere\.([a-zA-Z0-9_]+)\.material$`, tt.materialmSpheresmaterial)
			ctx.Step(`^material\.([a-zA-Z0-9_]+)\.ambient ← (.+)$`, tt.materialmambientAssign)
			ctx.Step(`^material\.([a-zA-Z0-9_]+)\.color = color\((.+), (.+), (.+)\)$`, tt.mcolorColor)
			ctx.Step(`^material\.([a-zA-Z0-9_]+)\.diffuse = (.+)$`, tt.mdiffuse)
			ctx.Step(`^material\.([a-zA-Z0-9_]+)\.shininess = (.+)$`, tt.mshininess)
			ctx.Step(`^material\.([a-zA-Z0-9_]+)\.specular = (.+)$`, tt.mspecular)
			ctx.Step(`^sphere\.([a-zA-Z0-9_]+)\.material ← material\.([a-zA-Z0-9_]+)$`, tt.spheresmaterialMaterialmAssign)
			ctx.Step(`^sphere\.([a-zA-Z0-9_]+)\.material = material\.([a-zA-Z0-9_]+)$`, tt.spheresmaterialMaterialmEquals)
			ctx.Step(`^colors\.([a-zA-Z0-9_]+) = color\((.+?), (.+?), (.+?)\)$`, tt.colorsresultColor)
			ctx.Step(`^colors\.([a-zA-Z0-9_]+) ← lighting\(material\.([a-zA-Z0-9_]+), light\.([a-zA-Z0-9_]+), tuple\.([a-zA-Z0-9_]+), tuple\.([a-zA-Z0-9_]+), tuple\.([a-zA-Z0-9_]+)\)$`, tt.colorsresultLightingmaterialmLightlightTuplepositionTupleeyevTuplenormalv)
			ctx.Step(`^light\.([a-zA-Z0-9_]+) ← point_light\(point\((.+?), (.+?), (.+?)\), color\((.+?), (.+?), (.+?)\)\)$`, tt.lightlightPoint_lightpointColor)

			// Chapter 7
			ctx.Step(`^world\.([a-zA-Z0-9_]+) contains no objects$`, tt.worldwContainsNoObjects)
			ctx.Step(`^world\.([a-zA-Z0-9_]+) has no light source$`, tt.worldwHasNoLightSource)
			ctx.Step(`^world\.([a-zA-Z0-9_]+) ← world\(\)$`, tt.worldwWorld)
			ctx.Step(`^sphere\.([a-zA-Z0-9_]+) ← sphere\(\) with:$`, tt.spheresSphereWith)
			ctx.Step(`^world\.([a-zA-Z0-9_]+) contains sphere\.([a-zA-Z0-9_]+)$`, tt.worldwContainsSpheres)
			ctx.Step(`^world\.([a-zA-Z0-9_]+) ← default_world\(\)$`, tt.worldwDefault_world)
			ctx.Step(`^world\.([a-zA-Z0-9_]+)\.light = light.([a-zA-Z0-9_]+)$`, tt.worldwlightLight)
			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+) ← intersect_world\(world\.([a-zA-Z0-9_]+), ray\.([a-zA-Z0-9_]+)\)$`, tt.arrayintersectionsxsIntersect_worldworldwRayr)
			ctx.Step(`^computes\.([a-zA-Z0-9_]+) ← prepare_computations\(intersection\.([a-zA-Z0-9_]+), ray\.([a-zA-Z0-9_]+)\)$`, tt.computescompsPrepare_computationsintersectioniRayr)
			ctx.Step(`^computes\.([a-zA-Z0-9_]+)\.eyev = vector\((.+?), (.+?), (.+?)\)$`, tt.computescompseyevVector)
			ctx.Step(`^computes\.([a-zA-Z0-9_]+)\.normalv = vector\((.+?), (.+?), (.+?)\)$`, tt.computescompsnormalvVector)
			ctx.Step(`^computes\.([a-zA-Z0-9_]+)\.object = intersection\.([a-zA-Z0-9_]+)\.object$`, tt.computescompsobjectIntersectioniobject)
			ctx.Step(`^computes\.([a-zA-Z0-9_]+)\.point = point\((.+?), (.+?), (.+?)\)$`, tt.computescompspointPoint)
			ctx.Step(`^computes\.([a-zA-Z0-9_]+)\.t = intersection\.([a-zA-Z0-9_]+)\.t$`, tt.computescompstIntersectionit)

			ctx.Step(`^computes\.([a-zA-Z0-9_]+)\.inside = (true|false)$`, tt.computescompsinsideBool)
			ctx.Step(`^colors\.([a-zA-Z0-9_]+) ← shade_hit\(world.([a-zA-Z0-9_]+), computes\.([a-zA-Z0-9_]+)\)$`, tt.colorscShade_hitwComputescomps)
			ctx.Step(`^sphere\.([a-zA-Z0-9_]+) ← the first object in world\.([a-zA-Z0-9_]+)$`, tt.sphereshapeTheFirstObjectInWorldw)
			ctx.Step(`^sphere\.([a-zA-Z0-9_]+) ← the second object in world\.([a-zA-Z0-9_]+)$`, tt.sphereshapeTheSecondObjectInWorldw)
			ctx.Step(`^world\.([a-zA-Z0-9_]+)\.light ← point_light\(point\((.+?), (.+?), (.+?)\), color\((.+?), (.+?), (.+?)\)\)$`, tt.worldwlightPoint_lightpointColor)
			ctx.Step(`^colors\.([a-zA-Z0-9_]+) ← color_at\(world\.([a-zA-Z0-9_]+), ray\.([a-zA-Z0-9_]+)\)$`, tt.colorscColor_atworldwRayr)

			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) = scaling\(([^,]+), ([^,]+), ([^)]+)\)$`, tt.matrixtScalingEqual)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) = translation\(([^,]+), ([^,]+), ([^)]+)\)$`, tt.matrixtTranslationEqual)
			ctx.Step(`^matrix\.([a-zA-Z0-9_]+) ← view_transform\(tuple.([a-zA-Z0-9_]+), tuple.([a-zA-Z0-9_]+), tuple.([a-zA-Z0-9_]+)\)$`, tt.matrixtView_transformfromToUpSet)

			ctx.Step(`^camera\.([a-zA-Z0-9_]+) ← camera\(([^,]+), ([^,]+), ([^)]+)\)$`, tt.cameracCamera)
			ctx.Step(`^camera\.([a-zA-Z0-9_]+)\.field_of_view = (.+)$`, tt.cameracfield_of_view)
			ctx.Step(`^camera\.([a-zA-Z0-9_]+)\.hsize = (.+)$`, tt.camerachsize)
			ctx.Step(`^camera\.([a-zA-Z0-9_]+)\.pixel_size = (.+)$`, tt.cameracpixel_size)
			ctx.Step(`^camera\.([a-zA-Z0-9_]+)\.transform = identity_matrix$`, tt.cameractransformIdentity_matrix)
			ctx.Step(`^camera\.([a-zA-Z0-9_]+)\.vsize = (.+)$`, tt.cameracvsize)

			ctx.Step(`^camera\.([a-zA-Z0-9_]+)\.transform ← rotation_y\((.*)\) \* translation\((.*), (.*), (.*)\)$`, tt.cameractransformRotation_yTranslation)
			ctx.Step(`^ray\.([a-zA-Z0-9_]+) ← ray_for_pixel\(camera\.([a-zA-Z0-9_]+), (\d+), (\d+)\)$`, tt.rayrRay_for_pixelcamerac)
			ctx.Step(`^ray\.([a-zA-Z0-9_]+)\.direction = vector\((.*), (.*), (.*)\)$`, tt.rayrdirectionVector)
			ctx.Step(`^camera\.([a-zA-Z0-9_]+)\.transform ← view_transform\(tuple\.([a-zA-Z0-9_]+), tuple\.([a-zA-Z0-9_]+), tuple\.([a-zA-Z0-9_]+)\)$`, tt.cameractransformView_transformtuplefromTupletoTupleup)
			ctx.Step(`^canvas\.([a-zA-Z0-9_]+) ← render\(camera\.([a-zA-Z0-9_]+), world\.([a-zA-Z0-9_]+)\)$`, tt.canvasimageRendercameracWorldw)
			ctx.Step(`^pixel_at\(canvas\.([a-zA-Z0-9_]+), (\d+), (\d+)\) = color\((.*), (.*), (.*)\)$`, tt.pixel_atcanvasimageColor)

			// CHAPTER 8

			ctx.Step(`^colors\.([a-zA-Z0-9_]+) ← lighting\(material\.([a-zA-Z0-9_]+), light\.([a-zA-Z0-9_]+), tuple\.([a-zA-Z0-9_]+), tuple\.([a-zA-Z0-9_]+), tuple\.([a-zA-Z0-9_]+), true\)$`, tt.colorsresultLightingmaterialmLightlightTuplepositionTupleeyevTuplenormalvTrue)
			ctx.Step(`^is_shadowed\(world\.([a-zA-Z0-9_]+), tuple\.([a-zA-Z0-9_]+)\) is (true|false)$`, tt.is_shadowedworldwTuplepIsFalse)
			ctx.Step(`^sphere\.([a-zA-Z0-9_]+) is added to world\.([a-zA-Z0-9_]+)$`, tt.spheresIsAddedToWorldw)
			ctx.Step(`^computes\.([a-zA-Z0-9_]+)\.over_point\.z < -EPSILON\/2$`, tt.computescompsover_pointzEPSILON)
			ctx.Step(`^computes\.([a-zA-Z0-9_]+)\.point\.z > computes\.([a-zA-Z0-9_]+)\.over_point\.z$`, tt.computescompspointzComputescompsover_pointz)

			// CHAPTER 9

			ctx.Step(`^shapes\.([a-zA-Z0-9_]+) ← test_shape\(\)$`, tt.shapessTest_shape)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+)\.transform = identity_matrix$`, tt.shapesstransformIdentity_matrix)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+)\.transform = translation\((.+), (.+), (.+)\)$`, tt.shapesstransformTranslation)

			ctx.Step(`^material\.([a-zA-Z0-9_]+) ← shapes\.([a-zA-Z0-9_]+)\.material$`, tt.materialmShapessmaterial)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+)\.material ← material\.([a-zA-Z0-9_]+)$`, tt.shapessmaterialMaterialmAssign)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+)\.material = material\.([a-zA-Z0-9_]+)$`, tt.shapessmaterialMaterialmEqual)

			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+) ← intersect\(shapes\.([a-zA-Z0-9_]+), ray\.([a-zA-Z0-9_]+)\)$`, tt.arrayintersectionsxsIntersectshapessRayr)
			ctx.Step(`^set_transform\(shapes\.([a-zA-Z0-9_]+), scaling\((.+), (.+), (.+)\)\)$`, tt.set_transformshapessScaling)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+)\.saved_ray\.direction = vector\((.+), (.+), (.+)\)$`, tt.shapesssaved_raydirectionVector)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+)\.saved_ray\.origin = point\((.+), (.+), (.+)\)$`, tt.shapesssaved_rayoriginPoint)

			ctx.Step(`^set_transform\(shapes\.([a-zA-Z0-9_]+), matrix\.([a-zA-Z0-9_]+)\)$`, tt.set_transformshapessMatrixm)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) ← normal_at\(shapes\.([a-zA-Z0-9_]+), point\((.*), (.*), (.*)\)\)$`, tt.tuplenNormal_atshapessPoint)

			// CHAPTER ?PLANE
			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+) is empty$`, tt.arrayintersectionsxsIsEmpty)
			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+) ← local_intersect\(shapes\.([a-zA-Z0-9_]+), ray\.([a-zA-Z0-9_]+)\)$`, tt.arrayintersectionsxsLocal_intersectplanepRayr)
			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+)\[(\d+)\]\.object = shapes\.([a-zA-Z0-9_]+)$`, tt.arrayintersectionsxsObjectPlanep)
			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+)\[(\d+)\]\.t = (\d+)$`, tt.arrayintersectionsxsT)
			ctx.Step(`^color\.black ← color\((\d+), (\d+), (\d+)\)$`, tt.colorblackColor)
			ctx.Step(`^color\.white ← color\((\d+), (\d+), (\d+)\)$`, tt.colorwhiteColor)
			ctx.Step(`^pattern\.([a-zA-Z0-9_]+) ← stripe_pattern\(color\.white, color\.black\)$`, tt.patternpatternStripe_patterncolorwhiteColorblack)
			ctx.Step(`^pattern\.([a-zA-Z0-9_]+)\.a = color\.white$`, tt.patternpatternaColorwhite)
			ctx.Step(`^pattern\.([a-zA-Z0-9_]+)\.b = color\.black$`, tt.patternpatternbColorblack)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+) ← plane\(\)$`, tt.planepPlane)
			ctx.Step(`^stripe_at\(pattern\.([a-zA-Z0-9_]+), point\((.*), (.*), (.*)\)\) = color\.black$`, tt.stripe_atpatternpatternPointColorblack)
			ctx.Step(`^stripe_at\(pattern\.([a-zA-Z0-9_]+), point\((.*), (.*), (.*)\)\) = color\.white$`, tt.stripe_atpatternpatternPointColorwhite)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) ← local_normal_at\(shapes\.([a-zA-Z0-9_]+), point\((.+), (.+), (.+)\)\)$`, tt.tuplenLocal_normal_atplanepPoint)
			ctx.Step(`^colors\.([a-zA-Z0-9_]+) ← lighting\(material\.([a-zA-Z0-9_]+), light\.([a-zA-Z0-9_]+), point\((.+), (.+), (.+)\), tuple\.([a-zA-Z0-9_]+), tuple\.([a-zA-Z0-9_]+), false\)$`, tt.colorscLightingmaterialmLightlightPointTupleeyevTuplenormalvFalse)
			ctx.Step(`^material\.([a-zA-Z0-9_]+)\.diffuse ← (.+)$`, tt.materialmdiffuse)
			ctx.Step(`^material\.([a-zA-Z0-9_]+)\.pattern ← stripe_pattern\(color\((.+), (.+), (.+)\), color\((.+), (.+), (.+)\)\)$`, tt.materialmpatternStripe_patterncolorColor)
			ctx.Step(`^material\.([a-zA-Z0-9_]+)\.specular ← (.+)$`, tt.materialmspecular)

			ctx.Step(`^colors\.([a-zA-Z0-9_]+) = colors\.([a-zA-Z0-9_]+)$`, tt.colorscColorswhite)
			ctx.Step(`^colors\.([a-zA-Z0-9_]+) ← stripe_at_object\(pattern\.([a-zA-Z0-9_]+), shapes\.([a-zA-Z0-9_]+), point\((.+), (.+), (.+)\)\)$`, tt.colorscStripe_at_objectpatternpatternSphereobjectPoint)
			ctx.Step(`^set_pattern_transform\(pattern\.([a-zA-Z0-9_]+), scaling\((.+), (.+), (.+)\)\)$`, tt.set_pattern_transformpatternpatternSpherescaling)
			ctx.Step(`^set_pattern_transform\(pattern\.([a-zA-Z0-9_]+), translation\((.+), (.+), (.+)\)\)$`, tt.set_pattern_transformpatternpatternTranslation)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+) ← sphere\(\)$`, tt.shapesobjectSphere)

			ctx.Step(`^pattern\.([a-zA-Z0-9_]+) ← test_pattern\(\)$`, tt.patternpatternTest_pattern)
			ctx.Step(`^pattern\.([a-zA-Z0-9_]+)\.transform = identity_matrix$`, tt.patternpatterntransformIdentity_matrix)
			ctx.Step(`^pattern\.([a-zA-Z0-9_]+)\.transform = translation\((.*), (.*), (.*)\)$`, tt.patternpatterntransformTranslation)
			ctx.Step(`^colors\.([a-zA-Z0-9_]+) ← pattern_at_shape\(pattern\.([a-zA-Z0-9_]+), shapes\.([a-zA-Z0-9_]+), point\((.*), (.*), (.*)\)\)$`, tt.colorscPattern_at_shapepatternpatternShapesshapePoint)

			ctx.Step(`^pattern_at\(pattern\.([a-zA-Z0-9_]+), point\((.*), (.*), (.*)\)\) = color\((.*), (.*), (.*)\)$`, tt.pattern_atpatternpatternPointColor)
			ctx.Step(`^pattern_at\(pattern\.([a-zA-Z0-9_]+), point\((.*), (.*), (.*)\)\) = white$`, tt.pattern_atpatternpatternPointWhite)
			ctx.Step(`^pattern\.([a-zA-Z0-9_]+) ← gradient_pattern\(white, black\)$`, tt.patternpatternGradient_patternwhiteBlack)

			ctx.Step(`^pattern_at\(pattern\.([a-zA-Z0-9_]+), point\((.*), (.*), (.*)\)\) = black$`, tt.pattern_atpatternpatternPointBlack)
			ctx.Step(`^pattern\.([a-zA-Z0-9_]+) ← ring_pattern\(white, black\)$`, tt.patternpatternRing_patternwhiteBlack)

			ctx.Step(`^pattern\.([a-zA-Z0-9_]+) ← checkers_pattern\(white, black\)$`, tt.patternpatternCheckers_patternwhiteBlack)
			ctx.Step(`^material\.([a-zA-Z0-9_]+)\.reflective = (.+)$`, tt.materialmreflective)

			// 11
			ctx.Step(`^computes\.([a-zA-Z0-9_]+)\.reflectv = vector\((.*), (.*), (.*)\)$`, tt.computescompsreflectvVector)
			ctx.Step(`^intersection\.([a-zA-Z0-9_]+) ← intersection\((.*), shapes\.([a-zA-Z0-9_]+)\)$`, tt.intersectioniIntersectionShapesshape)
			ctx.Step(`^ray\.([a-zA-Z0-9_]+) ← ray\(point\((.*), (.*), (.*)\), vector\((.*), (.*), (.*)\)\)$`, tt.rayrRaypointVector)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+) ← plane\(\)$`, tt.shapesshapePlane)

			ctx.Step(`^colors\.([a-zA-Z0-9_]+) ← reflected_color\(world\.([a-zA-Z0-9_]+), computes\.([a-zA-Z0-9_]+)\)$`, tt.colorscolorReflected_colorworldwComputescomps)
			ctx.Step(`^colors\.([a-zA-Z0-9_]+) ← reflected_color\(world\.([a-zA-Z0-9_]+), computes\.([a-zA-Z0-9_]+), 0\)$`, tt.colorscolorReflected_colorworldwComputescomps2)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+) ← the second object in world\.([a-zA-Z0-9_]+)$`, tt.shapesshapeTheSecondObjectInWorldw)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+)\.material\.ambient ← (\d+)$`, tt.shapesshapematerialambient)

			ctx.Step(`^shapes\.([a-zA-Z0-9_]+) is added to world\.([a-zA-Z0-9_]+)$`, tt.shapesshapeIsAddedToWorldw)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+) ← plane\(\) with:$`, tt.shapesshapePlaneWith)

			ctx.Step(`^color_at\(world\.([a-zA-Z0-9_]+), ray\.([a-zA-Z0-9_]+)\) should terminate successfully$`, tt.color_atworldwRayrShouldTerminateSuccessfully)

			ctx.Step(`^material\.([a-zA-Z0-9_]+)\.refractive_index = (.+)$`, tt.materialmrefractive_index)
			ctx.Step(`^material\.([a-zA-Z0-9_]+)\.transparency = (.+)$`, tt.materialmtransparency)

			ctx.Step(`^sphere\.([a-zA-Z0-9_]+) ← glass_sphere\(\)$`, tt.spheresGlass_sphere)
			ctx.Step(`^sphere\.([a-zA-Z0-9_]+)\.material\.refractive_index = (.+)$`, tt.spheresmaterialrefractive_index)
			ctx.Step(`^sphere\.([a-zA-Z0-9_]+)\.material\.transparency = (.+)$`, tt.spheresmaterialtransparency)

			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+) ← intersections\((.+):A, (.+):B, (.+):C, (.+):B, (.+):C, (.+):A\)$`, tt.arrayintersectionsxsIntersectionsABCBCA)
			ctx.Step(`^computes\.([a-zA-Z0-9_]+) ← prepare_computations\(arrayintersections\.([a-zA-Z0-9_]+)\[(\d+)\], ray\.([a-zA-Z0-9_]+), arrayintersections\.([a-zA-Z0-9_]+)\)$`, tt.computescompsPrepare_computationsarrayintersectionsxsRayrArrayintersectionsxs)
			ctx.Step(`^computes\.([a-zA-Z0-9_]+)\.n1 = (.+)$`, tt.computescompsn1)
			ctx.Step(`^computes\.([a-zA-Z0-9_]+)\.n2 = (.+)$`, tt.computescompsn2)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+) ← glass_sphere\(\) with:$`, tt.shapesGlass_sphereWith)

			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+) ← intersections\(intersection\.([a-zA-Z0-9_]+)\)$`, tt.arrayintersectionsxsIntersectionsintersectioni)
			ctx.Step(`^computes\.([a-zA-Z0-9_]+) ← prepare_computations\(intersection\.([a-zA-Z0-9_]+), ray\.([a-zA-Z0-9_]+), arrayintersections\.([a-zA-Z0-9_]+)\)$`, tt.computescompsPrepare_computationsintersectioniRayrArrayintersectionsxs)
			ctx.Step(`^computes\.([a-zA-Z0-9_]+)\.point\.z < computes\.([a-zA-Z0-9_]+)\.under_point\.z$`, tt.computescompspointzComputescompsunder_pointz)
			ctx.Step(`^computes\.([a-zA-Z0-9_]+)\.under_point\.z > EPSILON\/2$`, tt.computescompsunder_pointzEPSILON)

			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+) ← intersections\(([^:]+):shapes\.([a-zA-Z0-9_]+), ([^:]+):shapes\.([a-zA-Z0-9_]+)\)$`, tt.arrayintersectionsxsIntersectionsShapesshapeShapesshape)
			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+) ← intersections\(([^:]+):shapes\.([a-zA-Z0-9_]+), ([^:]+):shapes\.([a-zA-Z0-9_]+), ([^:]+):shapes\.([a-zA-Z0-9_]+), ([^:]+):shapes\.([a-zA-Z0-9_]+)\)$`, tt.arrayintersectionsxsIntersectionsShapesshapeShapesshape4)
			ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+) ← intersections\(([^:]+):shapes\.([a-zA-Z0-9_]+)\)$`, tt.arrayintersectionsxsIntersectionsShapesshapeShapesshape1)
			ctx.Step(`^colors\.([a-zA-Z0-9_]+) ← refracted_color\(world\.([a-zA-Z0-9_]+), computes\.([a-zA-Z0-9_]+), (\d+)\)$`, tt.colorscRefracted_colorworldwComputescomps)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+) ← the first object in world\.([a-zA-Z0-9_]+)$`, tt.shapesshapeTheFirstObjectInWorldw)

			ctx.Step(`^shapes\.([a-zA-Z0-9_]+) has:$`, tt.shapesshapeHas)

			ctx.Step(`^colors\.([a-zA-Z0-9_]+) ← shade_hit\(world\.([a-zA-Z0-9_]+), computes\.([a-zA-Z0-9_]+), (\d+)\)$`, tt.colorscolorShade_hitworldwComputescomps)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+) ← sphere\(\) with:$`, tt.shapesballSphereWith)
			ctx.Step(`^floats\.([a-zA-Z0-9_]+) = (.*)$`, tt.floatsreflectance)
			ctx.Step(`^floats\.([a-zA-Z0-9_]+) ← schlick\(computes\.([a-zA-Z0-9_]+)\)$`, tt.floatsreflectanceSchlickcomputescomps)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+) ← glass_sphere\(\)$`, tt.shapesshapeGlass_sphere)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+) ← cube\(\)$`, tt.shapescCube)

			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) ← local_normal_at\(shapes\.([a-zA-Z0-9_]+), tuple\.([a-zA-Z0-9_]+)\)$`, tt.tuplenormalLocal_normal_atshapescTuplep)

			// 12

			ctx.Step(`^ray\.([a-zA-Z0-9_]+) ← ray\(point\((.+), (.+), (.+)\), tuple\.([a-zA-Z0-9_]+)\)$`, tt.rayrRaypointTupledirection)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+) ← cylinder\(\)$`, tt.shapescylCylinder)
			ctx.Step(`^tuple\.([a-zA-Z0-9_]+) ← normalize\(vector\((.+), (.+), (.+)\)\)$`, tt.tupledirectionNormalizevector)

			//13
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+)\.maximum = infinity$`, tt.shapescylmaximumInfinity)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+)\.minimum = -infinity$`, tt.shapescylminimumInfinity)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+)\.maximum ← (.+)$`, tt.shapescylmaximum)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+)\.minimum ← (.+)$`, tt.shapescylminimum)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+)\.closed = (true|false)$`, tt.shapescylclosedFalse)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+)\.closed ← (true|false)$`, tt.shapescylclosedTrue)

			ctx.Step(`^shapes\.([a-zA-Z0-9_]+) ← cone\(\)$`, tt.shapesshapeCone)

			// 14
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+) ← group\(\)$`, tt.shapesgGroup)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+) is empty$`, tt.shapesgIsEmpty)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+)\.parent is nothing$`, tt.shapessparentIsNothing)
			ctx.Step(`^add_child\(shapes\.([a-zA-Z0-9_]+), shapes\.([a-zA-Z0-9_]+)\)$`, tt.add_childshapesgShapess)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+) includes shapes\.([a-zA-Z0-9_]+)$`, tt.shapesgIncludesShapess)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+) is not empty$`, tt.shapesgIsNotEmpty)
			ctx.Step(`^shapes\.([a-zA-Z0-9_]+)\.parent = shapes\.([a-zA-Z0-9_]+)$`, tt.shapessparentShapesg)
			ctx.Step(`^set_transform\(shapes\.([a-zA-Z0-9_]+), translation\((.+), (.+), (.+)\)\)$`, tt.set_transformshapessTranslation)
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

func (tt *tupletest) matrixtransformScaling(varName1 string, x, y, z string) error {
	tt.Matrices[varName1] = NewScaling(StringToFloat(x), StringToFloat(y), StringToFloat(z))
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

func (tt *tupletest) rayrRaypointVector(varName1 string, px, py, pz, vx, vy, vz string) error {
	tt.Rays[varName1] = NewRay(
		NewPoint(StringToFloat(px), StringToFloat(py), StringToFloat(pz)),
		NewVector(StringToFloat(vx), StringToFloat(vy), StringToFloat(vz)),
	)
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

func (tt *tupletest) slicexsDT(varName1 string, index int, expect string) error {
	a, ok := tt.ArrayIntersections[varName1]
	if !ok {
		return fmt.Errorf("IntersectionArray %s not available", varName1)
	}
	if epsilonEquals(a[index].T, StringToFloat(expect)) {
		return nil
	}
	return fmt.Errorf("IntersectionArray index %d fail, is %f not %s", index, a[index].T, expect)
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

func (tt *tupletest) tuplenNormal_atspheresPoint(varName1, varName2 string, x, y, z float64) error {
	a, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("Shape %s not available", varName2)
	}
	tt.Tuples[varName1] = a.NormalAt(NewPoint(x, y, z))
	return nil
}

func (tt *tupletest) tuplenNormal_atspheresPointSqrt(varName1, varName2 string, xn, xd, yn, yd, zn, zd float64) error {
	a, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("Shape %s not available", varName2)
	}
	tt.Tuples[varName1] = a.NormalAt(NewPoint(math.Sqrt(xn)/xd, math.Sqrt(yn)/yd, math.Sqrt(zn)/zd))
	return nil
}

func (tt *tupletest) tuplenNormalizen(varName1, varName2 string) error {
	a, ok := tt.Tuples[varName1]
	if !ok {
		return fmt.Errorf("Tuple %s not available", varName1)
	}
	b, ok := tt.Tuples[varName2]
	if !ok {
		return fmt.Errorf("Tuple %s not available", varName2)
	}
	if a.EqualsTuple(b.Normalize()) {
		return nil
	}
	return fmt.Errorf("Tuple isn't self normalised")
}

func (tt *tupletest) tuplenVector(varName1 string, xS, yS, zS string) error {
	a, ok := tt.Tuples[varName1]
	if !ok {
		return fmt.Errorf("Tuple %s not available", varName1)
	}
	if a.EqualsTuple(NewVector(StringToFloat(xS), StringToFloat(yS), StringToFloat(zS))) {
		return nil
	}
	return fmt.Errorf("tuple vector poop")
}

func StringToFloat(incomingString string) float64 {
	var newNom, newDom float64
	newFraction := strings.Split(incomingString, "/")
	newNom = SpecificStringCases(newFraction[0])

	if len(newFraction) > 1 {
		newDom = SpecificStringCases(newFraction[1])
	} else {
		newDom = 1
	}
	return newNom / newDom
}

func SpecificStringCases(strInt string) float64 {
	strInt = strings.Trim(strInt, " ,")
	if strInt[0:1] == "-" {
		return StringToFloat(strInt[1:]) * -1
	} else if len(strInt) > 1 {
		if strInt[0:2] == "π" {
			return math.Pi
		} else if len(strInt) > 2 && strInt[1:2] == "π" {
			return StringToFloat(strInt[0:1]) * math.Pi
		}
		if len(strInt) > 2 && strInt[0:3] == "√" {
			return math.Sqrt(StringToFloat(strInt[3:]))
		}
	}
	var err error
	var newNumber float64
	newNumber, err = strconv.ParseFloat(strInt, 64)
	if err != nil {
		log.Fatalf("to heck with it [%s] [%s] |%s|", strInt, strInt[0:3], err)
	}

	return newNumber
}

func (tt *tupletest) matrixmScalingRotation_z(varName1 string, xS, yS, zS, rS string) error {
	x := StringToFloat(xS)
	y := StringToFloat(yS)
	z := StringToFloat(zS)
	r := StringToFloat(rS)
	a := NewScaling(x, y, z)
	tt.Matrices[varName1] = a.MultiplyMatrix(NewRotationZ(r))
	return nil
}

func (tt *tupletest) matrixnNormal_atspheresPoint(varName1, varName2 string, xS, yS, zS string) error {
	b, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("Shape %s not available", varName2)
	}
	tt.Tuples[varName1] = b.NormalAt(NewPoint(StringToFloat(xS), StringToFloat(yS), StringToFloat(zS)))
	return nil
}

func (tt *tupletest) matrixnVector(varName1 string, xS, yS, zS string) error {
	a, ok := tt.Tuples[varName1]
	if !ok {
		return fmt.Errorf("Tuple %s not available", varName1)
	}
	if a.EqualsTuple(NewVector(StringToFloat(xS), StringToFloat(yS), StringToFloat(zS))) {
		return nil
	}
	x := StringToFloat(xS)
	y := StringToFloat(yS)
	z := StringToFloat(zS)
	return fmt.Errorf("Matrix N to vector fail %v,%v [%f,%f,%f]", a, NewVector(x, y, z), x, y, z)
}

func (tt *tupletest) set_transformspheresM(varName1, varName2 string) error {
	a, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("Shape %s not available", varName1)
	}
	b, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("Matrix %s not available", varName2)
	}
	a.SetTransform(b)
	return nil
}

func (tt *tupletest) set_transformsTranslation(varName1, xS, yS, zS string) error {
	a, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("Shape %s not available", varName1)
	}
	a.SetTransform(NewTranslation(StringToFloat(xS), StringToFloat(yS), StringToFloat(zS)))
	return nil
}

func (tt *tupletest) tuplenVectorAssign(varName1, xS, yS, zS string) error {
	tt.Tuples[varName1] = NewVector(StringToFloat(xS), StringToFloat(yS), StringToFloat(zS))
	return nil
}

func (tt *tupletest) tuplerReflectvN(varName1, varName2, varName3 string) error {
	b, ok := tt.Tuples[varName2]
	if !ok {
		return fmt.Errorf("Tuples %s not available", varName2)
	}
	c, ok := tt.Tuples[varName3]
	if !ok {
		return fmt.Errorf("Tuples %s not available", varName3)
	}
	tt.Tuples[varName1] = b.Reflect(c)
	return nil
}

func (tt *tupletest) lightlightPoint_lighttuplepositionColorsintensity(varName1, varName2, varName3 string) error {
	b, ok := tt.Tuples[varName2]
	if !ok {
		return fmt.Errorf("Tuples %s not available", varName2)
	}
	c, ok := tt.Colors[varName3]
	if !ok {
		return fmt.Errorf("Colors %s not available", varName3)
	}
	tt.Lights[varName1] = NewLight(b, c)
	return nil
}

func (tt *tupletest) lightlightintensityColorsintensity(varName1, varName2 string) error {
	b, ok := tt.Colors[varName2]
	if !ok {
		return fmt.Errorf("Color %s not available", varName2)
	}
	a, ok := tt.Lights[varName1]
	if !ok {
		return fmt.Errorf("Lights %s not available", varName1)
	}
	if b.Equals(a.Intensity) {
		return nil
	}
	return fmt.Errorf("Intensity didn't equallight")
}

func (tt *tupletest) lightlightpositionTupleposition(varName1, varName2 string) error {
	a, ok := tt.Lights[varName1]
	if !ok {
		return fmt.Errorf("Lights %s not available", varName1)
	}
	b, ok := tt.Tuples[varName2]
	if !ok {
		return fmt.Errorf("Tuple %s not available", varName2)
	}
	if b.EqualsTuple(a.Position) {
		return nil
	}
	return fmt.Errorf("Position didn't equal light")
}

func (tt *tupletest) mMaterial(varName1 string) error {
	tt.Materials[varName1] = NewMaterial()
	return nil
}
func (tt *tupletest) mAmbientEquals(varName1, ambient string) error {
	a, ok := tt.Materials[varName1]
	if !ok {
		return fmt.Errorf("Material %s not available", varName1)
	}
	if epsilonEquals(a.Ambient, StringToFloat(ambient)) {
		return nil
	}
	return fmt.Errorf("Ambient fail")
}

func (tt *tupletest) materialmMaterialEquals(varName1 string) error {
	a, ok := tt.Materials[varName1]
	if !ok {
		return fmt.Errorf("Material %s not available", varName1)
	}
	if a.Equals(NewMaterial()) {
		return nil
	}
	return fmt.Errorf("Material match nope")
}

func (tt *tupletest) materialmMaterialAssigned(varName1 string) error {
	tt.Materials[varName1] = NewMaterial()
	return nil
}
func (tt *tupletest) materialmSpheresmaterial(varName1, varName2 string) error {
	b, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("Shape %s not available", varName2)
	}
	tt.Materials[varName1] = b.GetMaterial()
	return nil
}

func (tt *tupletest) materialmambientAssign(varName1, ambient string) error {
	a, ok := tt.Materials[varName1]
	if !ok {
		return fmt.Errorf("Material %s not available for ambient", varName1)
	}
	a.Ambient = StringToFloat(ambient)
	tt.Materials[varName1] = a
	return nil
}
func (tt *tupletest) mcolorColor(varName1, r, g, b string) error {
	a, ok := tt.Materials[varName1]
	if !ok {
		return fmt.Errorf("Material %s not available for color", varName1)
	}
	a.Color = NewColor(StringToFloat(r), StringToFloat(g), StringToFloat(b))
	return nil
}
func (tt *tupletest) mdiffuse(varName1, diffuse string) error {
	a, ok := tt.Materials[varName1]
	if !ok {
		return fmt.Errorf("Material %s not available for diffuse", varName1)
	}
	a.Diffuse = StringToFloat(diffuse)
	return nil
}
func (tt *tupletest) mshininess(varName1, shininess string) error {
	a, ok := tt.Materials[varName1]
	if !ok {
		return fmt.Errorf("Material %s not available for shien", varName1)
	}
	a.Shininess = StringToFloat(shininess)
	return nil
}
func (tt *tupletest) mspecular(varName1, specular string) error {
	a, ok := tt.Materials[varName1]
	if !ok {
		return fmt.Errorf("Material %s not available for specuklar", varName1)
	}
	a.Specular = StringToFloat(specular)
	return nil
}
func (tt *tupletest) spheresmaterialMaterialmAssign(varName1, varName2 string) error {
	_, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("Shape %s not available", varName1)
	}
	b, ok := tt.Materials[varName2]
	if !ok {
		return fmt.Errorf("Material %s not available for var assign", varName2)
	}
	tt.Shapes[varName1].SetMaterial(b)
	return nil
}
func (tt *tupletest) spheresmaterialMaterialmEquals(varName1, varName2 string) error {
	a, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("Shape %s not available", varName1)
	}
	b, ok := tt.Materials[varName2]
	if !ok {
		return fmt.Errorf("Material %s not available for check", varName2)
	}
	if a.GetMaterial().Equals(b) {
		return nil
	}
	return fmt.Errorf("Material failed to match")
}

func (tt *tupletest) colorsresultColor(varName1, r, g, b string) error {
	a, ok := tt.Colors[varName1]
	if !ok {
		return fmt.Errorf("Color %s not available", varName1)
	}
	/*
		if STOPHERE {
			log.Fatalf(
				"\nPLANE %v\nTransform: %v\nTransparent: %v\nRefractive: %v\n\n"+
					"SPHERE\nColor: %v\nAmbient: %v\nTransform: %v\n\n"+
					"World: %v\n\n"+
					"Ray: %v\n\n"+
					"AI: %v\nID: %v\n\n"+
					"Comps: %v\n\n",
				tt.Shapes["floor"].GetID(),
				tt.Shapes["floor"].GetTransform(),
				tt.Shapes["floor"].GetMaterial().Transparency,
				tt.Shapes["floor"].GetMaterial().RefractiveIndex,
				tt.Shapes["ball"].GetMaterial().Color,
				tt.Shapes["ball"].GetMaterial().Ambient,
				tt.Shapes["ball"].GetTransform(),
				tt.Worlds["w"],
				tt.Rays["r"],
				tt.ArrayIntersections["xs"],
				tt.ArrayIntersections["xs"][0].Object.GetID(),
				tt.Computations["comps"],
			)
		}
	*/
	if a.Equals(NewColor(StringToFloat(r), StringToFloat(g), StringToFloat(b))) {
		return nil
	}

	return fmt.Errorf("Colour mismatch\n%v\n%v", a, NewColor(StringToFloat(r), StringToFloat(g), StringToFloat(b)))
}
func (tt *tupletest) colorsresultLightingmaterialmLightlightTuplepositionTupleeyevTuplenormalv(varName1, varNameM, varNameL, varName3, varName4, varName5 string) error {
	material, ok := tt.Materials[varNameM]
	if !ok {
		return fmt.Errorf("Material not avail")
	}
	light, ok := tt.Lights[varNameL]
	if !ok {
		return fmt.Errorf("Light not avail")
	}
	position, ok := tt.Tuples[varName3]
	if !ok {
		return fmt.Errorf("Position not avail")
	}
	eyev, ok := tt.Tuples[varName4]
	if !ok {
		return fmt.Errorf("EyeVector not avail")
	}
	normalv, ok := tt.Tuples[varName5]
	if !ok {
		return fmt.Errorf("NormalVector not avail")
	}
	inShadow := false
	tt.Colors[varName1] = Lighting(material, NewSphere(), light, position, eyev, normalv, inShadow)
	return nil
}
func (tt *tupletest) lightlightPoint_lightpointColor(varName1, x, y, z, r, g, b string) error {
	tt.Lights[varName1] = NewLight(NewPoint(StringToFloat(x), StringToFloat(y), StringToFloat(z)), NewColor(StringToFloat(r), StringToFloat(g), StringToFloat(b)))
	return nil
}

func (tt *tupletest) worldwContainsNoObjects(varName1 string) error {
	a, ok := tt.Worlds[varName1]
	if !ok {
		return fmt.Errorf("Color %s not available", varName1)
	}
	if len(a.Objects) == 0 {
		return nil
	}
	return fmt.Errorf("Howd this get objects?")
}
func (tt *tupletest) worldwHasNoLightSource(varName1 string) error {
	a, ok := tt.Worlds[varName1]
	if !ok {
		return fmt.Errorf("Color %s not available", varName1)
	}
	if len(a.Lights) == 0 {
		return nil
	}
	return fmt.Errorf("Howd this get lights?")
}
func (tt *tupletest) worldwWorld(varName1 string) error {
	tt.Worlds[varName1] = NewWorld()
	return nil
}

func (tt *tupletest) spheresSphereWith(varName1 string, tablevalue *godog.Table) error {
	sph := NewSphere()
	for _, row := range tablevalue.Rows {
		switch row.Cells[0].Value {
		case "color":
			rgb := strings.Split(strings.Trim(row.Cells[1].Value, " ()"), ",")
			sph.Material.Color = NewColor(
				StringToFloat(strings.Trim(rgb[0], " ")),
				StringToFloat(strings.Trim(rgb[1], " ")),
				StringToFloat(strings.Trim(rgb[2], " ")),
			)
		case "diffuse":
			sph.Material.Diffuse = StringToFloat(row.Cells[1].Value)
		case "specular":
			sph.Material.Specular = StringToFloat(row.Cells[1].Value)
		case "transform":
			bob := strings.Split(row.Cells[1].Value, "(")
			switch bob[0] {
			case "scaling":
				vals := strings.Split(strings.Trim(bob[1], ")"), ",")
				sph.Transform = NewScaling(
					StringToFloat(strings.Trim(vals[0], " ")),
					StringToFloat(strings.Trim(vals[1], " ")),
					StringToFloat(strings.Trim(vals[2], " ")))
			case "translation":
				vals := strings.Split(strings.Trim(bob[1], ")"), ",")
				sph.Transform = NewTranslation(
					StringToFloat(strings.Trim(vals[0], " ")),
					StringToFloat(strings.Trim(vals[1], " ")),
					StringToFloat(strings.Trim(vals[2], " ")))

			}
		}
	}
	tt.Shapes[varName1] = sph
	return nil
}
func (tt *tupletest) worldwContainsSpheres(varName1, varName2 string) error {

	a, ok := tt.Worlds[varName1]
	if !ok {
		return fmt.Errorf("World %s not available", varName1)
	}
	b, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("Shape %s not available for check", varName2)
	}

	if a.Contains(b) {
		return nil
	}
	return fmt.Errorf("Sphere %s doesn't contain %s", varName1, varName2)
}
func (tt *tupletest) worldwDefault_world(varName1 string) error {
	tt.Worlds[varName1] = DefaultWorld()
	return nil
}
func (tt *tupletest) worldwlightLight(varName1, varName2 string) error {
	_, ok := tt.Worlds[varName1]
	if !ok {
		return fmt.Errorf("World %s not available", varName1)
	}
	b, ok := tt.Lights[varName2]
	if !ok {
		return fmt.Errorf("light %s not available for check", varName2)
	}
	v := tt.Worlds[varName1]
	v.SetLight(b)
	tt.Worlds[varName1] = v
	return nil
}

func (tt *tupletest) arrayintersectionsxsIntersect_worldworldwRayr(varName1, varName2, varName3 string) error {
	a, ok := tt.Worlds[varName2]
	if !ok {
		return fmt.Errorf("World %s not available", varName2)
	}
	b, ok := tt.Rays[varName3]
	if !ok {
		return fmt.Errorf("light %s not available for check", varName3)
	}
	tt.ArrayIntersections[varName1] = a.Intersect(b)
	return nil
}

func (tt *tupletest) computescompsPrepare_computationsintersectioniRayr(varName1, varName2, varName3 string) error {
	b, ok := tt.Intersections[varName2]
	if !ok {
		return fmt.Errorf("Intersection %s not available", varName2)
	}
	c, ok := tt.Rays[varName3]
	if !ok {
		return fmt.Errorf("Ray %s not available for check", varName3)
	}
	tt.Computations[varName1] = b.PrepareComputations(c, map[int]Intersection{0: b})
	return nil
}

func (tt *tupletest) computescompseyevVector(varName1 string, x, y, z string) error {
	a, ok := tt.Computations[varName1]
	if !ok {
		return fmt.Errorf("Comp %s not available", varName1)
	}
	if a.Eyev.EqualsTuple(NewVector(StringToFloat(x), StringToFloat(y), StringToFloat(z))) {
		return nil
	}
	return fmt.Errorf("Comp Point :P %v,%v", a.Eyev, NewVector(StringToFloat(x), StringToFloat(y), StringToFloat(z)))
}

func (tt *tupletest) computescompsnormalvVector(varName1 string, x, y, z string) error {
	a, ok := tt.Computations[varName1]
	if !ok {
		return fmt.Errorf("Comp %s not available", varName1)
	}
	if a.Normalv.EqualsTuple(NewVector(StringToFloat(x), StringToFloat(y), StringToFloat(z))) {
		return nil
	}
	return fmt.Errorf("Normalv :P")
}

func (tt *tupletest) computescompsobjectIntersectioniobject(varName1 string, varName2 string) error {
	a, ok := tt.Computations[varName1]
	if !ok {
		return fmt.Errorf("Comp %s not available", varName1)
	}
	b, ok := tt.Intersections[varName2]
	if !ok {
		return fmt.Errorf("Intersection %s not available", varName2)
	}
	if a.Object.Equals(b.Object) {
		return nil
	}
	return fmt.Errorf("iObject :P")
}

func (tt *tupletest) computescompspointPoint(varName1 string, x, y, z string) error {
	a, ok := tt.Computations[varName1]
	if !ok {
		return fmt.Errorf("Comp %s not available", varName1)
	}
	if a.Point.EqualsTuple(NewPoint(StringToFloat(x), StringToFloat(y), StringToFloat(z))) {
		return nil
	}
	return fmt.Errorf("Comp Point :P %v,%v", a.Point, NewPoint(StringToFloat(x), StringToFloat(y), StringToFloat(z)))
}

func (tt *tupletest) computescompstIntersectionit(varName1 string, varName2 string) error {
	a, ok := tt.Computations[varName1]
	if !ok {
		return fmt.Errorf("Comp %s not available", varName1)
	}
	b, ok := tt.Intersections[varName2]
	if !ok {
		return fmt.Errorf("Intersection %s not available", varName2)
	}
	if epsilonEquals(a.T, b.T) {
		return nil
	}
	return fmt.Errorf("Comp T :P")
}

func (tt *tupletest) computescompsinsideBool(varName1 string, valString string) error {
	a, ok := tt.Computations[varName1]
	if !ok {
		return fmt.Errorf("Comp %s not available", varName1)
	}
	val := valString == "true"
	if a.Inside == val {
		return nil
	}
	return fmt.Errorf("Inside bool fail %v", val)
}

func (tt *tupletest) colorscShade_hitwComputescomps(varName1, varName2, varName3 string) error {
	b, ok := tt.Worlds[varName2]
	if !ok {
		return fmt.Errorf("Intersection %s not available", varName2)
	}
	c, ok := tt.Computations[varName3]
	if !ok {
		return fmt.Errorf("Intersection %s not available", varName3)
	}
	tt.Colors[varName1] = b.ShadeHit(c, maxReflects)
	return nil
}

func (tt *tupletest) sphereshapeTheFirstObjectInWorldw(varName1, varName2 string) error {
	b, ok := tt.Worlds[varName2]
	if !ok {
		return fmt.Errorf("Intersection %s not available", varName2)
	}
	tt.Shapes[varName1] = b.Objects[0]
	return nil
}

func (tt *tupletest) sphereshapeTheSecondObjectInWorldw(varName1, varName2 string) error {
	b, ok := tt.Worlds[varName2]
	if !ok {
		return fmt.Errorf("Intersection %s not available", varName2)
	}
	tt.Shapes[varName1] = b.Objects[1]
	return nil
}

func (tt *tupletest) worldwlightPoint_lightpointColor(varName1 string, x, y, z, r, g, b string) error {
	a := tt.Worlds[varName1]
	a.SetLight(NewLight(
		NewPoint(StringToFloat(x), StringToFloat(y), StringToFloat(z)),
		NewColor(StringToFloat(r), StringToFloat(g), StringToFloat(b))))
	tt.Worlds[varName1] = a
	return nil
}

func (tt *tupletest) colorscColor_atworldwRayr(varName1, varName2, varName3 string) error {
	b, ok := tt.Worlds[varName2]
	if !ok {
		return fmt.Errorf("World %s not available", varName2)
	}
	c, ok := tt.Rays[varName3]
	if !ok {
		return fmt.Errorf("Ray %s not available", varName3)
	}
	tt.Colors[varName1] = b.ColorAt(c, maxReflects)
	return nil
}
func (tt *tupletest) matrixtScalingEqual(varName1, x, y, z string) error {
	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("Matrix %s not available", varName1)
	}
	if a.EqualsMatrix(NewScaling(StringToFloat(x), StringToFloat(y), StringToFloat(z))) {
		return nil
	}
	return fmt.Errorf("Not scaling")
}
func (tt *tupletest) matrixtTranslationEqual(varName1, x, y, z string) error {
	a, ok := tt.Matrices[varName1]
	if !ok {
		return fmt.Errorf("Matrix %s not available", varName1)
	}
	if a.EqualsMatrix(NewTranslation(StringToFloat(x), StringToFloat(y), StringToFloat(z))) {
		return nil
	}
	return fmt.Errorf("Not translating")
}
func (tt *tupletest) matrixtView_transformfromToUpSet(varName1, varName2, varName3, varName4 string) error {
	b, ok := tt.Tuples[varName2]
	if !ok {
		return fmt.Errorf("Tuples %s not available", varName2)
	}
	c, ok := tt.Tuples[varName3]
	if !ok {
		return fmt.Errorf("Tuples %s not available", varName3)
	}
	d, ok := tt.Tuples[varName4]
	if !ok {
		return fmt.Errorf("Tuples %s not available", varName4)
	}

	tt.Matrices[varName1] = ViewTransform(b, c, d)
	return nil
}

func (tt *tupletest) cameracCamera(varName1 string, h, v int64, f string) error {
	tt.Cameras[varName1] = NewCamera(h, v, StringToFloat(f))
	return nil
}
func (tt *tupletest) cameracfield_of_view(varName1 string, f string) error {
	a, ok := tt.Cameras[varName1]
	if !ok {
		return fmt.Errorf("Camera %s not available", varName1)
	}
	if epsilonEquals(a.FieldOfView, StringToFloat(f)) {
		return nil
	}
	return fmt.Errorf("Bad FOF")
}
func (tt *tupletest) camerachsize(varName1 string, b int64) error {
	a, ok := tt.Cameras[varName1]
	if !ok {
		return fmt.Errorf("Camera %s not available", varName1)
	}
	if a.HSize == b {
		return nil
	}
	return fmt.Errorf("Bad H")
}
func (tt *tupletest) cameracpixel_size(varName1 string, f string) error {
	a, ok := tt.Cameras[varName1]
	if !ok {
		return fmt.Errorf("Camera %s not available", varName1)
	}
	if epsilonEquals(a.PixelSize, StringToFloat(f)) {
		return nil
	}
	return fmt.Errorf("Bad Pixels")
}
func (tt *tupletest) cameractransformIdentity_matrix(varName1 string) error {
	a, ok := tt.Cameras[varName1]
	if !ok {
		return fmt.Errorf("Camera %s not available", varName1)
	}
	if a.Transform.EqualsMatrix(IdentityMatrix()) {
		return nil
	}
	return fmt.Errorf("Bad Transform")
}
func (tt *tupletest) cameracvsize(varName1 string, b int64) error {
	a, ok := tt.Cameras[varName1]
	if !ok {
		return fmt.Errorf("Camera %s not available", varName1)
	}
	if a.VSize == b {
		return nil
	}
	return fmt.Errorf("Bad FOF")
}

func (tt *tupletest) cameractransformRotation_yTranslation(varName1, r, x, y, z string) error {
	a, ok := tt.Cameras[varName1]
	if !ok {
		return fmt.Errorf("Camera %s not available", varName1)
	}
	ry := NewRotationY(StringToFloat(r))
	a.SetTransform(ry.MultiplyMatrix(NewTranslation(StringToFloat(x), StringToFloat(y), StringToFloat(z))))
	tt.Cameras[varName1] = a
	return nil
}

func (tt *tupletest) rayrRay_for_pixelcamerac(varName1, varName2 string, x, y int64) error {
	b, ok := tt.Cameras[varName2]
	if !ok {
		return fmt.Errorf("Camera %s not available", varName2)
	}
	tt.Rays[varName1] = b.RayForPixel(x, y)
	return nil
}

func (tt *tupletest) rayrdirectionVector(varName1, x, y, z string) error {
	a, ok := tt.Rays[varName1]
	if !ok {
		return fmt.Errorf("Camera %s not available", varName1)
	}
	if a.Direction.EqualsTuple(NewVector(StringToFloat(x), StringToFloat(y), StringToFloat(z))) {
		return nil
	}
	return fmt.Errorf("Ray and direction fail")
}
func (tt *tupletest) cameractransformView_transformtuplefromTupletoTupleup(varName1, varName2, varName3, varName4 string) error {
	a, ok := tt.Cameras[varName1]
	if !ok {
		return fmt.Errorf("Camera %s not available", varName1)
	}

	b, ok := tt.Tuples[varName2]
	if !ok {
		return fmt.Errorf("Tuples %s not available", varName2)
	}

	c, ok := tt.Tuples[varName3]
	if !ok {
		return fmt.Errorf("Tuples %s not available", varName3)
	}

	d, ok := tt.Tuples[varName4]
	if !ok {
		return fmt.Errorf("Tuples %s not available", varName4)
	}

	a.Transform = ViewTransform(b, c, d)
	tt.Cameras[varName1] = a
	return nil
}
func (tt *tupletest) canvasimageRendercameracWorldw(varName1, varName2, varName3 string) error {
	b, ok := tt.Cameras[varName2]
	if !ok {
		return fmt.Errorf("Tuples %s not available", varName2)
	}

	c, ok := tt.Worlds[varName3]
	if !ok {
		return fmt.Errorf("Tuples %s not available", varName3)
	}

	tt.Canvases[varName1] = b.Render(c)
	return nil
}
func (tt *tupletest) pixel_atcanvasimageColor(varName1 string, x, y int64, r, g, b string) error {
	a, ok := tt.Canvases[varName1]
	if !ok {
		return fmt.Errorf("Canvases %s not available", varName1)
	}

	if a.PixelAt(int(x), int(y)).Equals(NewColor(StringToFloat(r), StringToFloat(g), StringToFloat(b))) {
		return nil
	}
	return fmt.Errorf("Pixel off %d,%d | %f,%f,%f", x, y, StringToFloat(r), StringToFloat(g), StringToFloat(b))

}

func (tt *tupletest) colorsresultLightingmaterialmLightlightTuplepositionTupleeyevTuplenormalvTrue(varName1, mName, lName, pName, eName, nName string) error {
	m, ok := tt.Materials[mName]
	if !ok {
		return fmt.Errorf("Canvases %s not available", mName)
	}

	l, ok := tt.Lights[lName]
	if !ok {
		return fmt.Errorf("Canvases %s not available", lName)
	}

	p, ok := tt.Tuples[pName]
	if !ok {
		return fmt.Errorf("Canvases %s not available", pName)
	}

	e, ok := tt.Tuples[eName]
	if !ok {
		return fmt.Errorf("Canvases %s not available", eName)
	}

	n, ok := tt.Tuples[nName]
	if !ok {
		return fmt.Errorf("Canvases %s not available", nName)
	}

	tt.Colors[varName1] = Lighting(m, NewSphere(), l, p, e, n, true)
	return nil
}

func (tt *tupletest) is_shadowedworldwTuplepIsFalse(varName1, varName2, truefalse string) error {
	w, ok := tt.Worlds[varName1]
	if !ok {
		return fmt.Errorf("World %s not available", varName1)
	}

	t, ok := tt.Tuples[varName2]
	if !ok {
		return fmt.Errorf("Tuple %s not available", varName2)
	}

	tf := true
	if truefalse == "false" {
		tf = false
	}

	if w.IsShadowed(t) && tf {
		return nil
	}
	if !w.IsShadowed(t) && !tf {
		return nil
	}
	return fmt.Errorf("False/true bad")

}

func (tt *tupletest) spheresIsAddedToWorldw(varName1, varName2 string) error {

	s, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("Shapes %s not available", varName1)
	}
	w, ok := tt.Worlds[varName2]
	if !ok {
		return fmt.Errorf("World %s not available", varName2)
	}

	w.Objects = append(tt.Worlds[varName2].Objects, s)
	tt.Worlds[varName2] = w
	return nil

}

func (tt *tupletest) computescompsover_pointzEPSILON(varName1 string) error {
	c, ok := tt.Computations[varName1]
	if !ok {
		return fmt.Errorf("Computation %s not available", varName1)
	}

	//a := tt.Intersections["i"]
	//b := tt.Rays["r"]
	//d := tt.Shapes["shape"]
	//log.Fatalf("\nInt: %v\nRay: %v\nShape: %v\nComp: %v\n\n%v", a, b, d, c, tt.Shapes)
	if c.OverPoint.Z < -epsilon/2 {
		return nil
	}
	return fmt.Errorf("Overpoint nerts %f|%f", c.OverPoint.Z, -epsilon/2)
}
func (tt *tupletest) computescompspointzComputescompsover_pointz(varName1, varName2 string) error {

	c1, ok := tt.Computations[varName1]
	if !ok {
		return fmt.Errorf("Computation %s not available", varName1)
	}

	c2, ok := tt.Computations[varName2]
	if !ok {
		return fmt.Errorf("Computation %s not available", varName2)
	}

	if c1.Point.Z > c2.OverPoint.Z {
		return nil
	}
	return fmt.Errorf("OverPoitn mismatch")
}

func (tt *tupletest) set_transformshapessTranslation(varName1, x, y, z string) error {
	_, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("Shape %s not available", varName1)
	}

	tt.Shapes[varName1].SetTransform(NewTranslation(StringToFloat(x), StringToFloat(y), StringToFloat(z)))
	return nil
}
func (tt *tupletest) shapessTest_shape(varName1 string) error {
	tt.Shapes[varName1] = NewTestShape()
	return nil
}
func (tt *tupletest) shapesstransformIdentity_matrix(varName1 string) error {
	t, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("Shape %s not available", varName1)
	}

	x := t.GetTransform()
	if x.EqualsMatrix(IdentityMatrix()) {
		return nil
	}
	return fmt.Errorf("Nope")
}
func (tt *tupletest) shapesstransformTranslation(varName1, x, y, z string) error {
	t, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("Computation %s not available", varName1)
	}

	x2 := t.GetTransform()
	if x2.EqualsMatrix(NewTranslation(StringToFloat(x), StringToFloat(y), StringToFloat(z))) {
		return nil
	}
	return fmt.Errorf("Bad trans \n%v\n%v", x2, NewTranslation(StringToFloat(x), StringToFloat(y), StringToFloat(z)))
}

func (tt *tupletest) materialmShapessmaterial(varName1, varName2 string) error {
	s, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("Shape %s not available", varName2)
	}
	tt.Materials[varName1] = s.GetMaterial()
	return nil
}
func (tt *tupletest) shapessmaterialMaterialmAssign(varName1, varName2 string) error {
	m, ok := tt.Materials[varName2]
	if !ok {
		return fmt.Errorf("Materials %s not available", varName2)
	}

	tt.Shapes[varName1].SetMaterial(m)
	return nil
}
func (tt *tupletest) shapessmaterialMaterialmEqual(varName1, varName2 string) error {
	s, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("Sha %s not available", varName1)
	}
	m, ok := tt.Materials[varName2]
	if !ok {
		return fmt.Errorf("Mat %s not available", varName2)
	}

	if s.GetMaterial().Equals(m) {
		return nil
	}
	return fmt.Errorf("Material mismatch")
}

func (tt *tupletest) arrayintersectionsxsIntersectshapessRayr(varName1, varName2, varName3 string) error {
	s, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("Sha %s not available", varName2)
	}
	r, ok := tt.Rays[varName3]
	if !ok {
		return fmt.Errorf("Ray %s not available", varName3)
	}

	tt.ArrayIntersections[varName1] = s.Intersects(r)
	return nil
}
func (tt *tupletest) set_transformshapessScaling(varName1, x, y, z string) error {
	s, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("Sha %s not available", varName1)
	}
	s.SetTransform(NewScaling(StringToFloat(x), StringToFloat(y), StringToFloat(z)))
	tt.Shapes[varName1] = s
	return nil
}
func (tt *tupletest) shapesssaved_raydirectionVector(varName1, x, y, z string) error {
	s, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("Sha %s not available", varName1)
	}
	if s.GetSavedRay().Direction.EqualsTuple(NewVector(StringToFloat(x), StringToFloat(y), StringToFloat(z))) {
		return nil
	}
	return fmt.Errorf("Saved ray fail2")
}
func (tt *tupletest) shapesssaved_rayoriginPoint(varName1, x, y, z string) error {
	s, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("Sha %s not available", varName1)
	}
	if s.GetSavedRay().Origin.EqualsTuple(NewPoint(StringToFloat(x), StringToFloat(y), StringToFloat(z))) {
		return nil
	}
	return fmt.Errorf("Saved ray fail %v", s.GetSavedRay())
}

func (tt *tupletest) set_transformshapessMatrixm(varName1, varName2 string) error {

	s, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("Sha %s not available", varName1)
	}

	m, ok := tt.Matrices[varName2]
	if !ok {
		return fmt.Errorf("MA %s not available", varName2)
	}
	s.SetTransform(m)
	tt.Shapes[varName1] = s
	return nil
}

func (tt *tupletest) tuplenNormal_atshapessPoint(varName1, varName2, x, y, z string) error {
	s, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("Sha %s not available", varName2)
	}

	tt.Tuples[varName1] = s.NormalAt(NewPoint(StringToFloat(x), StringToFloat(y), StringToFloat(z)))
	return nil
}

func (tt *tupletest) arrayintersectionsxsIsEmpty(varName1 string) error {

	ai, ok := tt.ArrayIntersections[varName1]
	if !ok {
		return fmt.Errorf("AI %s not available", varName1)
	}
	if len(ai) == 0 {
		return nil
	}
	return fmt.Errorf("Nope to empty AI")
}

func (tt *tupletest) arrayintersectionsxsLocal_intersectplanepRayr(varName1, varName2, varName3 string) error {
	pl, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("AI %s not available", varName2)
	}

	r, ok := tt.Rays[varName3]
	if !ok {
		return fmt.Errorf("AI %s not available", varName3)
	}

	tt.ArrayIntersections[varName1] = pl.LocalIntersects(r)
	return nil
}

//ctx.Step(`^arrayintersections\.([a-zA-Z0-9_]+)\[(\d+)\]\.object = plane\.([a-zA-Z0-9_]+)$`, tt.arrayintersectionsxsObjectPlanep)

func (tt *tupletest) arrayintersectionsxsObjectPlanep(varName1 string, indx int, varName2 string) error {
	ai, ok := tt.ArrayIntersections[varName1]
	if !ok {
		return fmt.Errorf("AI %s not available", varName1)
	}
	pl, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("PL%s not available", varName2)
	}
	if ai[indx].Object.Equals(pl) {
		return nil
	}
	bep := ""
	for _, p := range tt.Shapes["g"].GetShapes() {
		bep = fmt.Sprintf("%s\n%d", bep, p.GetID())
	}
	log.Fatal(ai[indx].Object.GetID(), pl.GetID(), bep)
	return fmt.Errorf("AI[%d].O failed\n%v\n%b", indx, ai[indx].Object.GetID(), pl.GetID())
}
func (tt *tupletest) arrayintersectionsxsT() error {
	return godog.ErrPending
}
func (tt *tupletest) colorblackColor() error {
	tt.Colors["black"] = NewColor(0, 0, 0)
	return nil
}
func (tt *tupletest) colorwhiteColor() error {
	tt.Colors["white"] = NewColor(1, 1, 1)
	return nil
}
func (tt *tupletest) patternpatternStripe_patterncolorwhiteColorblack(varName1 string) error {
	tt.Patterns[varName1] = NewStripePattern(NewColor(1, 1, 1), NewColor(0, 0, 0))
	return nil
}
func (tt *tupletest) patternpatternaColorwhite(varName1 string) error {
	pa, ok := tt.Patterns[varName1]
	if !ok {
		return fmt.Errorf("PA%s not available", varName1)
	}
	if pa.GetColorString("A").Equals(NewColor(1, 1, 1)) {
		return nil
	}
	return fmt.Errorf("Pattern A wasn't white, %f,%f,%f", pa.GetColorString("A").Red, pa.GetColorString("A").Green, pa.GetColorString("A").Blue)
}
func (tt *tupletest) patternpatternbColorblack(varName1 string) error {
	pa, ok := tt.Patterns[varName1]
	if !ok {
		return fmt.Errorf("PA%s not available", varName1)
	}
	if pa.GetColorString("B").Equals(NewColor(0, 0, 0)) {
		return nil
	}
	return fmt.Errorf("Pattern %s wasn't black, %f,%f,%f", varName1, pa.GetColorString("A").Red, pa.GetColorString("A").Green, pa.GetColorString("A").Blue)
}
func (tt *tupletest) planepPlane(varName1 string) error {
	tt.Shapes[varName1] = NewPlane()
	return nil
}
func (tt *tupletest) stripe_atpatternpatternPointColorblack(varName1, x, y, z string) error {
	pa, ok := tt.Patterns[varName1]
	if !ok {
		return fmt.Errorf("PA%s not available", varName1)
	}
	if pa.ColorAt(NewPoint(StringToFloat(x), StringToFloat(y), StringToFloat(z))).Equals(NewColor(0, 0, 0)) {
		return nil
	}
	return fmt.Errorf("Point %s,%s,%s was not black", x, y, z)
}
func (tt *tupletest) stripe_atpatternpatternPointColorwhite(varName1, x, y, z string) error {
	pa, ok := tt.Patterns[varName1]
	if !ok {
		return fmt.Errorf("PA%s not available", varName1)
	}
	if pa.ColorAt(NewPoint(StringToFloat(x), StringToFloat(y), StringToFloat(z))).Equals(NewColor(1, 1, 1)) {
		return nil
	}
	return fmt.Errorf("Point %s,%s,%s was not white", x, y, z)
}
func (tt *tupletest) tuplenLocal_normal_atplanepPoint(varName1, varName2, x, y, z string) error {
	sa, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("MA%s not available", varName2)
	}
	tt.Tuples[varName1] = sa.LocalNormalAt(NewPoint(StringToFloat(x), StringToFloat(y), StringToFloat(z)))
	return nil
}

func (tt *tupletest) colorscLightingmaterialmLightlightPointTupleeyevTuplenormalvFalse(varName1, varName2, varName3, x, y, z, varName4, varName5 string) error {
	ma, ok := tt.Materials[varName2]
	if !ok {
		return fmt.Errorf("MA%s not available", varName2)
	}
	li, ok := tt.Lights[varName3]
	if !ok {
		return fmt.Errorf("LI%s not available", varName3)
	}
	point := NewPoint(StringToFloat(x), StringToFloat(y), StringToFloat(z))
	ey, ok := tt.Tuples[varName4]
	if !ok {
		return fmt.Errorf("Ey%s not available", varName4)
	}
	no, ok := tt.Tuples[varName5]
	if !ok {
		return fmt.Errorf("no%s not available", varName5)
	}
	tt.Colors[varName1] = Lighting(ma, NewSphere(), li, point, ey, no, false)
	return nil
}
func (tt *tupletest) materialmdiffuse(varName1, v string) error {
	ma, ok := tt.Materials[varName1]
	if !ok {
		return fmt.Errorf("MA%s not available", varName1)
	}
	ma.Diffuse = StringToFloat(v)
	tt.Materials[varName1] = ma
	return nil
}
func (tt *tupletest) materialmpatternStripe_patterncolorColor(varName1, r1, g1, b1, r2, g2, b2 string) error {
	ma, ok := tt.Materials[varName1]
	if !ok {
		return fmt.Errorf("MA%s not available", varName1)
	}
	ma.SetPattern(
		NewStripePattern(
			NewColor(StringToFloat(r1), StringToFloat(g1), StringToFloat(b1)),
			NewColor(StringToFloat(r2), StringToFloat(g2), StringToFloat(b2)),
		),
	)
	tt.Materials[varName1] = ma
	return nil
}
func (tt *tupletest) materialmspecular(varName1, v string) error {
	ma, ok := tt.Materials[varName1]
	if !ok {
		return fmt.Errorf("MA%s not available", varName1)
	}
	ma.Specular = StringToFloat(v)
	tt.Materials[varName1] = ma
	return nil
}

func (tt *tupletest) colorscColorswhite(varName1, varName2 string) error {
	c1, ok := tt.Colors[varName1]
	if !ok {
		return fmt.Errorf("C%s not avail", varName1)
	}
	c2, ok := tt.Colors[varName2]
	if !ok {
		return fmt.Errorf("C%s not avail", varName2)
	}
	if c1.Equals(c2) {
		return nil
	}
	return fmt.Errorf("Colors mismatch %v,%v", c1, c2)
}
func (tt *tupletest) colorscStripe_at_objectpatternpatternSphereobjectPoint(varName1, varName2, varName3, x, y, z string) error {

	p1, ok := tt.Patterns[varName2]
	if !ok {
		return fmt.Errorf("pa%s not avail", varName2)
	}

	o1, ok := tt.Shapes[varName3]
	if !ok {
		return fmt.Errorf("pa%s not avail", varName3)
	}

	tt.Colors[varName1] = p1.ColorAtObject(o1, NewPoint(StringToFloat(x), StringToFloat(y), StringToFloat(z)))
	return nil
}
func (tt *tupletest) set_pattern_transformpatternpatternSpherescaling(varName1, x, y, z string) error {
	_, ok := tt.Patterns[varName1]
	if !ok {
		return fmt.Errorf("pa%s not avail", varName1)
	}
	tt.Patterns[varName1].SetTransform(NewScaling(StringToFloat(x), StringToFloat(y), StringToFloat(z)))
	return nil
}
func (tt *tupletest) set_pattern_transformpatternpatternTranslation(varName1, x, y, z string) error {
	_, ok := tt.Patterns[varName1]
	if !ok {
		return fmt.Errorf("pa%s not avail", varName1)
	}
	tt.Patterns[varName1].SetTransform(NewTranslation(StringToFloat(x), StringToFloat(y), StringToFloat(z)))
	return nil
}

func (tt *tupletest) shapesobjectSphere(varName1 string) error {
	tt.Shapes[varName1] = NewSphere()
	return nil
}

func (tt *tupletest) patternpatternTest_pattern(varName1 string) error {
	tt.Patterns[varName1] = NewTestPattern()
	return nil
}
func (tt *tupletest) patternpatterntransformIdentity_matrix(varName1 string) error {
	p1, ok := tt.Patterns[varName1]
	if !ok {
		return fmt.Errorf("pa%s not avail", varName1)
	}
	d := p1.GetTransform()
	if d.EqualsMatrix(IdentityMatrix()) {
		return nil
	}
	return fmt.Errorf("Not identity")
}
func (tt *tupletest) patternpatterntransformTranslation(varName1, x, y, z string) error {
	p1, ok := tt.Patterns[varName1]
	if !ok {
		return fmt.Errorf("pa%s not avail", varName1)
	}
	d := p1.GetTransform()
	if d.EqualsMatrix(NewTranslation(StringToFloat(x), StringToFloat(y), StringToFloat(z))) {
		return nil
	}
	return fmt.Errorf("Not tramsfpr,")
}

func (tt *tupletest) colorscPattern_at_shapepatternpatternShapesshapePoint(varName1, varName2, varName3, x, y, z string) error {
	pa, ok := tt.Patterns[varName2]
	if !ok {
		return fmt.Errorf("pa%s not avail", varName2)
	}

	sh, ok := tt.Shapes[varName3]
	if !ok {
		return fmt.Errorf("pa%s not avail", varName3)
	}
	tt.Colors[varName1] = pa.ColorAtObject(sh, NewPoint(StringToFloat(x), StringToFloat(y), StringToFloat(z)))
	return nil
}

func (tt *tupletest) pattern_atpatternpatternPointColor(varName1, x, y, z, r, g, b string) error {
	pa, ok := tt.Patterns[varName1]
	if !ok {
		return fmt.Errorf("pa%s not avail", varName1)
	}
	if pa.ColorAt(
		NewPoint(StringToFloat(x), StringToFloat(y), StringToFloat(z)),
	).
		Equals(
			NewColor(StringToFloat(r), StringToFloat(g), StringToFloat(b))) {
		return nil
	}
	return fmt.Errorf(
		"Color match fail\n%v\n%v",
		pa.ColorAt(NewPoint(StringToFloat(x), StringToFloat(y), StringToFloat(z))),
		NewColor(StringToFloat(r), StringToFloat(g), StringToFloat(b)),
	)
}
func (tt *tupletest) pattern_atpatternpatternPointWhite(varName1, x, y, z string) error {
	pa, ok := tt.Patterns[varName1]
	if !ok {
		return fmt.Errorf("pa%s not avail", varName1)
	}
	if pa.ColorAt(NewPoint(StringToFloat(x), StringToFloat(y), StringToFloat(z))).Equals(NewColor(1, 1, 1)) {
		return nil
	}
	return fmt.Errorf("White match fail %v", pa.ColorAt(NewPoint(StringToFloat(x), StringToFloat(y), StringToFloat(z))))
}
func (tt *tupletest) patternpatternGradient_patternwhiteBlack(varName1 string) error {
	tt.Patterns[varName1] = NewGradientPattern(NewColor(1, 1, 1), NewColor(0, 0, 0))
	return nil
}

func (tt *tupletest) pattern_atpatternpatternPointBlack(varName1, x, y, z string) error {
	pa, ok := tt.Patterns[varName1]
	if !ok {
		return fmt.Errorf("pa%s not avail", varName1)
	}
	if pa.ColorAt(NewPoint(StringToFloat(x), StringToFloat(y), StringToFloat(z))).Equals(NewColor(0, 0, 0)) {
		return nil
	}
	return fmt.Errorf("Black match fail %v", pa.ColorAt(NewPoint(StringToFloat(x), StringToFloat(y), StringToFloat(z))))
}
func (tt *tupletest) patternpatternRing_patternwhiteBlack(varName1 string) error {
	tt.Patterns[varName1] = NewRingPattern(NewColor(1, 1, 1), NewColor(0, 0, 0))
	return nil
}

func (tt *tupletest) patternpatternCheckers_patternwhiteBlack(varName1 string) error {
	tt.Patterns[varName1] = NewCheckerPattern(NewColor(1, 1, 1), NewColor(0, 0, 0))
	return nil
}

func (tt *tupletest) materialmreflective(varName1, trgt string) error {
	ma, ok := tt.Materials[varName1]
	if !ok {
		return fmt.Errorf("ma%s not avail", varName1)
	}
	if ma.Reflective == StringToFloat(trgt) {
		return nil
	}
	return fmt.Errorf("Reflect went nerts")

}

func (tt *tupletest) computescompsreflectvVector(varName1, x, y, z string) error {
	co, ok := tt.Computations[varName1]
	if !ok {
		return fmt.Errorf("co%s not avail", varName1)
	}
	if co.Reflectv.EqualsTuple(
		NewVector(StringToFloat(x), StringToFloat(y), StringToFloat(z)),
	) {
		return nil
	}
	return fmt.Errorf("Reflect didn't %v\n%v", co.Reflectv, NewVector(StringToFloat(x), StringToFloat(y), StringToFloat(z)))
}
func (tt *tupletest) intersectioniIntersectionShapesshape(varName1, t, varName2 string) error {
	sh, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("sh%s not avail", varName2)
	}
	tt.Intersections[varName1] = NewIntersection(StringToFloat(t), sh)
	return nil
}
func (tt *tupletest) shapesshapePlane(varName1 string) error {
	tt.Shapes[varName1] = NewPlane()
	return nil
}

func (tt *tupletest) colorscolorReflected_colorworldwComputescomps(varName1, varName2, varName3 string) error {
	wr, ok := tt.Worlds[varName2]
	if !ok {
		return fmt.Errorf("we%s not avail", varName2)
	}

	co, ok := tt.Computations[varName3]
	if !ok {
		return fmt.Errorf("sh%s not avail", varName3)
	}

	tt.Colors[varName1] = wr.ReflectedColor(co, maxReflects)
	return nil
}
func (tt *tupletest) shapesshapeTheSecondObjectInWorldw(varName1, varName2 string) error {
	wr, ok := tt.Worlds[varName2]
	if !ok {
		return fmt.Errorf("we%s not avail", varName2)
	}

	tt.Shapes[varName1] = wr.Objects[1]
	return nil
}
func (tt *tupletest) shapesshapematerialambient(varName1, amb string) error {
	sh, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("sh%s not avail", varName1)
	}
	m := sh.GetMaterial()
	m.Ambient = StringToFloat(amb)
	tt.Shapes[varName1].SetMaterial(m)
	return nil
}

func (tt *tupletest) shapesshapeIsAddedToWorldw(varName1, varName2 string) error {
	sh, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("sh%s not availw", varName1)
	}
	wr, ok := tt.Worlds[varName2]
	if !ok {
		return fmt.Errorf("we%s not avail", varName2)
	}
	wr.Objects = append(wr.Objects, sh)
	tt.Worlds[varName2] = wr
	return nil
}

func (tt *tupletest) shapesshapePlaneWith(varName1 string, arg1 *godog.Table) error {
	pl := NewPlane()
	for _, x := range arg1.Rows {
		switch x.Cells[0].Value {
		case "material.ambient":
			pl.Material.Ambient = StringToFloat(x.Cells[1].Value)
		case "material.pattern":
			if x.Cells[1].Value == "test_pattern()" {
				pl.Material.SetPattern(NewTestPattern())
			}
		case "material.reflective":
			pl.Material.Reflective = StringToFloat(x.Cells[1].Value)
		case "material.refractive_index":
			pl.Material.RefractiveIndex = StringToFloat(x.Cells[1].Value)
		case "material.transparency":
			pl.Material.Transparency = StringToFloat(x.Cells[1].Value)
		case "transform":
			funko := regexp.MustCompile(`^(.*)\((.*), (.*), (.*)\)$`)
			matches := funko.FindStringSubmatch(x.Cells[1].Value)
			switch matches[1] {
			case "translation":
				dude := pl.GetTransform()
				dude2 := dude.MultiplyMatrix(
					NewTranslation(
						StringToFloat(matches[2]),
						StringToFloat(matches[3]),
						StringToFloat(matches[4]),
					),
				)
				pl.SetTransform(dude2) // p159
			case "scaling":
				dude := pl.GetTransform()
				pl.SetTransform(
					dude.MultiplyMatrix(
						NewScaling(
							StringToFloat(matches[2]),
							StringToFloat(matches[3]),
							StringToFloat(matches[4]),
						),
					),
				)
			default:
				log.Fatal("Unknown transform")
			}
		default:
			log.Fatal("Unknown Aspect")
		}
	}
	tt.Shapes[varName1] = pl
	return nil
}
func (tt *tupletest) color_atworldwRayrShouldTerminateSuccessfully(varName1, varName2 string) error {
	wr, ok := tt.Worlds[varName1]
	if !ok {
		return fmt.Errorf("we%s not avail", varName1)
	}

	ry, ok := tt.Rays[varName2]
	if !ok {
		return fmt.Errorf("sh%s not avail", varName2)
	}

	wr.ColorAt(ry, 1)
	return nil
}

func (tt *tupletest) colorscolorReflected_colorworldwComputescomps2(varName1, varName2, varName3 string) error {
	wr, ok := tt.Worlds[varName2]
	if !ok {
		return fmt.Errorf("we%s not avail", varName2)
	}

	co, ok := tt.Computations[varName3]
	if !ok {
		return fmt.Errorf("sh%s not avail", varName3)
	}

	tt.Colors[varName1] = wr.ReflectedColor(co, 0)
	return nil
}

func (tt *tupletest) materialmrefractive_index(varName1, value string) error {
	ma, ok := tt.Materials[varName1]
	if !ok {
		return fmt.Errorf("ma%s not avail", varName1)
	}
	if epsilonEquals(ma.RefractiveIndex, StringToFloat(value)) {
		return nil
	}
	return fmt.Errorf("Bad refact")
}
func (tt *tupletest) materialmtransparency(varName1, value string) error {
	ma, ok := tt.Materials[varName1]
	if !ok {
		return fmt.Errorf("ma%s not avail", varName1)
	}
	if epsilonEquals(ma.Transparency, StringToFloat(value)) {
		return nil
	}
	return fmt.Errorf("Bad transp")
}

func (tt *tupletest) spheresGlass_sphere(varName1 string) error {
	tt.Shapes[varName1] = NewGlassSphere()
	return nil
}
func (tt *tupletest) spheresmaterialrefractive_index(varName1, value string) error {
	sh, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("P")
	}
	if epsilonEquals(sh.GetMaterial().RefractiveIndex, StringToFloat(value)) {
		return nil
	}
	return fmt.Errorf("Bad refr glass")
}
func (tt *tupletest) spheresmaterialtransparency(varName1, value string) error {
	sh, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("P")
	}
	if epsilonEquals(sh.GetMaterial().Transparency, StringToFloat(value)) {
		return nil
	}
	return fmt.Errorf("Bad transp glass")
}

func (tt *tupletest) arrayintersectionsxsIntersectionsABCBCA(varName1, val1, val2, val3, val4, val5, val6 string) error {
	A := tt.Shapes["A"]
	B := tt.Shapes["B"]
	C := tt.Shapes["C"]
	tt.ArrayIntersections[varName1] = Intersections(
		NewIntersection(StringToFloat(val1), A),
		NewIntersection(StringToFloat(val2), B),
		NewIntersection(StringToFloat(val3), C),
		NewIntersection(StringToFloat(val4), B),
		NewIntersection(StringToFloat(val5), C),
		NewIntersection(StringToFloat(val6), A),
	)
	return nil
}

var ThisIndex = -1

func (tt *tupletest) computescompsPrepare_computationsarrayintersectionsxsRayrArrayintersectionsxs(varName1, varName2 string, index1 int, varName3, varName4 string) error {
	xs, ok := tt.ArrayIntersections[varName2]
	if !ok {
		return fmt.Errorf("AI")
	}
	r, ok := tt.Rays[varName3]
	if !ok {
		return fmt.Errorf("RY")
	}
	x2, ok := tt.ArrayIntersections[varName4]
	if !ok {
		return fmt.Errorf("VN")
	}

	i := xs[index1]
	ThisIndex = index1
	tt.Computations[varName1] = i.PrepareComputations(r, x2)
	return nil
}
func (tt *tupletest) computescompsn1(varName2, value string) error {
	co, ok := tt.Computations[varName2]
	if !ok {
		return fmt.Errorf("co")
	}
	if epsilonEquals(co.N1, StringToFloat(value)) {
		return nil
	}
	return fmt.Errorf("[%d]N1 ffail %f should be %f", ThisIndex, co.N1, StringToFloat(value))
}
func (tt *tupletest) computescompsn2(varName2, value string) error {
	co, ok := tt.Computations[varName2]
	if !ok {
		return fmt.Errorf("co")
	}
	if epsilonEquals(co.N2, StringToFloat(value)) {
		return nil
	}
	return fmt.Errorf("[%d]N2 ffail %f should be %f", ThisIndex, co.N2, StringToFloat(value))
}
func (tt *tupletest) shapesGlass_sphereWith(varName1 string, arg1 *godog.Table) error {
	pl := NewSphere()
	for _, x := range arg1.Rows {
		switch x.Cells[0].Value {
		case "material.refractive_index":
			pl.Material.RefractiveIndex = StringToFloat(x.Cells[1].Value)
		case "material.reflective":
			pl.Material.Reflective = StringToFloat(x.Cells[1].Value)
		case "transform":
			funko := regexp.MustCompile(`^(.*)\((.*), (.*), (.*)\)$`)
			matches := funko.FindStringSubmatch(x.Cells[1].Value)
			switch matches[1] {
			case "translation":
				dude := pl.GetTransform()
				pl.SetTransform(
					dude.MultiplyMatrix(
						NewTranslation(
							StringToFloat(matches[2]),
							StringToFloat(matches[3]),
							StringToFloat(matches[4]),
						),
					),
				)
			case "scaling":
				dude := pl.GetTransform()
				pl.SetTransform(
					dude.MultiplyMatrix(
						NewScaling(
							StringToFloat(matches[2]),
							StringToFloat(matches[3]),
							StringToFloat(matches[4]),
						),
					),
				)
			}
		}
	}
	tt.Shapes[varName1] = pl
	return nil
}

func (tt *tupletest) arrayintersectionsxsIntersectionsintersectioni(varName1, varName2 string) error {
	i, ok := tt.Intersections[varName2]
	if !ok {
		return fmt.Errorf("X")
	}
	tt.ArrayIntersections[varName1] = Intersections(i)
	return nil
}
func (tt *tupletest) computescompsPrepare_computationsintersectioniRayrArrayintersectionsxs(varName1, varName2, varName3, varName4 string) error {
	i, ok := tt.Intersections[varName2]
	if !ok {
		return fmt.Errorf("X1")
	}

	r, ok := tt.Rays[varName3]
	if !ok {
		return fmt.Errorf("X2")
	}

	xs, ok := tt.ArrayIntersections[varName4]
	if !ok {
		return fmt.Errorf("X3")
	}

	tt.Computations[varName1] = i.PrepareComputations(r, xs)
	return nil
}
func (tt *tupletest) computescompspointzComputescompsunder_pointz(varName1, varName2 string) error {
	c, ok := tt.Computations[varName1]
	if !ok {
		return fmt.Errorf("Where are you")
	}
	if c.Point.Z < c.UnderPoint.Z {
		return nil
	}
	return fmt.Errorf("N")
}
func (tt *tupletest) computescompsunder_pointzEPSILON(varName1 string) error {
	c, ok := tt.Computations[varName1]
	if !ok {
		return fmt.Errorf("X")
	}
	if c.UnderPoint.Z > epsilon/2 {
		return nil
	}
	return fmt.Errorf("M")
}

func (tt *tupletest) arrayintersectionsxsIntersectionsShapesshapeShapesshape(varName1, inter1, varName2, inter2, varName3 string) error {
	sh1, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("X")
	}
	sh2, ok := tt.Shapes[varName3]
	if !ok {
		return fmt.Errorf("X")
	}
	tt.ArrayIntersections[varName1] = Intersections(
		NewIntersection(StringToFloat(inter1), sh1),
		NewIntersection(StringToFloat(inter2), sh2),
	)
	return nil
}
func (tt *tupletest) colorscRefracted_colorworldwComputescomps(varName1, varName2, varName3 string, max int) error {
	w, ok := tt.Worlds[varName2]
	if !ok {
		return fmt.Errorf("X")
	}
	co, ok := tt.Computations[varName3]
	if !ok {
		return fmt.Errorf("Y")
	}
	tt.Colors[varName1] = w.RefractedColor(co, max)
	return nil
}
func (tt *tupletest) shapesshapeTheFirstObjectInWorldw(varName1, varName2 string) error {
	w, ok := tt.Worlds[varName2]
	if !ok {
		return fmt.Errorf("X")
	}
	tt.Shapes[varName1] = w.Objects[0]
	return nil
}

func (tt *tupletest) shapesshapeHas(varName1 string, arg1 *godog.Table) error {
	sh1, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("X")
	}
	m := sh1.GetMaterial()

	for _, x := range arg1.Rows {
		switch x.Cells[0].Value {
		case "material.ambient":
			m.Ambient = StringToFloat(x.Cells[1].Value)
		case "material.pattern":
			if x.Cells[1].Value == "test_pattern()" {
				m.SetPattern(NewTestPattern())
			}
		case "material.reflective":
			m.Reflective = StringToFloat(x.Cells[1].Value)
		case "material.refractive_index":
			m.RefractiveIndex = StringToFloat(x.Cells[1].Value)
		case "material.transparency":
			m.Transparency = StringToFloat(x.Cells[1].Value)
		case "transform":
			funko := regexp.MustCompile(`^(.*)\((.*), (.*), (.*)\)$`)
			matches := funko.FindStringSubmatch(x.Cells[1].Value)
			switch matches[1] {
			case "translation":
				dude := sh1.GetTransform()
				sh1.SetTransform(
					dude.MultiplyMatrix(
						NewTranslation(
							StringToFloat(matches[2]),
							StringToFloat(matches[3]),
							StringToFloat(matches[4]),
						),
					),
				)
			case "scaling":
				dude := sh1.GetTransform()
				sh1.SetTransform(
					dude.MultiplyMatrix(
						NewScaling(
							StringToFloat(matches[2]),
							StringToFloat(matches[3]),
							StringToFloat(matches[4]),
						),
					),
				)
			}
		}
	}

	sh1.SetMaterial(m)
	tt.Shapes[varName1] = sh1
	return nil
}

func (tt *tupletest) arrayintersectionsxsIntersectionsShapesshapeShapesshape4(
	varName1,
	inter1, varName2,
	inter2, varName3,
	inter3, varName4,
	inter4, varName5 string) error {
	sh1, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("X")
	}
	sh2, ok := tt.Shapes[varName3]
	if !ok {
		return fmt.Errorf("X")
	}
	sh3, ok := tt.Shapes[varName4]
	if !ok {
		return fmt.Errorf("X")
	}
	sh4, ok := tt.Shapes[varName5]
	if !ok {
		return fmt.Errorf("X")
	}
	tt.ArrayIntersections[varName1] = Intersections(
		NewIntersection(StringToFloat(inter1), sh1),
		NewIntersection(StringToFloat(inter2), sh2),
		NewIntersection(StringToFloat(inter3), sh3),
		NewIntersection(StringToFloat(inter4), sh4),
	)
	return nil
}

func (tt *tupletest) arrayintersectionsxsIntersectionsShapesshapeShapesshape1(
	varName1,
	inter1, varName2 string) error {
	sh1, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("X")
	}
	tt.ArrayIntersections[varName1] = Intersections(
		NewIntersection(StringToFloat(inter1), sh1),
	)
	return nil
}

func (tt *tupletest) colorscolorShade_hitworldwComputescomps(varName1, varName2, varName3 string, max int) error {
	wld, ok := tt.Worlds[varName2]
	if !ok {
		return fmt.Errorf("X")
	}
	com, ok := tt.Computations[varName3]
	if !ok {
		return fmt.Errorf("X")
	}
	tt.Colors[varName1] = wld.ShadeHit(com, max)
	return nil
}
func (tt *tupletest) shapesballSphereWith(varName1 string, arg1 *godog.Table) error {
	sh1 := NewSphere()

	for _, x := range arg1.Rows {
		switch x.Cells[0].Value {
		case "material.ambient":
			sh1.Material.Ambient = StringToFloat(x.Cells[1].Value)
		case "material.color":
			funko := regexp.MustCompile(`^\((.*), (.*), (.*)\)$`)
			matches := funko.FindStringSubmatch(x.Cells[1].Value)
			sh1.Material.Color = NewColor(
				StringToFloat(matches[1]),
				StringToFloat(matches[2]),
				StringToFloat(matches[3]))
		case "material.pattern":
			if x.Cells[1].Value == "test_pattern()" {
				sh1.Material.SetPattern(NewTestPattern())
			}
		case "material.reflective":
			sh1.Material.Reflective = StringToFloat(x.Cells[1].Value)
		case "material.refractive_index":
			sh1.Material.RefractiveIndex = StringToFloat(x.Cells[1].Value)
		case "material.transparency":
			sh1.Material.Transparency = StringToFloat(x.Cells[1].Value)
		case "transform":
			funko := regexp.MustCompile(`^(.*)\((.*), (.*), (.*)\)$`)
			matches := funko.FindStringSubmatch(x.Cells[1].Value)
			switch matches[1] {
			case "translation":
				dude := sh1.GetTransform()
				sh1.SetTransform(
					dude.MultiplyMatrix(
						NewTranslation(
							StringToFloat(matches[2]),
							StringToFloat(matches[3]),
							StringToFloat(matches[4]),
						),
					),
				)
			case "scaling":
				dude := sh1.GetTransform()
				sh1.SetTransform(
					dude.MultiplyMatrix(
						NewScaling(
							StringToFloat(matches[2]),
							StringToFloat(matches[3]),
							StringToFloat(matches[4]),
						),
					),
				)
			}
		}
	}

	tt.Shapes[varName1] = sh1
	//STOPHERE = true
	return nil
}

func (tt *tupletest) floatsreflectance(varName1, val string) error {
	fl, ok := tt.Floats[varName1]
	if !ok {
		return fmt.Errorf("X")
	}
	if epsilonEquals(fl, StringToFloat(val)) {
		return nil
	}
	return fmt.Errorf("Casdf")
}

func (tt *tupletest) floatsreflectanceSchlickcomputescomps(varName1, varName2 string) error {
	cm, ok := tt.Computations[varName2]
	if !ok {
		return fmt.Errorf("X")
	}

	tt.Floats[varName1] = cm.Schlick()
	return nil
}

func (tt *tupletest) shapesshapeGlass_sphere(varName1 string) error {
	tt.Shapes[varName1] = NewGlassSphere()
	return nil
}

func (tt *tupletest) shapescCube(varName1 string) error {
	tt.Shapes[varName1] = NewCube()
	return nil
}

func (tt *tupletest) tuplenormalLocal_normal_atshapescTuplep(varName1, varName2, varName3 string) error {
	c, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("X")
	}
	p, ok := tt.Tuples[varName3]
	if !ok {
		return fmt.Errorf("X")
	}
	tt.Tuples[varName1] = c.LocalNormalAt(p)
	return nil
}

func (tt *tupletest) rayrRaypointTupledirection(varName1, x, y, z, varName3 string) error {
	q, ok := tt.Tuples[varName3]
	if !ok {
		return fmt.Errorf("X")
	}
	tt.Rays[varName1] = NewRay(NewPoint(StringToFloat(x), StringToFloat(y), StringToFloat(z)), q)

	return nil
}
func (tt *tupletest) shapescylCylinder(varName1 string) error {
	tt.Shapes[varName1] = NewCylinder()
	return nil
}
func (tt *tupletest) tupledirectionNormalizevector(varName1, x, y, z string) error {
	tt.Tuples[varName1] = NewVector(StringToFloat(x), StringToFloat(y), StringToFloat(z)).Normalize()
	return nil
}

func (tt *tupletest) shapescylmaximumInfinity(varName1 string) error {
	c, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("yyy")
	}
	if math.IsInf(c.GetMaximum(), 1) {
		return nil
	}
	return fmt.Errorf("Not inf")
}
func (tt *tupletest) shapescylminimumInfinity(varName1 string) error {
	c, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("yyy")
	}
	if math.IsInf(c.GetMinimum(), -1) {
		return nil
	}
	return fmt.Errorf("Not -inf")
}

func (tt *tupletest) shapescylmaximum(varName1, value string) error {
	_, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("zzz")
	}

	tt.Shapes[varName1].SetMaximum(StringToFloat(value))
	return nil
}
func (tt *tupletest) shapescylminimum(varName1, value string) error {
	_, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("zzz")
	}

	tt.Shapes[varName1].SetMinimum(StringToFloat(value))
	return nil
}

func (tt *tupletest) shapescylclosedFalse(varName1, value string) error {
	c, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("zzz")
	}
	switch value {
	case "false":
		if !c.GetClosed() {
			return nil
		}
	case "true":
		if c.GetClosed() {
			return nil
		}
	}
	return fmt.Errorf("DASDF")

}

func (tt *tupletest) shapescylclosedTrue(varName1, value string) error {
	STOPHERE = true
	_, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("zzz")
	}
	switch value {
	case "false":
		tt.Shapes[varName1].SetClosed(false)
	case "true":
		tt.Shapes[varName1].SetClosed(true)
	default:
		return fmt.Errorf("ASDF")
	}
	return nil

}

func (tt *tupletest) shapesshapeCone(varName1 string) error {
	tt.Shapes[varName1] = NewCone()
	return nil
}

func (tt *tupletest) shapesgGroup(varName1 string) error {
	x := NewGroup()
	tt.Shapes[varName1] = x
	return nil
}
func (tt *tupletest) shapesgIsEmpty(varName1 string) error {
	x, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("zzz")
	}
	if x.GetShapesCount() == 0 {
		return nil
	}
	return fmt.Errorf("How'd shapes get in there?")
}

func (tt *tupletest) shapessparentIsNothing(varName1 string) error {
	x, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("zzz")
	}
	if x.GetParent() == nil {
		return nil
	}
	return fmt.Errorf("damn")

}

var AddedShape = ""

func (tt *tupletest) add_childshapesgShapess(varName1, varName2 string) error {
	_, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("zzz")
	}
	s, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("zzz")
	}
	AddedShape += fmt.Sprintf("%s]%d\n", varName2, s.GetID())
	tt.Shapes[varName1].AddShape(&s)
	tt.Shapes[varName2] = s
	return nil
}
func (tt *tupletest) shapesgIncludesShapess(varName1, varName2 string) error {
	g, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("zzz")
	}
	s, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("zzz")
	}
	mike, ok := g.GetShapes()[s.GetID()]
	if ok {
		return nil
	}
	return fmt.Errorf("No child of that name %v", mike)
}
func (tt *tupletest) shapesgIsNotEmpty(varName1 string) error {
	g, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("zzz")
	}
	if g.GetShapesCount() != 0 {
		return nil
	}
	return fmt.Errorf("Found children")
}
func (tt *tupletest) shapessparentShapesg(varName1, varName2 string) error {
	g, ok := tt.Shapes[varName2]
	if !ok {
		return fmt.Errorf("zzz")
	}
	s, ok := tt.Shapes[varName1]
	if !ok {
		return fmt.Errorf("zzz")
	}
	p := s.GetParent()
	if p.Equals(g) {
		return nil
	}
	return fmt.Errorf("Wrong parent")
}
