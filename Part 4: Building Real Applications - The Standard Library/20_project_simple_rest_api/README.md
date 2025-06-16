# Project: Simple REST API

This project is a multi-file REST API for a "contact book" that handles GET and POST requests, serving data in JSON format. It is built using only Go's standard library.

## Concepts Demonstrated
*   **`net/http`**: Setting up a web server and routing requests.
*   **`encoding/json`**: Marshaling Go structs into JSON for responses and unmarshaling JSON from requests into Go structs.
*   **Project Structure**: Separating project files (this `README.md` and `main.go`).
*   **RESTful Principles**: Handling different HTTP methods (`GET`, `POST`) on a single resource URL (`/contacts`).
*   **Concurrency Safety**: Using a `sync.Mutex` to protect shared data (the in-memory map).

## File Structure
Use code with caution.
Markdown
20_project_simple_rest_api/
├── main.go # The main application logic and server setup.
└── README.md # This file - instructions and explanations.

## How to Run This Server

1.  **Navigate to the project directory** in your terminal:
    ```sh
    # Example path, adjust if necessary
    cd Part4_Building_Real_Applications/20_project_simple_rest_api/
    ```

2.  **Run the application** using the `go run` command:
    ```sh
    go run main.go
    ```
    The terminal will print `Server starting on port :8080...`. The server is now running and listening for requests.

## How to Use the API

You will need a tool like `curl` or Postman to interact with this API. The following examples use `curl`.

### 1. Get All Contacts (GET)

Open a **new terminal window** (while the server is running in the other) and run the following command:

```sh
curl http://localhost:8080/contacts
Use code with caution.
Expected Response:
You should see a JSON array containing the initial contacts stored in the server's memory.

[{"id":1,"name":"Alice","email":"alice@example.com","phone":"111-111-1111"},{"id":2,"name":"Bob","email":"bob@example.com","phone":"222-222-2222"}]
Use code with caution.
Json
2. Add a New Contact (POST)
Run the following curl command to send a POST request with a JSON body. This will create a new contact.

curl -X POST -H "Content-Type: application/json" -d '{"name":"Charlie","email":"charlie@example.com","phone":"333-333-3333"}' http://localhost:8080/contacts
Use code with caution.
Sh
Expected Response:
The server will respond with the contact that was just created, now including its new server-assigned id.

{"id":3,"name":"Charlie","email":"charlie@example.com","phone":"333-333-3333"}
Use code with caution.
Json
If you run the GET command again, you will now see Charlie in the list.

curl http://localhost:8080/contacts```
```json
[{"id":1,"name":"Alice","email":"alice@example.com","phone":"111-111-1111"},{"id":2,"name":"Bob","email":"bob@example.com","phone":"222-222-2222"},{"id":3,"name":"Charlie","email":"charlie@example.com","phone":"333-333-3333"}]