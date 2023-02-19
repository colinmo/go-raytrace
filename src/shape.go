package main

type Shaper interface {
	Equals(t Shaper) bool
	Intersects(r Ray) map[int]Intersection
	LocalIntersects(r Ray) map[int]Intersection
	GetID() int
	SetOrigin(t Tuple)
	GetOrigin() Tuple
	SetTransform(t Matrix)
	GetTransform() Matrix
	NormalAt(t Tuple) Tuple
	LocalNormalAt(t Tuple) Tuple
	GetSavedRay() Ray

	GetMaterial() Material
	SetMaterial(m Material)
	GetType() string
	GetMinimum() float64
	GetMaximum() float64
	SetMinimum(m float64)
	SetMaximum(m float64)
	GetClosed() bool
	SetClosed(b bool)
	GetShapesCount() int
	GetShapes() map[int]Shaper
	AddShape(s *Shaper)
	RemoveShape(s Shaper)
	GetParent() *Group
	SetParent(g *Group)
}

type TestShapeType struct {
	Shaper
	ID        int
	Origin    Tuple
	Radius    float64
	Transform Matrix
	Material  Material
	SavedRay  Ray
	Parent    *Group
}

var BaseTransform = IdentityMatrix()
var BaseOrigin = NewPoint(0, 0, 0)
var BaseMaterial = NewMaterial()

func NewTestShape() *TestShapeType {
	return &TestShapeType{
		Transform: IdentityMatrix(),
		Origin:    NewPoint(0, 0, 0),
		Material:  NewMaterial(),
		Parent:    nil,
	}
}

func (s *TestShapeType) Equals(t Shaper) bool {
	return false
}

func (s *TestShapeType) GetType() string { return "test" }

func (s *TestShapeType) GetSavedRay() Ray {
	return s.SavedRay
}
func (s *TestShapeType) Intersects(r Ray) map[int]Intersection {
	bob := s.GetTransform()
	localRay := r.Transform(bob.Inverse())
	s.SavedRay = localRay
	return s.LocalIntersects(localRay)
}

func Intersect(s Shaper, r Ray) map[int]Intersection {
	sTrans := s.GetTransform()
	localRay := r.Transform(sTrans.Inverse())
	return s.LocalIntersects(localRay)
}

func (s *TestShapeType) LocalIntersects(r Ray) map[int]Intersection {
	return make(map[int]Intersection)
}

func (s *TestShapeType) NormalAt(p Tuple) Tuple {
	return NormalAt(s, p)
}

func NormalAt(s Shaper, p Tuple) Tuple {
	v := s.GetTransform()
	v = v.Inverse()
	localPoint := v.MultiplyTuple(p)
	localNormal := s.LocalNormalAt(localPoint)
	trans := v.Transpose()
	worldNormal := trans.MultiplyTuple(localNormal)
	worldNormal.W = 0
	return worldNormal.Normalize()
}

func (s *TestShapeType) LocalNormalAt(p Tuple) Tuple {
	return NewVector(p.X, p.Y, p.Z)
}

func (s *TestShapeType) GetID() int {
	return 1
}

func (s *TestShapeType) SetOrigin(t Tuple) {
	s.Origin = t
}

func (s *TestShapeType) GetOrigin() Tuple {
	return s.Origin
}

func (s *TestShapeType) SetTransform(t Matrix) {
	s.Transform = t
}
func (s *TestShapeType) GetTransform() Matrix {
	return s.Transform
}

func (s *TestShapeType) GetMaterial() Material {
	return s.Material
}
func (s *TestShapeType) SetMaterial(m Material) {
	s.Material = m
}

func (s *TestShapeType) GetParent() *Group {
	return s.Parent
}

func (s *TestShapeType) SetParent(g *Group) {
	s.Parent = g
}
