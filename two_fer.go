// Package twofer has a single method for outputting a common phrase.
package twofer

// ShareWith outputs both an explicit and a generic two-fer phrase.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
