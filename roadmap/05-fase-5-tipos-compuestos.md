# Fase 5 — Tipos compuestos

> **Objetivo:** Modelar datos reales con slices, maps, structs e interfaces. Acá Go empieza a ser divertido.

## Lo que aprenderemos

- **`array`** (tamaño fijo) vs **`slice`** (dinámico, el pan de cada día de Go):
  ```go
  var colores [3]string         // array: fijo
  numeros := []int{1, 2, 3}     // slice: crece solo
  numeros = append(numeros, 4)  // así se agrega
  ```
- Operaciones sobre slices: `append`, `len`, `cap`, slicing (`s[1:3]`), copiado (`copy`)
- **`map`**: diccionarios clave-valor
  ```go
  edades := map[string]int{"Ana": 30, "Luis": 25}
  edad, existe := edades["Pedro"] // el segundo valor dice si la clave existe
  ```
- **`struct`**: define tu propio tipo de datos. La base de TODO backend Go
  ```go
  type Usuario struct {
      Nombre string
      Edad   int
      Email  string
  }
  ```
- **Métodos**: funciones que cuelgan de un tipo. La forma de Go de "programación orientada a objetos" (sin clases ni herencia)
  ```go
  func (u Usuario) Saludar() string {
      return "Hola, soy " + u.Nombre
  }
  ```
- **Receiver por valor vs puntero** (`func (u *Usuario) CumplirAnios()`) — fundamental, y el origen de 90% de los bugs de principiantes
- **`interface`**: describe comportamientos, no datos. El polimorfismo de Go
  ```go
  type Hablador interface {
      Hablar() string
  }
  ```
  (Si tiene el método, es del tipo — no hay que declarar implementación explícita, eso es lo genial)
- **`type`**: alias y definición de nuevos tipos
- **`nil`**: el "cero" de slices, maps, punteros e interfaces. Y por qué los `nil` causan pánicos

## Palabras reservadas de esta fase

| Palabra | Uso |
|---|---|
| `type` | Define nuevos tipos (structs, interfaces, alias) |
| `struct` | Tipo compuesto de campos |
| `interface` | Tipo que describe comportamientos |
| `map` | Tipo clave-valor |
| `nil` | (builtin, no palabra reservada) Valor cero de punteros/slices/maps/interfaces |

## Ejercicio 🗂️

**Agenda de contactos**:

- `type Contacto struct { Nombre string; Telefono string }`
- Un `map[string]Contacto` (clave = nombre)
- Métodos: `Agregar(nombre, telefono)`, `Buscar(nombre) (Contacto, bool)`, `Eliminar(nombre)`, `ListarTodos()`

```go
package main

import "fmt"

type Contacto struct {
	Nombre   string
	Telefono string
}

type Agenda struct {
	contactos map[string]Contacto
}

func (a *Agenda) Agregar(c Contacto) {
	a.contactos[c.Nombre] = c
}

func (a Agenda) Buscar(nombre string) (Contacto, bool) {
	c, existe := a.contactos[nombre]
	return c, existe
}

func main() {
	agenda := Agenda{contactos: make(map[string]Contacto)}
	agenda.Agregar(Contacto{Nombre: "Ana", Telefono: "555-1234"})

	c, ok := agenda.Buscar("Ana")
	if ok {
		fmt.Println("Encontrado:", c.Nombre, c.Telefono)
	}
}
```

> 💡 Fijate: `Agregar` usa puntero (`*Agenda`) porque modifica el map, `Buscar` no necesita modificar nada.

## Criterio para pasar a la Fase 6

- [ ] Explico la diferencia entre `array` y `slice` en una frase
- [ ] Usé `append` y el patrón `valor, ok := map[clave]`
- [ ] Definí un `struct` y le agregué un método
- [ ] Entiendo por qué `Agregar` necesita puntero y `Buscar` no
- [ ] Defín una `interface` y un tipo que la implemente
