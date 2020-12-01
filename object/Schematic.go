package object

import (
	"bufio"
	"encoding/json"
	"github.com/civiledcode/govok/tools"
	"io"
)

func LoadColorMap(file io.Reader) map[string]tools.Texture {
	var cMap map[string]tools.Texture

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	var texture tools.Texture

	for _, each_ln := range text {
		json.Unmarshal([]byte(each_ln), texture)
		cMap[texture.Name] = texture
	}

	return cMap
}
