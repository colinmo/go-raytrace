Feature: Planes

    Feature Description

    Scenario: The normal of a plane is constant everywhere
        Given shapes.p ← plane()
        When tuple.n1 ← local_normal_at(shapes.p, point(0, 0, 0))
        And tuple.n2 ← local_normal_at(shapes.p, point(10, 0, -10))
        And tuple.n3 ← local_normal_at(shapes.p, point(-5, 0, 150))
        Then tuple.n1 = vector(0, 1, 0)
        And tuple.n2 = vector(0, 1, 0)
        And tuple.n3 = vector(0, 1, 0)
    Scenario: Intersect with a ray parallel to the plane
        Given shapes.p ← plane()
        And ray.r ← ray(point(0, 10, 0), vector(0, 0, 1))
        When arrayintersections.xs ← local_intersect(shapes.p, ray.r)
        Then arrayintersections.xs is empty
    Scenario: Intersect with a coplanar ray
        Given shapes.p ← plane()
        And ray.r ← ray(point(0, 0, 0), vector(0, 0, 1))
        When arrayintersections.xs ← local_intersect(shapes.p, ray.r)
        Then arrayintersections.xs is empty
    Scenario: A ray intersecting a plane from above
        Given shapes.p ← plane()
        And ray.r ← ray(point(0, 1, 0), vector(0, -1, 0))
        When arrayintersections.xs ← local_intersect(shapes.p, ray.r)
        Then arrayintersections.xs.count = 1
        And arrayintersections.xs[0].t = 1
        And arrayintersections.xs[0].object = shapes.p
    Scenario: A ray intersecting a plane from below
        Given shapes.p ← plane()
        And ray.r ← ray(point(0, -1, 0), vector(0, 1, 0))
        When arrayintersections.xs ← local_intersect(shapes.p, ray.r)
        Then arrayintersections.xs.count = 1
        And arrayintersections.xs[0].t = 1
        And arrayintersections.xs[0].object = shapes.p