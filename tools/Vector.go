package tools

type Vector struct {
	X, Y, Z float64
}

func (v *Vector) Equals(h *Vector) bool {
	if h == nil {
		return false
	}
	if v.X == h.X && v.Y == h.Y && v.Z == h.Z {
		return true
	}
	return false
}
