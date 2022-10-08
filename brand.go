// Copyright (c) 2022 Hervé Gouchet. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package brandfetch

// Brand is a brand with basic assets.
type Brand struct {
	Name   *string `json:"name,omitempty"`
	Domain string  `json:"domain"`
	Icon   *string `json:"icon,omitempty"`
}

// String returns the "name" of the brand.
// It implements the fmt.stringer interface.
func (b Brand) String() string {
	if b.Name != nil {
		return *b.Name
	}
	return b.Domain
}

// URL returns the website URL.
func (b Brand) URL() string {
	if b.Domain == "" {
		return ""
	}
	return "https://" + b.Domain
}
