// Copyright (c) 2022 Hervé Gouchet. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package brandfetch_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/rvflash/brandfetch/v2"
)

func TestWarn_Error(t *testing.T) {
	t.Parallel()
	var (
		are = is.New(t)
		dt  = []struct {
			exp string
			got error
		}{
			{exp: "brandfetch: invalid http client", got: brandfetch.ErrHTTPClient},
			{exp: "brandfetch: no results for this query", got: brandfetch.ErrNoResults},
			{exp: "brandfetch: bad request body", got: brandfetch.ErrRequest},
			{exp: "brandfetch: unsupported response", got: brandfetch.ErrResponse},
		}
	)
	for _, tt := range dt {
		are.Equal(tt.exp, tt.got.Error()) // mismatch error message
	}
}
