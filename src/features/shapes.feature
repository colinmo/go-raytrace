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
    Scenario: Converting a point from world to object space
        Given shapes.g1 ← group()
        And set_transform(shapes.g1, rotation_y(π/2))
        And shapes.g2 ← group()
        And set_transform(shapes.g2, scaling(2, 2, 2))
        And add_child(shapes.g1, shapes.g2)
        And shapes.s ← sphere()
        And set_transform(shapes.s, translation(5, 0, 0))
        And add_child(shapes.g2, shapes.s)
        When tuple.p ← world_to_object(shapes.s, point(-2, 0, -10))
        Then tuple.p = point(0, 0, -1)
    Scenario: Converting a normal from object to world space
        Given shapes.g1 ← group()
        And set_transform(shapes.g1, rotation_y(π/2))
        And shapes.g2 ← group()
        And set_transform(shapes.g2, scaling(1, 2, 3))
        And add_child(shapes.g1, shapes.g2)
        And shapes.s ← sphere()
        And set_transform(shapes.s, translation(5, 0, 0))
        And add_child(shapes.g2, shapes.s)
        When tuple.n ← normal_to_world(shapes.s, vector(√3/3, √3/3, √3/3))
        Then tuple.n = vector(0.2857, 0.4286, -0.8571)
    Scenario: Finding the normal on a child object
        Given shapes.g1 ← group()
        And set_transform(shapes.g1, rotation_y(π/2))
        And shapes.g2 ← group()
        And set_transform(shapes.g2, scaling(1, 2, 3))
        And add_child(shapes.g1, shapes.g2)
        And shapes.s ← sphere()
        And set_transform(shapes.s, translation(5, 0, 0))
        And add_child(shapes.g2, shapes.s)
        When tuple.n ← normal_at(shapes.s, point(1.7321, 1.1547, -5.5774))
        Then tuple.n = vector(0.2857, 0.4286, -0.8571)