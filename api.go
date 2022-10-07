// Package brandfetch provides methods to get brand assets by names (logo, domain, etc.)
package brandfetch

import "context"

// BrandByName returns, if exists in the Brandfetch database, the brand behind this name.
func BrandByName(ctx context.Context, name string) (*Brand, error) {
	cli, err := Connect()
	if err != nil {
		return nil, err
	}
	return cli.BrandByName(ctx, name)
}

// BrandsByName returns the list of brands matching this name in the Brandfetch database.
func BrandsByName(ctx context.Context, name string) ([]Brand, error) {
	cli, err := Connect()
	if err != nil {
		return nil, err
	}
	return cli.BrandsByName(ctx, name)
}
