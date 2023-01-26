Feature: Canvas checks
    In order to see the output of our work
    As an image
    We need a canvas to have pixels and visible image formats

    Scenario: Creating a canvas
        Given canvas.c ← canvas(10, 20)
        Then canvas.c.width = 10
        And canvas.c.height = 20
        And every pixel of canvas.c is color(0, 0, 0)

    Scenario: Writing pixels to a canvas
        Given canvas.c ← canvas(10, 20)
        And colors.red ← color(1, 0, 0)
        When write_pixel(canvas.c, 2, 3, colors.red)
        Then pixel_at(canvas.c, 2, 3) = colors.red
    Scenario: Constructing the PPM header
        Given canvas.c ← canvas(5, 3)
        When ppm.ppm ← canvas_to_ppm(canvas.c)
        Then lines 1-3 of ppm.ppm are
            """
            P3
            5 3
            255
            """
    Scenario: Constructing the PPM pixel data
        Given canvas.c ← canvas(5, 3)
        And colors.c1 ← color(1.5, 0, 0)
        And colors.c2 ← color(0, 0.5, 0)
        And colors.c3 ← color(-0.5, 0, 1)
        When write_pixel(canvas.c, 0, 0, colors.c1)
        And write_pixel(canvas.c, 2, 1, colors.c2)
        And write_pixel(canvas.c, 4, 2, colors.c3)
        And ppm.ppm ← canvas_to_ppm(canvas.c)
        Then lines 4-6 of ppm.ppm are
            """
            255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
            0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
            0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
            """
    Scenario: Splitting long lines in PPM files
        Given canvas.c ← canvas(10, 2)
        When every pixel of canvas.c is set to color(1, 0.8, 0.6)
        And ppm.ppm ← canvas_to_ppm(canvas.c)
        Then lines 4-7 of ppm.ppm are
            """
            255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
            153 255 204 153 255 204 153 255 204 153 255 204 153
            255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
            153 255 204 153 255 204 153 255 204 153 255 204 153
            """
    Scenario: PPM files are terminated by a newline character
        Given canvas.c ← canvas(5, 3)
        When ppm.ppm ← canvas_to_ppm(canvas.c)
        Then ppm.ppm ends with a newline character