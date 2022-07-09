package maps

// Dictionary stores some words with their definitions.
type Dictionary map[string]string

// Search find a word in the dictionary.
func (d Dictionary) Search(word string) string {
	return d[word]
}
