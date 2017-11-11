package capillarity

import (
	"reflect"
	"testing"
)

func TestNewCapillarityWithOption(t *testing.T) {
	testCases := []struct {
		name     string
		actual   Capillarity
		expected Capillarity
	}{
		{
			name:   "base",
			actual: NewCapillarity(),
			expected: Capillarity{
				SliceItemNumber:     2,
				MapItemNumber:       2,
				DefaultString:       "foobar",
				DefaultNumber:       theAnswer,
				DefaultBool:         true,
				DefaultMapKeyPrefix: "name",
			},
		},
		{
			name:   "WithDefaultString",
			actual: NewCapillarity(WithDefaultString("go")),
			expected: Capillarity{
				SliceItemNumber:     2,
				MapItemNumber:       2,
				DefaultString:       "go",
				DefaultNumber:       theAnswer,
				DefaultBool:         true,
				DefaultMapKeyPrefix: "name",
			},
		},
		{
			name:   "WithDefaultBool",
			actual: NewCapillarity(WithDefaultBool(false)),
			expected: Capillarity{
				SliceItemNumber:     2,
				MapItemNumber:       2,
				DefaultString:       "foobar",
				DefaultNumber:       theAnswer,
				DefaultBool:         false,
				DefaultMapKeyPrefix: "name",
			},
		},
		{
			name:   "WithDefaultNumber",
			actual: NewCapillarity(WithDefaultNumber(6)),
			expected: Capillarity{
				SliceItemNumber:     2,
				MapItemNumber:       2,
				DefaultString:       "foobar",
				DefaultNumber:       6,
				DefaultBool:         true,
				DefaultMapKeyPrefix: "name",
			},
		},
		{
			name:   "WithSliceItemNumber",
			actual: NewCapillarity(WithSliceItemNumber(6)),
			expected: Capillarity{
				SliceItemNumber:     6,
				MapItemNumber:       2,
				DefaultString:       "foobar",
				DefaultNumber:       theAnswer,
				DefaultBool:         true,
				DefaultMapKeyPrefix: "name",
			},
		},
		{
			name:   "WithMapItemNumber",
			actual: NewCapillarity(WithMapItemNumber(6)),
			expected: Capillarity{
				SliceItemNumber:     2,
				MapItemNumber:       6,
				DefaultString:       "foobar",
				DefaultNumber:       theAnswer,
				DefaultBool:         true,
				DefaultMapKeyPrefix: "name",
			},
		},
		{
			name:   "WithDefaultMapKeyPrefix",
			actual: NewCapillarity(WithDefaultMapKeyPrefix("prefix")),
			expected: Capillarity{
				SliceItemNumber:     2,
				MapItemNumber:       2,
				DefaultString:       "foobar",
				DefaultNumber:       theAnswer,
				DefaultBool:         true,
				DefaultMapKeyPrefix: "prefix",
			},
		},
		{
			name:   "multiple options",
			actual: NewCapillarity(WithDefaultString("go"), WithDefaultNumber(6), WithDefaultMapKeyPrefix("prefix")),
			expected: Capillarity{
				SliceItemNumber:     2,
				MapItemNumber:       2,
				DefaultString:       "go",
				DefaultNumber:       6,
				DefaultBool:         true,
				DefaultMapKeyPrefix: "prefix",
			},
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if !reflect.DeepEqual(test.actual, test.expected) {
				t.Errorf("Got %+v, want %+v", test.actual, test.expected)
			}
		})
	}
}
