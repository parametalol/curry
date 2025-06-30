
{{- if ne .N 1}}
// region Reverse{{.N}}

// Reverse{{.N}}R the order of the parameters of a function returning one value.
func Reverse{{.N}}R[{{template "GenR" .}}](fn func({{ range $i, $a := .Args }}{{if $i}}, {{end}}T{{$a.I}}{{end}}) R) func({{ range $i, $a := .RArgs }}{{if $i}}, {{end}}T{{$a.I}}{{end}}) R {
	return func({{ range $i, $a := .RArgs }}{{if $i}}, {{end}}{{$a.Name}} T{{$a.I}}{{end}}) R {
		return fn({{ range $i, $a := .Args }}{{if $i}}, {{end}}{{$a.Name}}{{end}})
	}
}

// Reverse{{.N}}R2 the order of the parameters of a function returning two values.
func Reverse{{.N}}R2[{{template "GenR2" .}}](fn func({{ range $i, $a := .Args }}{{if $i}}, {{end}}T{{$a.I}}{{end}}) (R0, R1)) func({{ range $i, $a := .RArgs }}{{if $i}}, {{end}}T{{$a.I}}{{end}}) (R0, R1) {
	return func({{ range $i, $a := .RArgs }}{{if $i}}, {{end}}{{$a.Name}} T{{$a.I}}{{end}}) (R0, R1) {
		return fn({{ range $i, $a := .Args }}{{if $i}}, {{end}}{{$a.Name}}{{end}})
	}
}

// Reverse{{.N}} the order of the parameters of a function returning no values.
func Reverse{{.N}}[{{template "Gen" .}}](fn func({{ range $i, $a := .Args }}{{if $i}}, {{end}}T{{$a.I}}{{end}})) func({{ range $i, $a := .RArgs }}{{if $i}}, {{end}}T{{$a.I}}{{end}}) {
	return func({{ range $i, $a := .RArgs }}{{if $i}}, {{end}}{{$a.Name}} T{{$a.I}}{{end}}) {
		fn({{ range $i, $a := .Args }}{{if $i}}, {{end}}{{$a.Name}}{{end}})
	}
}
{{end -}}
