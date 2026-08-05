# Fase 1 — Setup y Hola Mundo

> **Objetivo:** Tener Go instalado, el entorno configurado y el primer programa corriendo.

## Lo que aprenderemos

- Instalación de **Go 1.26.x** (versión estable actual) desde [go.dev/dl](https://go.dev/dl/)
- Configuración de **VS Code** con la extensión oficial de Go (gopls)
- Qué es un **paquete** (`package`) y por qué todo programa arranca con `package main`
- La función **`func main()`** — el punto de entrada de todo programa
- Cómo **importar** paquetes de la librería estándar (`import`)
- Los 3 comandos esenciales:
  - `go run` → compila y ejecuta (para desarrollo)
  - `go build` → compila a un binario (para producción)
  - `go vet` → analiza el código en busca de errores
- Qué es el **`go.mod`** (el archivo que define el módulo del proyecto)
- El formato de código con `gofmt` (Go te obliga a código limpio, y es genial)

## Palabras reservadas de esta fase

| Palabra | Uso |
|---|---|
| `package` | Declara el paquete al que pertenece el archivo |
| `import` | Trae paquetes de la librería estándar o externos |
| `func` | Declara una función |
| `var` / `const` | Declaran variables y constantes (primer vistazo) |

## Ejercicio 🖨️

Programa que imprima tu nombre, edad y ciudad, cada dato en su propia línea:

```go
package main

import "fmt"

func main() {
	fmt.Println("Nombre: Arturo")
	fmt.Println("Edad: 24")
	fmt.Println("Ciudad: Buenos Aires")
}
```

Luego modificarlo para guardar cada dato en una variable antes de imprimir.

## Criterio para pasar a la Fase 2

- [ ] `go version` responde con la versión 1.26.x
- [ ] Creé un módulo con `go mod init` y sé para qué sirve
- [ ] Corrí mi programa con `go run` y lo compilé con `go build`
- [ ] Sé explicar con mis palabras qué hace `package main` y `func main()`
