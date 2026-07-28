package json_conversion

import (
	"encoding/json"
	"fmt"
)

func Encrypt(data any) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(b), nil
}

func Decrypt[T any](data string) (T, error) {
	var value T
	err := json.Unmarshal([]byte(data), &value)
	if err != nil {
		fmt.Println(err)
		var zero T
		return zero, err
	}
	return value, nil
}
