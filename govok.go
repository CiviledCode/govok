package govok

import (
	"github.com/civiledcode/govok/object"
	"github.com/civiledcode/govok/vox"
)

/**
Converts the vox file defined and writes it to a new file with the same name, but different extension

Returns the name of the new file created
*/
func ConvertAndWriteFile(file string) string {
	return "ConvertAndWriteFile(LoadFile(file), file)"
}

/**
ConvertFileToContent takes a vox file and converts it to the Content object
*/
func ConvertFileToContent(file string) object.Content {
	return object.Content{}
}

/**
Convert takes raw vox data and converts it to govok data
*/
func Convert(file string) string {
	tr, _ := vox.LoadVOXAndTriangulate(file)
	for i := 0; i < len(tr); i += 2 {

	}
	return ""
}

/**
LoadFileToContent takes a file and loads the content to a content object
*/
func LoadFileToContent(file string) object.Content {
	content := LoadFile(file)
	return LoadContent(content)
}

/**
LoadFile takes the content of the file and loads it as a string
*/
func LoadFile(file string) string {
	return file
}

/**
LoadContent converts a string into our content object
*/
func LoadContent(content string) object.Content {
	return object.Content{}
}
