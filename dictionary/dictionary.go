package dictionary

import "errors"

type Dictionary map[string]string

var (
	ErrWordExists       = errors.New("cannot add word because it already exists")
	ErrNotFound         = errors.New("searched word cannot be found")
	ErrWordDoesNotExist = errors.New("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// An interesting property of maps is that you can modify them without passing them as a pointer. This is because map is a reference type
// my own definition
// func (d Dictionary) Add(word, definition string) error {
// 	if _, ok := d[word]; ok {
// 		return ErrWordExists
// 	}
// 	d[word] = definition
// 	return nil
// }

// definition of Add from the tutorial
// 1 thing i don't think about this is the Error driven development, but i guess Error is a type in GoLang
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

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}
