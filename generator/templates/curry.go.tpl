{{- if ne $.N 1}}
// region: Curry{{.N}}

// Curry{{.N}}R curries a function of {{.N}} parameter(s), returning one value.
func Curry{{.N}}R[{{template "GenR" .}}](fn func({{template "Types" .}}) R){{range $i, $a := .Args}} func(T{{$i}}){{end}} R {
	{{if eq $.N 2 -}}
	return Head{{.N}}R(fn)
	{{- else -}}
	return func(arg0 T0){{range $i, $a := .Args}}{{if $i}} func(T{{$i}}){{end}}{{end}} R {
		return Curry{{.Last}}R(Head{{.N}}R(fn)(arg0))
	}
	{{- end}}
}

// Curry{{.N}} curries a function of {{.N}} parameter(s), returning no values.
func Curry{{.N}}[{{template "Gen" .}}](fn func({{template "Types" .}})){{range $i, $a := .Args}} func(T{{$i}}){{end}} {
	{{if eq $.N 2 -}}
	return Head{{.N}}(fn)
	{{- else -}}
	return func(arg0 T0){{range $i, $a := .Args}}{{if $i}} func(T{{$i}}){{end}}{{end}} {
		return Curry{{.Last}}(Head{{.N}}(fn)(arg0))
	}
	{{- end}}
}

// Curry{{.N}}R2 curries a function of {{.N}} parameter(s), returning two values.
func Curry{{.N}}R2[{{template "GenR2" .}}](fn func({{template "Types" .}}) (R0, R1)){{range $i, $a := .Args}} func(T{{$i}}){{end}} (R0, R1) {
	{{if eq $.N 2 -}}
	return Head{{.N}}R2(fn)
	{{- else -}}
	return func(arg0 T0){{range $i, $a := .Args}}{{if $i}} func(T{{$i}}){{end}}{{end}} (R0, R1) {
		return Curry{{.Last}}R2(Head{{.N}}R2(fn)(arg0))
	}
	{{- end}}
}

// Curry{{.N}}SR curries a function of {{.N}} parameter(s), returning one value.
func Curry{{.N}}SR[{{template "GenR" .}}](fn func({{template "TypesS" .}}) R){{range $i, $a := .Args}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}} R {
	{{if eq $.N 2 -}}
	return Head{{.N}}SR(fn)
	{{- else -}}
	return func(arg0 T0){{range $i, $a := .Args}}{{if $i}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}}{{end}} R {
		return Curry{{.Last}}SR(Head{{.N}}SR(fn)(arg0))
	}
	{{- end}}
}

// Curry{{.N}}S curries a function of {{.N}} parameter(s), returning no values.
func Curry{{.N}}S[{{template "Gen" .}}](fn func({{template "TypesS" .}})){{range $i, $a := .Args}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}} {
	{{if eq $.N 2 -}}
	return Head{{.N}}S(fn)
	{{- else -}}
	return func(arg0 T0){{range $i, $a := .Args}}{{if $i}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}}{{end}} {
		return Curry{{.Last}}S(Head{{.N}}S(fn)(arg0))
	}
	{{- end}}
}

// Curry{{.N}}SR2 curries a function of {{.N}} parameter(s), returning two values.
func Curry{{.N}}SR2[{{template "GenR2" .}}](fn func({{template "TypesS" .}}) (R0, R1)){{range $i, $a := .Args}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}} (R0, R1) {
	{{if eq $.N 2 -}}
	return Head{{.N}}SR2(fn)
	{{- else -}}
	return func(arg0 T0){{range $i, $a := .Args}}{{if $i}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}}{{end}} (R0, R1) {
		return Curry{{.Last}}SR2(Head{{.N}}SR2(fn)(arg0))
	}
	{{- end}}
}
{{end -}}
