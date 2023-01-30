Feature: Spheres

    In order to make spheres
    As a program
    I need to calculate sphere things

    Scenario: A ray intersects a sphere at two points
        Given ray.r ← ray(point(0, 0, -5), vector(0, 0, 1))
        And sphere.s ← sphere()
        When slice.xs ← intersect(sphere.s, ray.r)
        Then slice.xs.count = 2
        And slice.xs[0] = 4.0
        And slice.xs[1] = 6.0
    Scenario: A ray intersects a sphere at a tangent
        Given ray.r ← ray(point(0, 1, -5), vector(0, 0, 1))
        And sphere.s ← sphere()
        When slice.xs ← intersect(sphere.s, ray.r)
        Then slice.xs.count = 2
        And slice.xs[0] = 5.0
        And slice.xs[1] = 5.0
    Scenario: A ray misses a sphere
        Given ray.r ← ray(point(0, 2, -5), vector(0, 0, 1))
        And sphere.s ← sphere()
        When slice.xs ← intersect(sphere.s, ray.r)
        Then slice.xs.count = 0
    Scenario: A ray originates inside a sphere
        Given ray.r ← ray(point(0, 0, 0), vector(0, 0, 1))
        And sphere.s ← sphere()
        When slice.xs ← intersect(sphere.s, ray.r)
        Then slice.xs.count = 2
        And slice.xs[0] = -1.0
        And slice.xs[1] = 1.0
    Scenario: A sphere is behind a ray
        Given ray.r ← ray(point(0, 0, 5), vector(0, 0, 1))
        And sphere.s ← sphere()
        When slice.xs ← intersect(sphere.s, ray.r)
        Then slice.xs.count = 2
        And slice.xs[0] = -6.0
        And slice.xs[1] = -4.0

#    Scenario: An intersection encapsulates t and object
#        Given sphere.s ← sphere()
#        When intersections.i ← intersection(3.5, sphere.s)
#        Then intersections.i.t = 3.5
#        And intersections.i.object = sphere.s
#    Scenario: Intersect sets the object on the intersection
#        Given ray.r ← ray(point(0, 0, -5), vector(0, 0, 1))
#        And sphere.s ← sphere()
#        When slice.xs ← intersect(sphere.s, ray.r)
#        Then slice.xs.count = 2
#        And slice.xs[0].object = sphere.s
#        And slice.xs[1].object = sphere.s