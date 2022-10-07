# Brandfetch

[![GoDoc](https://godoc.org/github.com/rvflash/brandfetch?status.svg)](https://godoc.org/github.com/rvflash/brandfetch)
[![Build Status](https://github.com/rvflash/brandfetch/workflows/build/badge.svg)](https://github.com/rvflash/brandfetch/actions?workflow=build)
[![Code Coverage](https://codecov.io/gh/rvflash/brandfetch/branch/master/graph/badge.svg)](https://codecov.io/gh/rvflash/brandfetch)
[![Go Report Card](https://goreportcard.com/badge/github.com/rvflash/brandfetch?)](https://goreportcard.com/report/github.com/rvflash/brandfetch)

Unofficial Golang interface for the [Brandfetch](https://brandfetch.com/) API.

This package starts with the v2 tag to follow the API, see the module name. 

### Installation

```bash
$ go get -u github.com/rvflash/brandfetch/v2
```

Requirement: this package uses the `url.JoinPath` function incoming with Golang 1.19.

### Usage

```go
    import "github.com/rvflash/brandfetch/v2"
    // ...
    res, err := brandfetch.BrandByName(context.Background(), "example")
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(res.Domain)
	// Output: example.com
```

See API tests for more examples. 