{{- define "Types" -}}
{{ range $i, $a := .Args }}{{if $i}}, {{end}}T{{$i}}{{end}}
{{- end}}

{{- define "TypesS" -}}
{{ range $i, $a := .Args }}{{if $i}}, {{end}}{{if eq $i $.Last}}...{{end}}T{{$i}}{{end}}
{{- end}}


{{- define "Args" -}}
{{ range $i, $a := .Args }}{{if $i}}, {{end}}{{$a.Name}}{{end}}
{{- end}}

{{- define "ArgsS" -}}
{{ range $i, $a := .Args }}{{if $i}}, {{end}}{{$a.Name}}{{if eq $i $.Last}}...{{end}}{{end}}
{{- end}}

{{- define "Gen" -}}
{{ template "Types" . }} any
{{- end}}

{{- define "GenR" -}}
{{range $i, $a := .Args }}T{{$i}}, {{end}}R any
{{- end}}

{{- define "GenR2" -}}
{{ range $i, $a := .Args }}T{{$i}}, {{end}}R0, R1 any
{{- end}}
