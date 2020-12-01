package shape

import "github.com/civiledcode/govok/tools"

type Triangle struct {
	X, Y, Z tools.Vector

	Color tools.Texture
}

func (t *Triangle) ToStream() []*tools.Vector {
	var array = []*tools.Vector{&t.X, &t.Y, &t.Z}
	return array
}
