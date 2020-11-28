package govok

import (
	"github.com/civiledcode/govok/object"
	"github.com/civiledcode/govok/shape"
	"github.com/civiledcode/govok/vox"
)

//TODO: Turn this into a struct instead

type Govok struct {
	File string
}

func ConvertAndWrite(file string, outputFile string) (bool, error) {
	_, err := convert(file)

	if err != nil {
		return false, err
	}

	return true, nil
}

func ConvertFileToContent(file string) object.Content {
	return object.Content{}
}

func convert(file string) (string, error) {
	tr, _ := vox.LoadVOXAndTriangulate(file)

	for i := 0; i < len(tr); i += 2 {
		plane := shape.Plane{TriangleOne: tr[i], TriangleTwo: tr[i+1], Color: &tr[i].Color}

		plane.ConvertToArray()
	}
	return "", nil
}

func (g *Govok) LoadFile() string {
	return ""
}

func (g *Govok) LoadContent() object.Content {
	return object.Content{}
}
