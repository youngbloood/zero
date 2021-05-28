package zero

import (
	"reflect"
)

var Filter map[reflect.Kind]func(reflect.Value)

// Reset . reset object to zero-value
func Reset(v interface{}) {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.Array, reflect.String,
		reflect.Bool, reflect.Struct:
		return
	}

	rve := rv
	for rve.Kind() == reflect.Ptr {
		rve = rve.Elem()
	}

	switch rve.Kind() {
	case reflect.Struct:
		for i := 0; i < rve.NumField(); i++ {
			field := rve.Field(i)
			kind := field.Kind()
			if Filter != nil {
				if filt, ok := Filter[kind]; ok {
					filt(field)
					continue
				}
			}
			field.Set(reflect.Zero(field.Type()))
		}
	default:
		if !rve.CanAddr() {
			return
		}
		if Filter != nil {
			if filt, ok := Filter[rve.Kind()]; ok {
				filt(rve)
				break
			}
		}
		rve.Set(reflect.Zero(rve.Type()))
		return
	}
}
