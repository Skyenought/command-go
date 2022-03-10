package sql2struct

import (
	"fmt"
	"github.com/skyenought/command-go/internal/word"
	"os"
	"text/template"
)

const structTemplate = `type {{.TableName | ToCamelCase}} struct {
{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
	{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}}

func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

type StructTemplate struct {
	structTemplate string
}


type StructColumn struct {
	Name, Type, Tag, Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns [] *StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTemplate: structTemplate}
}

func (t *StructTemplate) AssemblyColumns(tableColumns []*TableColumn) []*StructColumn {
	templateColumns := make([]*StructColumn, 0, len(tableColumns))
	for _, column := range tableColumns {
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName)
		templateColumns = append(templateColumns, &StructColumn{
			column.ColumnName,
			DBTypeToStructType[column.DataType],
			tag,
			column.ColumnComment,
		})
	}
	
	return templateColumns
}

func (t *StructTemplate) Generate(tableName string, templateColumns []*StructColumn) error {
	_template := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToLowerCamelCase,
	}).Parse(t.structTemplate))
	templateDB := StructTemplateDB{
		TableName: tableName,
		Columns:   templateColumns,
	}
	if err := _template.Execute(os.Stdout, templateDB); err != nil {
		return err
	}
	return nil
}