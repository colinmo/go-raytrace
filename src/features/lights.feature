Feature: Lights

    Feature Description

    Scenario: A point light has a position and intensity
        Given colors.intensity ← color(1, 1, 1)
        And tuple.position ← point(0, 0, 0)
        When light.light ← point_light(tuple.position, colors.intensity)
        Then light.light.position = tuple.position
        And light.light.intensity = colors.intensity