// Part 2, Lesson 9: Maps
//
// Author: dunamismax
// Date: 06-15-2025
//
// This file serves as the lesson and demonstration for maps.
// It explains the core concepts through structured comments and
// provides a runnable example of their implementation.

/*
=====================================================================================
|                                   - LESSON START -                                  |
=====================================================================================

WHAT IS A MAP?
After learning about slices for ordered data, we now turn to MAPS for unordered
data. A MAP is Go's built-in data structure for storing KEY-VALUE pairs. [1]

If you've used other programming languages, you might know this concept by other
names like "dictionary" (Python), "hash map" (Java), or "associative array" (PHP).

A map associates a specific KEY with a specific VALUE. You can then use the key to
look up the value very quickly. The keys in a map must be unique.

Key characteristics:
-   **Unordered**: The items in a map are not stored in any particular order. When you
    iterate over a map, the order is not guaranteed to be the same every time. [3]
-   **Key Types**: Map keys must be a "comparable" type, meaning Go must be able
    to tell if two keys are equal using the `==` operator. Common key types are
    `string`, `int`, `float64`, and `bool`. Slices cannot be map keys.
-   **Zero Value**: The zero value of a map is `nil`. A `nil` map has no keys and
    you cannot add keys to it. You must initialize it first.
*/

package main

import "fmt"

func main() {
	// --- Part 1: Creating and Initializing a Map ---

	// The most common way to create a map is using the `make` function.
	// The syntax is `make(map[KeyType]ValueType)`.
	// Let's create a map to store the scores of players in a game.
	// The keys will be `string` (player names) and values will be `int` (their scores).
	playerScores := make(map[string]int)

	// Now we can add key-value pairs to the map.
	playerScores["alice"] = 95
	playerScores["bob"] = 82
	playerScores["charlie"] = 100

	fmt.Println("--- Creating and using a map ---")
	fmt.Printf("Player scores map: %v\n", playerScores)

	// --- Part 2: Map Literals ---

	// You can also create and initialize a map in one step using a "map literal".
	// This is often more convenient if you know some of the key-value pairs upfront.
	productPrices := map[string]float64{
		"Milk":  3.50,
		"Bread": 2.75,
		"Eggs":  4.25, // A trailing comma is required on the last element!
	}
	fmt.Println("\n--- Using a map literal ---")
	fmt.Printf("Product prices: %v\n", productPrices)

	// --- Part 3: Accessing, Updating, and Deleting ---

	// Accessing a value is done by using its key in square brackets.
	bobScore := playerScores["bob"]
	fmt.Printf("\nBob's score is: %d\n", bobScore)

	// Updating a value is the same as adding one.
	playerScores["bob"] = 87
	fmt.Printf("Bob's updated score is: %d\n", playerScores["bob"])

	// The built-in `delete` function removes a key-value pair from a map.
	delete(playerScores, "charlie")
	fmt.Printf("Map after deleting Charlie: %v\n", playerScores)

	// --- Part 4: The "Comma, Ok" Idiom ---

	// What happens if we try to access a key that doesn't exist?
	dianaScore := playerScores["diana"]
	fmt.Printf("\nTried to get Diana's score. Got: %d\n", dianaScore)
	// Notice it returns `0`, the zero value for the `int` type. This is a problem!
	// Did Diana score 0, or is she not in our map? We can't tell.

	// To solve this, Go provides a special two-value assignment for map access.
	// `value, ok := myMap[key]`
	// The second value, conventionally named `ok`, is a boolean.
	// It's `true` if the key exists, and `false` if it doesn't.
	score, ok := playerScores["diana"]
	if ok {
		fmt.Printf("Diana's score is %d.\n", score)
	} else {
		fmt.Println("The key 'diana' was not found in the map.")
	}

	// Let's check for a key that does exist.
	score, ok = playerScores["alice"]
	if ok {
		fmt.Printf("Alice's score is %d.\n", score)
	} else {
		fmt.Println("The key 'alice' was not found in the map.")
	}

	// --- Part 5: Iterating Over a Map ---

	// We can loop through all the key-value pairs in a map using a `for range` loop.
	fmt.Println("\n--- Iterating over the product prices map ---")
	for key, value := range productPrices {
		fmt.Printf("  Product: %s, Price: $%.2f\n", key, value)
	}
	// Remember: The order of iteration is NOT guaranteed!
}

/*
=====================================================================================
|                                    - LESSON END -                                   |
=====================================================================================

KEY TAKEAWAYS:
- A MAP is an unordered collection of KEY-VALUE pairs.
- Use `make(map[KeyType]ValueType)` to create an empty map.
- Use `map[KeyType]ValueType{...}` for a map literal with initial values.
- Access, add, and update with `myMap[key] = value`.
- Remove elements with `delete(myMap, key)`.
- Accessing a non-existent key returns a zero value, which can be misleading.
- ALWAYS use the "comma, ok" idiom (`value, ok := myMap[key]`) to safely check
  for a key's existence. [9]
- Use a `for range` loop to iterate over a map's key-value pairs.

HOW TO RUN THIS CODE:

1.  Open a terminal or command prompt.
2.  Navigate to the directory where you saved this file.
3.  Use the `go run` command to compile and execute the file:
    `go run 9_maps.go`
*/
