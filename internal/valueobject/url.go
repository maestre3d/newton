package valueobject

import "strings"

const (
	urlMinLength = 5
	urlMaxLength = 2000
)

func ensureValidURL(v string, err error) error {
	uriVec := strings.Split(v, ".")
	if len(uriVec) < 2 {
		return err
	} else if uriJoin := strings.Join(uriVec[1:], "."); !strings.Contains(uriJoin, "/") {
		// verify if URI path was given (e.g. aka.ms/book.pdf)
		return err
	}
	return nil
}

func ensureURLLength(v string, err error) error {
	if len(v) < urlMinLength || len(v) > urlMaxLength {
		return err
	}
	return nil
}
