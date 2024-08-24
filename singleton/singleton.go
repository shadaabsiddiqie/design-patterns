package singleton

import (
	"fmt"
	"sync"
)

type someObject struct {
	name string
}

var once sync.Once
var obj *someObject

func initSingletonObj() *someObject {
	once.Do(func() {
		obj = &someObject{
			name: "initial value",
		}
	})
	return obj
}

func CreatAnSingletonObject() {
	obj1 := initSingletonObj()
	fmt.Println(obj1.name)
	obj1.name = "next value"
	fmt.Println(obj1.name)
	obj2 := initSingletonObj()
	fmt.Println(obj2.name)
}
