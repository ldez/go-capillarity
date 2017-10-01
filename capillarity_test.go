package capillarity

import (
	"reflect"
	"testing"
)

type myString string

func TestNewCapillarity(t *testing.T) {

	capillarity := NewCapillarity()

	testCase := []struct {
		name     string
		base     interface{}
		expected interface{}
	}{
		{
			name: "string",
			base: &struct {
				Bar string
			}{},
			expected: &struct {
				Bar string
			}{
				Bar: "foobar",
			},
		},
		{
			name: "int",
			base: &struct {
				Bar int
			}{},
			expected: &struct {
				Bar int
			}{
				Bar: 42,
			},
		},
		{
			name: "bool",
			base: &struct {
				Bar bool
			}{},
			expected: &struct {
				Bar bool
			}{
				Bar: true,
			},
		},
		{
			name: "struct",
			base: &struct {
				Bar struct {
					One string
					Two string
				}
			}{},
			expected: &struct {
				Bar struct {
					One string
					Two string
				}
			}{
				Bar: struct {
					One string
					Two string
				}{
					One: "foobar",
					Two: "foobar",
				},
			},
		},
		{
			name: "struct non exported field",
			base: &struct {
				Bar struct {
					One string
					two string
				}
			}{},
			expected: &struct {
				Bar struct {
					One string
					two string
				}
			}{
				Bar: struct {
					One string
					two string
				}{
					One: "foobar",
				},
			},
		},
		{
			name: "pointer on struct",
			base: &struct {
				Bar *struct {
					One string
					Two string
				}
			}{},
			expected: &struct {
				Bar *struct {
					One string
					Two string
				}
			}{
				Bar: &struct {
					One string
					Two string
				}{
					One: "foobar",
					Two: "foobar",
				},
			},
		},
		{
			name: "slice string",
			base: &struct {
				Bar []string
			}{},
			expected: &struct {
				Bar []string
			}{
				Bar: []string{"foobar", "foobar"},
			},
		},
		{
			name: "map string/string",
			base: &struct {
				Bar map[string]string
			}{},
			expected: &struct {
				Bar map[string]string
			}{
				Bar: map[string]string{
					"name0": "foobar",
					"name1": "foobar",
				},
			},
		},
		{
			name: "map string/int",
			base: &struct {
				Bar map[string]int
			}{},
			expected: &struct {
				Bar map[string]int
			}{
				Bar: map[string]int{
					"name0": 42,
					"name1": 42,
				},
			},
		},
		{
			name: "map string/string slice",
			base: &struct {
				Bar map[string][]string
			}{},
			expected: &struct {
				Bar map[string][]string
			}{
				Bar: map[string][]string{
					"name0": {"foobar", "foobar"},
					"name1": {"foobar", "foobar"},
				},
			},
		},
		{
			name: "map string/struct",
			base: &struct {
				Bar map[string]struct {
					One string
					Two string
				}
			}{},
			expected: &struct {
				Bar map[string]struct {
					One string
					Two string
				}
			}{
				Bar: map[string]struct {
					One string
					Two string
				}{
					"name0": {
						One: "foobar",
						Two: "foobar",
					},
					"name1": {
						One: "foobar",
						Two: "foobar",
					},
				},
			},
		},
		{
			name: "custom type",
			base: &struct {
				Bar myString
			}{},
			expected: &struct {
				Bar myString
			}{
				Bar: "foobar",
			},
		},
	}

	for _, test := range testCase {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := capillarity.Fill(test.base)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(test.base, test.expected) {
				t.Errorf("Got %+v, want %+v", test.base, test.expected)
			}
		})
	}
}
