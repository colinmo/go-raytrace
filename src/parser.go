package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Parser struct {
	Output       string
	LinesParsed  int
	LinesSkipped int
	Vertices     []Tuple
	Groups       map[string]*Group
	ActiveGroup  string
}

func NewParser() *Parser {
	return &Parser{
		Output:       "",
		LinesParsed:  0,
		LinesSkipped: 0,
		Vertices:     []Tuple{NewTuple(0, 0, 0, 0)},
		Groups:       map[string]*Group{"Default": NewGroup()},
		ActiveGroup:  "Default",
	}
}

func NewParserFromFile(filename string) *Parser {
	b, e := os.ReadFile(filename)
	if e != nil {
		log.Fatalf("failed to read %s", filename)
	}

	p := Parser{
		Output:       "",
		LinesParsed:  0,
		LinesSkipped: 0,
		Vertices:     []Tuple{NewTuple(0, 0, 0, 0)},
		Groups:       map[string]*Group{"Default": NewGroup()},
		ActiveGroup:  "Default",
	}
	p.Parse(string(b))
	return &p
}

func (p *Parser) Parse(b string) {
	lines := strings.Split(b, "\n")
	repl := regexp.MustCompile(" +")
	for _, l := range lines {
		fmt.Printf("Line: %s\n", l)
		commandAndParams := strings.Split(repl.ReplaceAllString(l, ` `), " ")
		switch commandAndParams[0] {
		case "v":
			p.Vertices = append(
				p.Vertices,
				NewPoint(
					StringToFloat(commandAndParams[1]),
					StringToFloat(commandAndParams[2]),
					StringToFloat(commandAndParams[3]),
				))
		case "f":
			for i := 2; i < len(commandAndParams)-1; i++ {
				ll, _ := strconv.Atoi(strings.Split(commandAndParams[1], "/")[0])
				P1 := p.Vertices[ll]
				ll, _ = strconv.Atoi(strings.Split(commandAndParams[i], "/")[0])
				P2 := p.Vertices[ll]
				ll, _ = strconv.Atoi(strings.Split(commandAndParams[i+1], "/")[0])
				P3 := p.Vertices[ll]

				t := NewTriangle(P1, P2, P3)
				p.Groups[p.ActiveGroup].AddTriangle(t)
			}
		case "g":
			p.ActiveGroup = commandAndParams[1]
			p.Groups[p.ActiveGroup] = NewGroup()
		default:
			p.LinesSkipped++
		}
	}
}

func (p *Parser) ToGroup() *Group {
	g := NewGroup()
	for _, g2 := range p.Groups {
		g.Shapes = append(g.Shapes, g2.Shapes...)
	}
	return g
}

func StringToFloat(incomingString string) float64 {
	var newNom, newDom float64
	newFraction := strings.Split(incomingString, "/")
	newNom = SpecificStringCases(newFraction[0])

	if len(newFraction) > 1 {
		newDom = SpecificStringCases(newFraction[1])
	} else {
		newDom = 1
	}
	return newNom / newDom
}

func SpecificStringCases(strInt string) float64 {
	strInt = strings.Trim(strInt, " ,")
	if strInt[0:1] == "-" {
		return StringToFloat(strInt[1:]) * -1
	} else if len(strInt) > 1 {
		if strInt[0:2] == "π" {
			return math.Pi
		} else if len(strInt) > 2 && strInt[1:2] == "π" {
			return StringToFloat(strInt[0:1]) * math.Pi
		}
		if len(strInt) > 2 && strInt[0:3] == "√" {
			return math.Sqrt(StringToFloat(strInt[3:]))
		}
	}
	var err error
	var newNumber float64
	newNumber, err = strconv.ParseFloat(strInt, 64)
	if err != nil {
		log.Fatalf("to heck with it [%s] [%s] |%s|", strInt, strInt[0:3], err)
	}

	return newNumber
}
