package shape

import "github.com/civiledcode/MagicaToOBJ/tools"

type Triangle struct {
	X, Y, Z tools.Vector

	Color tools.Color
}

func (t *Triangle) ToStream() []*tools.Vector {
	var array = []*tools.Vector{&t.X, &t.Y, &t.Z}
	return array
}
