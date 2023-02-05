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
    Scenario: The normal on a sphere at a point on the x axis
        Given sphere.s ← sphere()
        When tuple.n ← normal_at(sphere.s, point(1, 0, 0))
        Then tuple.n = vector(1, 0, 0)
    Scenario: The normal on a sphere at a point on the y axis
        Given sphere.s ← sphere()
        When tuple.n ← normal_at(sphere.s, point(0, 1, 0))
        Then tuple.n = vector(0, 1, 0)
    Scenario: The normal on a sphere at a point on the z axis
        Given sphere.s ← sphere()
        When tuple.n ← normal_at(sphere.s, point(0, 0, 1))
        Then tuple.n = vector(0, 0, 1)
    Scenario: The normal on a sphere at a nonaxial point
        Given sphere.s ← sphere()
        When tuple.n ← normal_at(sphere.s, point(√3/3, √3/3, √3/3))
        Then tuple.n = vector(√3/3, √3/3, √3/3)
    Scenario: The normal is a normalized vector
        Given sphere.s ← sphere()
        When tuple.n ← normal_at(sphere.s, point(√3/3, √3/3, √3/3))
        Then tuple.n = normalize(n)
    Scenario: Computing the normal on a translated sphere
        Given sphere.s ← sphere()
        And set_transform(sphere.s, translation(0, 1, 0))
        When matrix.n ← normal_at(sphere.s, point(0, 1.70711, -0.70711))
        Then matrix.n = vector(0, 0.70711, -0.70711)
    Scenario: Computing the normal on a transformed sphere
        Given sphere.s ← sphere()
        And matrix.m ← scaling(1, 0.5, 1) * rotation_z(π/5)
        And set_transform(sphere.s, matrix.m)
        When matrix.n ← normal_at(sphere.s, point(0, √2/2, -√2/2))
        Then matrix.n = vector(0, 0.97014, -0.24254)
    Scenario: A sphere has a default material
        Given sphere.s ← sphere()
        When material.m ← sphere.s.material
        Then material.m = material()
    Scenario: A sphere may be assigned a material
        Given sphere.s ← sphere()
        And material.m ← material()
        And material.m.ambient ← 1
        When sphere.s.material ← material.m
        Then sphere.s.material = material.m