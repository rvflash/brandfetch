// Copyright (c) 2022 Hervé Gouchet. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package brandfetch_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/matryer/is"
	"github.com/rvflash/brandfetch/v2"
)

func TestConnect(t *testing.T) {
	t.Parallel()
	var (
		are = is.New(t)
		dt  = map[string]struct {
			// inputs
			opts []brandfetch.Configurator
			// outputs
			err error
		}{
			"Default": {},
			"Invalid HTTP Client": {
				opts: []brandfetch.Configurator{brandfetch.SetHTTPClient(nil)},
				err:  brandfetch.ErrHTTPClient,
			},
		}
	)
	for name, tc := range dt {
		tt := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			c, err := brandfetch.Connect(tt.opts...)
			are.True(errors.Is(err, tt.err)) // mismatch error
			are.Equal(c != nil, err == nil)  // wrong return
		})
	}
}

func TestClient_BrandByName(t *testing.T) {
	t.Parallel()
	var (
		are = is.New(t)
		dt  = map[string]struct {
			// inputs
			client brandfetch.HTTPClient
			// outputs
			res *brandfetch.Brand
			err error
		}{
			"Default": {client: &mockClient{}, err: brandfetch.ErrResponse},
			"No results": {
				client: &mockClient{
					filename: "not_found.json",
					reqURL:   "https://api.brandfetch.io/v2/search/example",
				},
				err: brandfetch.ErrNoResults,
			},
			"OK": {
				client: &mockClient{
					filename: "ok.json",
					reqURL:   "https://api.brandfetch.io/v2/search/example",
				},
				res: &brandfetch.Brand{
					Name:   newString("Example"),
					Domain: "example.com",
				},
			},
		}
	)
	for name, tc := range dt {
		tt := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			c, err := brandfetch.Connect(brandfetch.SetHTTPClient(tt.client))
			are.NoErr(err) // unexpected connect error
			b, err := c.BrandByName(context.Background(), "example")
			are.True(errors.Is(err, tt.err)) // mismatch error
			are.Equal(tt.res, b)             // mismatch result
		})
	}
}

func TestClient_BrandsByName(t *testing.T) {
	t.Parallel()
	var (
		are = is.New(t)
		dt  = map[string]struct {
			// inputs
			client brandfetch.HTTPClient
			ctx    context.Context
			name   string
			// outputs
			res []brandfetch.Brand
			err error
			msg string
		}{
			"Default": {err: context.Canceled, client: &mockClient{}},
			"Missing name": {
				ctx:    context.Background(),
				name:   "  ",
				client: &mockClient{},
				err:    brandfetch.ErrRequest,
				msg:    "missing query",
			},
			"Spaced name": {
				ctx:  context.Background(),
				name: " example ",
				client: &mockClient{
					filename: "ok.json",
					reqURL:   "https://api.brandfetch.io/v2/search/example",
				},
				res: []brandfetch.Brand{
					{
						Name:   newString("Example"),
						Domain: "example.com",
					},
					{
						Name:   newString("Examples"),
						Domain: "examples.com",
						Icon:   newString("https://asset.brandfetch.io/idnt7TCUjo/idREMYWRWm_s.webp"),
					},
					{
						Domain: "example-sentences.com",
						Icon:   newString("https://asset.brandfetch.io/idTipDPkyD/idZpkiw4k1_s.webp"),
					},
				},
			},
		}
	)
	for name, tc := range dt {
		tt := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			c, err := brandfetch.Connect(brandfetch.SetHTTPClient(tt.client))
			are.NoErr(err) // unexpected connect error
			res, err := c.BrandsByName(tt.ctx, tt.name)
			are.True(errors.Is(err, tt.err))                              // mismatch error
			are.True(err == nil || strings.Contains(err.Error(), tt.msg)) // mismatch error message
			are.Equal("", cmp.Diff(tt.res, res))                          // mismatch result
		})
	}
}

func newString(s string) *string {
	s2 := s
	return &s2
}

type mockClient struct {
	filename string
	reqURL   string
	err      error
}

func (m *mockClient) Do(req *http.Request) (*http.Response, error) {
	if m.err != nil {
		return nil, m.err
	}
	if m.reqURL != req.URL.String() {
		return nil, fmt.Errorf("unexpected request URL: exp: %q, got: %q", m.reqURL, req.URL)
	}
	b, err := os.Open(filepath.Join("testdata", m.filename))
	if err != nil {
		return nil, err
	}
	resp := &http.Response{
		Status:     http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
		Body:       b,
		Request:    req,
	}
	return resp, nil
}
