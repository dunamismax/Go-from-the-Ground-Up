# Go From The Ground Up

<p align="left">
  <b>A complete, open-source curriculum that teaches you Go by building a versatile portfolio of high-performance backend services and command-line tools.</b>
</p>
<p align="left">
  This course uses a unique, hands-on teaching method: <b>the lesson is in the code</b>. You'll learn every concept from structured comments inside a single, runnable Go file for each topic.
</p>
<p align="center">
  <a href="https://go.dev/"><img src="https://img.shields.io/badge/Language-Go-blue.svg" alt="Go"></a>
  <a href="https://go.dev/doc/toolchain"><img src="https://img.shields.io/badge/Tooling-Go%20Modules-green.svg" alt="Go Modules"></a>
  <a href="https://github.com/dunamismax/Go-from-the-Ground-Up/blob/main/LICENSE"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License: MIT"></a>
  <a href="https://github.com/dunamismax/Go-from-the-Ground-Up/pulls"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square" alt="PRs Welcome"></a>
  <a href="https://github.com/dunamismax/Go-from-the-Ground-Up/stargazers"><img src="https://img.shields.io/github/stars/dunamismax/Go-from-the-Ground-Up?style=social" alt="GitHub Stars"></a>
</p>

---

Welcome to the ultimate learn-by-doing course for the Go programming language! This isn't just a tutorial; it's a complete educational journey designed to build deep, practical knowledge from absolute zero. We start with the basics—assuming you've never written a line of code—and guide you to building an impressive portfolio of projects, including a concurrent website checker, a JSON-based REST API, and a full URL shortener microservice.

> Learning Go is about more than just syntax; it's about learning to write simple, reliable, and efficient software. This course teaches you to write clean, idiomatic Go, empowering you to build high-performance applications with confidence and clarity.

## ✨ Why This Approach?

*   📖 **Learn Directly In The Code**: Forget switching between a book and your editor. Every lesson is taught directly within the comments of a single, runnable Go file. The code *is* the textbook.
*   👨‍💻 **True Beginners Welcome**: This course has zero prerequisites. We'll guide you through your first `fmt.Println("Hello, World!")` before gradually introducing the powerful concepts like concurrency and interfaces that make Go a world-class language for backend and systems development.
*   🚀 **From "Hello, World!" to Concurrent Microservices**: The curriculum is carefully structured to build your skills layer by layer. You'll master Go fundamentals, then use them to build CLI tools, handle JSON, and build high-performance, concurrent web services using only the standard library.
*   🛠️ **Build a Versatile Portfolio**: The entire course is oriented around practical projects. Every concept, from `structs` to `goroutines`, is a building block for real-world applications like a command-line quiz, a concurrent website checker, and a complete REST API.
*   💪 **Master Professional Practices**: We don't just teach syntax. You will gain true confidence with Go's unique approach to Object-Oriented design with interfaces, its world-class concurrency model, built-in testing, and professional project structure with Go modules.
*   🌍 **Open Source & Community Driven**: This curriculum is for everyone. Contributions, suggestions for improvement, bug fixes, and new project ideas are highly encouraged!

---

## 💻 Tech Stack & Prerequisites

You don't need any programming knowledge to start, but you will need a few standard, free tools.

*   **Go** (the latest version is recommended).
*   **The Go Toolchain** (this is included with your Go installation and provides `go run`, `go build`, `go mod`, etc.).
*   **Git** for cloning this repository to your computer.
*   A good Text Editor or IDE (**Visual Studio Code** with the official Go extension is a fantastic, free choice).

---

## 🚀 How to Use This Course

Each lesson folder contains a single `.go` file (or multiple files for larger projects). This file is both the complete, runnable program and the full lesson text.

1.  **Read the Lesson:** Navigate to a lesson folder (e.g., `Part1_The_Foundations/01_HelloWorld/`) and open the `1_hello_world.go` file. Read the comments from top to bottom to understand the concepts.

2.  **Run the Code:** To see the lesson's concepts in action, you'll need to run the Go script from your terminal using the built-in `go run` command.

    First, clone the repository (you only need to do this once):
    ```sh
    git clone https://github.com/dunamismax/Go-from-the-Ground-Up.git
    cd Go-from-the-Ground-Up
    ```

    Then, for each lesson, navigate to its folder and use the `go run` command.
    ```sh
    # Example for the first lesson
    cd Part1_The_Foundations/01_HelloWorld/

    # Run the Go script
    go run 1_hello_world.go
    ```
    *Note: Later projects will require initializing a Go module. In those cases, specific instructions will be provided in the project's README.*

---

## 📚 The Curriculum

The curriculum is divided into five parts, taking you from a blank text file to a proficient Go developer capable of building high-performance backend applications.

<details>
<summary><strong>Part 1: The Foundations - Core Go Syntax</strong></summary>
<br>
<i>(Focus: The basic building blocks of Go, including its unique approach to types, functions, and errors.)</i>

| Lesson                           | Key Concepts                                              | Description                                                              |
| -------------------------------- | --------------------------------------------------------- | ------------------------------------------------------------------------ |
| `1_hello_world.go`               | `package main`, `import`, `fmt.Println`, `go run`         | The essential first step: running your very first Go program.            |
| `2_variables_and_types.go`       | `var`, `:=`, `string`, `int`, `float64`, `bool`           | Learn to store, manage, and work with statically typed information.      |
| `3_packages_and_imports.go`      | `math`, `strings`, standard library                       | Explore Go's powerful standard library to perform common tasks.          |
| `4_functions.go`                 | `func`, typed arguments, multiple `return` values         | Organize code into reusable blocks and return multiple results with ease.|
| `5_control_flow.go`              | `if/else`, `switch`, the unified `for` loop               | Give your program a brain by letting it execute code based on conditions.|
| `6_pointers.go`                  | `*`, `&`, memory addresses, value sharing                 | A gentle introduction to how Go manages memory and shares data.        |
| `7_error_handling.go`            | `(value, error)` return pattern                           | Master the idiomatic, robust way Go handles errors without exceptions.   |

</details>

<details>
<summary><strong>Part 2: Structuring Data - Go's Type System</strong></summary>
<br>
<i>(Focus: Moving beyond primitive types to create and manage complex data structures.)</i>

| Lesson                           | Key Concepts                                              | Description                                                              |
| -------------------------------- | --------------------------------------------------------- | ------------------------------------------------------------------------ |
| `8_arrays_and_slices.go`         | Fixed-size arrays vs. dynamic slices, `len()`, `cap()`    | Learn the difference between fixed arrays and Go's powerful slices.      |
| `9_maps.go`                      | `make()`, key-value pairs, checking for existence         | Master Go's built-in key-value data structure for fast lookups.        |
| `10_structs.go`                  | `type`, custom data structures                            | Create your own custom data types by composing fields together.          |
| `11_project_simple_cli_quiz.go`  | **Project:** Structs, slices, `fmt.Scanln`                | Build your first interactive command-line application: a simple quiz.    |

</details>

<details>
<summary><strong>Part 3: The Go Way - Methods, Interfaces, and Concurrency</strong></summary>
<br>
<i>(Focus: Mastering the core concepts that make Go unique, powerful, and fun to use.)</i>

| Lesson                                     | Key Concepts                                              | Description                                                              |
| ------------------------------------------ | --------------------------------------------------------- | ------------------------------------------------------------------------ |
| `12_methods_on_structs.go`                 | Pointer vs. value receivers                               | Attach functions (methods) directly to your custom data types.           |
| `13_interfaces.go`                         | Implicit satisfaction, polymorphism                       | Unlock the power of Go's flexible and decoupled design philosophy.       |
| `14_goroutines.go`                         | The `go` keyword, lightweight concurrency                 | Learn to run functions concurrently with incredible ease.                |
| `15_channels.go`                           | `make(chan)`, `<-`, thread-safe communication             | Communicate safely between goroutines, a cornerstone of Go concurrency.|
| `16_project_concurrent_web_checker.go`     | **Project:** Goroutines, channels, standard library         | Build a tool that checks a list of websites concurrently and efficiently.|

</details>

<details>
<summary><strong>Part 4: Building Real Applications - The Standard Library</strong></summary>
<br>
<i>(Focus: Leveraging Go's exceptional standard library to build practical, real-world tools.)</i>

| Lesson                           | Key Concepts                                              | Description                                                              |
| -------------------------------- | --------------------------------------------------------- | ------------------------------------------------------------------------ |
| `17_working_with_files.go`       | `os`, `io`, `bufio` packages                              | Persist data beyond program execution by reading from and writing to files.|
| `18_handling_json.go`            | `encoding/json`, struct tags, `Marshal`/`Unmarshal`       | Master JSON, the universal language of web APIs and data exchange.       |
| `19_intro_to_net_http.go`        | `http.ListenAndServe`, `http.HandleFunc`                  | Build a production-ready web server using *only* Go's standard library.|
| `20_project_simple_rest_api/`    | **Project:** `net/http`, JSON, structs, project structure | Build a multi-file REST API for a contact book with GET and POST routes. |

</details>

<details>
<summary><strong>Part 5: The Go Ecosystem - Dependency Management and Beyond</strong></summary>
<br>
<i>(Focus: Professional practices for creating scalable, testable, and distributable applications.)</i>

| Lesson                                     | Key Concepts                                              | Description                                                              |
| ------------------------------------------ | --------------------------------------------------------- | ------------------------------------------------------------------------ |
| `21_go_modules_and_dependencies.go`| `go mod init`, `go.mod`, `go get`                         | Learn the professional standard for managing project dependencies.       |
| `22_testing_in_go.go`                | `testing` package, `go test`, table-driven tests          | Write unit tests with Go's simple, built-in testing framework.         |
| `23_structuring_a_larger_project/`   | `internal/`, `cmd/`, project layout                       | Learn how to organize code for large, scalable, and maintainable apps.   |
| `24_capstone_url_shortener_service/` | **Final Capstone Project**                                | A complete URL shortener microservice using a web server and a map store.|
| `25_what_next.go`                    | Guidance document                                         | A commented guide to further learning in gRPC, DevOps, Gin, and more.  |

</details>

---

## ⭐ Show Your Support

If this course helps you become a better developer, please **give this repository a star!** It helps the project reach more aspiring programmers and encourages us to keep creating great, free educational content.

## 🤝 Connect & Contribute

This project is for the community. Have an idea for a new feature, a better way to explain a concept, or find a bug? Feel free to [open an issue](https://github.com/dunamismax/Go-from-the-Ground-Up/issues) or [submit a pull request](https://github.com/dunamismax/Go-from-the-Ground-Up/pulls)!

Connect with the author, **dunamismax**, on:

*   **Twitter:** [@dunamismax](https://twitter.com/dunamismax)
*   **Bluesky:** [@dunamismax.bsky.social](https://bsky.app/profile/dunamismax.bsky.social)
*   **Reddit:** [u/dunamismax](https://www.reddit.com/user/dunamismax)
*   **Discord:** `dunamismax`
*   **Signal:** `dunamismax.66`

## 📜 License

This project is licensed under the **MIT License**. See the `LICENSE` file for details.