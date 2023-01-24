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