package gen

import "confMQTT/internal/pkg/tools"

var (
	Devices = Data[ObjectOptions]{
		PackageName: DevicesFolder,
		Imports: []string{
			EntriesPackage,
			JsonHelpersPackage,
		},
	}
)

func init() {
	Devices.Values = tools.MustJSON[[]ObjectOptions](DeviceList, Devices.Values)

	for i, v := range Devices.Values {
		deviceName := "Device" + tools.FormatTypeName(v.Name)
		Devices.Values[i].TypeName = deviceName
	}
}
