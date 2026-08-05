# Fase 7 — Backend: HTTP y JSON 🌐

> **Objetivo:** Servir una API REST con la librería estándar. Sin frameworks — para entender cómo funciona todo por dentro.

## Lo que aprenderemos

- **`net/http`**: crear un servidor con `http.ListenAndServe`
  ```go
  http.HandleFunc("/", manejador)
  http.ListenAndServe(":8080", nil)
  ```
- **Handlers**: funciones que reciben `http.ResponseWriter` y `*http.Request`
  ```go
  func manejador(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Content-Type", "application/json")
      w.Write([]byte(`{"mensaje":"hola"}`))
  }
  ```
- **Verbos HTTP** en Go: `r.Method` — distinguir GET, POST, PUT, DELETE
- Parámetros de ruta y query: `r.URL.Path`, `r.URL.Query()`, `r.PathValue("id")` (Go 1.22+ con el mux mejorado: `/usuarios/{id}`)
- **`encoding/json`**: la librería de serialización
  - `json.Marshal` / `json.NewEncoder(w).Encode(...)` → Go → JSON
  - `json.Unmarshal` / `json.NewDecoder(r.Body).Decode(...)` → JSON → Go
  - **Tags de struct**: `json:"nombre"` para controlar el nombre en JSON
- **Middleware**: funciones que envuelven handlers (logging, CORS, auth)
  ```go
  func loggingMiddleware(siguiente http.HandlerFunc) http.HandlerFunc {
      return func(w http.ResponseWriter, r *http.Request) {
          fmt.Printf("%s %s\n", r.Method, r.URL.Path)
          siguiente(w, r)
      }
  }
  ```
- Códigos de estado: `w.WriteHeader(http.StatusNotFound)`, `http.StatusOK`, etc.
- **`http.Server`** con timeouts bien configurados (ReadTimeout, WriteTimeout) — detalle de profesional

## Palabras reservadas / paquetes de esta fase

| Elemento | Uso |
|---|---|
| `net/http` | Servidor HTTP completo de la stdlib |
| `encoding/json` | Serialización/deserialización JSON |
| Tags de struct | `json:"campo"` — controlan el JSON |
| `io` | Lectura de cuerpos (`io.ReadAll`) |

## Ejercicio 🌐

**API REST de una lista en memoria** — CRUD completo sin base de datos:

- `GET /items` → devuelve todos los items (JSON)
- `GET /items/{id}` → devuelve un item
- `POST /items` → crea un item (recibe JSON en el body)
- `PUT /items/{id}` → actualiza un item
- `DELETE /items/{id}` → elimina un item

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Item struct {
	ID   int    `json:"id"`
	Nombre string `json:"nombre"`
}

var items = []Item{{ID: 1, Nombre: "primer item"}}
var nextID = 2

func listaItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func crearItem(w http.ResponseWriter, r *http.Request) {
	var nuevo Item
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nuevo.ID = nextID
	nextID++
	items = append(items, nuevo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevo)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /items", listaItems)
	mux.HandleFunc("POST /items", crearItem)

	fmt.Println("Servidor en :8080")
	http.ListenAndServe(":8080", mux)
}
```

> 💡 Fijate el mux mejorado de Go 1.22+: `"GET /items"` ya distingue método y ruta.

## Criterio para pasar a la Fase 8

- [ ] Mi API responde los 5 verbos HTTP correctamente (probado con curl o Postman)
- [ ] Uso tags `json:"..."` en mis structs
- [ ] Devuelvo códigos de estado correctos (201 al crear, 404 si no existe, 400 si body inválido)
- [ ] Escribí UN middleware y lo apliqué
- [ ] Explico la diferencia entre `Marshal` y `NewEncoder`
