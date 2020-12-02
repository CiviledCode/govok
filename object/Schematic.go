package object

import (
	"bufio"
	"encoding/json"
	"github.com/civiledcode/govok/tools"
	"io"
)

/**
	LoadColorMap loads all of the texture data of a specific schematic using json.Unmarshal.
	The name is always the namespace ID of the block we are mapping the texture to

	file: The io.Reader pointing to the location of the file
 */
func LoadColorMap(file io.Reader) map[string]tools.Texture {
	var cMap map[string]tools.Texture

	// Open our scanner
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Go through all of the lines of the file and map the unmarshalled line to block ids in the map
	var texture tools.Texture
	for scanner.Scan() {
		json.Unmarshal([]byte(scanner.Text()), texture)
		cMap[texture.Name] = texture
	}

	return cMap
}
