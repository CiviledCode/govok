package shape

import (
	"github.com/civiledcode/govok/tools"
)

/*
	Plane is a 4 point shape made up of two triangles. We can default that all values are planes as there is no irregular shapes as everything is made of rectangles
 */
type Plane struct {
	TriangleOne, TriangleTwo *Triangle

	Color *tools.Texture
}

/*
	ConvertToArray converts the two triangles into a point list with the
	[t1u.x, t1u.y, t1u.z, c1.x, c1.y, c1.z, t2u.x, t2u.y, t2u.z, c2.x, c2.y, c2.z] format

	t1u represents triangle ones unique value
	c1 represents the first common points value
	t2u represents triangle twos unique value
	c2 represents the second common points value
 */
func (p *Plane) ConvertToArray() []float32 {
	array := make([]float32, 12)

	values := toPointArray(p.TriangleOne, p.TriangleTwo)

	array[0] = values[0].X
	array[1] = values[0].Y
	array[2] = values[0].Z

	array[3] = values[1].X
	array[4] = values[1].Y
	array[5] = values[1].Z

	array[6] = values[2].X
	array[7] = values[2].Y
	array[8] = values[2].Z

	array[9] = values[3].X
	array[10] = values[3].Y
	array[11] = values[3].Z

	return array
}

/*
	toPointArray takes two triangles and finds the common and unique values and returns them in a list as follows: [t1 unique, common, t2 unique, common]
 */
func toPointArray(t1 *Triangle, t2 *Triangle) []*tools.Vector {
	vals := make([]*tools.Vector, 4)
	var value *tools.Vector

	// Find the common values between the triangles and the unique value of triangle one and add it to the list
	for _, value = range t1.ToArray() {

		if value.Equals(&t2.X) || value.Equals(&t2.Y) || value.Equals(&t2.Z) {
			if vals[1] == nil {
				vals[1] = value
			} else {
				vals[3] = value
			}
		} else {
			if vals[0] == nil {
				vals[0] = value
			} else {
				vals[2] = value
			}
		}
	}

	// Find the unique value of triangle two
	for _, value = range t2.ToArray() {

		if !value.Equals(&t1.X) || !value.Equals(&t1.Y) || !value.Equals(&t1.Z) {
			vals[2] = value
		}
	}
	return vals
}
