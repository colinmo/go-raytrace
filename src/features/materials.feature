Feature: Materials

    Feature Description

    Background:
        Given material.m ← material()
        And tuple.position ← point(0, 0, 0)

    Scenario: The default material
        Given material.m ← material()
        Then material.m.color = color(1, 1, 1)
        And material.m.ambient = 0.1
        And material.m.diffuse = 0.9
        And material.m.specular = 0.9
        And material.m.shininess = 200.0

    Scenario: Lighting with the eye between the light and the surface
        Given tuple.eyev ← vector(0, 0, -1)
        And tuple.normalv ← vector(0, 0, -1)
        And light.light ← point_light(point(0, 0, -10), color(1, 1, 1))
        When colors.result ← lighting(material.m, light.light, tuple.position, tuple.eyev, tuple.normalv)
        Then colors.result = color(1.9, 1.9, 1.9)
    Scenario: Lighting with the eye between light and surface, eye offset 45°
        Given tuple.eyev ← vector(0, √2/2, -√2/2)
        And tuple.normalv ← vector(0, 0, -1)
        And light.light ← point_light(point(0, 0, -10), color(1, 1, 1))
        When colors.result ← lighting(material.m, light.light, tuple.position, tuple.eyev, tuple.normalv)
        Then colors.result = color(1.0, 1.0, 1.0)
    Scenario: Lighting with eye opposite surface, light offset 45°
        Given tuple.eyev ← vector(0, 0, -1)
        And tuple.normalv ← vector(0, 0, -1)
        And light.light ← point_light(point(0, 10, -10), color(1, 1, 1))
        When colors.result ← lighting(material.m, light.light, tuple.position, tuple.eyev, tuple.normalv)
        Then colors.result = color(0.7364, 0.7364, 0.7364)
    Scenario: Lighting with eye in the path of the reflection vector
        Given tuple.eyev ← vector(0, -√2/2, -√2/2)
        And tuple.normalv ← vector(0, 0, -1)
        And light.light ← point_light(point(0, 10, -10), color(1, 1, 1))
        When colors.result ← lighting(material.m, light.light, tuple.position, tuple.eyev, tuple.normalv)
        Then colors.result = color(1.6364, 1.6364, 1.6364)
    Scenario: Lighting with the light behind the surface
        Given tuple.eyev ← vector(0, 0, -1)
        And tuple.normalv ← vector(0, 0, -1)
        And light.light ← point_light(point(0, 0, 10), color(1, 1, 1))
        When colors.result ← lighting(material.m, light.light, tuple.position, tuple.eyev, tuple.normalv)
        Then colors.result = color(0.1, 0.1, 0.1)