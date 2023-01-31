Feature: Spheres

    In order to make spheres
    As a program
    I need to calculate sphere things

    Scenario: A ray intersects a sphere at two points
        Given ray.r ← ray(point(0, 0, -5), vector(0, 0, 1))
        And sphere.s ← sphere()
        When arrayintersections.xs ← intersect(sphere.s, ray.r)
        Then arrayintersections.xs.count = 2
        And arrayintersections.xs[0] = 4.0
        And arrayintersections.xs[1] = 6.0
    Scenario: A ray intersects a sphere at a tangent
        Given ray.r ← ray(point(0, 1, -5), vector(0, 0, 1))
        And sphere.s ← sphere()
        When arrayintersections.xs ← intersect(sphere.s, ray.r)
        Then arrayintersections.xs.count = 2
        And arrayintersections.xs[0] = 5.0
        And arrayintersections.xs[1] = 5.0
    Scenario: A ray misses a sphere
        Given ray.r ← ray(point(0, 2, -5), vector(0, 0, 1))
        And sphere.s ← sphere()
        When arrayintersections.xs ← intersect(sphere.s, ray.r)
        Then arrayintersections.xs.count = 0
    Scenario: A ray originates inside a sphere
        Given ray.r ← ray(point(0, 0, 0), vector(0, 0, 1))
        And sphere.s ← sphere()
        When arrayintersections.xs ← intersect(sphere.s, ray.r)
        Then arrayintersections.xs.count = 2
        And arrayintersections.xs[0] = -1.0
        And arrayintersections.xs[1] = 1.0
    Scenario: A sphere is behind a ray
        Given ray.r ← ray(point(0, 0, 5), vector(0, 0, 1))
        And sphere.s ← sphere()
        When arrayintersections.xs ← intersect(sphere.s, ray.r)
        Then arrayintersections.xs.count = 2
        And arrayintersections.xs[0] = -6.0
        And arrayintersections.xs[1] = -4.0

    Scenario: Intersect sets the object on the intersection
        Given ray.r ← ray(point(0, 0, -5), vector(0, 0, 1))
        And sphere.s ← sphere()
        When arrayintersections.xs ← intersect(sphere.s, ray.r)
        Then arrayintersections.xs.count = 2
        And arrayintersections.xs[0].object = sphere.s
        And arrayintersections.xs[1].object = sphere.s

    Scenario: A sphere's default transformation
        Given sphere.s ← sphere()
        Then sphere.s.transform = identity_matrix
    Scenario: Changing a sphere's transformation
        Given sphere.s ← sphere()
        And matrix.t ← translation(2, 3, 4)
        When set_transform(sphere.s, matrix.t)
        Then sphere.s.transform = matrix.t
    Scenario: Intersecting a scaled sphere with a ray
        Given ray.r ← ray(point(0, 0, -5), vector(0, 0, 1))
        And sphere.s ← sphere()
        When set_transform(sphere.s, scaling(2, 2, 2))
        And arrayintersections.xs ← intersect(sphere.s, ray.r)
        Then arrayintersections.xs.count = 2
        And arrayintersections.xs[0].t = 3
        And arrayintersections.xs[1].t = 7
    Scenario: Intersecting a translated sphere with a ray
        Given ray.r ← ray(point(0, 0, -5), vector(0, 0, 1))
        And sphere.s ← sphere()
        When set_transform(sphere.s, translation(5, 0, 0))
        And arrayintersections.xs ← intersect(sphere.s, ray.r)
        Then arrayintersections.xs.count = 0