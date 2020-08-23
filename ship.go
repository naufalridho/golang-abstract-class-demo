package golang_abstract_class_demo

import (
	"fmt"
)

type Vehicle interface {
	move()
	dock()
}

type abstractPiece struct{
	name string
	maxSpeed int
	yearBuilt int
	capacity int

	move func()
}
func newAbstractPiece() *abstractPiece {
	return &abstractPiece{}
}
func (p *abstractPiece) dock() {
	fmt.Println("dock")
}
