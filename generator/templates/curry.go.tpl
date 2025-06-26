{{- define "Tail" -}}
{{range $i, $a := .Args}}{{if ne $i 0}}{{if ne $i 1}}, {{end}}T{{$i}}{{end}}{{end}}
{{- end -}}

{{if ne $.N 1}}

func Head{{.N}}R[{{template "GenR" .}}](fn func({{template "Types" .}}) R) func(T0) func({{template "Tail" .}}) R {
	return func(arg0 T0) func({{template "Tail" .}}) R {
		return func({{range $i, $a := .Args}}{{if ne $i 0}}{{if ne $i 1}}, {{end}}{{$a.Name}} T{{$i}}{{end}}{{end}}) R {
			return fn({{template "Args" .}})
		}
	}
}


func Curry{{.N}}R[{{template "GenR" .}}](fn func({{template "Types" .}}) R){{range $i, $a := .Args}} func(T{{$i}}){{end}} R {
	return func(arg0 T0){{range $i, $a := .Args}}{{if ne $i 0}} func(T{{$i}}){{end}}{{end}} R {
		return {{range $i, $a := .Args -}}
        {{if gt $i 1}}Head{{$i}}R({{end}}
        {{- end -}}
        {{if gt $.N 1}}Head{{$.N}}R({{end -}}
            fn)(arg0){{range $i, $a := .Args}}{{if gt $i 1}}){{end}}{{end}}
	}
}

{{end}}
