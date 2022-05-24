package tools

import "encoding/json"

func MustJSON[T any](data []byte, into T) T {
	if err := json.Unmarshal(data, &into); err != nil {
		panic(err.Error())
	}
	return into
}
