package gen

import (
	"github.com/yanxin666/goctlpri/model/sql/template"
	"github.com/yanxin666/goctlpri/util"
	"github.com/yanxin666/goctlpri/util/pathx"
	"github.com/yanxin666/goctlpri/util/stringx"
)

func genTypes(table Table, methods string, withCache bool) (string, error) {
	fields := table.Fields
	fieldsString, err := genFields(table, fields)
	if err != nil {
		return "", err
	}

	text, err := pathx.LoadTemplate(category, typesTemplateFile, template.Types)
	if err != nil {
		return "", err
	}

	camel := table.Name.ToCamel()
	upCamel := camel
	if table.typ != "" {
		upCamel = table.typ
	}

	output, err := util.With("types").
		Parse(text).
		Execute(map[string]any{
			"withCache":             withCache,
			"method":                methods,
			"upperStartCamelObject": upCamel,
			"lowerStartCamelObject": stringx.From(camel).Untitle(),
			"fields":                fieldsString,
			"data":                  table,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
