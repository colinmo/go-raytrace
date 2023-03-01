Feature: Triangle

    Yay triangle

    Scenario: Constructing a triangle
        Given tuple.p1 ← point(0, 1, 0)
        And tuple.p2 ← point(-1, 0, 0)
        And tuple.p3 ← point(1, 0, 0)
        And shapes.t ← triangle(p1, p2, p3)
        Then shapes.t.p1 = tuple.p1
        And shapes.t.p2 = tuple.p2
        And shapes.t.p3 = tuple.p3
        And shapes.t.e1 = vector(-1, -1, 0)
        And shapes.t.e2 = vector(1, -1, 0)
        And shapes.t.normal = vector(0, 0, -1)
    Scenario: Finding the normal on a triangle
        Given shapes.t ← triangle(point(0, 1, 0), point(-1, 0, 0), point(1, 0, 0))
        When tuple.n1 ← local_normal_at(shapes.t, point(0, 0.5, 0))
        And tuple.n2 ← local_normal_at(shapes.t, point(-0.5, 0.75, 0))
        And tuple.n3 ← local_normal_at(shapes.t, point(0.5, 0.25, 0))
        Then tuple.n1 = shapes.t.normal
        And tuple.n2 = shapes.t.normal
        And tuple.n3 = shapes.t.normal
    Scenario: Intersecting a ray parallel to the triangle
        Given shapes.t ← triangle(point(0, 1, 0), point(-1, 0, 0), point(1, 0, 0))
        And ray.r ← ray(point(0, -1, -2), vector(0, 1, 0))
        When arrayintersections.xs ← local_intersect(shapes.t, ray.r)
        Then arrayintersections.xs is empty
    Scenario: A ray misses the p1-p3 edge
        Given shapes.t ← triangle(point(0, 1, 0), point(-1, 0, 0), point(1, 0, 0))
        And ray.r ← ray(point(1, 1, -2), vector(0, 0, 1))
        When arrayintersections.xs ← local_intersect(shapes.t, ray.r)
        Then arrayintersections.xs is empty

    Scenario: A ray misses the p1-p2 edge
        Given shapes.t ← triangle(point(0, 1, 0), point(-1, 0, 0), point(1, 0, 0))
        And ray.r ← ray(point(-1, 1, -2), vector(0, 0, 1))
        When arrayintersections.xs ← local_intersect(shapes.t, ray.r)
        Then arrayintersections.xs is empty
    Scenario: A ray misses the p2-p3 edge
        Given shapes.t ← triangle(point(0, 1, 0), point(-1, 0, 0), point(1, 0, 0))
        And ray.r ← ray(point(0, -1, -2), vector(0, 0, 1))
        When arrayintersections.xs ← local_intersect(shapes.t, ray.r)
        Then arrayintersections.xs is empty
    Scenario: A ray strikes a triangle
        Given shapes.t ← triangle(point(0, 1, 0), point(-1, 0, 0), point(1, 0, 0))
        And ray.r ← ray(point(0, 0.5, -2), vector(0, 0, 1))
        When arrayintersections.xs ← local_intersect(shapes.t, ray.r)
        Then arrayintersections.xs.count = 1
        And arrayintersections.xs[0].t = 2