package main

import (
	"confMQTT/internal/pkg/gen"
	"fmt"
	"path"
)

func main() {
	fmt.Println("--- Generating code for device discovery configurations. ---")

	gen.RecreateDir(path.Join("pkg", gen.EntriesFolder))
	gen.RecreateDir(path.Join("pkg", gen.DevicesFolder))

	gen.WriteTemplate(gen.KeysPath, gen.KeysTemplateName, gen.Keys)
	gen.WriteTemplate(gen.DevicesPath, gen.DeviceTemplateName, gen.Devices)
	gen.WriteTemplate(gen.ValuesPath, gen.ValueTemplateName, gen.Values)

	fmt.Println("--- Done! ---")
}
