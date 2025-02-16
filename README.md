# Calculator Module

Este módulo es una prueba de calculadora básica escrita en Go. Proporciona funciones para realizar operaciones simples como suma, resta, multiplicación y división.

## Instalación

Para instalar el módulo, puedes clonar el repositorio y luego compilar el código:

```bash
git clone <URL_DEL_REPOSITORIO>
cd calculator
go build
```

## Uso

Puedes importar el módulo en tu proyecto de Go y utilizar las funciones disponibles:

```go
package main

import (
    "fmt"
    "calculator"
)

func main() {
    a := 10
    b := 5

    fmt.Println("Suma:", calculator.Sumar(a, b))
    fmt.Println("Resta:", calculator.Restar(a, b))
    fmt.Println("Multiplicación:", calculator.Multiplicar(a, b))
    fmt.Println("División:", calculator.Dividir(a, b))
}
```

## Funciones

- `Sumar(a int, b int) int`: Devuelve la suma de `a` y `b`.
- `Restar(a int, b int) int`: Devuelve la resta de `b` de `a`.
- `Multiplicar(a int, b int) int`: Devuelve el producto de `a` y `b`.
- `Dividir(a int, b int) (int, error)`: Devuelve el cociente de `a` y `b`. Si `b` es 0, devuelve un error.

## Contribuciones

Las contribuciones son bienvenidas. Por favor, abre un issue o un pull request para discutir cualquier cambio que te gustaría realizar.