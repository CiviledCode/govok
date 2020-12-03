package shape

import "github.com/civiledcode/govok/tools"

/*
	Triangle is a 3 tools.Vector shape mapped with X, Y, Z
 */
type Triangle struct {
	X, Y, Z tools.Vector

	Color tools.Texture
}

/*
	ToArray takes the vectors inside of the Triangle and returns them in the X, Y, Z order
 */
func (t *Triangle) ToArray() []*tools.Vector {
	return []*tools.Vector{&t.X, &t.Y, &t.Z}
}