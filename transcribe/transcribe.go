package transcribe

import (
	"reflect"
	"unsafe"
)

// Transcribe copy src in any type to an interface
func Transcribe(src interface{}) interface{} {
	v := reflect.ValueOf(src)
	return copyAny(v).Interface()
}

func copyAny(v reflect.Value) reflect.Value {
	switch v.Kind() {
	case reflect.Slice:
		return copySlice(v)
	case reflect.Array:
		return copyArray(v)
	case reflect.Struct:
		return copyStruct(v)
	case reflect.Map:
		return copyMap(v)
	case reflect.Ptr:
		return copyPtr(v)
	case reflect.Interface:
		return copyInterface(v)
	default:
		return v
	}
}

func copySlice(v reflect.Value) reflect.Value {
	d := reflect.MakeSlice(v.Type(), v.Len(), v.Cap())
	for i := 0; i < v.Len(); i++ {
		d.Index(i).Set(copyAny(v.Index(i)))
	}
	return d
}

func copyArray(v reflect.Value) reflect.Value {
	d := reflect.New(v.Type()).Elem()
	for i := 0; i < v.Len(); i++ {
		d.Index(i).Set(copyAny(v.Index(i)))
	}
	return d
}

func copyStruct(v reflect.Value) reflect.Value {
	d := reflect.New(v.Type()).Elem()
	for i := 0; i < v.NumField(); i++ {
		ft := v.Type().Field(i)
		fv := v.Field(i)

		dst := d.FieldByName(ft.Name)
		if ft.PkgPath == "" {
			dst.Set(copyAny(fv))
		} else {
			// unexported field
			unexported := reflect.NewAt(fv.Type(), unsafe.Pointer(fv.UnsafeAddr())).Elem()
			dst.Set(copyAny(unexported))
		}
	}
	return d
}

func copyMap(v reflect.Value) reflect.Value {
	d := reflect.MakeMap(v.Type())
	for _, key := range v.MapKeys() {
		d.SetMapIndex(key, copyAny(v.MapIndex(key)))
	}
	return d
}

func copyPtr(v reflect.Value) reflect.Value {
	return copyAny(v.Elem()).Addr()
}

func copyInterface(v reflect.Value) reflect.Value {
	return copyAny(v.Elem())
}
