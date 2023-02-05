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
    Scenario: There is no shadow when nothing is collinear with point and light
        Given world.w ← default_world()
        And tuple.p ← point(0, 10, 0)
        Then is_shadowed(world.w, tuple.p) is false
    Scenario: The shadow when an object is between the point and the light
        Given world.w ← default_world()
        And tuple.p ← point(10, -10, 10)
        Then is_shadowed(world.w, tuple.p) is true
    Scenario: There is no shadow when an object is behind the light
        Given world.w ← default_world()
        And tuple.p ← point(-20, 20, -20)
        Then is_shadowed(world.w, tuple.p) is false
    Scenario: There is no shadow when an object is behind the point
        Given world.w ← default_world()
        And tuple.p ← point(-2, 2, -2)
        Then is_shadowed(world.w, tuple.p) is false
    Scenario: shade_hit() is given an intersection in shadow
        Given world.w ← world()
        And world.w.light ← point_light(point(0, 0, -10), color(1, 1, 1))
        And sphere.s1 ← sphere()
        And sphere.s1 is added to world.w
        And sphere.s2 ← sphere() with:
            | transform | translation(0, 0, 10) |
        And sphere.s2 is added to world.w
        And ray.r ← ray(point(0, 0, 5), vector(0, 0, 1))
        And intersection.i ← intersection(4, sphere.s2)
        When computes.comps ← prepare_computations(intersection.i, ray.r)
        And colors.c ← shade_hit(world.w, computes.comps)
        Then colors.c = color(0.1, 0.1, 0.1)