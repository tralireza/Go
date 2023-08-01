package lsync

import (
	"log"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	i := 42
	v, p := reflect.ValueOf(i), reflect.ValueOf(&i)
	log.Printf("%T %[1]v  %T %[2]v", v, p)
}
