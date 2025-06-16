# Lesson 23: Structuring a Larger Project

Welcome to one of the most important lessons for building professional Go applications: project structure. As your programs grow, putting everything into a single `main.go` file becomes messy and difficult to manage. A well-organized project is easier to navigate, maintain, and scale.

This lesson introduces a standard, widely-accepted layout for Go applications using `cmd/` and `internal/` directories, matching the structure you have in this folder.

---

## A Standard Go Project Layout

Your current directory structure is a common and effective way to organize a Go application:
Use code with caution.
Markdown
/23_structuring_a_larger_project/
├── /cmd/
│ └── /myapp/
│ └── main.go <-- The main entry point for your application.
├── /internal/
│ └── /greeter/
│ └── greeter.go <-- A private "internal" package for business logic.
└── README.md <-- This lesson file.

### The `cmd/` Directory

The `cmd/` directory is where you place the `main` packages for your applications.
*   **Purpose:** It contains the entry points for the executable programs your project builds. A project might produce multiple binaries (e.g., a web server and a separate admin tool). Each would get its own subdirectory within `cmd/`.
*   **Role:** The code here should be minimal. Its main job is to "wire up" and run the core logic from other packages.

### The `internal/` Directory

The `internal/` directory is special, with a rule enforced by the Go compiler.
*   **Purpose:** It contains all the private, internal code for your application. This is where your core business logic lives.
*   **The Rule:** Packages inside `internal/` **cannot be imported by any external project**. They can only be imported by code inside the same parent module.
*   **Why it's useful:** This prevents other developers from depending on parts of your code that you don't want to support as a public API. It gives you the freedom to refactor and change your internal code without worrying about breaking someone else's project.

---

## How to Build and Run This Project

Follow these steps precisely from within the `23_structuring_a_larger_project` directory.

1.  **Initialize the Go Module:**
    Navigate your terminal to this project's root directory (`23_structuring_a_larger_project/`). To make the package imports work, you must first initialize a Go module. Let's name our module `structuredapp`.

    ```sh
    # Make sure you are inside the '23_structuring_a_larger_project' directory
    go mod init structuredapp
    ```
    This command creates a `go.mod` file, which is essential for tracking your project's packages and dependencies.

2.  **Build the Application:**
    Now, from the same root directory, you can build the executable. The `go build` command can take a path to a `main` package.

    ```sh
    go build ./cmd/myapp
    ```
    After running this, you will see a new executable file named `myapp` (or `myapp.exe` on Windows) appear in the `23_structuring_a_larger_project/` directory.

3.  **Run the Executable:**
    Execute the file you just built.

    ```sh
    # On Linux or macOS
    ./myapp

    # On Windows
    .\myapp.exe
    ```
    You should see the greeting printed to your console. Congratulations! You have successfully built and run a well-structured Go application.