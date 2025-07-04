
// regionDrop*Of{{.N}}

// DropFirstOf{{.N}} drops the first value from a tuple of {{.N}}.
func DropFirstOf{{.N}}[{{template "Gen" .}}]({{range $i, $a := .Args}}{{if eq $i 0}}_ T0{{else}}, {{$a.Name}} T{{$i}}{{end}}{{end}}){{if ne $.N 1}} {{end}}{{if gt $.N 2}}({{end}}{{- range $i, $a := .Args}}{{if gt $i 0}}{{if ne $i 1}}, {{end}}T{{$i}}{{end}}{{end}}{{if gt $.N 2}}){{end}} {
	return{{range $i, $a := .Args}}{{if $i}}{{if ne $i 1}},{{end}} {{$a.Name}}{{end}}{{end}}
}

{{- if ne $.N 1}}

// DropLastOf{{.N}} drops the last value from a tuple of {{.N}}.
func DropLastOf{{.N}}[{{template "Gen" .}}]({{range $i, $a := .Args}}{{if $i}}, {{end}}{{$a.Name}} T{{$i}}{{end}}){{if ne $.N 1}} {{end}}{{if gt $.N 2}}({{end}}{{- range $i, $a := .Args}}{{if ne $i $.Last}}{{if $i}}, {{end}}T{{$i}}{{end}}{{end}}{{if gt $.N 2}}){{end}} {
	return{{range $i, $a := .Args}}{{if ne $i $.Last}}{{if $i}},{{end}} {{$a.Name}}{{end}}{{end}}
}{{end}}
