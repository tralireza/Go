package lsync

import (
	"log"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	n := 42
	v, p := reflect.ValueOf(n), reflect.ValueOf(&n)
	log.Printf("%T %[1]v %v | %T %[3]v %v", v, v.Kind(), p, p.Kind())

	var i interface{}
	i = new(int)
	log.Printf("interface{} = new(int): %v, %v, %v", reflect.TypeOf(i), reflect.TypeOf(i).Kind(), reflect.ValueOf(i).Elem())

	i = new(struct{ a, b int })
	log.Printf("interface{} = new(struct{}): %v, %v, %v", reflect.TypeOf(i), reflect.TypeOf(i).Kind(), reflect.ValueOf(i).Elem())
}
