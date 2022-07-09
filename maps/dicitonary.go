package maps

import "errors"

// Dictionary stores some words with their definitions.
type Dictionary map[string]string

var (
	// ErrNotFound means the definition could not be found for the given word.
	ErrNotFound = errors.New("could not find the word you were looking for")
	// ErrWordExists means the word that is trying to add already exists.
	ErrWordExists = errors.New("cannot add word because it already exists")
)

// Search find a word in the dictionary.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add inserts a word and definition into the dictionary if the word doesn't exist.
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
