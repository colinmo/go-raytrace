Feature: Rays

    We need rays

    Scenario: Creating and querying a ray
        Given tuple.origin ← point(1, 2, 3)
        And tuple.direction ← vector(4, 5, 6)
        When ray.r ← ray(tuple.origin, tuple.direction)
        Then ray.r.origin = tuple.origin
        And ray.r.direction = tuple.direction

    Scenario: Computing a point from a distance
        Given ray.r ← ray(point(2, 3, 4), vector(1, 0, 0))
        Then position(ray.r, 0) = point(2, 3, 4)
        And position(ray.r, 1) = point(3, 3, 4)
        And position(ray.r, -1) = point(1, 3, 4)
        And position(ray.r, 2.5) = point(4.5, 3, 4)
    Scenario: Translating a ray
        Given ray.r ← ray(point(1, 2, 3), vector(0, 1, 0))
        And matrix.m ← translation(3, 4, 5)
        When ray.r2 ← transform(ray.r, matrix.m)
        Then ray.r2.origin = point(4, 6, 8)
        And ray.r2.direction = vector(0, 1, 0)
    Scenario: Scaling a ray
        Given ray.r ← ray(point(1, 2, 3), vector(0, 1, 0))
        And matrix.m ← scaling(2, 3, 4)
        When ray.r2 ← transform(ray.r, matrix.m)
        Then ray.r2.origin = point(2, 6, 12)
        And ray.r2.direction = vector(0, 3, 0)