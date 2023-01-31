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