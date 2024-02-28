//Se requiere inicializar el proyecto como módulo con el comando go mod init main

package main

import (
	"fmt"
	"reflect"
)

// funciones en GO
func main() {
	fmt.Println("Hello, World!")

	// STRINGS
	var miString string = "Hola mundo variable"
	fmt.Println(miString)

	// NUMBERS

	var miInt int = 10
	miInt = miInt + 1
	fmt.Println(miInt)

	// FLOATS
	var miFloat32 float32 = 3.14 * 9999999999999
	var miFloat64 float64 = 3.14 * 9999999999999

	fmt.Println(miFloat32)
	fmt.Println(miFloat64)

	// BOOLEAN
	var miBool bool = true
	fmt.Println(miBool)

	// CONSTANTES
	const pi float64 = 3.14
	fmt.Println(pi)

	// Controles de flujo

	// IF

	if miInt >= 10 {
		fmt.Println("Es mayor o igual a 10")
	} else {
		fmt.Println("Es menor a 10")
	}

	// ARRAYS
	var miArray [10]int
	miArray[0] = 2
	fmt.Println(miArray)

	// MAPS

	// Maps se inicializan con un key : value, indicando que la key es de tipo string y el value es de tipo int
	miMap := make(map[string]int)

	miMap["Edad Vicente Jorquera"] = 21
	fmt.Println(miMap)
	fmt.Println(miMap["Edad Vicente Jorquera"])
	// Verficar el tipo de dato con el que se está trabajando

	fmt.Println(reflect.TypeOf(miInt))
	fmt.Println(reflect.TypeOf(miString))
	fmt.Println(reflect.TypeOf(miFloat32))
	fmt.Println(reflect.TypeOf(miFloat64))
	fmt.Println(reflect.TypeOf(miBool))
	fmt.Println(reflect.TypeOf(pi))
}
