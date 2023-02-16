Feature: Cylinders

    Feature Description

    Scenario Outline: A ray misses a cylinder
        Given shapes.cyl ← cylinder()
        And tuple.direction ← normalize(<direction>)
        And ray.r ← ray(<origin>, tuple.direction)
        When arrayintersections.xs ← local_intersect(shapes.cyl, ray.r)
        Then arrayintersections.xs.count = 0
        Examples:
            | origin          | direction       |
            | point(1, 0, 0)  | vector(0, 1, 0) |
            | point(0, 0, 0)  | vector(0, 1, 0) |
            | point(0, 0, -5) | vector(1, 1, 1) |
    Scenario Outline: A ray strikes a cylinder
        Given shapes.cyl ← cylinder()
        And tuple.direction ← normalize(<direction>)
        And ray.r ← ray(<origin>, tuple.direction)
        When arrayintersections.xs ← local_intersect(shapes.cyl, ray.r)
        Then arrayintersections.xs.count = 2
        And arrayintersections.xs[0].t = <t0>
        And arrayintersections.xs[1].t = <t1>
        Examples:
            | origin            | direction         | t0      | t1      |
            | point(1, 0, -5)   | vector(0, 0, 1)   | 5       | 5       |
            | point(0, 0, -5)   | vector(0, 0, 1)   | 4       | 6       |
            | point(0.5, 0, -5) | vector(0.1, 1, 1) | 6.80798 | 7.08872 |
    Scenario Outline: Normal vector on a cylinder
        Given shapes.cyl ← cylinder()
        When tuple.n ← local_normal_at(shapes.cyl, <point>)
        Then tuple.n = <normal>
        Examples:
            | point           | normal           |
            | point(1, 0, 0)  | vector(1, 0, 0)  |
            | point(0, 5, -1) | vector(0, 0, -1) |
            | point(0, -2, 1) | vector(0, 0, 1)  |
            | point(-1, 1, 0) | vector(-1, 0, 0) |
    Scenario: The default minimum and maximum for a cylinder
        Given shapes.cyl ← cylinder()
        Then shapes.cyl.minimum = -infinity
        And shapes.cyl.maximum = infinity
    Scenario Outline: Intersecting a constrained cylinder
        Given shapes.cyl ← cylinder()
        And shapes.cyl.minimum ← 1
        And shapes.cyl.maximum ← 2
        And tuple.direction ← normalize(<direction>)
        And ray.r ← ray(<point>, tuple.direction)
        When arrayintersections.xs ← local_intersect(shapes.cyl, ray.r)
        Then arrayintersections.xs.count = <count>
        Examples:
            |   | point             | direction         | count |
            | 1 | point(0, 1.5, 0)  | vector(0.1, 1, 0) | 0     |
            | 2 | point(0, 3, -5)   | vector(0, 0, 1)   | 0     |
            | 3 | point(0, 0, -5)   | vector(0, 0, 1)   | 0     |
            | 4 | point(0, 2, -5)   | vector(0, 0, 1)   | 0     |
            | 5 | point(0, 1, -5)   | vector(0, 0, 1)   | 0     |
            | 6 | point(0, 1.5, -2) | vector(0, 0, 1)   | 2     |
    Scenario: The default closed value for a cylinder
        Given shapes.cyl ← cylinder()
        Then shapes.cyl.closed = false
    Scenario Outline: Intersecting the caps of a closed cylinder
        Given shapes.cyl ← cylinder()
        And shapes.cyl.minimum ← 1
        And shapes.cyl.maximum ← 2
        And shapes.cyl.closed ← true
        And tuple.direction ← normalize(<direction>)
        And ray.r ← ray(<point>, tuple.direction)
        When arrayintersections.xs ← local_intersect(shapes.cyl, ray.r)
        Then arrayintersections.xs.count = <count>
        Examples:
            |   | point            | direction        | count |
            | 1 | point(0, 3, 0)   | vector(0, -1, 0) | 2     |
            | 2 | point(0, 3, -2)  | vector(0, -1, 2) | 2     |
            | 3 | point(0, 4, -2)  | vector(0, -1, 1) | 2     |
            | 4 | point(0, 0, -2)  | vector(0, 1, 2)  | 2     |
            | 5 | point(0, -1, -2) | vector(0, 1, 1)  | 2     |
    Scenario Outline: The normal vector on a cylinder's end caps
        Given shapes.cyl ← cylinder()
        And shapes.cyl.minimum ← 1
        And shapes.cyl.maximum ← 2
        And shapes.cyl.closed ← true
        When tuple.n ← local_normal_at(shapes.cyl, <point>)
        Then tuple.n = <normal>
        Examples:
            | point            | normal           |
            | point(0, 1, 0)   | vector(0, -1, 0) |
            | point(0.5, 1, 0) | vector(0, -1, 0) |
            | point(0, 1, 0.5) | vector(0, -1, 0) |
            | point(0, 2, 0)   | vector(0, 1, 0)  |
            | point(0.5, 2, 0) | vector(0, 1, 0)  |
            | point(0, 2, 0.5) | vector(0, 1, 0)  |