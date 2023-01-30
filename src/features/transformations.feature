Feature: Transformation checks
    In order to move objects around the scene from origin
    As a component
    We need transformations on points/ vectors/ matrices

    Scenario: Multiplying by a translation matrix
        Given matrix.transform ← translation(5, -3, 2)
        And tuple.p ← point(-3, 4, 5)
        Then matrix.transform * tuple.p = point(2, 1, 7)
    Scenario: Multiplying by the inverse of a translation matrix
        Given matrix.transform ← translation(5, -3, 2)
        And matrix.inv ← inverse(matrix.transform)
        And tuple.p ← point(-3, 4, 5)
        Then matrix.inv * tuple.p = point(-8, 7, 3)
    Scenario: Translation does not affect vectors
        Given matrix.transform ← translation(5, -3, 2)
        And tuple.v ← vector(-3, 4, 5)
        Then matrix.transform * tuple.v = tuple.v
    Scenario: A scaling matrix applied to a point
        Given matrix.transform ← scaling(2, 3, 4)
        And tuple.p ← point(-4, 6, 8)
        Then matrix.transform * tuple.p = point(-8, 18, 32)
    Scenario: A scaling matrix applied to a vector
        Given matrix.transform ← scaling(2, 3, 4)
        And tuple.v ← vector(-4, 6, 8)
        Then matrix.transform * tuple.v = vector(-8, 18, 32)
    Scenario: Multiplying by the inverse of a scaling matrix
        Given matrix.transform ← scaling(2, 3, 4)
        And matrix.inv ← inverse(matrix.transform)
        And tuple.v ← vector(-4, 6, 8)
        Then matrix.inv * tuple.v = vector(-2, 2, 2)
    Scenario: Reflection is scaling by a negative value
        Given matrix.transform ← scaling(-1, 1, 1)
        And tuple.p ← point(2, 3, 4)
        Then matrix.transform * tuple.p = point(-2, 3, 4)
    Scenario: Rotating a point around the x axis
        Given tuple.p ← point(0, 1, 0)
        And matrix.half_quarter ← rotation_x(π / 4)
        And matrix.full_quarter ← rotation_x(π / 2)
        Then matrix.half_quarter * tuple.p = point(0, √2/2, √2/2)
        And matrix.full_quarter * tuple.p = point(0, 0, 1)
    Scenario: The inverse of an x-rotation rotates in the opposite direction
        Given tuple.p ← point(0, 1, 0)
        And matrix.half_quarter ← rotation_x(π / 4)
        And matrix.inv ← inverse(matrix.half_quarter)
        Then matrix.inv * tuple.p = point(0, √2/2, -√2/2)
    Scenario: Rotating a point around the y axis
        Given tuple.p ← point(0, 0, 1)
        And matrix.half_quarter ← rotation_y(π / 4)
        And matrix.full_quarter ← rotation_y(π / 2)
        Then matrix.half_quarter * tuple.p = point(√2/2, 0, √2/2)
        And matrix.full_quarter * tuple.p = point(1, 0, 0)
    Scenario: Rotating a point around the z axis
        Given tuple.p ← point(0, 1, 0)
        And matrix.half_quarter ← rotation_z(π / 4)
        And matrix.full_quarter ← rotation_z(π / 2)
        Then matrix.half_quarter * tuple.p = point(-√2/2, √2/2, 0)
        And matrix.full_quarter * tuple.p = point(-1, 0, 0)
    Scenario: A shearing transformation moves x in proportion to y
        Given matrix.transform ← shearing(1, 0, 0, 0, 0, 0)
        And tuple.p ← point(2, 3, 4)
        Then matrix.transform * tuple.p = point(5, 3, 4)
    Scenario: A shearing transformation moves x in proportion to z
        Given matrix.transform ← shearing(0, 1, 0, 0, 0, 0)
        And tuple.p ← point(2, 3, 4)
        Then matrix.transform * tuple.p = point(6, 3, 4)
    Scenario: A shearing transformation moves y in proportion to x
        Given matrix.transform ← shearing(0, 0, 1, 0, 0, 0)
        And tuple.p ← point(2, 3, 4)
        Then matrix.transform * tuple.p = point(2, 5, 4)
    Scenario: A shearing transformation moves y in proportion to z
        Given matrix.transform ← shearing(0, 0, 0, 1, 0, 0)
        And tuple.p ← point(2, 3, 4)
        Then matrix.transform * tuple.p = point(2, 7, 4)
    Scenario: A shearing transformation moves z in proportion to x
        Given matrix.transform ← shearing(0, 0, 0, 0, 1, 0)
        And tuple.p ← point(2, 3, 4)
        Then matrix.transform * tuple.p = point(2, 3, 6)
    Scenario: A shearing transformation moves z in proportion to y
        Given matrix.transform ← shearing(0, 0, 0, 0, 0, 1)
        And tuple.p ← point(2, 3, 4)
        Then matrix.transform * tuple.p = point(2, 3, 7)
    Scenario: Individual transformations are applied in sequence
        Given tuple.p ← point(1, 0, 1)
        And matrix.A ← rotation_x(π / 2)
        And matrix.B ← scaling(5, 5, 5)
        And matrix.C ← translation(10, 5, 7)
        # apply rotation first
        When tuple.p2 ← matrix.A * tuple.p
        Then tuple.p2 = point(1, -1, 0)
        # then apply scaling
        When tuple.p3 ← matrix.B * tuple.p2
        Then tuple.p3 = point(5, -5, 0)
        # then apply translation
        When tuple.p4 ← matrix.C * tuple.p3
        Then tuple.p4 = point(15, 0, 7)
    Scenario: Chained transformations must be applied in reverse order
        Given tuple.p ← point(1, 0, 1)
        And matrix.A ← rotation_x(π / 2)
        And matrix.B ← scaling(5, 5, 5)
        And matrix.C ← translation(10, 5, 7)
        When matrix.T ← matrix.C * matrix.B * matrix.A
        Then matrix.T * tuple.p = point(15, 0, 7)