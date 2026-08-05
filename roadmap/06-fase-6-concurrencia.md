# Fase 6 — Concurrencia 🚀

> **Objetivo:** Ejecutar cosas en paralelo de forma simple. Esto es lo que hace famoso a Go y lo distingue del resto.

## Lo que aprenderemos

- **Goroutine** (`go`): ejecutar una función en paralelo con una sola palabra
  ```go
  go procesarArchivo("a.txt")
  go procesarArchivo("b.txt")
  // ambas corren "al mismo tiempo"
  ```
  (Una goroutine cuesta ~2KB de memoria, contra ~1MB de un hilo del sistema. Podés tener miles.)
- **`chan` (channel)**: el tubo por donde las goroutines se comunican y sincronizan
  ```go
  resultados := make(chan int)
  go func() { resultados <- 42 }()  // envía al canal
  valor := <-resultados             // recibe del canal (bloquea hasta que llegue algo)
  ```
- Canales **buffered** vs **unbuffered** (¿cuándo bloquea el envío?)
- **`select`**: esperar a varios canales a la vez — elegir el primero que responda
  ```go
  select {
  case r := <-canalA:
      fmt.Println("Llegó de A:", r)
  case <-time.After(2 * time.Second):
      fmt.Println("Timeout!")
  }
  ```
- Cerrar canales (`close`) y el patrón `for range canal`
- **`sync.WaitGroup`**: esperar a que todas las goroutines terminen
- **`sync.Mutex`**: proteger datos compartidos de accesos simultáneos (evitar race conditions)
- **Race detector**: `go run -race` — la herramienta que detecta condiciones de carrera
- **`context`**: cómo cancelar goroutines (tiempos de espera, cancelación en cascada) — CRÍTICO para backends con peticiones de usuarios

## Palabras reservadas de esta fase

| Palabra | Uso |
|---|---|
| `go` | Lanza una goroutine |
| `chan` | Define un tipo de canal |
| `select` | Espera a múltiples canales a la vez |
| `range` | Itera sobre un canal (recibe hasta que se cierre) |
| `sync` | (paquete, no palabra reservada) WaitGroup, Mutex |

## Ejercicio ⏳

**Suma paralela**: 5 goroutines, cada una suma números de su rango, y envían su resultado por un canal. El `main` recibe los 5 resultados y suma el total.

```go
package main

import (
	"fmt"
	"sync"
)

func sumarRango(inicio, fin int, canal chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	total := 0
	for i := inicio; i <= fin; i++ {
		total += i
	}
	canal <- total
}

func main() {
	canal := make(chan int, 5)
	var wg sync.WaitGroup

	wg.Add(5)
	go sumarRango(1, 20, canal, &wg)
	go sumarRango(21, 40, canal, &wg)
	go sumarRango(41, 60, canal, &wg)
	go sumarRango(61, 80, canal, &wg)
	go sumarRango(81, 100, canal, &wg)

	wg.Wait()
	close(canal)

	total := 0
	for r := range canal {
		total += r
	}
	fmt.Println("Total 1..100 =", total) // debe ser 5050
}
```

Extra: corré `go run -race` y probá agregar un `sleep` aleatorio para ver que el orden de llegada no importa.

## Criterio para pasar a la Fase 7

- [ ] Lancé goroutines con `go` y las sincronicé con `WaitGroup`
- [ ] Comuniqué goroutines con un `chan` (envío y recepción)
- [ ] Usé `select` con un `time.After` como timeout
- [ ] Corrí `go run -race` y sé qué detecta
- [ ] Entiendo (con palabras simples) por qué un `Mutex` protege datos compartidos
