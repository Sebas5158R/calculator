package calculator

// variable privada (inicia con minuscula), solo se puede usar dentro del paquete
var logMessage = "[LOG]"

// variable publica (inicia con mayuscula)
var Version = "1.0.0"

// Funcion privada (inicia con minuscula), solo se puede usar dentro del paquete
// internalSum recibe un numero entero y devuelve el numero - 1
func internalSum(number int) int {
	return number - 1
}

// Funcion publica (inicia con mayuscula)
// Sum recibe dos numeros enteros y devuelve la suma de ellos
func Sum(number1, number2 int) int {
	return number1 + number2
}
