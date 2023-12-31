package gen

import (
	"fmt"
	"sort"
	"strings"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/yanxin666/goctlpri/model/sql/template"
	"github.com/yanxin666/goctlpri/util"
	"github.com/yanxin666/goctlpri/util/pathx"
	"github.com/yanxin666/goctlpri/util/stringx"
)

func genVars(table Table, withCache, postgreSql bool) (string, error) {
	keys := make([]string, 0)
	keys = append(keys, table.PrimaryCacheKey.VarExpression)
	for _, v := range table.UniqueCacheKey {
		keys = append(keys, v.VarExpression)
	}

	camel := table.Name.ToCamel()
	upCamel := camel
	if table.typ != "" {
		upCamel = table.typ
	}
	text, err := pathx.LoadTemplate(category, varTemplateFile, template.Vars)
	if err != nil {
		return "", err
	}

	output, err := util.With("var").Parse(text).
		GoFmt(true).Execute(map[string]any{
		"lowerStartCamelObject": stringx.From(camel).Untitle(),
		"upperStartCamelObject": upCamel,
		"cacheKeys":             strings.Join(keys, "\n"),
		"autoIncrement":         table.PrimaryKey.AutoIncrement,
		"originalPrimaryKey":    wrapWithRawString(table.PrimaryKey.Name.Source(), postgreSql),
		"withCache":             withCache,
		"postgreSql":            postgreSql,
		"data":                  table,
		"ignoreColumns": func() string {
			var set = collection.NewSet()
			for _, c := range table.ignoreColumns {
				if postgreSql {
					set.AddStr(fmt.Sprintf(`"%s"`, c))
				} else {
					set.AddStr(fmt.Sprintf("\"`%s`\"", c))
				}
			}
			list := set.KeysStr()
			sort.Strings(list)
			return strings.Join(list, ", ")
		}(),
	})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
