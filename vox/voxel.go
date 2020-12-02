package vox

import (
	"github.com/civiledcode/govok/shape"
	"github.com/civiledcode/govok/tools"
)

// This code was forked from https://github.com/fogleman/fauxgl

type Voxel struct {
	X, Y, Z int
	Color   tools.Texture
}

type voxelAxis int

const (
	_ voxelAxis = iota
	voxelX
	voxelY
	voxelZ
)

type voxelNormal struct {
	Axis voxelAxis
	Sign int
}

var (
	voxelPosX = voxelNormal{voxelX, 1}
	voxelNegX = voxelNormal{voxelX, -1}
	voxelPosY = voxelNormal{voxelY, 1}
	voxelNegY = voxelNormal{voxelY, -1}
	voxelPosZ = voxelNormal{voxelZ, 1}
	voxelNegZ = voxelNormal{voxelZ, -1}
)

type voxelPlane struct {
	Normal   voxelNormal
	Position int
	Color    tools.Texture
}

type voxelFace struct {
	I0, J0 int
	I1, J1 int
}

func TriangulateVoxels(voxels []Voxel) []*shape.Triangle {
	type key struct {
		X, Y, Z int
	}

	// create lookup table
	lookup := make(map[key]bool)
	for _, v := range voxels {
		lookup[key{v.X, v.Y, v.Z}] = true
	}

	// find exposed faces
	planeFaces := make(map[voxelPlane][]voxelFace)
	for _, v := range voxels {
		if !lookup[key{v.X + 1, v.Y, v.Z}] {
			plane := voxelPlane{voxelPosX, v.X, v.Color}
			face := voxelFace{v.Y, v.Z, v.Y, v.Z}
			planeFaces[plane] = append(planeFaces[plane], face)
		}
		if !lookup[key{v.X - 1, v.Y, v.Z}] {
			plane := voxelPlane{voxelNegX, v.X, v.Color}
			face := voxelFace{v.Y, v.Z, v.Y, v.Z}
			planeFaces[plane] = append(planeFaces[plane], face)
		}
		if !lookup[key{v.X, v.Y + 1, v.Z}] {
			plane := voxelPlane{voxelPosY, v.Y, v.Color}
			face := voxelFace{v.X, v.Z, v.X, v.Z}
			planeFaces[plane] = append(planeFaces[plane], face)
		}
		if !lookup[key{v.X, v.Y - 1, v.Z}] {
			plane := voxelPlane{voxelNegY, v.Y, v.Color}
			face := voxelFace{v.X, v.Z, v.X, v.Z}
			planeFaces[plane] = append(planeFaces[plane], face)
		}
		if !lookup[key{v.X, v.Y, v.Z + 1}] {
			plane := voxelPlane{voxelPosZ, v.Z, v.Color}
			face := voxelFace{v.X, v.Y, v.X, v.Y}
			planeFaces[plane] = append(planeFaces[plane], face)
		}
		if !lookup[key{v.X, v.Y, v.Z - 1}] {
			plane := voxelPlane{voxelNegZ, v.Z, v.Color}
			face := voxelFace{v.X, v.Y, v.X, v.Y}
			planeFaces[plane] = append(planeFaces[plane], face)
		}
	}

	var triangles []*shape.Triangle

	// find large rectangles, triangulate and outline
	for plane, faces := range planeFaces {
		faces = combineVoxelFaces(faces)
		triangles = append(triangles, triangulateVoxelFaces(plane, faces)...)
	}

	return triangles
}

func combineVoxelFaces(faces []voxelFace) []voxelFace {
	// determine bounding box
	i0 := faces[0].I0
	j0 := faces[0].J0
	i1 := faces[0].I1
	j1 := faces[0].J1
	for _, f := range faces {
		if f.I0 < i0 {
			i0 = f.I0
		}
		if f.J0 < j0 {
			j0 = f.J0
		}
		if f.I1 > i1 {
			i1 = f.I1
		}
		if f.J1 > j1 {
			j1 = f.J1
		}
	}
	// create arrays
	nj := j1 - j0 + 1
	ni := i1 - i0 + 1
	a := make([][]int, nj)
	w := make([][]int, nj)
	h := make([][]int, nj)
	for j := range a {
		a[j] = make([]int, ni)
		w[j] = make([]int, ni)
		h[j] = make([]int, ni)
	}
	// populate array
	count := 0
	for _, f := range faces {
		for j := f.J0; j <= f.J1; j++ {
			for i := f.I0; i <= f.I1; i++ {
				a[j-j0][i-i0] = 1
				count++
			}
		}
	}
	// find rectangles
	var result []voxelFace
	for count > 0 {
		var maxArea int
		var maxFace voxelFace
		for j := 0; j < nj; j++ {
			for i := 0; i < ni; i++ {
				if a[j][i] == 0 {
					continue
				}
				if j == 0 {
					h[j][i] = 1
				} else {
					h[j][i] = h[j-1][i] + 1
				}
				if i == 0 {
					w[j][i] = 1
				} else {
					w[j][i] = w[j][i-1] + 1
				}
				minw := w[j][i]
				for dh := 0; dh < h[j][i]; dh++ {
					if w[j-dh][i] < minw {
						minw = w[j-dh][i]
					}
					area := (dh + 1) * minw
					if area > maxArea {
						maxArea = area
						maxFace = voxelFace{
							i0 + i - minw + 1, j0 + j - dh, i0 + i, j0 + j}
					}
				}
			}
		}
		result = append(result, maxFace)
		for j := maxFace.J0; j <= maxFace.J1; j++ {
			for i := maxFace.I0; i <= maxFace.I1; i++ {
				a[j-j0][i-i0] = 0
				count--
			}
		}
		for j := 0; j < nj; j++ {
			for i := 0; i < ni; i++ {
				w[j][i] = 0
				h[j][i] = 0
			}
		}
	}
	return result
}

func triangulateVoxelFaces(plane voxelPlane, faces []voxelFace) []*shape.Triangle {
	triangles := make([]*shape.Triangle, len(faces)*2)
	k := float32(plane.Position) + float32(plane.Normal.Sign)*0.5
	for i, face := range faces {
		i0 := float32(face.I0) - 0.5
		j0 := float32(face.J0) - 0.5
		i1 := float32(face.I1) + 0.5
		j1 := float32(face.J1) + 0.5
		var p1, p2, p3, p4 tools.Vector
		switch plane.Normal.Axis {
		case voxelX:
			p1 = tools.Vector{k, i0, j0}
			p2 = tools.Vector{k, i1, j0}
			p3 = tools.Vector{k, i1, j1}
			p4 = tools.Vector{k, i0, j1}
		case voxelY:
			p1 = tools.Vector{i0, k, j1}
			p2 = tools.Vector{i1, k, j1}
			p3 = tools.Vector{i1, k, j0}
			p4 = tools.Vector{i0, k, j0}
		case voxelZ:
			p1 = tools.Vector{i0, j0, k}
			p2 = tools.Vector{i1, j0, k}
			p3 = tools.Vector{i1, j1, k}
			p4 = tools.Vector{i0, j1, k}
		}
		if plane.Normal.Sign < 0 {
			p1, p2, p3, p4 = p4, p3, p2, p1
		}
		t1 := shape.Triangle{p1, p2, p3, plane.Color}
		t2 := shape.Triangle{p1, p3, p4, plane.Color}
		triangles[i*2+0] = &t1
		triangles[i*2+1] = &t2
	}
	return triangles
}
