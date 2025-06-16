# Capstone Project: URL Shortener Microservice

This is it—the final capstone project for **Go From The Ground Up**! This project is the culmination of everything you've learned. We will build a complete, production-ready URL shortener microservice from scratch, using only Go's standard library.

This isn't just a simple script; it's a well-architected application that demonstrates professional Go development practices you can apply directly to your own projects.

## ✨ Key Concepts Applied

This project synthesizes core concepts from the entire course into a single, cohesive application:

- ✅ **Professional Project Structure**: We use the standard `cmd/` and `internal/` directory layout to organize a scalable and maintainable application.
- ✅ **Go Modules**: The project is structured as a proper Go module, with import paths based on the repository URL, enabling the Go toolchain to manage dependencies and builds correctly.
- ✅ **Clean API Design**: A simple, intuitive JSON API for creating short links (`POST /api/shorten`) and a clean redirect mechanism (`GET /{shortCode}`).
- ✅ **Dependency Injection**: We create our dependencies (like the logger and data store) in `main` and pass them into our handlers, making our code decoupled and easy to test.
- ✅ **Concurrency Safety**: Our in-memory store uses a `sync.RWMutex` to handle many concurrent reads (redirects) and writes (creations) safely, preventing race conditions.
- ✅ **Structured Logging**: We set up a simple but effective logger to provide insight into the server's operations.
- ✅ **Standard Library Power**: We build the entire web server, router, and logic using only Go's powerful standard library packages like `net/http`, `encoding/json`, `log`, and `sync`.

## 📁 Go Modules & Project Structure

A critical part of this lesson is understanding how professional Go projects are organized. Our project is defined as a Go module in the `go.mod` file at the repository's root. This allows Go's tooling to correctly resolve the `import` paths for our internal packages.

The project uses a standard layout to separate concerns:
content_copy
download
Use code with caution.
Markdown
.
├── go.mod <-- Defines the module path for the whole repository
└── Part5_The_Go_Ecosystem
└── 24_capstone_url_shortener_service/
├── cmd/
│ └── urlshortener/
│ └── main.go # Entry point: configuration and wiring
├── internal/
│ ├── handler/
│ │ └── handler.go # API Layer: HTTP request/response logic
│ ├── shortener/
│ │ └── shortener.go # Business Logic: Generating short codes
│ └── store/
│ └── store.go # Data Layer: Storing and retrieving data
└── README.md

## 🚀 How to Run the Service

**Prerequisites**: You must have Go installed and have cloned the `Go-from-the-Ground-Up` repository. A `go.mod` file must exist at the root of the repository.

1.  **Navigate to the Command Directory**:
    In Go, we run applications from their `main` package, which is typically located in a `cmd/` directory. Open your terminal and `cd` into this project's `cmd` directory.

    ```sh
    # This path must match your local file system structure.
    cd Part5_The_Go_Ecosystem/24_capstone_url_shortener_service/cmd/urlshortener/
    ```

2.  **Run the Server**:
    Use the `go run .` command. This tells Go to compile and run the `main` package in the current directory. It will automatically find the project's other internal packages using the module information from `go.mod`.

    ```sh
    go run .
    ```

    You should see the server's startup log message:
    `[INFO] Server starting on http://localhost:8080`

## ⚙️ API Reference

You can interact with the running API using a tool like `curl` or an API client like Postman.

| Endpoint       | Method | Description                                                                      | Example `curl` Command                                                                                                                  |
| :------------- | :----- | :------------------------------------------------------------------------------- | :-------------------------------------------------------------------------------------------------------------------------------------- |
| `/api/shorten` | `POST` | Takes a long URL and returns its shortened version. This endpoint is idempotent. | `curl -i -X POST -H "Content-Type: application/json" -d '{"url": "https://go.dev/doc/effective_go"}' http://localhost:8080/api/shorten` |
| `/{shortCode}` | `GET`  | Redirects the browser to the original long URL associated with the short code.   | `curl -i -L http://localhost:8080/{shortCode}` (Replace `{shortCode}` with one you created)                                             |
| `/`            | `GET`  | Displays a simple welcome message for users who visit the root URL.              | `curl http://localhost:8080`                                                                                                            |

---

Congratulations on completing the capstone! You've built a robust, real-world application and are now well-equipped to build your own high-performance services in Go.
content_copy
download
Use code with caution.
