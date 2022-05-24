package gen

import (
	"confMQTT/internal/pkg/tools"
	"errors"
)

var (
	Keys = Data[KeyOptions]{
		PackageName: EntriesFolder,
		Imports: []string{
			JsonHelpersPackage,
		},
	}
	keyMap = map[string]KeyOptions{}
)

func init() {
	Keys.Values = tools.MustJSON[[]KeyOptions](KeyList, Keys.Values)

	for _, k := range Keys.Values {
		keyMap[k.TypeName] = k
	}
}

func GetKey(s string) *KeyOptions {
	k, ok := keyMap[s]
	if !ok {
		return nil
	}
	return &k
}

func keyProperties(keys []string) ([]KeyOptions, error) {
	var out []KeyOptions
	for _, k := range keys {
		if key := GetKey(k); key != nil {
			out = append(out, *key)
			continue
		}
		return nil, errors.New("key does not exist " + k)
	}
	return out, nil
}
