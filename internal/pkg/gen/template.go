package gen

import (
	"bytes"
	_ "embed"
	"go/format"
	"io/ioutil"
	"os"
	"text/template"
)

const (
	DeviceTemplateName = "devices"
	ValueTemplateName  = "values"
	KeysTemplateName   = "keys"
)

var (
	Template *template.Template
)

func init() {
	initTemplate()
}

func initTemplate() {
	Template = template.Must(template.New("header").Parse(HeaderTemplate))

	Template = template.Must(Template.New("object").Funcs(map[string]any{
		"keyProperties": keyProperties,
	}).Parse(ObjectTemplate))

	Template = template.Must(Template.New(KeysTemplateName).Parse(KeyTemplate))

	Template = template.Must(Template.New(DeviceTemplateName).Funcs(
		map[string]any{
			"noDeviceTypeName": func(tn string) string {
				return tn[6:]
			},
		}).Parse(DeviceTemplate))

	Template = template.Must(Template.New(ValueTemplateName).Parse(ValueTemplate))
}

func WriteTemplate(filename string, templateName string, data any) {
	var buf bytes.Buffer
	if err := Template.ExecuteTemplate(&buf, templateName, data); err != nil {
		panic(err.Error())
	}

	b, err := format.Source(buf.Bytes())
	if err != nil {
		_ = ioutil.WriteFile("panic-format.go", buf.Bytes(), os.ModePerm)
		panic(err.Error())
	}

	if err = ioutil.WriteFile(filename, b, os.ModePerm); err != nil {
		panic(err.Error())
	}
}
