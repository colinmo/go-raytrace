Feature: Object files

    Parsing Obj files

    Scenario: Ignoring unrecognized lines
        Given files.gibberish ← a file containing:
            """
            There was a young lady named Bright
            who traveled much faster than light.
            She set out one day
            in a relative way,
            and came back the previous night.
            """
        When parsers.parser ← parse_obj_file(files.gibberish)
        Then parsers.parser should have ignored 5 lines
    Scenario: Vertex records
        Given files.file ← a file containing:
            """
            v -1 1 0
            v -1.0000 0.5000 0.0000
            v 1 0 0
            v 1 1 0
            """
        When parsers.parser ← parse_obj_file(files.file)
        Then parsers.parser.vertices[1] = point(-1, 1, 0)
        And parsers.parser.vertices[2] = point(-1, 0.5, 0)
        And parsers.parser.vertices[3] = point(1, 0, 0)
        And parsers.parser.vertices[4] = point(1, 1, 0)
    Scenario: Parsing triangle faces
        Given files.file ← a file containing:
            """
            v -1 1 0
            v -1 0 0
            v 1 0 0
            v 1 1 0
            f 1 2 3
            f 1 3 4
            """
        When parsers.parser ← parse_obj_file(files.file)
        And shapes.g ← parsers.parser.default_group
        And shapes.t1 ← first child of shapes.g
        And shapes.t2 ← second child of shapes.g
        Then shapes.t1.p1 = parsers.parser.vertices[1]
        And shapes.t1.p2 = parsers.parser.vertices[2]
        And shapes.t1.p3 = parsers.parser.vertices[3]
        And shapes.t2.p1 = parsers.parser.vertices[1]
        And shapes.t2.p2 = parsers.parser.vertices[3]
        And shapes.t2.p3 = parsers.parser.vertices[4]
    Scenario: Triangulating polygons
        Given files.file ← a file containing:
            """
            v -1 1 0
            v -1 0 0
            v 1 0 0
            v 1 1 0
            v 0 2 0
            f 1 2 3 4 5
            """
        When parsers.parser ← parse_obj_file(files.file)
        And shapes.g ← parsers.parser.default_group
        And shapes.t1 ← first child of shapes.g
        And shapes.t2 ← second child of shapes.g
        And shapes.t3 ← third child of shapes.g
        Then shapes.t1.p1 = parsers.parser.vertices[1]
        And shapes.t1.p2 = parsers.parser.vertices[2]
        And shapes.t1.p3 = parsers.parser.vertices[3]
        And shapes.t2.p1 = parsers.parser.vertices[1]
        And shapes.t2.p2 = parsers.parser.vertices[3]
        And shapes.t2.p3 = parsers.parser.vertices[4]
        And shapes.t3.p1 = parsers.parser.vertices[1]
        And shapes.t3.p2 = parsers.parser.vertices[4]
        And shapes.t3.p3 = parsers.parser.vertices[5]
    Scenario: Triangles in groups
        Given files.file ← the file "triangles.obj"
        When parsers.parser ← parse_obj_file(files.file)
        And shapes.g1 ← "FirstGroup" from parsers.parser
        And shapes.g2 ← "SecondGroup" from parsers.parser
        And shapes.t1 ← first child of shapes.g1
        And shapes.t2 ← first child of shapes.g2
        Then shapes.t1.p1 = parsers.parser.vertices[1]
        And shapes.t1.p2 = parsers.parser.vertices[2]
        And shapes.t1.p3 = parsers.parser.vertices[3]
        And shapes.t2.p1 = parsers.parser.vertices[1]
        And shapes.t2.p2 = parsers.parser.vertices[3]
        And shapes.t2.p3 = parsers.parser.vertices[4]
Scenario: Converting an OBJ file to a group
Given files.file ← the file "triangles.obj"
And parsers.parser ← parse_obj_file(files.file)
When shapes.g ← obj_to_group(parsers.parser)
Then shapes.g includes "FirstGroup" from parsers.parser
And shapes.g includes "SecondGroup" from parsers.parser        