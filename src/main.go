package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hola Mundo")

	//Declaracion de constantes
	const constante1 = 50
	const constante2 int = 100

	fmt.Println(constante1, constante2)

	//Declaracion de variables

	variable1 := 14.5             //Se declara por primera vez la variable y el sistema asigna el tipo
	var variable2 = 50            //Se declara la variable y el sistema asigna el tipo
	var variable3 float64 = 63.50 //Se declara la variable y se declara el tipo de variable
	var variable4 string          //se declara pero no se inicializa

	fmt.Println(variable1, variable2, variable3, variable4)

	var base int = 10
	var altura int = 25

	areaTriangulo := base * altura / 2

	fmt.Println("AreaTriangulo:", areaTriangulo)

	//Zero Values, aqui no se usa el null, se asigna un valor por defecto
	var zeroInt int       //Valor 0
	var zeroFloat float64 // valor 0
	var zeroString string //Valor " " Espacio vacio
	var zeroBool bool     //Valor false

	fmt.Println(zeroBool, zeroFloat, zeroString, zeroInt)

	//Operaciones artimeticas
	var a int = 10
	var b int = 50

	//Suma
	resultado := a + b
	fmt.Println(resultado)

	//Resta
	resultado = a - b
	fmt.Println(resultado)

	//Multiplicacion
	resultado = a * b
	fmt.Println(resultado)

	//Division
	resultado = a / b
	fmt.Println(resultado)

	//Residuo o Modulo
	resultado = a % b
	fmt.Println(resultado)

	//Area del trapecio
	var baseMayor float64 = 9.5
	var baseMenor float64 = 3.5
	var alturaTrapecio float64 = 4

	var areaTrapecio float64 = (baseMayor + baseMenor) / 2 * alturaTrapecio
	fmt.Println("Area de Trapecio:", areaTrapecio)

	//Area del circulo
	var radio float64 = 10

	var areaCirculo = (3.1416) * radio * radio
	fmt.Println("Area del circulo:", areaCirculo)

	/* Uso de libreria fmt */

	//printf
	var nombre string = "Cesar"
	var edad int = 21

	fmt.Printf("%s tiene la edad de %d\n", nombre, edad)
	fmt.Printf("%v tiene la edad de %v\n", nombre, edad) //Se usa %v cuando no se sabe el tipo de dato a imprimir

	//Sprintf
	var stringPrint string
	stringPrint = fmt.Sprintf("%s tiene la edad de %d", nombre, edad)

	fmt.Println(stringPrint)

	//Saber el tipo de dato %T
	fmt.Printf("El tipo de dato de nombre es: %T\n", nombre)
	fmt.Printf("El tipo de dato de edad es: %T\n", edad)

	/* Uso de funciones */
	funcionSimple()

	funcionConParametros(5, 10, "Hola")

	funcionConParametrosAlt(10, 40, 45, "Hoola", "Que tal")

	returnFuncion := funcionConParametrosYReturnSimple(5, 10, "Cesar")
	fmt.Println(returnFuncion)

	return1, return2 := funcionConParametrosYReturnDoble(30, 50, "Hooola")
	fmt.Println("Doble return:", return1, return2)

	returnSolo, _ := funcionConParametrosYReturnDoble(50, 47, "Hola")
	fmt.Println("Return doble y solo agarras un return", returnSolo)

	_, returnSolo2 := funcionConParametrosYReturnDoble(50, 47, "Hola")
	fmt.Println("Return doble y solo agarras un return", returnSolo2)

	trianguloArea := areaDeTriangulo(5, 10)
	circuloArea := areaDeCirculo(10)
	trapecioArea := areaDeTrapecio(5, 10, 15)

	fmt.Println("Area de Triangulo de 5 y 10:", trianguloArea)
	fmt.Println("Area de Circulo de radio 10:", circuloArea)
	fmt.Println("Area de Trapecio:", trapecioArea)

}

/* Funcion simple sin parametros ni return */
func funcionSimple() {
	fmt.Println("Este es el uso de una funcion simple")
}

/* Funcion con parametros declarados con su tipo en cada uno */
func funcionConParametros(a int, b int, c string) {
	fmt.Println("Esta es funcion con parametros", a, b, c)
}

/* Funcion con parametros declarados con su tipo en uno por cada tipo */
func funcionConParametrosAlt(a, b, f int, c, r string) {
	fmt.Println("Esta es funcion con parametros simplificado", a, b, c)
}

/* Funcion con parametros y un return */
func funcionConParametrosYReturnSimple(a int, b int, c string) int {
	fmt.Println("Esta es funcion con parametros y un return simple", a, b, c)
	return a + b
}

/* Funcion con parametros y doble return  */
func funcionConParametrosYReturnDoble(a int, b int, c string) (return1, return2 int) {
	fmt.Println("Esta es funcion con parametros", a, b, c)
	fmt.Println("Se realiza la suma de a y b")
	return a, a + b
}

func areaDeTriangulo(base, altura float64) float64 {
	area := base * altura / 2
	return area
}

func areaDeCirculo(radio float64) float64 {
	area := math.Pow(radio, 2) * math.Pi
	return area
}

func areaDeTrapecio(baseMenor, baseMayor, altura float64) float64 {
	area := (baseMayor + baseMenor) / 2 * altura
	return area
}
