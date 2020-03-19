package main

import (
	"fmt"
	"reflect"
)

var otherThing float32 = -32.34

func init() {

}

func main() {
	fmt.Println(doAnotherThing(otherThing))
}

func doAnotherThing(float float32) string {
	fmt.Println("Hillary is so cool")
	return fmt.Sprintf("%f - is the %x", float, reflect.TypeOf(float))
}
