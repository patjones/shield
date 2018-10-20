package bus

import (
	"reflect"
)

func marshal(thing interface{}) interface{} {
	t := reflect.TypeOf(thing)
	v := reflect.ValueOf(thing)

	for t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = v.Type()
	}

	if t.Kind() != reflect.Struct {
		panic("bus.event() only operates on structures")
	}

	return reflectOn(t, &v)
}

func reflectOn(t reflect.Type, v *reflect.Value) interface{} {
	switch t.Kind() {
	default:
		return v.Interface()

	case reflect.Slice:
		l := make([]interface{}, v.Len())

		for i := 0; i < v.Len(); i++ {
			v2 := v.Index(i)
			l[i] = reflectOn(v2.Type(), &v2)
		}
		return l

	case reflect.Struct:
		m := make(map[string] interface{})
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.PkgPath != "" {
				continue
			}
			tag, set := field.Tag.Lookup("mbus")
			if !set {
				continue
			}

			v2 := v.Field(i)
			m[tag] = reflectOn(v2.Type(), &v2)
		}
		return m
	}
}
