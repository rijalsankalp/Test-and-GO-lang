package maps

type Dictionary map[string]string

const (
	ErrNotFound   = DictError("Word not in dict.")
	ErrWordExists = DictError("Word already exists in dict.")
)

type DictError string

func (e DictError) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	meaning, message := d[word]

	if !message {
		return "", ErrNotFound
	}
	return meaning, nil
}

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

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrNotFound
	case nil:
		d[word] = definition
	default:
		return err
	}
	return ErrWordExists
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}