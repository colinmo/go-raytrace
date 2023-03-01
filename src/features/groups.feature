Feature: Groups

    Feature Description

    Scenario: Creating a new group
        Given shapes.g ← group()
        Then shapes.g.transform = identity_matrix
        And shapes.g is empty
    Scenario: A shape has a parent attribute
        Given shapes.s ← test_shape()
        Then shapes.s.parent is nothing
#    Scenario: Adding a child to a group
#        Given shapes.g ← group()
#        And shapes.s ← test_shape()
#        When add_child(shapes.g, shapes.s)
#        Then shapes.g is not empty
#        And shapes.g includes shapes.s
#        And shapes.s.parent = shapes.g
    Scenario: Intersecting a ray with an empty group
        Given shapes.g ← group()
        And ray.r ← ray(point(0, 0, 0), vector(0, 0, 1))
        When arrayintersections.xs ← local_intersect(shapes.g, ray.r)
        Then arrayintersections.xs is empty

    Scenario: Intersecting a ray with a nonempty group
        Given shapes.g ← group()
        And shapes.s1 ← sphere()
        And shapes.s2 ← sphere()
        And set_transform(shapes.s2, translation(0, 0, -3))
        And shapes.s3 ← sphere()
        And set_transform(shapes.s3, translation(5, 0, 0))
        And add_child(shapes.g, shapes.s1)
        And add_child(shapes.g, shapes.s2)
        And add_child(shapes.g, shapes.s3)
        When ray.r ← ray(point(0, 0, -5), vector(0, 0, 1))
        And arrayintersections.xs ← local_intersect(shapes.g, ray.r)
        Then arrayintersections.xs.count = 4
        And arrayintersections.xs[0].object = shapes.s2
        And arrayintersections.xs[1].object = shapes.s2
        And arrayintersections.xs[2].object = shapes.s1
        And arrayintersections.xs[3].object = shapes.s1
Scenario: Intersecting a transformed group
Given shapes.g ← group()
And set_transform(shapes.g, scaling(2, 2, 2))
And shapes.s ← sphere()
And set_transform(shapes.s, translation(5, 0, 0))
And add_child(shapes.g, shapes.s)
When ray.r ← ray(point(10, 0, -10), vector(0, 0, 1))
And arrayintersections.xs ← intersect(shapes.g, ray.r)
Then arrayintersections.xs.count = 2        