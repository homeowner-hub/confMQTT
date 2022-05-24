package gen

import (
	"path"
	"path/filepath"
)

const (
	LibraryName = "confMQTT"

	ConfFolder    = ""
	EntriesFolder = "confEntry"
	DevicesFolder = "confDevice"

	KeysFileName    = "keys_gen.go"
	ValuesFileName  = "values_gen.go"
	DevicesFileName = "devices_gen.go"
)

var (
	BasePackage        = path.Join(LibraryName, "pkg")
	InternalPackage    = path.Join(LibraryName, "internal", "pkg")
	JsonHelpersPackage = path.Join(InternalPackage, "tools")
	EntriesPackage     = path.Join(BasePackage, ConfFolder, EntriesFolder)
	DevicesPackage     = path.Join(BasePackage, ConfFolder, DevicesFolder)

	KeysPath    = filepath.Join("pkg", EntriesFolder, KeysFileName)
	ValuesPath  = filepath.Join("pkg", EntriesFolder, ValuesFileName)
	DevicesPath = filepath.Join("pkg", DevicesFolder, DevicesFileName)
)
