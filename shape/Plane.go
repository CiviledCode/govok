package shape

import (
	"github.com/civiledcode/govok/tools"
)

type Plane struct {
	TriangleOne, TriangleTwo *Triangle

	Color *tools.Color
}

const (
	UNIQUE     = 0
	COMMON     = 1
	COMMON_NIL = -1
)

func (p *Plane) ConvertToArray() []float32 {
	var array []float32

	values, mapp := toPointArray(p.TriangleOne, p.TriangleTwo)

	var needed uint8

	for i := 0; i < len(mapp); i++ {
		switch needed {
		case 0, 2:
			if mapp[i] == UNIQUE {
				if needed == 2 {
					needed = 0
					continue
				}
				array = append(array, values[i].ToArray()...)
			}
		case 1, 3:
			if mapp[i] == COMMON {
				if needed == 3 {
					needed = 1
					continue
				}
				array = append(array, values[i].ToArray()...)
			}
		}
	}

	return array
}

func toPointArray(t1 *Triangle, t2 *Triangle) ([]*tools.Vector, [6]int) {
	array := t1.ToStream()
	for _, item := range t2.ToStream() {
		array = append(array, item)
	}

	var list [6]int
	for outside := 0; outside < len(array); outside++ {
		if list[outside] == COMMON_NIL {
			continue
		}

		for inside := outside + 1; inside < len(array); inside++ {
			if array[outside].Equals(array[inside]) {
				list[outside] = COMMON
				list[inside] = COMMON_NIL
				break
			}
		}
	}

	return array, list
}
