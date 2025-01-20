package dictionary

const (
	ErrNotFound = DictErr("could not find the word you were looking for")
	ErrWordExists = DictErr("could not add word because it already exists in the Dictionary")
	ErrWordDoesNotExist = DictErr("Could not find the word to perform this action")
)

type DictErr string

func (e DictErr) Error() string{
	return string(e)
}

type Dict map[string]string

func(d Dict)Search(word string) (string, error){
	def, ok := d[word]

	if !ok{
		return "", ErrNotFound
	}
	
	return def, nil
}

func (d Dict)Add(word, def string) error {
	_, err := d.Search(word)

	switch err{
	case ErrNotFound:
		d[word] = def
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dict)Update(word, newDef string) error {
	_, err := d.Search(word)

	switch err{
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDef
	default:
		return err
	}
	return nil
}


func (d Dict)Delete(word string) error{
	_, err := d.Search(word)

	switch err{
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}