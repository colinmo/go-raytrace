Feature: Intersections

    Feature Description

    Scenario: An intersection encapsulates t and object
        Given sphere.s ← sphere()
        When intersection.i ← intersection(3.5, sphere.s)
        Then intersection.i.t = 3.5
        And intersection.i.object = sphere.s

    Scenario: Aggregating intersections
        Given sphere.s ← sphere()
        And intersection.i1 ← intersection(1, sphere.s)
        And intersection.i2 ← intersection(2, sphere.s)
        When arrayintersections.xs ← intersections(intersection.i1, intersection.i2)
        Then arrayintersections.xs.count = 2
        And arrayintersections.xs[0].t = 1
        And arrayintersections.xs[1].t = 2
    Scenario: The hit, when all intersections have positive t
        Given sphere.s ← sphere()
        And intersection.i1 ← intersection(1, sphere.s)
        And intersection.i2 ← intersection(2, sphere.s)
        And arrayintersections.xs ← intersections(intersection.i2, intersection.i1)
        When intersection.i ← hit(arrayintersections.xs)
        Then intersection.i = intersection.i1
    Scenario: The hit, when some intersections have negative t
        Given sphere.s ← sphere()
        And intersection.i1 ← intersection(-1, sphere.s)
        And intersection.i2 ← intersection(1, sphere.s)
        And arrayintersections.xs ← intersections(intersection.i2, intersection.i1)
        When intersection.i ← hit(arrayintersections.xs)
        Then intersection.i = intersection.i2
    Scenario: The hit, when all intersections have negative t
        Given sphere.s ← sphere()
        And intersection.i1 ← intersection(-2, sphere.s)
        And intersection.i2 ← intersection(-1, sphere.s)
        And arrayintersections.xs ← intersections(intersection.i2, intersection.i1)
        When intersection.i ← hit(arrayintersections.xs)
        Then intersection.i is nothing
    Scenario: The hit is always the lowest nonnegative intersection
        Given sphere.s ← sphere()
        And intersection.i1 ← intersection(5, sphere.s)
        And intersection.i2 ← intersection(7, sphere.s)
        And intersection.i3 ← intersection(-3, sphere.s)
        And intersection.i4 ← intersection(2, sphere.s)
        And arrayintersections.xs ← intersections(intersection.i1, intersection.i2, intersection.i3, intersection.i4)
        When intersection.i ← hit(arrayintersections.xs)
        Then intersection.i = intersection.i4
    Scenario: Precomputing the state of an intersection
        Given ray.r ← ray(point(0, 0, -5), vector(0, 0, 1))
        And sphere.shape ← sphere()
        And intersection.i ← intersection(4, sphere.shape)
        When computes.comps ← prepare_computations(intersection.i, ray.r)
        Then computes.comps.t = intersection.i.t
        And computes.comps.object = intersection.i.object
        And computes.comps.point = point(0, 0, -1)
        And computes.comps.eyev = vector(0, 0, -1)
        And computes.comps.normalv = vector(0, 0, -1)
    Scenario: The hit, when an intersection occurs on the outside
        Given ray.r ← ray(point(0, 0, -5), vector(0, 0, 1))
        And sphere.shape ← sphere()
        And intersection.i ← intersection(4, sphere.shape)
        When computes.comps ← prepare_computations(intersection.i, ray.r)
        Then computes.comps.inside = false
    Scenario: The hit, when an intersection occurs on the inside
        Given ray.r ← ray(point(0, 0, 0), vector(0, 0, 1))
        And sphere.shape ← sphere()
        And intersection.i ← intersection(1, sphere.shape)
        When computes.comps ← prepare_computations(intersection.i, ray.r)
        Then computes.comps.point = point(0, 0, 1)
        And computes.comps.eyev = vector(0, 0, -1)
        And computes.comps.inside = true
        # normal would have been (0, 0, 1), but is inverted!
        And computes.comps.normalv = vector(0, 0, -1)
    Scenario: Shading an intersection
        Given world.w ← default_world()
        And ray.r ← ray(point(0, 0, -5), vector(0, 0, 1))
        And sphere.shape ← the first object in world.w
        And intersection.i ← intersection(4, sphere.shape)
        When computes.comps ← prepare_computations(intersection.i, ray.r)
        And colors.c ← shade_hit(world.w, computes.comps)
        Then colors.c = color(0.38066, 0.47583, 0.2855)
    Scenario: Shading an intersection from the inside
        Given world.w ← default_world()
        And world.w.light ← point_light(point(0, 0.25, 0), color(1, 1, 1))
        And ray.r ← ray(point(0, 0, 0), vector(0, 0, 1))
        And sphere.shape ← the second object in world.w
        And intersection.i ← intersection(0.5, sphere.shape)
        When computes.comps ← prepare_computations(intersection.i, ray.r)
        And colors.c ← shade_hit(world.w, computes.comps)
        Then colors.c = color(0.90498, 0.90498, 0.90498)
    Scenario: The hit should offset the point
        Given ray.r ← ray(point(0, 0, -5), vector(0, 0, 1))
        And sphere.shape ← sphere() with:
            | transform | translation(0, 0, 1) |
        And intersection.i ← intersection(5, sphere.shape)
        When computes.comps ← prepare_computations(intersection.i, ray.r)
        Then computes.comps.over_point.z < -EPSILON/2
        And computes.comps.point.z > computes.comps.over_point.z
    Scenario: Precomputing the reflection vector
        Given shapes.shape ← plane()
        And ray.r ← ray(point(0, 1, -1), vector(0, -√2/2, √2/2))
        And intersection.i ← intersection(√2, shapes.shape)
        When computes.comps ← prepare_computations(intersection.i, ray.r)
        Then computes.comps.reflectv = vector(0, √2/2, √2/2)
    Scenario Outline: Finding n1 and n2 at various intersections
        Given shapes.A ← glass_sphere() with:
            | transform                 | scaling(2, 2, 2) |
            | material.refractive_index | 1.5              |
        And shapes.B ← glass_sphere() with:
            | transform                 | translation(0, 0, -0.25) |
            | material.refractive_index | 2.0                      |
        And shapes.C ← glass_sphere() with:
            | transform                 | translation(0, 0, 0.25) |
            | material.refractive_index | 2.5                     |
        And ray.r ← ray(point(0, 0, -4), vector(0, 0, 1))
        And arrayintersections.xs ← intersections(2:A, 2.75:B, 3.25:C, 4.75:B, 5.25:C, 6:A)
        When computes.comps ← prepare_computations(arrayintersections.xs[<index>], ray.r, arrayintersections.xs)
        Then computes.comps.n1 = <n1>
        And computes.comps.n2 = <n2>

        Examples:
            | index | n1  | n2  |
            | 0     | 1.0 | 1.5 |
            | 1     | 1.5 | 2.0 |
            | 2     | 2.0 | 2.5 |
            | 3     | 2.5 | 2.5 |
            | 4     | 2.5 | 1.5 |
            | 5     | 1.5 | 1.0 |
    Scenario: The under point is offset below the surface
        Given ray.r ← ray(point(0, 0, -5), vector(0, 0, 1))
        And shapes.shape ← glass_sphere() with:
            | transform | translation(0, 0, 1) |
        And intersection.i ← intersection(5, shapes.shape)
        And arrayintersections.xs ← intersections(intersection.i)
        When computes.comps ← prepare_computations(intersection.i, ray.r, arrayintersections.xs)
        Then computes.comps.under_point.z > EPSILON/2
        And computes.comps.point.z < computes.comps.under_point.z
    Scenario: The Schlick approximation under total internal reflection
        Given shapes.shape ← glass_sphere()
        And ray.r ← ray(point(0, 0, √2/2), vector(0, 1, 0))
        And arrayintersections.xs ← intersections(-√2/2:shapes.shape, √2/2:shapes.shape)
        When computes.comps ← prepare_computations(arrayintersections.xs[1], ray.r, arrayintersections.xs)
        And floats.reflectance ← schlick(computes.comps)
        Then floats.reflectance = 1.0
    Scenario: The Schlick approximation with a perpendicular viewing angle
        Given shapes.shape ← glass_sphere()
        And ray.r ← ray(point(0, 0, 0), vector(0, 1, 0))
        And arrayintersections.xs ← intersections(-1:shapes.shape, 1:shapes.shape)
        When computes.comps ← prepare_computations(arrayintersections.xs[1], ray.r, arrayintersections.xs)
        And floats.reflectance ← schlick(computes.comps)
        Then floats.reflectance = 0.04
    Scenario: The Schlick approximation with small angle and n2 > n1
        Given shapes.shape ← glass_sphere()
        And ray.r ← ray(point(0, 0.99, -2), vector(0, 0, 1))
        And arrayintersections.xs ← intersections(1.8589:shapes.shape)
        When computes.comps ← prepare_computations(arrayintersections.xs[0], ray.r, arrayintersections.xs)
        And floats.reflectance ← schlick(computes.comps)
        Then floats.reflectance = 0.48873