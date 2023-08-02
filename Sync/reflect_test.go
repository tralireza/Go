package lsync

import (
	"log"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	func() {
		n := 42
		v, p := reflect.ValueOf(n), reflect.ValueOf(&n)
		log.Printf("%T %[1]v %v | %T %[3]v %v", v, v.Kind(), p, p.Kind())
	}()

	var i interface{}
	i = new(int)
	log.Printf("interface{} = new(int) -> %v, %v, %v", reflect.TypeOf(i), reflect.TypeOf(i).Kind(), reflect.ValueOf(i).Elem())

	i = new(struct{ a, b int })
	log.Printf("interface{} = new(struct{}) -> %v, %v, %v", reflect.TypeOf(i), reflect.TypeOf(i).Kind(), reflect.ValueOf(i).Elem())
	log.Printf("%v | %+v", reflect.ValueOf(i).Type(), reflect.ValueOf(i).Interface())

	log.Printf("CanSet: %v", reflect.ValueOf(i).CanSet())
	go func() {
		defer func() {
			log.Print("[panic] -> ", recover())
		}()
		reflect.ValueOf(i).SetInt(1)
	}()

	i = new(int)
	var e reflect.Value = reflect.ValueOf(&i).Elem()
	log.Printf("CanSet of (%v) -> ValueOf(): %v | ValueOf().Elem(): %v", reflect.TypeOf(&i), reflect.ValueOf(&i).CanSet(), e.CanSet())
}
