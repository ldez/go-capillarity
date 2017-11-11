package capillarity

// Option of Capillarity
type Option func(*Capillarity)

// WithSliceItemNumber The number of generated slice items
func WithSliceItemNumber(number int) Option {
	return func(capil *Capillarity) {
		capil.SliceItemNumber = number
	}
}

// WithMapItemNumber The number of generated map items
func WithMapItemNumber(number int) Option {
	return func(capil *Capillarity) {
		capil.MapItemNumber = number
	}
}

// WithDefaultString The default value for the type 'string'
func WithDefaultString(value string) Option {
	return func(capil *Capillarity) {
		capil.DefaultString = value
	}
}

// WithDefaultNumber The default value for the types [int*, uint*, float*]
func WithDefaultNumber(number int) Option {
	return func(capil *Capillarity) {
		capil.DefaultNumber = number
	}
}

// WithDefaultBool The default value for the type 'bool'
func WithDefaultBool(value bool) Option {
	return func(capil *Capillarity) {
		capil.DefaultBool = value
	}
}

// WithDefaultMapKeyPrefix The default map key prefix
func WithDefaultMapKeyPrefix(value string) Option {
	return func(capil *Capillarity) {
		capil.DefaultMapKeyPrefix = value
	}
}
