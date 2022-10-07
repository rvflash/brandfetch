package brandfetch

// Brand is a brand with basic assets.
type Brand struct {
	Name   string `json:"name,omitempty"`
	Domain string `json:"domain"`
	Icon   string `json:"icon,omitempty"`
}
