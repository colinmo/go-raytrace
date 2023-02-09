Feature: Patterns

    Feature Description

    Background:
        Given color.black ← color(0, 0, 0)
        And color.white ← color(1, 1, 1)
    Scenario: Creating a stripe pattern
        Given pattern.pattern ← stripe_pattern(color.white, color.black)
        Then pattern.pattern.a = color.white
        And pattern.pattern.b = color.black
    Scenario: A stripe pattern is constant in y
        Given pattern.pattern ← stripe_pattern(color.white, color.black)
        Then stripe_at(pattern.pattern, point(0, 0, 0)) = color.white
        And stripe_at(pattern.pattern, point(0, 1, 0)) = color.white
        And stripe_at(pattern.pattern, point(0, 2, 0)) = color.white
    Scenario: A stripe pattern is constant in z
        Given pattern.pattern ← stripe_pattern(color.white, color.black)
        Then stripe_at(pattern.pattern, point(0, 0, 0)) = color.white
        And stripe_at(pattern.pattern, point(0, 0, 1)) = color.white
        And stripe_at(pattern.pattern, point(0, 0, 2)) = color.white
    Scenario: A stripe pattern alternates in x
        Given pattern.pattern ← stripe_pattern(color.white, color.black)
        Then stripe_at(pattern.pattern, point(0, 0, 0)) = color.white
        And stripe_at(pattern.pattern, point(0.9, 0, 0)) = color.white
        And stripe_at(pattern.pattern, point(1, 0, 0)) = color.black
        And stripe_at(pattern.pattern, point(-0.1, 0, 0)) = color.black
        And stripe_at(pattern.pattern, point(-1, 0, 0)) = color.black
        And stripe_at(pattern.pattern, point(-1.1, 0, 0)) = color.white
    Scenario: Stripes with an object transformation
        Given shapes.object ← sphere()
        And set_transform(shapes.object, scaling(2, 2, 2))
        And pattern.pattern ← stripe_pattern(color.white, color.black)
        When colors.c ← stripe_at_object(pattern.pattern, shapes.object, point(1.5, 0, 0))
        Then colors.c = colors.white
    Scenario: Stripes with a pattern transformation
        Given shapes.object ← sphere()
        And pattern.pattern ← stripe_pattern(color.white, color.black)
        And set_pattern_transform(pattern.pattern, scaling(2, 2, 2))
        When colors.c ← stripe_at_object(pattern.pattern, shapes.object, point(1.5, 0, 0))
        Then colors.c = colors.white
    Scenario: Stripes with both an object and a pattern transformation
        Given shapes.object ← sphere()
        And set_transform(shapes.object, scaling(2, 2, 2))
        And pattern.pattern ← stripe_pattern(color.white, color.black)
        And set_pattern_transform(pattern.pattern, translation(0.5, 0, 0))
        When colors.c ← stripe_at_object(pattern.pattern, shapes.object, point(2.5, 0, 0))
        Then colors.c = colors.white
    Scenario: The default pattern transformation
        Given pattern.pattern ← test_pattern()
        Then pattern.pattern.transform = identity_matrix
    Scenario: Assigning a transformation
        Given pattern.pattern ← test_pattern()
        When set_pattern_transform(pattern.pattern, translation(1, 2, 3))
        Then pattern.pattern.transform = translation(1, 2, 3)
    Scenario: A pattern with an object transformation
        Given shapes.shape ← sphere()
        And set_transform(shapes.shape, scaling(2, 2, 2))
        And pattern.pattern ← test_pattern()
        When colors.c ← pattern_at_shape(pattern.pattern, shapes.shape, point(2, 3, 4))
        Then colors.c = color(1, 1.5, 2)
    Scenario: A pattern with a pattern transformation
        Given shapes.shape ← sphere()
        And pattern.pattern ← test_pattern()
        And set_pattern_transform(pattern.pattern, scaling(2, 2, 2))
        When colors.c ← pattern_at_shape(pattern.pattern, shapes.shape, point(2, 3, 4))
        Then colors.c = color(1, 1.5, 2)
    Scenario: A pattern with both an object and a pattern transformation
        Given shapes.shape ← sphere()
        And set_transform(shapes.shape, scaling(2, 2, 2))
        And pattern.pattern ← test_pattern()
        And set_pattern_transform(pattern.pattern, translation(0.5, 1, 1.5))
        When colors.c ← pattern_at_shape(pattern.pattern, shapes.shape, point(2.5, 3, 3.5))
        Then colors.c = color(0.75, 0.5, 0.25)
    Scenario: A gradient linearly interpolates between colors
        Given pattern.pattern ← gradient_pattern(white, black)
        Then pattern_at(pattern.pattern, point(0, 0, 0)) = white
        And pattern_at(pattern.pattern, point(0.25, 0, 0)) = color(0.75, 0.75, 0.75)
        And pattern_at(pattern.pattern, point(0.5, 0, 0)) = color(0.5, 0.5, 0.5)
        And pattern_at(pattern.pattern, point(0.75, 0, 0)) = color(0.25, 0.25, 0.25)
    Scenario: A ring should extend in both x and z
        Given pattern.pattern ← ring_pattern(white, black)
        Then pattern_at(pattern.pattern, point(0, 0, 0)) = white
        And pattern_at(pattern.pattern, point(1, 0, 0)) = black
        And pattern_at(pattern.pattern, point(0, 0, 1)) = black
        # 0.708 = just slightly more than √2/2
        And pattern_at(pattern.pattern, point(0.708, 0, 0.708)) = black

    Scenario: Checkers should repeat in x
        Given pattern.pattern ← checkers_pattern(white, black)
        Then pattern_at(pattern.pattern, point(0, 0, 0)) = white
        And pattern_at(pattern.pattern, point(0.99, 0, 0)) = white
        And pattern_at(pattern.pattern, point(1.01, 0, 0)) = black
    Scenario: Checkers should repeat in y
        Given pattern.pattern ← checkers_pattern(white, black)
        Then pattern_at(pattern.pattern, point(0, 0, 0)) = white
        And pattern_at(pattern.pattern, point(0, 0.99, 0)) = white
        And pattern_at(pattern.pattern, point(0, 1.01, 0)) = black
    Scenario: Checkers should repeat in z
        Given pattern.pattern ← checkers_pattern(white, black)
        Then pattern_at(pattern.pattern, point(0, 0, 0)) = white
        And pattern_at(pattern.pattern, point(0, 0, 0.99)) = white
        And pattern_at(pattern.pattern, point(0, 0, 1.01)) = black