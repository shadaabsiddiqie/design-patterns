// creates a object of a single interface type
package factory

import "fmt"

type shape interface {
	draw()
}
type shapeCircle struct {
	radius int
}
type shapeSqure struct {
	side int
}

func (m shapeCircle) draw() {
	fmt.Println("drawing a circle with radius", m.radius)
}
func (m shapeSqure) draw() {
	fmt.Println("drawing a squre with side", m.side)
}

type shapeFactory interface {
	createShape(info int) shape
}

type circleFactory struct{}
type squreFactory struct{}

func (cf circleFactory) createShape(info int) shape {
	return shapeCircle{radius: info}
}

func (sf squreFactory) createShape(info int) shape {
	return shapeSqure{side: info}
}

func shapeFactoryRegistery(needShape string) shapeFactory {
	if needShape == "circle" {
		return circleFactory{}
	}
	if needShape == "squre" {
		return squreFactory{}
	}
	return nil
}

func FactoryMethodRun() {
	cf := shapeFactoryRegistery("circle")
	sf := shapeFactoryRegistery("squre")
	c := cf.createShape(5)
	s := sf.createShape(8)
	c.draw()
	s.draw()
}
