package dictionary

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("no such word")
var WordExists = errors.New("word exists")

// func Search(dict map[string]string, word string) string {
// 	return dict[word]
// }

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if ok {
		return definition, nil
	}
	return "", ErrNotFound
}

func (d Dictionary) Add(key string, word string) error {
	_, error := d.Search(key)

	switch error {
	case ErrNotFound:
		d[key] = word
	case nil:
		return WordExists
	default:
		return error
	}
	return nil
}

func (d Dictionary) Update(key string, word string) error {
	_, error := d.Search(key)

	switch error {
	case nil:
		d[key] = word
	case ErrNotFound:
		return error
	}
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, error := d.Search(key)

	switch error {
	case ErrNotFound:
		return error
	}

	delete(d, key)
	return nil
}
