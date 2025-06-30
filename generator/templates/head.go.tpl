{{- define "Tail"}}{{if $.Last}} func({{range $i, $a := .Args}}{{if $i}}{{if ne $i 1}}, {{end}}T{{$i}}{{end}}{{end}}){{end}}{{end}}
{{- define "TailS"}}{{if $.Last}} func({{range $i, $a := .Args}}{{if $i}}{{if ne $i 1}}, {{end}}{{if eq $i $.Last}}...{{end}}T{{$i}}{{end}}{{end}}){{end}}{{end}}

{{- if ne $.N 1}}
// region Head{{.N}}

// Head{{.N}}R returns a partially curried function of {{.N}} parameter(s), returning one value.
func Head{{.N}}R[{{template "GenR" .}}](fn func({{template "Types" .}}) R) func(T0) {{- template "Tail" .}} R {
	return func(arg0 T0) {{- template "Tail" .}} R {
		return func({{range $i, $a := .Args}}{{if $i}}{{if ne $i 1}}, {{end}}{{$a.Name}} T{{$i}}{{end}}{{end}}) R {
			return fn({{template "Args" .}})
		}
	}
}

// Head{{.N}} returns a partially curried function of {{.N}} parameter(s), returning no values.
func Head{{.N}}[{{template "Gen" .}}](fn func({{template "Types" .}})) func(T0) {{- template "Tail" .}} {
	return func(arg0 T0) {{- template "Tail" .}} {
		return func({{range $i, $a := .Args}}{{if $i}}{{if ne $i 1}}, {{end}}{{$a.Name}} T{{$i}}{{end}}{{end}}) {
			fn({{template "Args" .}})
		}
	}
}

// Head{{.N}}R2 returns a partially curried function of {{.N}} parameter(s), returning one value.
func Head{{.N}}R2[{{template "GenR2" .}}](fn func({{template "Types" .}}) (R0, R1)) func(T0) {{- template "Tail" .}} (R0, R1) {
	return func(arg0 T0) {{- template "Tail" .}} (R0, R1) {
		return func({{range $i, $a := .Args}}{{if $i}}{{if ne $i 1}}, {{end}}{{$a.Name}} T{{$i}}{{end}}{{end}}) (R0, R1) {
			return fn({{template "Args" .}})
		}
	}
}

// Head{{.N}}SR returns a partially curried function of {{.N}} parameter(s), returning one value.
func Head{{.N}}SR[{{template "GenR" .}}](fn func({{template "TypesS" .}}) R) func(T0) {{- template "TailS" .}} R {
	return func(arg0 T0) {{- template "TailS" .}} R {
		return func({{range $i, $a := .Args}}{{if $i}}{{if ne $i 1}}, {{end}}{{$a.Name}} {{if eq $i $.Last}}...{{end}}T{{$i}}{{end}}{{end}}) R {
			return fn({{template "ArgsS" .}})
		}
	}
}

// Head{{.N}}S returns a partially curried function of {{.N}} parameter(s), returning no values.
func Head{{.N}}S[{{template "Gen" .}}](fn func({{template "TypesS" .}})) func(T0) {{- template "TailS" .}} {
	return func(arg0 T0) {{- template "TailS" .}} {
		return func({{range $i, $a := .Args}}{{if $i}}{{if ne $i 1}}, {{end}}{{$a.Name}} {{if eq $i $.Last}}...{{end}}T{{$i}}{{end}}{{end}}) {
			fn({{template "ArgsS" .}})
		}
	}
}

// Head{{.N}}SR2 returns a partially curried function of {{.N}} parameter(s), returning two values.
func Head{{.N}}SR2[{{template "GenR2" .}}](fn func({{template "TypesS" .}}) (R0, R1)) func(T0) {{- template "TailS" .}} (R0, R1) {
	return func(arg0 T0) {{- template "TailS" .}} (R0, R1) {
		return func({{range $i, $a := .Args}}{{if $i}}{{if ne $i 1}}, {{end}}{{$a.Name}} {{if eq $i $.Last}}...{{end}}T{{$i}}{{end}}{{end}}) (R0, R1) {
			return fn({{template "ArgsS" .}})
		}
	}
}
{{end -}}
