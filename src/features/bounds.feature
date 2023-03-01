Feature: Bounds
    BOUNDS

    Scenario: Can make a bounds
        Given bounds.b ← new_bounds()
        Then bounds.b.maxpoint = point(-inf,-inf,-inf)
        And bounds.b.minpoint = point(inf, inf, inf)
