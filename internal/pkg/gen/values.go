package gen

import (
	"confMQTT/internal/pkg/tools"
)

var (
	Values = Data[ObjectOptions]{
		PackageName: EntriesFolder,
		Imports: []string{
			JsonHelpersPackage,
		},
	}
	valueSet = map[string]bool{}
)

func init() {
	Values.Values = tools.MustJSON[[]ObjectOptions](ValueList, Values.Values)

	for i, v := range Values.Values {
		typeName := "Value" + tools.FormatTypeName(v.Name)
		Values.Values[i].TypeName = typeName
		valueSet[typeName] = true
	}
}

func IsValue(s string) bool {
	return valueSet[s]
}
