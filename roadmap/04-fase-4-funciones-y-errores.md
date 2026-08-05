# Fase 4 — Funciones y errores

> **Objetivo:** Escribir funciones reutilizables y manejar errores como lo hace Go: sin excepciones, con valores.

## Lo que aprenderemos

- **`func`**: parámetros, tipos de retorno, funciones con nombre y funciones anónimas
- **Retornos múltiples**: la joya de Go — una función puede devolver varios valores a la vez
  ```go
  func dividir(a, b int) (int, error) {
      if b == 0 {
          return 0, errors.New("no se puede dividir por cero")
      }
      return a / b, nil
  }
  ```
- **`error` como valor**: Go NO tiene try/catch. El error es un valor más que se retorna y se chequea:
  ```go
  resultado, err := dividir(10, 2)
  if err != nil {
      fmt.Println("Error:", err)
      return
  }
  fmt.Println("Resultado:", resultado)
  ```
  (Este patrón `if err != nil` va a estar en el 90% de tu código. Abrazalo.)
- Retornos nombrados (named returns) — y por qué conviene evitarlos salvo casos puntuales
- **`defer`**: programa una función para ejecutarse AL FINAL de la función actual. Perfecto para cerrar archivos, conexiones, etc.
  ```go
  f, _ := os.Open("archivo.txt")
  defer f.Close() // se cierra solo cuando termine la función
  ```
- **`panic`** / **`recover`**: el manejo de "pánico" de Go. Cuándo SÍ usarlo (errores irrecuperables) y cuándo NO (nunca para validar entrada del usuario)
- Parámetros variádicos (`...int`)

## Palabras reservadas de esta fase

| Palabra | Uso |
|---|---|
| `func` | Declara funciones (y métodos, en la Fase 5) |
| `return` | Devuelve valores de una función |
| `defer` | Ejecuta algo al final de la función |
| `panic` | (builtin, no palabra reservada) Detiene el programa con un error |
| `recover` | (builtin) Captura un panic en curso |

## Ejercicio 📊

**Conversor de temperaturas** con validación de errores:

- `celsiusAFahrenheit(c float64) float64`
- `fahrenheitACelsius(f float64) float64`
- `temperaturaValida(temp float64, escala string) (bool, error)` — valida que la escala sea "C" o "F" y que la temperatura esté en rango físico

```go
package main

import (
	"errors"
	"fmt"
)

func celsiusAFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func fahrenheitACelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	fmt.Println("100°C =", celsiusAFahrenheit(100), "°F")
	fmt.Println("212°F =", fahrenheitACelsius(212), "°C")

	// Crea una función que valide escala y devuelva error
	_, err := errors.New("placeholder") // reemplazar con tu implementación
	fmt.Println(err)
}
```

Extra: agregá una función `convertir()` que reciba temperatura, escala origen y escala destino, y devuelva `(resultado, error)` — usando los retornos múltiples.

## Criterio para pasar a la Fase 5

- [ ] Escribí una función con retorno múltiple `(valor, error)`
- [ ] Manejé un error con `if err != nil` sin copiar código
- [ ] Usé `defer` para cerrar algo (aunque sea `fmt.Println("fin")`)
- [ ] Explico la diferencia entre `error` y `panic` en una frase
- [ ] Sé cuándo NO usar `panic` (pista: casi siempre)
