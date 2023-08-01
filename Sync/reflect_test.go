package lsync

import (
	"log"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	i := 42
	v, p := reflect.ValueOf(i), reflect.ValueOf(&i)
	log.Printf("%T %[1]v %v  %T %[3]v %v", v, v.Kind(), p, p.Kind())

}
