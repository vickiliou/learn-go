package maps

// Dictionary stores some words with their definitions.
type Dictionary map[string]string

const (
	// ErrNotFound means the definition could not be found for the given word.
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists means the word that is trying to add already exists.
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	// ErrWordDoesNotExist means the word that is trying to update doesn't exist.
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

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

// Update changes the definition of a given word.
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}
