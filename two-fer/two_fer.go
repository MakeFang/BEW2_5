// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

func ShareWith(name string) string {
	// check if name is empty. If name is empty use "you", otherwise use the name.
	if name == "" {
		return "One for you, one for me."
	} else {
		return "One for " + name + ", one for me."
	}
}
