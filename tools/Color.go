package tools


type Texture struct {
	R float32 `json:"r"`

	G float32 `json:"g"`

	B float32 `json:"b"`

	A float32 `json:"a"`

	Emission float32 `json:"emission"`

	Roughness float32 `json:"roughness"`

	Name string `json:"id"`
}
