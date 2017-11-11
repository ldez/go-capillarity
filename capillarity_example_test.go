package capillarity_test

import (
	"fmt"
	"log"

	"github.com/ldez/go-capillarity"
)

func ExampleNewCapillarity() {
	type MyStruct struct {
		Foo string
		Bar struct {
			One string
			Two string
		}
	}

	myStruct := MyStruct{}

	capil := capillarity.NewCapillarity()
	err := capil.Fill(&myStruct)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", myStruct)

	// output: {Foo:foobar Bar:{One:foobar Two:foobar}}
}

func ExampleNewCapillarity_options() {
	type MyStruct struct {
		Foo string
		Bar struct {
			One string
			Two int
		}
	}

	myStruct := MyStruct{}

	capil := capillarity.NewCapillarity(capillarity.WithDefaultString("go"), capillarity.WithDefaultNumber(6))
	err := capil.Fill(&myStruct)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", myStruct)

	// output: {Foo:go Bar:{One:go Two:6}}
}
