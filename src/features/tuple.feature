Feature: Tuple checks
    In order to control things in 3d space
    As a codebase
    We need tuples to function as expected

    Scenario: A tuple with w=1.0 is a point
        Given tuple.a ← tuple(4.3, -4.2, 3.1, 1.0)
        Then tuple.a.x = 4.3
        And tuple.a.y = -4.2
        And tuple.a.z = 3.1
        And tuple.a.w = 1.0
        And tuple.a is a point
        And tuple.a is not a vector

    Scenario: A tuple with w=0 is a vector
        Given tuple.a ← tuple(4.3, -4.2, 3.1, 0.0)
        Then tuple.a.x = 4.3
        And tuple.a.y = -4.2
        And tuple.a.z = 3.1
        And tuple.a.w = 0.0
        And tuple.a is not a point
        And tuple.a is a vector

    Scenario: point() creates tuples with w=1
        Given tuple.p ← point(4, -4, 3)
        Then tuple.p = tuple(4, -4, 3, 1)

    Scenario: vector() creates tuples with w=0
        Given tuple.v ← vector(4, -4, 3)
        Then tuple.v = tuple(4, -4, 3, 0)

    Scenario: Adding two tuples
        Given tuple.a1 ← tuple(3, -2, 5, 1)
        And tuple.a2 ← tuple(-2, 3, 1, 0)
        Then tuple.a1 + tuple.a2 = tuple(1, 1, 6, 1)

    Scenario: Subtracting two points
        Given tuple.p1 ← point(3, 2, 1)
        And tuple.p2 ← point(5, 6, 7)
        Then tuple.p1 - tuple.p2 = vector(-2, -4, -6)

    Scenario: Subtracting a vector from a point
        Given tuple.p ← point(3, 2, 1)
        And tuple.v ← vector(5, 6, 7)
        Then tuple.p - tuple.v = point(-2, -4, -6)

    Scenario: Subtracting two vectors
        Given tuple.v1 ← vector(3, 2, 1)
        And tuple.v2 ← vector(5, 6, 7)
        Then tuple.v1 - tuple.v2 = vector(-2, -4, -6)

    Scenario: Subtracting a vector from the zero vector
        Given tuple.zero ← vector(0, 0, 0)
        And tuple.v ← vector(1, -2, 3)
        Then tuple.zero - tuple.v = vector(-1, 2, -3)

    Scenario: Negating a tuple
        Given tuple.a ← tuple(1, -2, 3, -4)
        Then -tuple.a = tuple(-1, 2, -3, 4)

    Scenario: Multiplying a tuple by a scalar
        Given tuple.a ← tuple(1, -2, 3, -4)
        Then tuple.a * 3.5 = tuple(3.5, -7, 10.5, -14)
    Scenario: Multiplying a tuple by a fraction
        Given tuple.a ← tuple(1, -2, 3, -4)
        Then tuple.a * 0.5 = tuple(0.5, -1, 1.5, -2)