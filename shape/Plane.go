package shape

import (
	"github.com/civiledcode/govok/tools"
)

/*
	Plane is a 4 point shape made up of 4 vectors. We can default to 4 points only because we are only handling regular rectangles, not irregular shapes with non-90 degree angles
 */

type Plane struct {
	X, Y, Z, W tools.Vector

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

	array[0] = p.X.X
	array[1] = p.X.Y
	array[2] = p.X.Z

	array[3] = p.Y.X
	array[4] = p.Y.Y
	array[5] = p.Y.Z

	array[6] = p.Z.X
	array[7] = p.Z.Y
	array[8] = p.Z.Z

	array[9] = p.W.X
	array[10] = p.W.Y
	array[11] = p.W.Z

	return array
}