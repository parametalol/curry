{{- define "Tail"}}{{range $i, $a := .Args}}{{if $i}}{{if ne $i 1}}, {{end}}T{{$i}}{{end}}{{end}}{{end}}
{{- define "TailS"}}{{range $i, $a := .Args}}{{if $i}}{{if ne $i 1}}, {{end}}{{if eq $i $.Last}}...{{end}}T{{$i}}{{end}}{{end}}{{end}}

{{- if ne $.N 1}}
// region: Head{{.N}}

// Head{{.N}}R returns a partially curried function of {{.N}} parameter(s), returning one value.
func Head{{.N}}R[{{template "GenR" .}}](fn func({{template "Types" .}}) R) func(T0) func({{template "Tail" .}}) R {
	return func(arg0 T0) func({{template "Tail" .}}) R {
		return func({{range $i, $a := .Args}}{{if $i}}{{if ne $i 1}}, {{end}}{{$a.Name}} T{{$i}}{{end}}{{end}}) R {
			return fn({{template "Args" .}})
		}
	}
}

// Head{{.N}} returns a partially curried function of {{.N}} parameter(s), returning no values.
func Head{{.N}}[{{template "Gen" .}}](fn func({{template "Types" .}})) func(T0) func({{template "Tail" .}}) {
	return func(arg0 T0) func({{template "Tail" .}}) {
		return func({{range $i, $a := .Args}}{{if $i}}{{if ne $i 1}}, {{end}}{{$a.Name}} T{{$i}}{{end}}{{end}}) {
			fn({{template "Args" .}})
		}
	}
}

// Head{{.N}}R2 returns a partially curried function of {{.N}} parameter(s), returning one value.
func Head{{.N}}R2[{{template "GenR2" .}}](fn func({{template "Types" .}}) (R0, R1)) func(T0) func({{template "Tail" .}}) (R0, R1) {
	return func(arg0 T0) func({{template "Tail" .}}) (R0, R1) {
		return func({{range $i, $a := .Args}}{{if $i}}{{if ne $i 1}}, {{end}}{{$a.Name}} T{{$i}}{{end}}{{end}}) (R0, R1) {
			return fn({{template "Args" .}})
		}
	}
}

// Head{{.N}}SR returns a partially curried function of {{.N}} parameter(s), returning one value.
func Head{{.N}}SR[{{template "GenR" .}}](fn func({{template "TypesS" .}}) R) func(T0) func({{template "TailS" .}}) R {
	return func(arg0 T0) func({{template "TailS" .}}) R {
		return func({{range $i, $a := .Args}}{{if $i}}{{if ne $i 1}}, {{end}}{{$a.Name}} {{if eq $i $.Last}}...{{end}}T{{$i}}{{end}}{{end}}) R {
			return fn({{template "ArgsS" .}})
		}
	}
}

// Head{{.N}}S returns a partially curried function of {{.N}} parameter(s), returning no values.
func Head{{.N}}S[{{template "Gen" .}}](fn func({{template "TypesS" .}})) func(T0) func({{template "TailS" .}}) {
	return func(arg0 T0) func({{template "TailS" .}}) {
		return func({{range $i, $a := .Args}}{{if $i}}{{if ne $i 1}}, {{end}}{{$a.Name}} {{if eq $i $.Last}}...{{end}}T{{$i}}{{end}}{{end}}) {
			fn({{template "ArgsS" .}})
		}
	}
}

// Head{{.N}}SR2 returns a partially curried function of {{.N}} parameter(s), returning two values.
func Head{{.N}}SR2[{{template "GenR2" .}}](fn func({{template "TypesS" .}}) (R0, R1)) func(T0) func({{template "TailS" .}}) (R0, R1) {
	return func(arg0 T0) func({{template "TailS" .}}) (R0, R1) {
		return func({{range $i, $a := .Args}}{{if $i}}{{if ne $i 1}}, {{end}}{{$a.Name}} {{if eq $i $.Last}}...{{end}}T{{$i}}{{end}}{{end}}) (R0, R1) {
			return fn({{template "ArgsS" .}})
		}
	}
}

// region: Curry{{.N}}

// Curry{{.N}}R curries a function of {{.N}} parameter(s), returning one value.
func Curry{{.N}}R[{{template "GenR" .}}](fn func({{template "Types" .}}) R){{range $i, $a := .Args}} func(T{{$i}}){{end}} R {
	return func(arg0 T0){{range $i, $a := .Args}}{{if $i}} func(T{{$i}}){{end}}{{end}} R {
		{{if eq $.N 2 -}}
		return Head{{.N}}R(fn)(arg0)
		{{- else -}}
		return Curry{{.Last}}R(Head{{.N}}R(fn)(arg0))
		{{- end}}
	}
}

// Curry{{.N}} curries a function of {{.N}} parameter(s), returning no values.
func Curry{{.N}}[{{template "Gen" .}}](fn func({{template "Types" .}})){{range $i, $a := .Args}} func(T{{$i}}){{end}} {
	return func(arg0 T0){{range $i, $a := .Args}}{{if $i}} func(T{{$i}}){{end}}{{end}} {
		{{if eq $.N 2 -}}
		return Head{{.N}}(fn)(arg0)
		{{- else -}}
		return Curry{{.Last}}(Head{{.N}}(fn)(arg0))
		{{- end}}
	}
}

// Curry{{.N}}R2 curries a function of {{.N}} parameter(s), returning two values.
func Curry{{.N}}R2[{{template "GenR2" .}}](fn func({{template "Types" .}}) (R0, R1)){{range $i, $a := .Args}} func(T{{$i}}){{end}} (R0, R1) {
	return func(arg0 T0){{range $i, $a := .Args}}{{if $i}} func(T{{$i}}){{end}}{{end}} (R0, R1) {
		{{if eq $.N 2 -}}
		return Head{{.N}}R2(fn)(arg0)
		{{- else -}}
		return Curry{{.Last}}R2(Head{{.N}}R2(fn)(arg0))
		{{- end}}
	}
}

// Curry{{.N}}SR curries a function of {{.N}} parameter(s), returning one value.
func Curry{{.N}}SR[{{template "GenR" .}}](fn func({{template "TypesS" .}}) R){{range $i, $a := .Args}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}} R {
	return func(arg0 T0){{range $i, $a := .Args}}{{if $i}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}}{{end}} R {
		{{if eq $.N 2 -}}
		return Head{{.N}}SR(fn)(arg0)
		{{- else -}}
		return Curry{{.Last}}SR(Head{{.N}}SR(fn)(arg0))
		{{- end}}
	}
}

// Curry{{.N}}S curries a function of {{.N}} parameter(s), returning no values.
func Curry{{.N}}S[{{template "Gen" .}}](fn func({{template "TypesS" .}})){{range $i, $a := .Args}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}} {
	return func(arg0 T0){{range $i, $a := .Args}}{{if $i}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}}{{end}} {
		{{if eq $.N 2 -}}
		return Head{{.N}}S(fn)(arg0)
		{{- else -}}
		return Curry{{.Last}}S(Head{{.N}}S(fn)(arg0))
		{{- end}}
	}
}

// Curry{{.N}}SR2 curries a function of {{.N}} parameter(s), returning two values.
func Curry{{.N}}SR2[{{template "GenR2" .}}](fn func({{template "TypesS" .}}) (R0, R1)){{range $i, $a := .Args}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}} (R0, R1) {
	return func(arg0 T0){{range $i, $a := .Args}}{{if $i}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}}{{end}} (R0, R1) {
		{{if eq $.N 2 -}}
		return Head{{.N}}SR2(fn)(arg0)
		{{- else -}}
		return Curry{{.Last}}SR2(Head{{.N}}SR2(fn)(arg0))
		{{- end}}
	}
}
{{end -}}
