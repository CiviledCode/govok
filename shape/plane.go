package shape

import (
	"github.com/civiledcode/MagicaToOBJ/shape"
	"github.com/civiledcode/MagicaToOBJ/tools"
)

type Plane struct {
	TriangleOne, TriangleTwo *shape.Triangle

	Color tools.Color
}

func (p *Plane) ConvertToArray() [12]float64 {
	var array = [12]float64{}
	common, unique := toPointArray(p.TriangleOne, p.TriangleTwo)
	array[0] = unique[0].X
	array[1] = unique[0].Y
	array[2] = unique[0].Z

	array[3] = common[0].X
	array[4] = common[0].Y
	array[5] = common[0].Z

	array[6] = unique[1].X
	array[7] = unique[1].Y
	array[8] = unique[1].Z

	array[9] = common[1].X
	array[10] = common[1].Y
	array[11] = common[1].Z
	return array
}

//TODO: Optimize the fuck outta this
func toPointArray(t1 *shape.Triangle, t2 *shape.Triangle) ([]*tools.Vector, []*tools.Vector) {
	array := t1.ToStream()
	for _, item := range t2.ToStream() {
		array = append(array, item)
	}
	var common, unique []*tools.Vector

	for outside := 0; outside < len(array); outside++ {
		if array[outside] == nil {
			continue
		}
		for inside := outside + 1; inside < len(array); inside++ {
			if array[outside].Equals(array[inside]) {
				common = append(common, array[outside])
				array[inside] = nil
				array[outside] = nil
				break
			}
		}
	}

	for _, val := range array {
		if val != nil {
			unique = append(unique, val)
		}
	}

	return common, unique
}
