// Part 1, Lesson 5: Control Flow
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file serves as the lesson and demonstration for control flow.
// It explains the core concepts through structured comments and
// provides a runnable example of their implementation.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

WHAT IS CONTROL FLOW?
So far, our programs have executed from top to bottom, one line after another.
CONTROL FLOW statements allow us to change that. They give our program a "brain"
by letting it make decisions and repeat actions.

We will learn the three fundamental types of control flow in Go:
1. `if / else` - To execute code based on a condition.
2. `switch` - A more powerful way to select one of many code blocks to run.
3. `for` - The single, unified way to loop (repeat code) in Go.
*/

package main

import "fmt"

func main() {
	fmt.Println("Starting the Control Flow Lesson!")

	// --- Part 1: `if`, `else if`, and `else` Statements ---

	// The `if` statement checks if a condition is true. If it is, the code
	// inside the curly braces `{}` is executed.
	userAge := 20

	fmt.Println("\n--- Demonstrating if/else ---")
	if userAge >= 18 {
		fmt.Println("Condition is true: You are old enough to vote.")
	}

	// You can add an `else` block to run code when the condition is false.
	if userAge < 13 {
		fmt.Println("You are a child.")
	} else {
		fmt.Println("Condition is false: You are not a child.")
	}

	// You can chain multiple conditions together with `else if`.
	score := 85

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B") // This block will run.
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: D or F")
	}

	// Go's `if` statement has a handy feature: you can include a short
	// initialization statement before the condition. The variable created
	// (`isValid`) only exists within the scope of the if-else block.
	if isValid := true; isValid {
		fmt.Println("This is a common and idiomatic Go pattern!")
	}

	// --- Part 2: The `switch` Statement ---

	// A `switch` statement is often a cleaner way to write a long `if-else if-else` chain.
	// It compares an expression against a series of `case` values.
	fmt.Println("\n--- Demonstrating switch ---")
	dayOfWeek := "Wednesday"

	switch dayOfWeek {
	case "Monday":
		fmt.Println("It's the start of the work week.")
	case "Tuesday", "Wednesday", "Thursday": // You can have multiple values per case.
		fmt.Println("It's a weekday.")
	case "Friday":
		fmt.Println("TGIF!")
	default:
		// The `default` case runs if no other case matches.
		fmt.Println("It's the weekend!")
	}

	// Unlike many other languages, Go's `switch` cases do not "fall through"
	// automatically. A `break` is implied at the end of each case.

	// A `switch` can also be used without an expression. This makes it a very
	// clean way to write complex `if-else` logic.
	hourOfDay := 14
	switch { // Note: no variable after `switch`
	case hourOfDay < 12:
		fmt.Println("Good morning!")
	case hourOfDay < 18:
		fmt.Println("Good afternoon!") // This will match.
	default:
		fmt.Println("Good evening!")
	}

	// --- Part 3: The Unified `for` Loop ---

	// Go has only one looping construct: the `for` loop. But it's very versatile
	// and can be used in several ways.

	fmt.Println("\n--- Demonstrating for loops ---")

	// 1. The "standard" loop (like C or Java): `for init; condition; post {}`
	// `init`: runs once before the loop.
	// `condition`: checked before each iteration. The loop stops when it's false.
	// `post`: runs after each iteration.
	fmt.Println("Standard 'for' loop:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("  Iteration %d\n", i)
	}

	// 2. The "while" loop: `for condition {}`
	// You can omit the `init` and `post` statements to make it behave like a `while` loop.
	fmt.Println("\n'While' style 'for' loop:")
	n := 1
	for n <= 5 {
		fmt.Printf("  Iteration %d\n", n)
		n++ // We must increment `n` manually inside the loop.
	}

	// 3. The "infinite" loop: `for {}`
	// Omitting everything creates a loop that runs forever. We need a way to get out!
	// `break` exits the loop immediately.
	// `continue` skips the rest of the current iteration and goes to the next one.
	fmt.Println("\nInfinite loop with 'break' and 'continue':")
	count := 0
	for { // Infinite loop
		count++
		if count > 10 {
			break // Exit the loop when count is greater than 10.
		}
		if count%2 != 0 {
			continue // If count is odd, skip the Println and start the next iteration.
		}
		// This line will only run for even numbers (2, 4, 6, 8, 10).
		fmt.Printf("  Processing even number: %d\n", count)
	}
	fmt.Println("Loop finished.")
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

KEY TAKEAWAYS:
- `if`/`else if`/`else`: The classic way to make decisions based on conditions.
- `switch`: A powerful and clean alternative for comparing a value against many cases.
- `for`: Go's single, flexible looping keyword. It can act like a standard `for` loop,
  a `while` loop, or an infinite loop.
- `break`: Exits a loop entirely.
- `continue`: Skips the current iteration and proceeds to the next.

HOW TO RUN THIS CODE:

1.  Open a terminal or command prompt.
2.  Navigate to the directory where you saved this file.
3.  Use the `go run` command to compile and execute the file:
    `go run 5_control_flow.go`
*/
