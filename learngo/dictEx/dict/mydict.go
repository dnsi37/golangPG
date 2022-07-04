package mydict

import (
	"errors"
)

//Dicationary type
type Dictionary map[string]string

var errNotFound = errors.New("not Found")
var errWordExist = errors.New("that word exist")
var errCantUpdate = errors.New("can't update non-existing word")

// Search for a word , type 에 method 추가 가능
func (d Dictionary) Search(word string) (string, error) {
	// key 의 여부 exist 확인가능 ( Bool )
	value, exist := d[word]
	if exist {
		//fmt.Println(exist)
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExist
	}
	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil

}

// Delete a word
func (d Dictionary) Delete(word string) error {
	delete(d, word)
	return nil
}
