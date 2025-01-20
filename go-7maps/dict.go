package dictionary

import(
	"errors"
)

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dict map[string]string

func(d Dict)Search(word string) (string, error){
	def, ok := d[word]

	if !ok{
		return "", ErrNotFound
	}
	
	return def, nil
}

func (d Dict)Add(word, def string) {
	d[word] = def
}