Feature: Cones

    Feature Description

    Scenario Outline: Intersecting a cone with a ray
        Given shapes.shape ← cone()
        And tuple.direction ← normalize(<direction>)
        And ray.r ← ray(<origin>, tuple.direction)
        When arrayintersections.xs ← local_intersect(shapes.shape, ray.r)
        Then arrayintersections.xs.count = 2
        And arrayintersections.xs[0].t = <t0>
        And arrayintersections.xs[1].t = <t1>
        Examples:
            | origin          | direction           | t0      | t1       |
            | point(0, 0, -5) | vector(0, 0, 1)     | 5       | 5        |
            | point(0, 0, -5) | vector(1, 1, 1)     | 8.66025 | 8.66025  |
            | point(1, 1, -5) | vector(-0.5, -1, 1) | 4.55006 | 49.44994 |
    Scenario: Intersecting a cone with a ray parallel to one of its halves
        Given shapes.shape ← cone()
        And tuple.direction ← normalize(vector(0, 1, 1))
        And ray.r ← ray(point(0, 0, -1), tuple.direction)
        When arrayintersections.xs ← local_intersect(shapes.shape, ray.r)
        Then arrayintersections.xs.count = 1
        And arrayintersections.xs[0].t = 0.35355
    Scenario Outline: Intersecting a cone's end caps
        Given shapes.shape ← cone()
        And shapes.shape.minimum ← -0.5
        And shapes.shape.maximum ← 0.5
        And shapes.shape.closed ← true
        And tuple.direction ← normalize(<direction>)
        And ray.r ← ray(<origin>, tuple.direction)
        When arrayintersections.xs ← local_intersect(shapes.shape, ray.r)
        Then arrayintersections.xs.count = <count>
        Examples:
            | origin             | direction       | count |
            | point(0, 0, -5)    | vector(0, 1, 0) | 0     |
            | point(0, 0, -0.25) | vector(0, 1, 1) | 2     |
            | point(0, 0, -0.25) | vector(0, 1, 0) | 4     |
    Scenario Outline: Computing the normal vector on a cone
        Given shapes.shape ← cone()
        When tuple.n ← local_normal_at(shapes.shape, <point>)
        Then tuple.n = <normal>
        Examples:
            | point            | normal            |
            | point(0, 0, 0)   | vector(0, 0, 0)   |
            | point(1, 1, 1)   | vector(1, -√2, 1) |
            | point(-1, -1, 0) | vector(-1, 1, 0)  |