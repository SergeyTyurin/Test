package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

import (
	"fmt"
	"strings"
)

type BuildingRoad interface {
	build(route *Route)
}

type Bycicle struct {
}

func (b *Bycicle) build(route *Route) {
	fmt.Println("Bike route built:")
	for _, el := range route.pointList {
		fmt.Println(el.name)
	}
}

type Car struct {
}

func (c *Car) build(route *Route) {
	fmt.Println("Car route built:")
	for _, el := range route.pointList {
		fmt.Println(el.name)
	}
}

type OnFoot struct {
}

func (o *OnFoot) build(route *Route) {
	fmt.Println("Walking route built:")
	for _, el := range route.pointList {
		fmt.Println(el.name)
	}
}

type Point struct {
	x    int
	y    int
	name string
}

func initPoint(x int, y int, name string) *Point {
	return &Point{
		x:    x,
		y:    y,
		name: name,
	}
}

type Route struct {
	pointList []Point
	road      BuildingRoad
}

func (r *Route) addPoint(p *Point) {
	r.pointList = append(r.pointList, *p)
}

func (r *Route) delPoint(p *Point) {
	var start bool
	for i, el := range r.pointList {
		if strings.EqualFold(el.name, p.name) {
			start = true
		}
		if start {
			r.pointList = append(r.pointList[:i], r.pointList[i+1:]...)
			break
		}
	}
	if !start {
		fmt.Println("Point not found")
	}
}

func (r *Route) setRoad(m BuildingRoad) {
	r.road = m
}

func (r *Route) buildingRoute() {
	r.road.build(r)
}

// func main() {
// 	p1 := Point{name: "A", x: 2, y: 4}
// 	p2 := Point{name: "B", x: 5, y: 13}
// 	p3 := Point{name: "C", x: 8, y: 28}

// 	route := Route{pointList: []Point{p1, p2, p3}}

// 	road1 := Bycicle{}

// 	route.setRoad(&road1)
// 	route.buildingRoute()

// 	route.delPoint(&p2)
// 	road2 := Car{}
// 	route.setRoad(&road2)
// 	route.buildingRoute()

// 	p4 := Point{name: "D", x: 10, y: 20}
// 	route.addPoint(&p4)
// 	road3 := OnFoot{}
// 	route.setRoad(&road3)
// 	route.buildingRoute()
// }
