package main

import "fmt"

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
}
