package tools

/*
	Vector stores 3 float32 values to be used in 3D coordinate operations. These values are mapped to X, Y, Z
 */
type Vector struct {
	X, Y, Z float32
}

/*
	Equals checks if the current Vector and another vector are the same
 */
func (v *Vector) Equals(h *Vector) bool {
	if h == nil {
		return false
	}
	if v.X == h.X && v.Y == h.Y && v.Z == h.Z {
		return true
	}
	return false
}

/*
	ToArray converts the float32 list starting with X and ending with Z
 */
func (v *Vector) ToArray() []float32 {
	return []float32{v.X, v.Y, v.Z}
}
