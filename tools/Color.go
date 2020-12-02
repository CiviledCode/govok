package tools

type Color struct {
	// TODO: Add reflection, roughness, and diffusion
	R, G, B, A float32
}

type Texture struct {
	R uint8 `json:"r"`

	G uint8 `json:"g"`

	B uint8 `json:"b"`

	A uint8 `json:"a"`

	Emission uint8 `json:"emission"`

	Roughness uint8 `json:"roughness"`

	Name string `json:"id"`
}
