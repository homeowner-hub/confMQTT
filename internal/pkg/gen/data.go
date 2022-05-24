package gen

import (
	_ "embed"
)

var (
	//go:embed templates/keys.go.tmpl
	KeyTemplate string
	//go:embed templates/devices.go.tmpl
	DeviceTemplate string
	//go:embed templates/values.go.tmpl
	ValueTemplate string
	//go:embed templates/header.go.tmpl
	HeaderTemplate string
	//go:embed templates/object.go.tmpl
	ObjectTemplate string

	//go:embed definitions/devices.json
	DeviceList []byte
	//go:embed definitions/keys.json
	KeyList []byte
	//go:embed definitions/values.json
	ValueList []byte
)

type (
	Data[T any] struct {
		PackageName string
		Imports     []string
		Values      []T
	}

	ObjectOptions struct {
		Name     string
		TypeName string
		Keys     []string
	}

	KeyOptions struct {
		TypeName  string
		LongName  string
		ShortName string

		Optional bool

		Type1 string
		Type2 string
		BiVal bool

		Def     bool
		DefType string
		DefVal  string
	}
)
