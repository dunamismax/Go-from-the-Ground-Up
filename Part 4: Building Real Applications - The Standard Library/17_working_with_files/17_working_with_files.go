// Part 4, Lesson 17: Working With Files
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file serves as the lesson and demonstration for file I/O (Input/Output).
// It explains the core concepts through structured comments and
// provides a runnable example of their implementation.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

Welcome to Part 4! So far, our programs have worked only with data in memory.
When the program finishes, all the information (variables, structs, etc.) is
lost. To create truly useful applications, we need to save data permanently.
This is called PERSISTENCE.

The most common way to persist data is by reading from and writing to FILES.
Whether it's saving user progress, loading a configuration file, or exporting
data, file I/O is a fundamental skill.

In Go, the `os` package provides simple, powerful functions for interacting
with the operating system, including its file system. We will explore the
three most common file operations: writing, reading, and deleting.
*/

package main

import (
	"fmt"
	"os" // The `os` package is our gateway to file system operations.
)

// We will create a file, read from it, and then clean it up.
// It's good practice to define filenames as constants to avoid typos.
const filename = "lesson.txt"

func main() {
	// --- Part 1: Writing to a File ---

	// Writing to a file often means replacing its contents entirely.
	// The `os.WriteFile` function is the simplest way to do this.
	// It handles opening the file, writing the data, and closing the file for us.
	//
	// `os.WriteFile` requires three arguments:
	// 1. FILENAME: A string with the name of the file (e.g., "lesson.txt").
	// 2. DATA: The content to write, as a slice of bytes (`[]byte`).
	// 3. PERMISSIONS: A number that tells the OS who can read, write, or execute the file.
	//
	// WHAT IS A BYTE SLICE?
	// Files fundamentally store BYTES, not text. A byte is a standard unit of digital
	// information. To write a string to a file, we must first convert it into a
	// slice of bytes. We do this with a simple type conversion: `[]byte("my string")`.
	//
	// WHAT ARE PERMISSIONS?
	// `0644` is a standard file permission setting. It means:
	// - The owner of the file can READ and WRITE (6).
	// - Users in the same group can READ (4).
	// - Everyone else can READ (4).
	// This is a very common and safe default for text files. [19, 20]

	content := "Hello from Go!\nThis is a line written to a file."
	// Convert the string to a byte slice for writing.
	data := []byte(content)

	fmt.Println("Attempting to write to file:", filename)

	// `os.WriteFile` will create the file if it doesn't exist, or completely
	// overwrite it if it does. [2, 10, 15]
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		// If something goes wrong (e.g., we don't have permission to write here),
		// the program will print the error and exit.
		fmt.Println("Error writing file:", err)
		return // Exit main if we can't proceed.
	}

	fmt.Println("File written successfully.")

	// --- Part 2: Reading from a File ---

	// Now that we've written the file, let's read its contents back.
	// The `os.ReadFile` function is the counterpart to `os.WriteFile`.
	// It handles opening, reading the entire content, and closing the file.
	//
	// It takes one argument: the FILENAME string.
	// It returns two values: a `[]byte` slice with the file's content and an `error`. [1, 7, 8]

	fmt.Println("\nAttempting to read from file:", filename)

	readData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// The data is read as a `[]byte` slice. To print it as text, we must
	// convert it back to a `string`.
	readContent := string(readData)

	fmt.Println("File read successfully. Content:")
	fmt.Println("---")
	fmt.Println(readContent)
	fmt.Println("---")

	// --- Part 3: Cleaning Up ---

	// In many programs, you create temporary files that should be deleted when
	// you are done with them.
	// The `os.Remove` function deletes a file. [5, 14, 16]

	fmt.Println("\nAttempting to delete file:", filename)
	err = os.Remove(filename)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}

	fmt.Println("File deleted successfully.")
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

HOW TO RUN THIS CODE:

1.  Open a terminal or command prompt.
2.  Navigate to the directory where you saved this file.
3.  Use the `go run` command to compile and execute the file:
    `go run 17_working_with_files.go`

When you run this program, it will create a file named `lesson.txt` in the same
directory, write to it, read it back to the console, and then delete it.
You probably won't even see the file appear unless you are watching the folder
very closely!
*/
