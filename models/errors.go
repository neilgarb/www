package models

type ValidationError map[string][]string

func NewValidationError() *ValidationError {
	return &ValidationError{}
}

func (self *ValidationError) Error() string {
	return "Please make the following corrections."
}

func (self *ValidationError) Add(field string, err string) {
	fieldErrors, found := (*self)[field]
	if !found {
		(*self)[field] = []string{err}
	} else {
		(*self)[field] = append(fieldErrors, err)
	}
}

func (self *ValidationError) HasErrors() bool {
	return len(*self) > 0
}
