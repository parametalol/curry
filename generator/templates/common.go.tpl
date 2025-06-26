{{- define "Types" -}}
{{ range $i, $a := .Args }}{{if ne $i 0}}, {{end}}T{{$i}}{{end}}
{{- end}}

{{- define "Args" -}}
{{ range $i, $a := .Args }}{{if ne $i 0}}, {{end}}{{$a.Name}}{{end}}
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
