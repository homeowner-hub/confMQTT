// Code generated by gen-confs DO NOT EDIT.
package confEntry

import (
	. "confMQTT/internal/pkg/tools"
)

type ValueDevice struct {
	KeyConfigurationUrl KeyConfigurationUrl
	KeyConnections      KeyConnections
	KeyIdentifiers      KeyIdentifiers
	KeyManufacturer     KeyManufacturer
	KeyModel            KeyModel
	KeyName             KeyName
	KeySuggestedArea    KeySuggestedArea
	KeySwVersion        KeySwVersion
	KeyViaDevice        KeyViaDevice
}

func (tn *ValueDevice) MarshalJSON() ([]byte, error) {
	return MarshalBiKeys(
		&PseudoKey[string, string]{K: &tn.KeyConfigurationUrl},
		&PseudoKey[[][]string, map[string]string]{K: &tn.KeyConnections},
		&PseudoKey[string, []string]{K: &tn.KeyIdentifiers},
		&PseudoKey[string, string]{K: &tn.KeyManufacturer},
		&PseudoKey[string, string]{K: &tn.KeyModel},
		&PseudoKey[string, string]{K: &tn.KeyName},
		&PseudoKey[string, string]{K: &tn.KeySuggestedArea},
		&PseudoKey[string, string]{K: &tn.KeySwVersion},
		&PseudoKey[string, string]{K: &tn.KeyViaDevice},
	)
}

func (tn *ValueDevice) UnmarshalJSON(b []byte) error {
	return UnmarshalBiKeys("ValueDevice", b,
		&PseudoKey[string, string]{K: &tn.KeyConfigurationUrl},
		&PseudoKey[[][]string, map[string]string]{K: &tn.KeyConnections},
		&PseudoKey[string, []string]{K: &tn.KeyIdentifiers},
		&PseudoKey[string, string]{K: &tn.KeyManufacturer},
		&PseudoKey[string, string]{K: &tn.KeyModel},
		&PseudoKey[string, string]{K: &tn.KeyName},
		&PseudoKey[string, string]{K: &tn.KeySuggestedArea},
		&PseudoKey[string, string]{K: &tn.KeySwVersion},
		&PseudoKey[string, string]{K: &tn.KeyViaDevice},
	)
}

type ValueAvailability struct {
	KeyPayloadAvailable    KeyPayloadAvailable
	KeyPayloadNotAvailable KeyPayloadNotAvailable
	KeyTopic               KeyTopic
	KeyValueTemplate       KeyValueTemplate
}

func (tn *ValueAvailability) MarshalJSON() ([]byte, error) {
	return MarshalBiKeys(
		&PseudoKey[string, string]{K: &tn.KeyPayloadAvailable},
		&PseudoKey[string, string]{K: &tn.KeyPayloadNotAvailable},
		&PseudoKey[string, string]{K: &tn.KeyTopic},
		&PseudoKey[string, string]{K: &tn.KeyValueTemplate},
	)
}

func (tn *ValueAvailability) UnmarshalJSON(b []byte) error {
	return UnmarshalBiKeys("ValueAvailability", b,
		&PseudoKey[string, string]{K: &tn.KeyPayloadAvailable},
		&PseudoKey[string, string]{K: &tn.KeyPayloadNotAvailable},
		&PseudoKey[string, string]{K: &tn.KeyTopic},
		&PseudoKey[string, string]{K: &tn.KeyValueTemplate},
	)
}
