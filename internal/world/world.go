package world

import (
	"SoB/internal/common"
	"SoB/internal/entity/entity"
	"math"
)

type Map struct {
	spaces map[common.Coordinates]*Space // The coordinates of the Space's on the map tile
}

func (d *Map) GetSpace(coordinates common.Coordinates) *Space {
	return d.spaces[coordinates]
}

func (d *Map) SetSpace(coordinates common.Coordinates, space Space) {
	d.spaces[coordinates] = &space
}

func NewMap() *Map {
	return &Map{spaces: make(map[common.Coordinates]*Space)}
}

type Space struct {
	wall     bool
	occupant *entity.Entity // Any entity occupying the space
}

func (d *Map) HasLineOfSight(from, to common.Coordinates) bool {
	points := getLinePoints(from, to)

	for i := 1; i < len(points)-1; i++ { // Skip start and end points
		if space := d.GetSpace(points[i]); space != nil && space.wall {
			return false
		}
	}
	return true
}

func getLinePoints(from, to common.Coordinates) []common.Coordinates {
	var points []common.Coordinates

	x0, y0 := from.X, from.Y
	x1, y1 := to.X, to.Y

	dx := math.Abs(float64(x1) - float64(x0))
	dy := math.Abs(float64(y1) - float64(y0))

	var sx, sy int
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}
	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}

	err := dx - dy
	x, y := x0, y0

	for {
		points = append(points, common.Coordinates{X: x, Y: y})

		if x == x1 && y == y1 {
			break
		}

		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x += sx
		}
		if e2 < dx {
			err += dx
			y += sy
		}
	}

	return points
}
