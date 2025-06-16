// Part 2, Lesson 11: Project: Simple CLI Quiz
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file is your first project! It combines concepts like structs, slices,
// loops, and user input to create a functioning command-line quiz application.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

PROJECT BRIEF: SIMPLE COMMAND-LINE QUIZ

Welcome to your first project! The goal is to build a simple, interactive quiz
that runs in your computer's terminal. This project will test your understanding of
the foundational concepts we've covered so far:
-   STRUCTS: To create a custom `Question` type.
-   SLICES: To hold the list of all the questions in our quiz.
-   CONTROL FLOW: A `for` loop to go through the questions and `if/else` to check answers.
-   FUNCTIONS: `main` will be our primary function to run the quiz.
-   PACKAGES: We'll use `fmt` for printing and getting input, `bufio` and `os` for
    more robust input handling, `strings` to clean it up, and `strconv` to
    convert the user's text input into a number.

By the end of this lesson, you will have built a complete, runnable Go program
that feels like a real application.
*/

package main

import (
	"bufio"   // For buffered I/O, allowing us to read lines of text.
	"fmt"     // For formatted input/output (printing and scanning).
	"os"      // Provides a way to access operating system functionality, like standard input.
	"strconv" // Provides string conversion functions (e.g., string to integer).
	"strings" // Provides string manipulation functions.
)

// --- Part 1: Defining the Data Structure ---

// First, we need a blueprint for what a single quiz question looks like.
// A `struct` is perfect for this. Each question will have the question text,
// a list of possible answers, and the index of the correct answer.
type Question struct {
	Text               string   // The question itself.
	Options            []string // A slice of possible answers (strings).
	CorrectAnswerIndex int      // The 0-based index of the correct answer in the Options slice.
}

func main() {
	// --- Part 2: Creating the Quiz Content ---

	// Now, let's create the actual content for our quiz using the `Question` struct.
	// We'll store all our questions in a slice of `Question`.
	questions := []Question{
		{
			Text:               "What keyword is used to declare a variable that can be changed?",
			Options:            []string{"const", "var", "let", "def"},
			CorrectAnswerIndex: 1,
		},
		{
			Text:               "Which is NOT a built-in type in Go?",
			Options:            []string{"int", "float64", "string", "long"},
			CorrectAnswerIndex: 3,
		},
		{
			Text:               "What is the name of Go's unified looping construct?",
			Options:            []string{"while", "loop", "for", "foreach"},
			CorrectAnswerIndex: 2,
		},
		{
			Text:               "How does Go handle errors?",
			Options:            []string{"By returning an error value", "With try/catch blocks", "By crashing the program", "By ignoring them"},
			CorrectAnswerIndex: 0,
		},
	}

	// --- Part 3: Running the Quiz ---

	score := 0
	// `bufio.NewReader` creates a new reader that gets input from the standard input (the keyboard).
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("==============================")
	fmt.Println("  Welcome to the Go Quiz!   ")
	fmt.Println("==============================")
	fmt.Println("Answer the following questions.")

	// We use a `for range` loop to iterate over our slice of questions.
	// `i` will be the index (0, 1, 2, ...) and `q` will be the question itself.
	for i, q := range questions {
		fmt.Printf("\n--- Question #%d ---\n", i+1)
		fmt.Println(q.Text)

		// Loop through the options for the current question and print them.
		for j, option := range q.Options {
			fmt.Printf("  %d) %s\n", j+1, option)
		}

		// Prompt for and get user input.
		fmt.Print("Your answer (1-4): ")
		userInput, _ := reader.ReadString('\n')  // Read until the user hits Enter.
		userInput = strings.TrimSpace(userInput) // Remove whitespace and newline characters.
		answer, err := strconv.Atoi(userInput)   // Convert the string input to an integer.

		// Check if the input was valid.
		// Was it a number? Was it in the valid range of options?
		if err != nil || answer < 1 || answer > len(q.Options) {
			fmt.Println("Invalid input. The correct answer was:", q.Options[q.CorrectAnswerIndex])
			continue // Skip to the next question.
		}

		// Check if the answer was correct.
		// We subtract 1 from the user's answer because our slice is 0-indexed.
		if (answer - 1) == q.CorrectAnswerIndex {
			fmt.Println("Correct! Great job.")
			score++
		} else {
			fmt.Println("Wrong. The correct answer was:", q.Options[q.CorrectAnswerIndex])
		}
	}

	// --- Part 4: Displaying the Final Score ---
	fmt.Println("\n==============================")
	fmt.Println("         Quiz Complete!       ")
	fmt.Println("==============================")
	fmt.Printf("Your final score is: %d out of %d\n", score, len(questions))
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

HOW TO RUN THIS CODE:

1.  Open a terminal or command prompt.
2.  Navigate to the directory where you saved this file.
3.  Use the `go run` command to compile and execute the file:
    `go run 11_project_simple_cli_quiz.go`

You have now built a complete, interactive application. Try adding your own
questions to the `questions` slice to customize the quiz!
*/
