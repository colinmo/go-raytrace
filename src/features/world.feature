Feature: TOMORROW THE WORLD

    Feature Description

    Scenario: Creating a world
        Given world.w ← world()
        Then world.w contains no objects
        And world.w has no light source
    Scenario: The default world
        Given light.light ← point_light(point(-10, 10, -10), color(1, 1, 1))
        And sphere.s1 ← sphere() with:
            | color    | (0.8, 1.0, 0.6) |
            | diffuse  | 0.7             |
            | specular | 0.2             |
        And sphere.s2 ← sphere() with:
            | transform | scaling(0.5, 0.5, 0.5) |
        When world.w ← default_world()
        Then world.w.light = light.light
        And world.w contains sphere.s1
        And world.w contains sphere.s2
    Scenario: Intersect a world with a ray
        Given world.w ← default_world()
        And ray.r ← ray(point(0, 0, -5), vector(0, 0, 1))
        When arrayintersections.xs ← intersect_world(world.w, ray.r)
        Then arrayintersections.xs.count = 4
        And arrayintersections.xs[0].t = 4
        And arrayintersections.xs[1].t = 4.5
        And arrayintersections.xs[2].t = 5.5
        And arrayintersections.xs[3].t = 6
    Scenario: The color when a ray misses
        Given world.w ← default_world()
        And ray.r ← ray(point(0, 0, -5), vector(0, 1, 0))
        When colors.c ← color_at(world.w, ray.r)
        Then colors.c = color(0, 0, 0)
    Scenario: The color when a ray hits
        Given world.w ← default_world()
        And ray.r ← ray(point(0, 0, -5), vector(0, 0, 1))
        When colors.c ← color_at(world.w, ray.r)
        Then colors.c = color(0.38066, 0.47583, 0.2855)