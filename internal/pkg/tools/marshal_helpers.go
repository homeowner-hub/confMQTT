package tools

import "encoding/json"

func MarshalBiKeys(keys ...Key[any, any]) ([]byte, error) {
	out := map[string]interface{}{}
	for _, k := range keys {
		if v, ok := k.Value(); ok {
			out[k.LongName()] = v
		} else if v, ok := k.Value2(); ok {
			out[k.LongName()] = v
		}
	}
	return json.Marshal(out)
}

type (
	Key[T, V any] interface {
		LongName() string
		ShortName() string
		Optional() bool
		Value() (T, bool)
		Value2() (V, bool)
	}

	PseudoKey[T, V any] struct {
		K Key[T, V]
	}
)

func (p *PseudoKey[T, V]) LongName() string {
	return p.K.LongName()
}

func (p *PseudoKey[T, V]) ShortName() string {
	return p.K.ShortName()
}

func (p *PseudoKey[T, V]) Optional() bool {
	return p.K.Optional()
}

func (p *PseudoKey[T, V]) Value() (any, bool) {
	return p.K.Value()
}

func (p *PseudoKey[T, V]) Value2() (any, bool) {
	return p.K.Value2()
}
