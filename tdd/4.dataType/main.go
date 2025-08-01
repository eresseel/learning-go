// signed integers
// | **Típus** | **Méret**              | **Tartomány**                                      |
// |----------:|------------------------|----------------------------------------------------|
// | `int`     | 32 bit / 64 bit        | -2,147,483,648 … 2,147,483,647 *(32 bit)*<br>-9,223,372,036,854,775,808 … 9,223,372,036,854,775,807 *(64 bit)* |
// | `int8`    | 8 bit / 1 byte         | -128 … 127                                         |
// | `int16`   | 16 bit / 2 byte        | -32,768 … 32,767                                   |
// | `int32`   | 32 bit / 4 byte        | -2,147,483,648 … 2,147,483,647                     |
// | `int64`   | 64 bit / 8 byte        | -9,223,372,036,854,775,808 … 9,223,372,036,854,775,807 |
// unsigned integers
// | **Típus** | **Méret**              | **Tartomány**                                      |
// |----------:|------------------------|----------------------------------------------------|
// | `uint`    | 32 bit / 64 bit        | 0 … 4,294,967,295 *(32 bit)*<br>0 … 18,446,744,073,709,551,615 *(64 bit)* |
// | `uint8`   | 8 bit / 1 byte         | 0 … 255                                            |
// | `uint16`  | 16 bit / 2 byte        | 0 … 65,535                                         |
// | `uint32`  | 32 bit / 4 byte        | 0 … 4,294,967,295                                  |
// | `uint64`  | 64 bit / 8 byte        | 0 … 18,446,744,073,709,551,615                     |
// float
// | **Típus**   | **Méret**     | **Tartomány**                      |
// |------------:|---------------|------------------------------------|
// | `float32`   | 32 bit        | -3.4e+38 … 3.4e+38                 |
// | `float64`   | 64 bit        | -1.7e+308 … 1.7e+308               |
package main

import (
	"fmt"
	"reflect"
)

func DetectType(value any) string {
	t := reflect.TypeOf(value).Kind()

	switch t {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Float64:
		return "float64"
	case reflect.Float32:
		return "float32"
	case reflect.Bool:
		return "bool"
	case reflect.Slice:
		return "slice"
	case reflect.Map:
		return "map"
	default:
		return "unknown"
	}
}

func main() {
	values := []any{5, "hello", 3.14, true, []int{1, 2}, map[string]int{"x": 10}}

	for _, v := range values {
		fmt.Printf("Value: %v → Type: %s\n", v, DetectType(v))
	}
}
