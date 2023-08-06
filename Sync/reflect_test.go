package lsync

import (
	"log"
	"reflect"
	"sync"
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
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			log.Print("[panic] -> ", recover())
		}()
		reflect.ValueOf(i).SetInt(1)
	}()
	wg.Wait()

	func() {
		i := int64(0)
		var e reflect.Value = reflect.ValueOf(&i).Elem()
		log.Printf("CanSet of (%v) -> ValueOf(): %v | ValueOf().Elem(): %v", reflect.TypeOf(&i), reflect.ValueOf(&i).CanSet(), e.CanSet())

		e.SetInt(42)
		log.Printf("SetInt: %v", i)
	}()

	r := reflect.TypeOf(int64(24))
	log.Printf("%T %[1]v %T %[2]v", r, reflect.New(r))
}

type RefPerson struct {
	Name string `json:"name" xml:"-" tst:"name,qfr1,qfr2"`
	Year int    `json:"year,omitempty" xml:"-"`
}

func TestRefFieldTags(t *testing.T) {
	v := reflect.ValueOf(RefPerson{"Mr Reflection", 2006})
	r := v.Type()

	log.Printf("%+v %v", v, r)
	for i := 0; i < r.NumField(); i++ {
		log.Printf("- %v: %v", r.Field(i).Name, v.Field(i))

		f := r.Field(i)
		log.Printf(":Tag:  %v | %v | %v", f.Tag.Get("json"), f.Tag.Get("xml"), f.Tag.Get("tst"))
	}
}

func TestRefMap(t *testing.T) {
	m := map[string]int{"a": 0, "b": 1, "c": 2, "z": 25}
	v := reflect.ValueOf(m)
	r := reflect.TypeOf(m)

	log.Printf("%v, Kind: %v, %v, %v", m, v.Kind(), v, r)

	for _, k := range v.MapKeys() {
		log.Print(k, v.MapIndex(k))
	}

	v.SetMapIndex(reflect.ValueOf("y"), reflect.ValueOf(24))

	itr := v.MapRange()
	log.Printf("%T -> %[1]v", itr)
	for itr.Next() {
		log.Print(itr.Key(), itr.Value())
	}
}

// 791
func TestCustomSortString(t *testing.T) {
	customSortString := func(order string, s string) string {
		m := map[byte]int{}
		for _, r := range s {
			m[byte(r)]++
		}
		log.Print(m)

		k, bs := 0, make([]byte, len(s))
		for _, r := range order {
			for f := m[byte(r)]; f > 0; f-- {
				bs[k] = byte(r)
				k++
			}
			m[byte(r)] = 0
		}
		log.Print(m)

		for b, f := range m {
			for ; f > 0; f-- {
				bs[k] = b
				k++
			}
		}
		return string(bs)
	}

	log.Print("abcd -> ", customSortString("bcafg", "abcd"))
}
