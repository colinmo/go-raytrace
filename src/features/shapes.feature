Feature: Shapes

    Feature Description

    Scenario: The default transformation
        Given shapes.s ← test_shape()
        Then shapes.s.transform = identity_matrix
    Scenario: Assigning a transformation
        Given shapes.s ← test_shape()
        When set_transform(shapes.s, translation(2, 3, 4))
        Then shapes.s.transform = translation(2, 3, 4)
    Scenario: The default material
        Given shapes.s ← test_shape()
        When material.m ← shapes.s.material
        Then material.m = material()
    Scenario: Assigning a material
        Given shapes.s ← test_shape()
        And material.m ← material()
        And material.m.ambient ← 1
        When shapes.s.material ← material.m
        Then shapes.s.material = material.m
    Scenario: Intersecting a scaled shape with a ray
        Given ray.r ← ray(point(0, 0, -5), vector(0, 0, 1))
        And shapes.s ← test_shape()
        When set_transform(shapes.s, scaling(2, 2, 2))
        And arrayintersections.xs ← intersect(shapes.s, ray.r)
        Then shapes.s.saved_ray.origin = point(0, 0, -2.5)
        And shapes.s.saved_ray.direction = vector(0, 0, 0.5)
    Scenario: Intersecting a translated shape with a ray
        Given ray.r ← ray(point(0, 0, -5), vector(0, 0, 1))
        And shapes.s ← test_shape()
        When set_transform(shapes.s, translation(5, 0, 0))
        And arrayintersections.xs ← intersect(shapes.s, ray.r)
        Then shapes.s.saved_ray.origin = point(-5, 0, -5)
        And shapes.s.saved_ray.direction = vector(0, 0, 1)
    Scenario: Computing the normal on a translated shape
        Given shapes.s ← test_shape()
        When set_transform(shapes.s, translation(0, 1, 0))
        And tuple.n ← normal_at(shapes.s, point(0, 1.70711, -0.70711))
        Then tuple.n = vector(0, 0.70711, -0.70711)
    Scenario: Computing the normal on a transformed shape
        Given shapes.s ← test_shape()
        And matrix.m ← scaling(1, 0.5, 1) * rotation_z(π/5)
        When set_transform(shapes.s, matrix.m)
        And tuple.n ← normal_at(shapes.s, point(0, √2/2, -√2/2))
        Then tuple.n = vector(0, 0.97014, -0.24254)