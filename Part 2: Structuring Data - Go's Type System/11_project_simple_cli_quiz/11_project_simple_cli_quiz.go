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

PROJECT BRIEF: SIMPLE COMMAND-LINE QUIZ (V2 - ENHANCED)

Welcome to your first project! We're building on the original idea to create a
more robust and engaging interactive quiz. This project will test your understanding
of foundational concepts and introduce some powerful new ideas.

NEW FEATURES IN THIS VERSION:
-   CODE STRUCTURE: We've broken the program into smaller, reusable functions.
    This makes the code cleaner, easier to read, and more maintainable.
-   QUESTION SHUFFLING: The questions are now shuffled at the start of each
    quiz, making it more fun to replay.
-   QUIZ TIMER: We'll time how long you take to complete the quiz!
-   ROBUST INPUT: Instead of failing on bad input, we'll now loop until the
    user enters a valid answer.

This project uses:
-   STRUCTS, SLICES, and CONTROL FLOW (for, if/else).
-   FUNCTIONS: To organize our code into logical blocks.
-   PACKAGES: `fmt`, `bufio`, `os`, `strings`, `strconv`, and now `math/rand`
    and `time` for our new features.
*/

package main

import (
	"bufio"     // For buffered I/O, allowing us to read lines of text.
	"fmt"       // For formatted input/output (printing and scanning).
	"math/rand" // For generating random numbers to shuffle our questions.
	"os"        // Provides access to OS functionality, like standard input.
	"strconv"   // Provides string conversion functions (e.g., string to integer).
	"strings"   // Provides string manipulation functions.
	"time"      // Provides time functionality, used for the timer and for seeding our random number generator.
)

// --- Part 1: Defining the Data Structures ---

// The Question struct is the blueprint for a single quiz question.
type Question struct {
	Text               string
	Options            []string
	CorrectAnswerIndex int
}

// The Quiz struct will hold all our questions and the user's score.
type Quiz struct {
	Questions []Question
	Score     int
}

// --- Part 2: Main Application Logic ---

// The main function is now the "orchestrator." Its job is to set up the
// application and run the primary logic from other functions. This makes it
// very clear what our program does at a high level.
func main() {
	// Create a new quiz instance with our questions.
	quiz := newQuiz()

	// Shuffle the questions to make it interesting on every run.
	quiz.shuffleQuestions()

	// Display the welcome message.
	displayWelcomeMessage(len(quiz.Questions))

	// Run the quiz and get the time it took.
	quiz.run()

	// Display the final results.
	quiz.displayFinalScore()
}

// --- Part 3: Helper & Logic Functions ---

// newQuiz creates and returns a new Quiz with a predefined list of questions.
func newQuiz() Quiz {
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
		{
			Text:               "What symbol is used to get the memory address of a variable?",
			Options:            []string{"*", "&", "$", "@"},
			CorrectAnswerIndex: 1,
		},
	}
	return Quiz{Questions: questions}
}

// shuffleQuestions randomizes the order of the questions in the quiz.
func (q *Quiz) shuffleQuestions() {
	// We need to "seed" the random number generator. If we don't, it will
	// produce the same "random" sequence every time. Using the current time
	// ensures a different sequence for each run.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(q.Questions), func(i, j int) {
		q.Questions[i], q.Questions[j] = q.Questions[j], q.Questions[i]
	})
}

// run contains the main quiz loop, a timer, and calls askQuestion for each question.
func (q *Quiz) run() {
	reader := bufio.NewReader(os.Stdin)
	startTime := time.Now() // Start the timer!

	// Iterate through the (now shuffled) questions.
	for i, question := range q.Questions {
		isCorrect := q.askQuestion(question, i+1, reader)
		if isCorrect {
			q.Score++
		}
	}
	duration := time.Since(startTime) // Stop the timer.
	q.displayFinalScore(duration)
}

// askQuestion handles the logic for a single question: displaying it, getting
// user input, and checking the answer. It returns true if the answer is correct.
func (q *Quiz) askQuestion(question Question, questionNumber int, reader *bufio.Reader) bool {
	fmt.Printf("\n--- Question #%d ---\n", questionNumber)
	fmt.Println(question.Text)

	// Display the options for the current question.
	for j, option := range question.Options {
		fmt.Printf("  %d) %s\n", j+1, option)
	}

	// This is our robust input loop. It will continue forever until the user
	// provides a valid answer.
	for {
		fmt.Printf("Your answer (1-%d): ", len(question.Options))
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)
		answer, err := strconv.Atoi(userInput)

		// Check if the input was valid (a number within the range of options).
		if err == nil && answer >= 1 && answer <= len(question.Options) {
			// Input is valid, now check if it's correct.
			if (answer - 1) == question.CorrectAnswerIndex {
				fmt.Println("Correct! Great job.")
				return true // Exit the function, returning true.
			} else {
				fmt.Println("Wrong. The correct answer was:", question.Options[question.CorrectAnswerIndex])
				return false // Exit the function, returning false.
			}
		}

		// If we reach here, the input was invalid. The loop will repeat.
		fmt.Println("Invalid input. Please enter a number corresponding to your choice.")
	}
}

// --- Part 4: UI and Display Functions ---

// displayWelcomeMessage prints the initial welcome banner.
func displayWelcomeMessage(totalQuestions int) {
	fmt.Println("=====================================")
	fmt.Println("      Welcome to the Go Quiz!      ")
	fmt.Println("=====================================")
	fmt.Printf("There are %d questions. Good luck!\n", totalQuestions)
}

// displayFinalScore prints the final results of the quiz.
func (q *Quiz) displayFinalScore(duration time.Duration) {
	fmt.Println("\n=====================================")
	fmt.Println("           Quiz Complete!          ")
	fmt.Println("=====================================")
	fmt.Printf("Final Score: %d out of %d\n", q.Score, len(q.Questions))
	// Round the duration to a more readable format.
	fmt.Printf("Total Time:  %.2f seconds\n", duration.Seconds())
	fmt.Println("=====================================")
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

You have now built a much more sophisticated application! By breaking the code
into smaller functions, you've made it easier to read, debug, and add new
features to in the future. This is a huge step forward in your journey as a
programmer.
*/
