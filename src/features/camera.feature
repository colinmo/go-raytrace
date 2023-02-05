Feature: Camera

    Feature Description

    Scenario: Constructing a camera
        Given camera.c ← camera(160, 120, π/2)
        Then camera.c.hsize = 160
        And camera.c.vsize = 120
        And camera.c.field_of_view = π/2
        And camera.c.transform = identity_matrix
    Scenario: The pixel size for a horizontal canvas
        Given camera.c ← camera(200, 125, π/2)
        Then camera.c.pixel_size = 0.01
    Scenario: The pixel size for a vertical canvas
        Given camera.c ← camera(125, 200, π/2)
        Then camera.c.pixel_size = 0.01
    Scenario: Constructing a ray through the center of the canvas
        Given camera.c ← camera(201, 101, π/2)
        When ray.r ← ray_for_pixel(camera.c, 100, 50)
        Then ray.r.origin = point(0, 0, 0)
        And ray.r.direction = vector(0, 0, -1)
    Scenario: Constructing a ray through a corner of the canvas
        Given camera.c ← camera(201, 101, π/2)
        When ray.r ← ray_for_pixel(camera.c, 0, 0)
        Then ray.r.origin = point(0, 0, 0)
        And ray.r.direction = vector(0.66519, 0.33259, -0.66851)
    Scenario: Constructing a ray when the camera is transformed
        Given camera.c ← camera(201, 101, π/2)
        When camera.c.transform ← rotation_y(π/4) * translation(0, -2, 5)
        And ray.r ← ray_for_pixel(camera.c, 100, 50)
        Then ray.r.origin = point(0, 2, -5)
        And ray.r.direction = vector(√2/2, 0, -√2/2)
    Scenario: Rendering a world with a camera
        Given world.w ← default_world()
        And camera.c ← camera(11, 11, π/2)
        And tuple.from ← point(0, 0, -5)
        And tuple.to ← point(0, 0, 0)
        And tuple.up ← vector(0, 1, 0)
        And camera.c.transform ← view_transform(tuple.from, tuple.to, tuple.up)
        When canvas.image ← render(camera.c, world.w)
        Then pixel_at(canvas.image, 5, 5) = color(0.38066, 0.47583, 0.2855)