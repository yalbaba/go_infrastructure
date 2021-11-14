package {{.Models}}

{{$m := .Models}}

{{$ilen := len .Imports}}
{{if gt $ilen 0}}
import (
	{{range .Imports}}"{{.}}"{{end}}
)
{{end}}

{{range .Tables}}
type {{ .Name | UpFirst}} struct {
{{$table := .}}
{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}	{{Mapper $col.Name}}	{{Type $col}} {{Tag $table $col}}
{{end}}
}

func (*{{ .Name | UpFirst}}) TableName() string {
    return "{{$m}}.{{ .Name}}"
}

{{end}}
