// Part 3, Lesson 16: Project: Concurrent Web Checker
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file is a project that demonstrates the power of goroutines and channels
// by building a tool that checks the status of a list of websites concurrently.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

PROJECT BRIEF: CONCURRENT WEBSITE STATUS CHECKER (V2)

In our last two lessons, we learned about GOROUTINES for running tasks at the
same time and CHANNELS for communicating safely between them. Now, we'll build a
more advanced version of our web checker that is both more powerful and better-written.

THE GOAL:
Build a command-line tool that takes a list of website URLs and checks each one,
reporting on its status, status code, and how long the check took.

THE CHALLENGE:
Checking websites sequentially is slow. The total time is the sum of all response
times. We need to do this concurrently. Furthermore, just getting an "OK" or "Down"
status isn't enough. We want structured data: what was the URL, what was the
exact error or status code, and how fast was the response?

THE SOLUTION:
We will use goroutines to check each website concurrently. But instead of sending
a simple string over our channel, we will define a `CheckResult` STRUCT. This
allows us to send rich, structured data back from each goroutine. The main
function will then collect all these structs and print a nicely formatted report.

This project will combine:
-   GOROUTINES: To run each check concurrently.
-   CHANNELS: To receive structured `CheckResult` data.
-   STRUCTS: To define the `CheckResult` type.
-   SLICES: To hold our list of websites and the results.
-   PACKAGES: `fmt`, `net/http`, and `time`.
*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

// --- Part 1: Defining Our Data Structure ---

// Instead of passing a simple string over the channel, we'll use a struct.
// This is a much better practice as it bundles all the related information
// for a single check into one neat package.
type CheckResult struct {
	URL        string        // The URL that was checked.
	Status     string        // A user-friendly status: "OK", "ERROR", etc.
	StatusCode int           // The HTTP status code (e.g., 200, 404, 500).
	Duration   time.Duration // How long the check took.
}

// --- Part 2: The Worker Function ---

// This function will be run in its own goroutine for each website.
// It now takes a channel that accepts `CheckResult` structs.
func checkWebsite(url string, c chan CheckResult) {
	// Record the start time of the check.
	start := time.Now()

	// The `http.Get` function makes an HTTP request to the given URL.
	// This is a "blocking" call; the goroutine will pause here until the
	// request completes or times out.
	resp, err := http.Get(url)

	// Record the total duration of the request.
	duration := time.Since(start)

	// If there's a connection error (e.g., DNS lookup failure, site is down),
	// the `err` variable will not be `nil`.
	if err != nil {
		// We create a CheckResult with the error details and send it.
		c <- CheckResult{
			URL:      url,
			Status:   "ERROR: " + err.Error(), // Include the specific error message.
			Duration: duration,
		}
		return
	}

	// It's very important to close the response body to free up system resources.
	// `defer` ensures this happens at the end of the function, even if errors occur later.
	defer resp.Body.Close()

	// If we successfully get a response, we create a CheckResult with the
	// details from the response and send it back on the channel.
	c <- CheckResult{
		URL:        url,
		Status:     resp.Status, // e.g., "200 OK", "404 Not Found"
		StatusCode: resp.StatusCode,
		Duration:   duration,
	}
}

func main() {
	// --- Part 3: Setting Up the Work ---

	fmt.Println("Starting concurrent website check...")

	// An expanded list of websites to check.
	// We've included some that will fail or return specific HTTP error codes.
	websites := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
		"https://github.com",
		"https://twitter.com",
		"https://youtube.com",
		"https://wikipedia.org",
		"https://httpstat.us/500",                    // A site that returns a 500 Internal Server Error.
		"https://httpstat.us/404",                    // A site that returns a 404 Not Found.
		"https://nonexistent-domain-for-testing.xyz", // This domain should not exist.
	}

	// Create a channel that can transport `CheckResult` structs.
	channel := make(chan CheckResult)

	// --- Part 4: Launching the Goroutines ---

	// We use a `for...range` loop to iterate over our slice of websites.
	for _, url := range websites {
		// For each URL, we launch a NEW GOROUTINE using the `go` keyword.
		// The `main` function does NOT wait; it immediately starts the next one.
		go checkWebsite(url, channel)
	}

	// --- Part 5: Collecting and Displaying the Results ---

	// We'll create a slice to hold the results as they come in.
	var results []CheckResult

	// We know exactly how many results to expect: one for each website.
	// So we will loop that many times, receiving one result per iteration.
	for i := 0; i < len(websites); i++ {
		// The `<-channel` expression BLOCKS and waits until a value is sent
		// to the channel. As soon as a goroutine finishes, we receive its result.
		result := <-channel
		fmt.Printf("Checked: %s\n", result.URL) // Give real-time feedback.
		results = append(results, result)
	}

	fmt.Println("\n======================================================================================")
	fmt.Println("|                                 - CHECK RESULTS -                                  |")
	fmt.Println("======================================================================================")
	// Print a formatted header for our results table.
	// `%-50s` means a left-aligned string padded to 50 characters.
	// `%10s` means a right-aligned string padded to 10 characters.
	fmt.Printf("%-50s | %-25s | %10s\n", "URL", "STATUS", "DURATION")
	fmt.Println("--------------------------------------------------------------------------------------")

	// Loop through our collected results and print each one in a formatted row.
	for _, result := range results {
		// If the status is an error, we don't want to print the whole long message in the table.
		statusText := result.Status
		if result.StatusCode == 0 { // An easy way to check if it was a connection error.
			statusText = "CONNECTION ERROR"
		}
		fmt.Printf("%-50s | %-25s | %10s\n", result.URL, statusText, result.Duration.Round(time.Millisecond))
	}
	fmt.Println("======================================================================================")

	fmt.Println("\nAll websites have been checked.")
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

HOW TO RUN THIS CODE:

1.  Open a terminal or command prompt.
2.  Navigate to the directory where you saved this file.
3.  Use the `go run` command to compile and execute the file:
    `go run 16_project_concurrent_web_checker.go`

Congratulations! You've built a high-performance, concurrent tool that provides
rich, structured feedback.

By doing this concurrently, the total time is roughly the time it takes for the
SLOWEST website to respond, not the sum of all of them. This is the power of
Go's concurrency model in action. Using a struct for communication made the final
reporting step much cleaner and more powerful.
*/
