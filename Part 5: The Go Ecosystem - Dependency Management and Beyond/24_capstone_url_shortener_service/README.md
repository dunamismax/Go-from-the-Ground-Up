# Lesson 24: Final Capstone - URL Shortener Service

Congratulations on reaching the final project! This capstone is where we bring everything you've learned together to build a complete, real-world microservice: a URL Shortener.

This service will take a long URL and return a short, unique code. When a user navigates to the short URL, they will be redirected to the original long URL. This project demonstrates the power and simplicity of Go for building high-performance network services.

---

## ✨ Concepts You Will Master

*   **Professional Project Structure:** Using the `cmd/` and `internal/` layout for clean separation of concerns.
*   **Web Services with `net/http`:** Building a robust HTTP server to handle API requests.
*   **RESTful API Design:** Creating endpoints for creating (`POST`) and retrieving (`GET`) resources.
*   **JSON Data Handling:** Encoding and decoding JSON for API communication.
*   **Concurrency-Safe State:** Using a `sync.Mutex` to safely handle in-memory data storage across multiple requests.
*   **Go Modules:** Organizing your code into reusable internal packages.
*   **Building and Running an Application:** The complete workflow from `go mod init` to a running executable.

---

## 🏗️ Project Architecture

We will use the professional project layout taught in the previous lesson. This separates our code based on its purpose, making it clean and maintainable.
Use code with caution.
Markdown
/24_capstone_url_shortener_service/
├── /cmd/
│ └── /urlshortener/
│ └── main.go <-- The application's entry point. Initializes and starts everything.
├── /internal/
│ ├── /handler/ <-- Contains the HTTP handlers that process requests.
│ │ └── routes.go
│ └── /store/ <-- Manages the in-memory data storage safely.
│ └── store.go
└── README.md <-- This lesson and instruction file.

*   `store.go`: Will define a `URLStore` struct responsible for saving and retrieving URLs. It will use a map and a mutex to be thread-safe.
*   `routes.go`: Will define the functions that handle the HTTP logic, such as parsing JSON from a request and redirecting users.
*   `main.go`: Will be the "glue". It will create the store, set up the HTTP routes with their handlers, and start the web server.

---

## 🚀 How to Build and Run the Service

Follow these steps precisely from within the `24_capstone_url_shortener_service/` directory.

1.  **Create the Directory Structure:**
    If you haven't already, create the necessary folders.
    ```sh
    # Create nested directories in one command
    mkdir -p cmd/urlshortener internal/handler internal/store
    ```

2.  **Copy the Code:**
    Place the code from the blocks below into their corresponding files.

3.  **Initialize the Go Module:**
    In your terminal, navigate to the root of this project (`24_capstone_url_shortener_service/`) and initialize the Go module. Let's call our module `urlshortener`.

    ```sh
    # Make sure you are in the '24_capstone_url_shortener_service' directory
    go mod init urlshortener
    ```
    This creates the `go.mod` file, which allows our `main` package to import the `internal` packages correctly.

4.  **Build the Application:**
    From the same root directory, build the executable.

    ```sh
    go build -o urlshortener-service ./cmd/urlshortener
    ```
    The `-o urlshortener-service` flag tells the compiler to name the output executable `urlshortener-service`.

5.  **Run the Service:**
    Execute the file you just built.

    ```sh
    ./urlshortener-service
    ```
    The server is now running on `http://localhost:8080`.

---

##  interagindo com a API

You will need a tool like `curl` to interact with the API. Open a new terminal window while the server is running.

### 1. Shorten a URL (POST)

Send a POST request with a JSON body containing the URL you want to shorten.

```sh
curl -X POST -H "Content-Type: application/json" -d '{"url":"https://www.google.com/search?q=golang"}' http://localhost:8080/shorten
Use code with caution.
Expected Response:
The server will respond with a JSON object containing the short code.

{"short_code":"XXXXXX"}
Use code with caution.
Json
(The XXXXXX will be a random 6-character string).

2. Use the Shortened URL (GET)
Now, use curl with the -v flag to see the redirect in action. Replace XXXXXX with the code you received.

curl -v http://localhost:8080/XXXXXX
Use code with caution.
Sh
Expected Response:
You will see a 302 Found status code and a Location header pointing to the original URL. A web browser would automatically redirect to this location.

*   Trying 127.0.0.1:8080...
> GET /XXXXXX HTTP/1.1
>
< HTTP/1.1 302 Found
< Location: https://www.google.com/search?q=golang
< Date: ...
< Content-Length: 0