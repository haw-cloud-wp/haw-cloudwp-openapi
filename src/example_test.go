package main

import (
	"log"
	"testing"
)

type Cubicle interface {
	getVolume() int
	getWidth() int
	getHeight() int
	getDepth() int
}

type Cube struct {
	sideLength int
}

func (c Cube) getWidth() int {
	return c.sideLength
}

func (c Cube) getHeight() int {
	return c.sideLength
}

func (c Cube) getDepth() int {
	return c.sideLength
}

func (c Cube) getVolume() int {
	return c.sideLength * c.sideLength *c.sideLength
}


type RectCube struct {
	width int
	height int
	depth int
}

func (r RectCube) getVolume() int {
	return r.width*r.height*r.depth
}

func (r RectCube) getWidth() int {
	return r.width
}

func (r RectCube) getHeight() int {
	return r.height
}

func (r RectCube) getDepth() int {
	return r.depth
}


func TestCubeicle(t *testing.T) {
	cubearr := make([]Cubicle, 10)
	for i := 0; i <5; i++ {
		cubearr[i] = RectCube{
			width:   10,
			height:  20,
			depth:   25,
		}
	}
	for i := 5; i <10; i++ {
		cubearr[i] = Cube{sideLength: 20}
	}

	for _, cubicle := range cubearr {
		log.Printf("Volume is: %d", cubicle.getVolume())
	}
}




