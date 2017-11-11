package capillarity

import (
	"fmt"
	"reflect"
)

const theAnswer = 42

// Capillarity configuration
type Capillarity struct {
	ExcludedFieldNames  []string
	SliceItemNumber     int
	MapItemNumber       int
	DefaultString       string
	DefaultNumber       int
	DefaultBool         bool
	DefaultMapKeyPrefix string
}

// NewCapillarity Create a new Capillarity
func NewCapillarity(options ...Option) Capillarity {
	capil := &Capillarity{
		SliceItemNumber:     2,
		MapItemNumber:       2,
		DefaultString:       "foobar",
		DefaultNumber:       theAnswer,
		DefaultBool:         true,
		DefaultMapKeyPrefix: "name",
	}

	for _, opt := range options {
		opt(capil)
	}

	return *capil
}

// Fill an object
func (c Capillarity) Fill(element interface{}) error {
	field := reflect.ValueOf(element)
	return c.fill(field)
}

func (c Capillarity) fill(field reflect.Value) error {
	switch field.Kind() {
	case reflect.Struct:
		err := c.setStruct(field)
		if err != nil {
			return err
		}
	case reflect.Ptr:
		err := c.setPointer(field)
		if err != nil {
			return err
		}
	case reflect.Slice:
		err := c.setSlice(field)
		if err != nil {
			return err
		}
	case reflect.Map:
		err := c.setMap(field)
		if err != nil {
			return err
		}
	case reflect.Interface:
		if err := c.fill(field.Elem()); err != nil {
			return err
		}
	case reflect.String:
		c.setTyped(field, c.DefaultString)
	case reflect.Int:
		c.setTyped(field, c.DefaultNumber)
	case reflect.Int8:
		c.setTyped(field, int8(c.DefaultNumber))
	case reflect.Int16:
		c.setTyped(field, int16(c.DefaultNumber))
	case reflect.Int32:
		c.setTyped(field, int32(c.DefaultNumber))
	case reflect.Int64:
		c.setTyped(field, int64(c.DefaultNumber))
	case reflect.Uint:
		c.setTyped(field, uint(c.DefaultNumber))
	case reflect.Uint8:
		c.setTyped(field, uint8(c.DefaultNumber))
	case reflect.Uint16:
		c.setTyped(field, uint16(c.DefaultNumber))
	case reflect.Uint32:
		c.setTyped(field, uint32(c.DefaultNumber))
	case reflect.Uint64:
		c.setTyped(field, uint64(c.DefaultNumber))
	case reflect.Bool:
		c.setTyped(field, c.DefaultBool)
	case reflect.Float32:
		c.setTyped(field, float32(c.DefaultNumber))
	case reflect.Float64:
		c.setTyped(field, float64(c.DefaultNumber))
	}

	return nil
}

func (c Capillarity) setTyped(field reflect.Value, i interface{}) {
	baseValue := reflect.ValueOf(i)
	if field.Kind().String() == field.Type().String() {
		field.Set(baseValue)
	} else {
		field.Set(baseValue.Convert(field.Type()))
	}
}

func (c Capillarity) setMap(field reflect.Value) error {
	field.Set(reflect.MakeMap(field.Type()))

	for i := 0; i < c.MapItemNumber; i++ {
		// TODO support only string... must be fixed
		//fmt.Println(field.Type().Key())
		baseKeyName := c.makeKeyName(field.Type().Elem())
		key := reflect.ValueOf(fmt.Sprintf("%s%d", baseKeyName, i))

		// generate value
		ptrType := reflect.PtrTo(field.Type().Elem())
		ptrValue := reflect.New(ptrType)
		if err := c.fill(ptrValue); err != nil {
			return err
		}
		value := ptrValue.Elem().Elem()

		field.SetMapIndex(key, value)
	}
	return nil
}

func (c Capillarity) makeKeyName(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.Ptr:
		name := typ.Elem().Name()
		if name == "" {
			return c.DefaultMapKeyPrefix
		}
		return name
	case reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Bool, reflect.Float32, reflect.Float64:
		return c.DefaultMapKeyPrefix
	default:
		name := typ.Name()
		if name == "" {
			return c.DefaultMapKeyPrefix
		}
		return name
	}
}

func (c Capillarity) setStruct(field reflect.Value) error {
	for i := 0; i < field.NumField(); i++ {
		fld := field.Field(i)
		stFld := field.Type().Field(i)

		if !isExported(stFld) || c.isExcluded(stFld) {
			continue
		}

		if err := c.fill(fld); err != nil {
			return err
		}
	}
	return nil
}

func (c Capillarity) setSlice(field reflect.Value) error {
	field.Set(reflect.MakeSlice(field.Type(), c.SliceItemNumber, c.SliceItemNumber))
	for j := 0; j < field.Len(); j++ {
		if err := c.fill(field.Index(j)); err != nil {
			return err
		}
	}
	return nil
}

func (c Capillarity) setPointer(field reflect.Value) error {
	if field.IsNil() {
		field.Set(reflect.New(field.Type().Elem()))
		if err := c.fill(field.Elem()); err != nil {
			return err
		}
	} else {
		if err := c.fill(field.Elem()); err != nil {
			return err
		}
	}
	return nil
}

// isExported return true is a struct field is exported, else false
func isExported(f reflect.StructField) bool {
	if f.PkgPath != "" && !f.Anonymous {
		return false
	}
	return true
}

func (c Capillarity) isExcluded(field reflect.StructField) bool {
	for _, name := range c.ExcludedFieldNames {
		if field.Name == name {
			return true
		}
	}
	return false
}
