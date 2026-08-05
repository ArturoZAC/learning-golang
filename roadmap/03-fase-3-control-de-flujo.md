# Fase 3 — Control de flujo

> **Objetivo:** Tomar decisiones y repetir código. Sin esto, un programa es una lista de instrucciones aburrida.

## Lo que aprenderemos

- **`if` / `else`** con la sintaxis especial de Go: inicializar una variable en la propia condición
  ```go
  if nota := 85; nota >= 60 {
      fmt.Println("Aprobado")
  } else {
      fmt.Println("Desaprobado")
  }
  ```
  (La variable `nota` solo vive dentro del if — scope reducido)
- **`for`**: el ÚNICO loop de Go. No existe `while` ni `do-while` — todo se hace con `for`:
  - `for i := 0; i < 10; i++` → loop clásico
  - `for condicion` → actúa como `while`
  - `for` solo → loop infinito (se corta con `break`)
- **`range`**: iterar sobre slices, maps y strings (primer vistazo)
- **`switch` / `case` / `default`** — y la gran diferencia con otros lenguajes: en Go NO hace falta `break` en cada case (el break es implícito)
- **`fallthrough`**: la rara palabra que fuerza a seguir al siguiente case (rara vez se usa, pero hay que conocerla)
- **`break`** y **`continue`** para controlar los loops

## Palabras reservadas de esta fase

| Palabra | Uso |
|---|---|
| `if` / `else` | Condicional |
| `for` | El único loop de Go |
| `range` | Itera sobre colecciones |
| `switch` / `case` / `default` | Condicional múltiple |
| `break` | Sale del loop o switch |
| `continue` | Salta a la siguiente iteración |
| `fallthrough` | Continúa al siguiente case |
| `goto` | Salto incondicional (la palabra más odiada de Go — solo se menciona para saber que existe) |

## Ejercicio 🔢

**FizzBuzz** — el clásico de entrevistas:

- Para números del 1 al 100:
  - Múltiplo de 3 → imprimir "Fizz"
  - Múltiplo de 5 → imprimir "Buzz"
  - Múltiplo de ambos → imprimir "FizzBuzz"
  - Otro → imprimir el número

```go
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
```

> 💡 Fijate el `switch` sin expresión: es la forma de Go de hacer `if/else if/else` elegante.

## Criterio para pasar a la Fase 4

- [ ] Escribí FizzBuzz SIN mirar el código de arriba
- [ ] Sé hacer los 3 tipos de `for` (clásico, condición, infinito)
- [ ] Explico por qué en Go el `switch` no necesita `break`
- [ ] Sé qué hace `range` sobre un slice
- [ ] Usé `break` y `continue` en un loop
