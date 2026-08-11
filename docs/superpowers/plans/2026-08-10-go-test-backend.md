# Go Test Backend Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Create a minimal Go HTTP API in `exercises/test-backend` that returns a JSON greeting from `GET /hola`, then compile and test its execution on the Linux shared hosting account.

**Architecture:** A single `main.go` will own the small smoke-test server. It will use Go's standard `net/http` package, read `PORT` when present, default to `8080`, and expose only `GET /hola`. Deployment is deliberately separated into local Linux-targeted compilation, FileZilla upload, SSH execution, and later domain routing verification.

**Tech Stack:** Go 1.26.5, `net/http`, `encoding/json`, Linux `amd64` binary, Namecheap cPanel/SSH.

## Global Constraints

- Keep the test API isolated in `exercises/test-backend`.
- Do not modify existing Node.js applications or production domains.
- Do not add a database until the HTTP/domain smoke test works.
- Compile the deployment binary with `GOOS=linux`, `GOARCH=amd64`, and `CGO_ENABLED=0`.
- Use `api.azacode.dev` only after confirming the current domain mapping is not serving production traffic.

---

### Task 1: Create the minimal HTTP API

**Files:**
- Create: `exercises/test-backend/main.go`

**Interfaces:**
- Produces `GET /hola` with `Content-Type: application/json` and body `{"mensaje":"Hola AZAC"}`.

- [ ] **Step 1: Create the server source file**

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type saludoResponse struct {
	Mensaje string `json:"mensaje"`
}

func holaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(saludoResponse{Mensaje: "Hola AZAC"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hola", holaHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("API escuchando en :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
```

- [ ] **Step 2: Format and verify locally**

Run from the repository root:

```bash
gofmt -w ./exercises/test-backend/main.go
go run ./exercises/test-backend
```

In another terminal, verify:

```bash
curl http://127.0.0.1:8080/hola
```

Expected response:

```json
{"mensaje":"Hola AZAC"}
```

- [ ] **Step 3: Stop the local server**

Press `Ctrl+C` in the terminal running the server.

### Task 2: Build the Linux deployment binary

**Files:**
- Create: local untracked artifact `go-backend-test` beside the project build command; do not commit the binary.

- [ ] **Step 1: Compile for the Namecheap host**

Run from the repository root in Bash:

```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o go-backend-test ./exercises/test-backend
```

- [ ] **Step 2: Verify the artifact exists**

```bash
ls -lh go-backend-test
```

### Task 3: Upload and execute the HTTP binary over SSH

**Files:**
- Create on hosting: `/home/azacexhp/go-tests/go-backend-test`

- [ ] **Step 1: Create the private test directory over SSH**

```bash
mkdir -p ~/go-tests
```

- [ ] **Step 2: Upload `go-backend-test` with FileZilla**

Upload it to:

```text
/home/azacexhp/go-tests/
```

- [ ] **Step 3: Grant execute permission**

```bash
chmod +x ~/go-tests/go-backend-test
```

- [ ] **Step 4: Start the API on the hosting account**

```bash
PORT=8765 nohup ~/go-tests/go-backend-test >~/go-tests/go-backend-test.log 2>&1 &
echo $!
```

- [ ] **Step 5: Test the API from inside SSH**

```bash
curl http://127.0.0.1:8765/hola
```

Expected response:

```json
{"mensaje":"Hola AZAC"}
```

### Task 4: Verify whether the domain can route to the Go process

**Files:**
- No repository files.

- [ ] **Step 1: Confirm the process is running**

```bash
ps -fu "$USER" | grep '[g]o-backend-test'
```

- [ ] **Step 2: Check the cPanel domain/application configuration**

Do not change the existing Node.js applications. Verify whether cPanel exposes a custom application/reverse-proxy configuration for `api.azacode.dev`.

- [ ] **Step 3: Test the domain only if a proxy mapping is available**

```bash
curl -i https://api.azacode.dev/hola
```

Expected result: HTTP 200 with the JSON greeting. If the local `curl` works but the domain fails, Go is running and the remaining problem is routing/proxy configuration, not compilation.

### Task 5: Keep deployment artifacts out of git

**Files:**
- Modify: `.gitignore` only if `go-backend-test` or other local binaries are not already ignored.

- [ ] **Step 1: Check repository status**

```bash
git status --short
```

- [ ] **Step 2: Remove only the local binary if it appears as an untracked file**

```bash
rm -f go-backend-test
```
