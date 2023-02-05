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