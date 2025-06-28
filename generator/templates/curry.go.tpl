{{- if ne $.N 1}}
// region: F{{.N}}

// F{{.N}}R curries a function of {{.N}} parameter(s), returning one value.
func F{{.N}}R[{{template "GenR" .}}](fn func({{template "Types" .}}) R){{range $i, $a := .Args}} func(T{{$i}}){{end}} R {
	{{if eq $.N 2 -}}
	return Head{{.N}}R(fn)
	{{- else -}}
	return func(arg0 T0){{range $i, $a := .Args}}{{if $i}} func(T{{$i}}){{end}}{{end}} R {
		return F{{.Last}}R(Head{{.N}}R(fn)(arg0))
	}
	{{- end}}
}

// F{{.N}} curries a function of {{.N}} parameter(s), returning no values.
func F{{.N}}[{{template "Gen" .}}](fn func({{template "Types" .}})){{range $i, $a := .Args}} func(T{{$i}}){{end}} {
	{{if eq $.N 2 -}}
	return Head{{.N}}(fn)
	{{- else -}}
	return func(arg0 T0){{range $i, $a := .Args}}{{if $i}} func(T{{$i}}){{end}}{{end}} {
		return F{{.Last}}(Head{{.N}}(fn)(arg0))
	}
	{{- end}}
}

// F{{.N}}R2 curries a function of {{.N}} parameter(s), returning two values.
func F{{.N}}R2[{{template "GenR2" .}}](fn func({{template "Types" .}}) (R0, R1)){{range $i, $a := .Args}} func(T{{$i}}){{end}} (R0, R1) {
	{{if eq $.N 2 -}}
	return Head{{.N}}R2(fn)
	{{- else -}}
	return func(arg0 T0){{range $i, $a := .Args}}{{if $i}} func(T{{$i}}){{end}}{{end}} (R0, R1) {
		return F{{.Last}}R2(Head{{.N}}R2(fn)(arg0))
	}
	{{- end}}
}

// F{{.N}}SR curries a function of {{.N}} parameter(s), returning one value.
func F{{.N}}SR[{{template "GenR" .}}](fn func({{template "TypesS" .}}) R){{range $i, $a := .Args}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}} R {
	{{if eq $.N 2 -}}
	return Head{{.N}}SR(fn)
	{{- else -}}
	return func(arg0 T0){{range $i, $a := .Args}}{{if $i}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}}{{end}} R {
		return F{{.Last}}SR(Head{{.N}}SR(fn)(arg0))
	}
	{{- end}}
}

// F{{.N}}S curries a function of {{.N}} parameter(s), returning no values.
func F{{.N}}S[{{template "Gen" .}}](fn func({{template "TypesS" .}})){{range $i, $a := .Args}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}} {
	{{if eq $.N 2 -}}
	return Head{{.N}}S(fn)
	{{- else -}}
	return func(arg0 T0){{range $i, $a := .Args}}{{if $i}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}}{{end}} {
		return F{{.Last}}S(Head{{.N}}S(fn)(arg0))
	}
	{{- end}}
}

// F{{.N}}SR2 curries a function of {{.N}} parameter(s), returning two values.
func F{{.N}}SR2[{{template "GenR2" .}}](fn func({{template "TypesS" .}}) (R0, R1)){{range $i, $a := .Args}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}} (R0, R1) {
	{{if eq $.N 2 -}}
	return Head{{.N}}SR2(fn)
	{{- else -}}
	return func(arg0 T0){{range $i, $a := .Args}}{{if $i}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}}{{end}} (R0, R1) {
		return F{{.Last}}SR2(Head{{.N}}SR2(fn)(arg0))
	}
	{{- end}}
}
{{end -}}
