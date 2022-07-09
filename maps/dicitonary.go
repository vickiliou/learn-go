package maps

// Search find a word in the dictionary.
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
