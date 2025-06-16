# Project: A Structured REST API in Go

Welcome to the capstone project for Part 4 of "Go From The Ground Up"! This project demonstrates how to build a simple but complete REST API for a "contact book" using only Go's standard library. It moves beyond a single-file script to showcase a structured, multi-file layout that separates concerns—a critical practice for building maintainable, real-world applications.

You'll combine everything you've learned about `net/http`, JSON handling, structs, and concurrency to create a functional and professional-looking backend service.

## ✨ What You'll Learn

- ✅ **Full CRUD API**: Implement all four `Create`, `Read`, `Update`, and `Delete` operations.
- 🗂️ **Professional Project Structure**: Organize code by function (handlers, storage, routing) to create a clean and scalable layout.
- 🔒 **Concurrency Safety**: Use a `sync.Mutex` to protect shared data and prevent race conditions in a concurrent environment.
- 🛣️ **HTTP Routing Demystified**: See how a basic HTTP router works under the hood by building one from scratch.
- 📝 **Structured JSON Handling**: Create helper functions to send consistent and predictable JSON responses for both successes and errors.
- ✔️ **Request Validation**: Implement simple server-side validation for incoming data.

## 📁 Project Structure

This project is intentionally split into multiple files to teach the principle of **separation of concerns**.

- `main.go`: The application's entry point. Responsible for initializing the data store, setting up the router, and starting the HTTP server.
- `store.go`: The data layer. Defines the `Contact` struct and a `ContactStore` that manages all interactions with our in-memory map (the "database").
- `handlers.go`: The API logic layer. Contains the HTTP handler functions responsible for parsing requests, calling the store, and crafting JSON responses.
- `router.go`: The routing layer. Defines a basic HTTP router that directs incoming requests to the correct handler based on the URL path and HTTP method.

## 🚀 Running the API

1.  **Navigate to the Project Directory**:
    Open your terminal and `cd` into this project's directory.

    ```sh
    # Example path
    cd Part4_Building_Real_Applications/20_project_simple_rest_api/
    ```

2.  **Run the Server**:
    Use the `go run .` command. This automatically finds, compiles, and runs all `.go` files in the current directory.
    ```sh
    go run .
    ```
    You should see the message: `Server starting on port :8080...` The API is now live and ready to accept requests!

## ⚙️ API Reference

You can interact with the running API using a tool like `curl` or an API client like Postman.

_**Note**: In the examples below, `{id}` should be replaced with an actual contact ID (e.g., `1`)._

| Endpoint         | Method   | Description                                                                                       | Example `curl` Command                                                                                                                                                          |
| :--------------- | :------- | :------------------------------------------------------------------------------------------------ | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `/contacts`      | `GET`    | Retrieves a JSON array of all contacts.                                                           | `curl -i http://localhost:8080/contacts`                                                                                                                                        |
| `/contacts`      | `POST`   | Creates a new contact. The request body must be a valid `Contact` JSON object.                    | `curl -i -X POST -H "Content-Type: application/json" -d '{"name": "Charlie", "email": "charlie@example.com", "phone": "333-333-3333"}' http://localhost:8080/contacts`          |
| `/contacts/{id}` | `GET`    | Retrieves a single contact by its unique ID.                                                      | `curl -i http://localhost:8080/contacts/1`                                                                                                                                      |
| `/contacts/{id}` | `PUT`    | Updates an existing contact's details. The request body must be a complete `Contact` JSON object. | `curl -i -X PUT -H "Content-Type: application/json" -d '{"name": "Alice Smith", "email": "alice.smith@example.com", "phone": "111-555-4444"}' http://localhost:8080/contacts/1` |
| `/contacts/{id}` | `DELETE` | Deletes a contact by its unique ID. Returns a `204 No Content` status on success.                 | `curl -i -X DELETE http://localhost:8080/contacts/2`                                                                                                                            |
