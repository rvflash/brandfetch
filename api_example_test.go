package brandfetch_test

import (
	"context"
	"fmt"
	"log"

	"github.com/rvflash/brandfetch/v2"
)

func ExampleBrandByName() {
	res, err := brandfetch.BrandByName(context.Background(), "example")
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(res.Domain)
	// Output: example.com
}

func ExampleBrandsByName() {
	res, err := brandfetch.BrandsByName(context.Background(), "example")
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(len(res))
	// Output: 3
}
