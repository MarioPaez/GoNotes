package maps

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExist        = DictionaryErr("the word already exist in the dictionary")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

type DictionaryErr string

// We implement the error interface, so DictionaryErr is a error
func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, ok := d[word]
	if ok {
		return ErrWordExist
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, ok := d[word]
	if !ok {
		return ErrWordDoesNotExist
	}
	d[word] = newDefinition
	return nil
}
