// Copyright (c) 2022 Hervé Gouchet. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package brandfetch_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/rvflash/brandfetch/v2"
)

const (
	srvName   = "example"
	srvDomain = "example.com"
)

func TestBrand_String(t *testing.T) {
	t.Parallel()
	var (
		are = is.New(t)
		srv = srvName
		dt  = map[string]struct {
			in  brandfetch.Brand
			out string
		}{
			"Default":     {},
			"Domain only": {in: brandfetch.Brand{Domain: srvDomain}, out: srvDomain},
			"Name only":   {in: brandfetch.Brand{Name: &srv}, out: srvName},
			"Both":        {in: brandfetch.Brand{Name: &srv, Domain: srvDomain}, out: srvName},
		}
	)
	for name, tc := range dt {
		tt := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			are.Equal(tt.out, tt.in.String())
		})
	}
}

func TestBrand_URL(t *testing.T) {
	t.Parallel()
	var (
		are = is.New(t)
		dt  = map[string]struct {
			in  brandfetch.Brand
			out string
		}{
			"Default": {},
			"OK":      {in: brandfetch.Brand{Domain: srvDomain}, out: "https://" + srvDomain},
		}
	)
	for name, tc := range dt {
		tt := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			are.Equal(tt.out, tt.in.URL())
		})
	}
}
