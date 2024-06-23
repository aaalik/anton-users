package validator

import "sync"

var rv = map[string]Validate{}

type (
	Validate func(value interface{}) error
)

func RegisterValidator(name string, v Validate) {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	rv[name] = v
}

func ValidateByName(input interface{}, name string) error {
	if v, ok := rv[name]; ok {
		return v(input)
	}
	return nil
}
