package golang_united_school_homework

import (
	"errors"
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes   []Shape
	capacity int // Maximum quantity of shapes that can be inside the box.
}

var (
	errorBoxIsFull         = errors.New("box is full")
	errorIndexOutOfRange   = errors.New("index out of range")
	errorShapeDoesNotExist = errors.New("shape does not exist")
	errorNoCircles         = errors.New("there is no circles in the box")
)

// NewBox creates new instance of box
func NewBox(capacity int) *box {
	return &box{
		shapes:   make([]Shape, 0, capacity),
		capacity: capacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the capacity range.
func (b *box) AddShape(shape Shape) error {
	if b.capacity < 1 {
		return fmt.Errorf("AddShape: adding shape %v to the box: %w", shape, errorBoxIsFull)
	}
	b.shapes = append(b.shapes, shape)
	b.capacity--
	return nil
}

// checkIndex checks whether the shape by index doesn't exist in the box
// or index went out of range, and returns an appropriate error.
func (b box) checkIndex(funcName string, i int) error {
	var action string
	switch funcName {
	case "GetByIndex":
		action = "getting"
	case "ExtractByIndex":
		action = "extracting"
	case "ReplaceByIndex":
		action = "replacing"
	}
	if l := len(b.shapes); i >= l || i < 0 {
		return fmt.Errorf("%s: %s shape by index %d from the box: %w", funcName, action, i, errorIndexOutOfRange)
	}
	if b.shapes[i] == nil {
		return fmt.Errorf("%s: %s shape by index %d from the box: %w", funcName, action, i, errorShapeDoesNotExist)
	}
	return nil
}

// GetByIndex allows getting shape by index.
func (b *box) GetByIndex(i int) (Shape, error) {
	if err := b.checkIndex("GetByIndex", i); err != nil {
		return nil, err
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if err := b.checkIndex("ExtractByIndex", i); err != nil {
		return nil, err
	}
	shape := b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if err := b.checkIndex("ReplaceByIndex", i); err != nil {
		return nil, err
	}
	toRemove := b.shapes[i]
	b.shapes[i] = shape
	return toRemove, nil
}

// SumPerimeter provides perimeter sum of all shapes in the list.
func (b *box) SumPerimeter() (sum float64) {
	for _, shape := range b.shapes {
		sum += shape.CalcPerimeter()
	}
	return
}

// SumArea provides area sum of all shapes in the list.
func (b *box) SumArea() (sum float64) {
	for _, shape := range b.shapes {
		sum += shape.CalcArea()
	}
	return
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var circles bool
	for i := 0; i < len(b.shapes); i++ {
		if _, ok := b.shapes[i].(*Circle); ok {
			circles, b.shapes = true, append(b.shapes[:i], b.shapes[i+1:]...)
			i--
		}
	}
	if circles {
		return nil
	}
	return fmt.Errorf("RemoveAllCircles: removing all circles from the box: %w", errorNoCircles)
}
