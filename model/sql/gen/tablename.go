package gen

import (
	"github.com/yanxin666/goctlpri/model/sql/template"
	"github.com/yanxin666/goctlpri/util"
	"github.com/yanxin666/goctlpri/util/pathx"
)

func genTableName(table Table) (string, error) {
	text, err := pathx.LoadTemplate(category, tableNameTemplateFile, template.TableName)
	if err != nil {
		return "", err
	}

	camel := table.Name.ToCamel()
	if table.typ != "" {
		camel = table.typ
	}

	output, err := util.With("tableName").
		Parse(text).
		Execute(map[string]any{
			"tableName":             table.Name.Source(),
			"upperStartCamelObject": camel,
		})
	if err != nil {
		return "", nil
	}

	return output.String(), nil
}
