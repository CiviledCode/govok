package govok

import (
	"github.com/civiledcode/govok/object"
	"github.com/civiledcode/govok/shape"
	"github.com/civiledcode/govok/vox"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"github.com/df-mc/schematic"
	"os"
	"strconv"
)

type Govok struct {
	File, TextureFile string

	Format FileFormat
}

type FileFormat int

const (
	SCHEMATIC FileFormat = 1
	VOX       FileFormat = 2
)

func (g *Govok) ConvertAndWrite(outputFile string) (bool, error) {
	planes, err := g.convert()

	if err != nil {
		return false, err
	}

	file, _ := os.Create(outputFile)
	write(planes, file)

	return true, nil
}

func (g *Govok) ConvertAndLoad() object.Content {
	return object.Content{}
}

func (g *Govok) convert() ([]*shape.Plane, error) {
	switch g.Format {
	case VOX:
		tr, _ := vox.LoadVOXAndTriangulate(g.File)

		var planes []*shape.Plane
		for i := 0; i < len(tr); i += 2 {
			plane := shape.Plane{TriangleOne: tr[i], TriangleTwo: tr[i+1], Color: &tr[i].Color}
			planes = append(planes, &plane)
		}

		return planes, nil
	case SCHEMATIC:
		file, _ := os.Open(g.File)
		s, _ := schematic.FromReader(file)
		dim := s.Dimensions()
		file, _ = os.Open(g.TextureFile)
		textureMap := object.LoadColorMap(file)
		var voxels []vox.Voxel

		for x := 0; x < dim[0]; x++ {
			for y := 0; y < dim[1]; y++ {
				for z := 0; z < dim[2]; z++ {
					block := s.At(x, y, z, func(int, int, int) world.Block { return nil })

					if block != nil {
						name, _ := block.EncodeBlock()
						if name != "minecraft:air" {
							voxels = append(voxels, vox.Voxel{X: x, Y: y, Z: z, Color: textureMap[name]})
						}
					}
				}
			}
		}
		tr, _ := vox.TriangulateVoxels(voxels)
		var planes []*shape.Plane
		for i := 0; i < len(tr); i += 2 {
			plane := shape.Plane{TriangleOne: tr[i], TriangleTwo: tr[i+1], Color: &tr[i].Color}
			planes = append(planes, &plane)
		}

		return planes, nil
	}

	return nil, nil
}

func (g *Govok) LoadFile() string {
	return ""
}

func (g *Govok) LoadContent() object.Content {
	return object.Content{}
}

func write(planes []*shape.Plane, file *os.File) error {
	var li []float32
	for _, plane := range planes {
		li = plane.ConvertToArray()
		var err error
		for _, ind := range li {
			_, err = file.WriteString(strconv.FormatFloat(float64(ind), 'f', 1, 32) + ",")

			if err != nil {
				return err
			}
		}
		_, err = file.WriteString(plane.Color.Name + "\n")

		if err != nil {
			return err
		}
	}

	return nil
}
