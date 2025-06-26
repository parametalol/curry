
// region:Drop*Of{{.N}}

// DropFirstOf{{.N}} drops the first value from a tuple of {{.N}}.
func DropFirstOf{{.N}}[{{range $i, $a := .Args}}{{if ne $i 0}}, {{end}}T{{$i}}{{end}} any]({{range $i, $a := .Args}}{{if eq $i 0}}_ T0{{else}}, {{$a.Name}} T{{$i}}{{end}}{{end}}){{if ne $.N 1}} {{end}}{{if gt $.N 2}}({{end}}{{- range $i, $a := .Args}}{{if gt $i 0}}{{if ne $i 1}}, {{end}}T{{$i}}{{end}}{{end}}{{if gt $.N 2}}){{end}} {
	return {{range $i, $a := .Args}}{{if ne $i 0}}{{if ne $i 1}}, {{end}}{{$a.Name}}{{end}}{{end}}
}

{{- if ne $.N 1}}

// DropLastOf{{.N}} drops the last value from a tuple of {{.N}}.
func DropLastOf{{.N}}[{{range $i, $a := .Args}}{{if ne $i 0}}, {{end}}T{{$i}}{{end}} any]({{range $i, $a := .Args}}{{if ne $i 0}}, {{end}}{{$a.Name}} T{{$i}}{{end}}){{if ne $.N 1}} {{end}}{{if gt $.N 2}}({{end}}{{- range $i, $a := .Args}}{{if ne $i $.Last}}{{if ne $i 0}}, {{end}}T{{$i}}{{end}}{{end}}{{if gt $.N 2}}){{end}} {
	return {{range $i, $a := .Args}}{{if ne $i $.Last}}{{if ne $i 0}}, {{end}}{{$a.Name}}{{end}}{{end}}
}{{- end}}
