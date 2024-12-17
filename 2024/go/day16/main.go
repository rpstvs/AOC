package main

type Point struct {
	x, y int
}

type Direction uint8

type Node struct {
	Pos       Point
	Direction Direction
}
