package maps

import "errors"

// Dictionary stores some words with their definitions.
type Dictionary map[string]string

// ErrNotFound means the definition could not be found for the given word.
var ErrNotFound = errors.New("could not find the word you were looking for")

// Search find a word in the dictionary.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add inserts a word and definition into the dictionary.
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
