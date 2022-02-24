package mydict

import "errors"

// Dictionary Type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not found.")
	errWordExists = errors.New("The word is existed already.")
)

// Search a word.
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word and its definition.
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil

	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExists
	// }
	// return nil
}

// Update a dictionary.
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return err
	}
	return nil
}

// Delete a word from a dictionary.
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
