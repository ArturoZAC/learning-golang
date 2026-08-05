# Fase 2 — Fundamentos del lenguaje

> **Objetivo:** Dominar variables, tipos y operadores. La base sobre la que se construye todo.

## Lo que aprenderemos

- Declaración de variables con **`var`** y la forma corta **`:=`** (el operador estrella de Go)
- Tipos básicos: `int`, `float64`, `string`, `bool`
- Constantes con **`const`** (y la diferencia con `var`)
- Ceros por defecto: Go inicializa toda variable (0, "", false) — nunca hay "undefined"
- Conversión de tipos (Go NO convierte automáticamente: `float64(x)` a la fuerza)
- Operadores aritméticos, de comparación y lógicos
- El paquete **`fmt`**: `Println`, `Printf` (formato con `%d`, `%s`, `%v`, etc.) y `Sprintf`
- Reglas de alcance (scope): variables a nivel de paquete vs dentro de funciones
- ¿Qué pasa con las variables declaradas y nunca usadas? (Go no te deja compilar 😤)

## Palabras reservadas de esta fase

| Palabra | Uso |
|---|---|
| `var` | Declara una variable con su tipo |
| `const` | Declara un valor inmutable |
| `type` | Primera mención: define tipos (la usamos en serio en Fase 5) |

## Ejercicio 🧮

Calculadora de propina: pedir (o declarar) el precio de una comida y un porcentaje de propina, y calcular:

- El monto de la propina
- El total a pagar

```go
package main

import "fmt"

func main() {
	precio := 125.50
	porcentaje := 15.0

	propina := precio * (porcentaje / 100)
	total := precio + propina

	fmt.Printf("Precio: $%.2f\n", precio)
	fmt.Printf("Propina (%v%%): $%.2f\n", porcentaje, propina)
	fmt.Printf("Total: $%.2f\n", total)
}
```

Variantes: agregar descuento, probar con `int` y `float64` mezclados (¿qué error te da?).

## Criterio para pasar a la Fase 3

- [ ] Explico la diferencia entre `var` y `:=`
- [ ] Sé los 4 tipos básicos y sus ceros por defecto
- [ ] Sé convertir entre `int` y `float64` sin errores
- [ ] Usé `Printf` con al menos `%s`, `%d`, `%v` y `%.2f`
- [ ] Entiendo por qué Go me obliga a usar todas las variables que declaro
