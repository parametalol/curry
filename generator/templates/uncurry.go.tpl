{{- if ne $.N 1}}
// region {{.N}}

// Un{{.N}}R transforms a curried function with {{.N}} parameter(s) back into its original form.
func Un{{.N}}R[{{template "GenR" .}}](fn{{range $i, $a := .Args}} func(T{{$i}}){{end}} R) func({{template "Types" .}}) R {
	return func({{ range $i, $a := .Args }}{{if $i}}, {{end}}{{$a.Name }} T{{$i}}{{end}}) R {
		return fn{{ range $i, $a := .Args }}({{$a.Name}}){{end}}
	}
}

// Un{{.N}} transforms a curried function with {{.N}} parameter(s) back into its original form.
func Un{{.N}}[{{template "Gen" .}}](fn{{range $i, $a := .Args}} func(T{{$i}}){{end}}) func({{template "Types" .}}) {
	return func({{ range $i, $a := .Args }}{{if $i}}, {{end}}{{$a.Name }} T{{$i}}{{end}}) {
		fn{{ range $i, $a := .Args }}({{$a.Name}}){{end}}
	}
}

// Un{{.N}}R2 transforms a curried function with {{.N}} parameter(s) back into its original form.
func Un{{.N}}R2[{{template "GenR2" .}}](fn{{range $i, $a := .Args}} func(T{{$i}}){{end}} (R0, R1)) func({{template "Types" .}}) (R0, R1) {
	return func({{ range $i, $a := .Args }}{{if $i}}, {{end}}{{$a.Name }} T{{$i}}{{end}}) (R0, R1) {
		return fn{{ range $i, $a := .Args }}({{$a.Name}}){{end}}
	}
}

// Un{{.N}}SR transforms a curried function with {{.N}} parameter(s) back into its original form.
func Un{{.N}}SR[{{template "GenR" .}}](fn{{range $i, $a := .Args}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}} R) func({{template "TypesS" .}}) R {
	return func({{ range $i, $a := .Args }}{{if $i}}, {{end}}{{$a.Name }} {{if eq $i $.Last}}...{{end}}T{{$i}}{{end}}) R {
		return fn{{ range $i, $a := .Args }}({{$a.Name}}{{if eq $i $.Last}}...{{end}}){{end}}
	}
}

// Un{{.N}}S transforms a curried function with {{.N}} parameter(s) back into its original form.
func Un{{.N}}S[{{template "Gen" .}}](fn{{range $i, $a := .Args}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}}) func({{template "TypesS" .}}) {
	return func({{ range $i, $a := .Args }}{{if $i}}, {{end}}{{$a.Name }} {{if eq $i $.Last}}...{{end}}T{{$i}}{{end}}) {
		fn{{ range $i, $a := .Args }}({{$a.Name}}{{if eq $i $.Last}}...{{end}}){{end}}
	}
}

// Un{{.N}}SR2 transforms a curried function with {{.N}} parameter(s) back into its original form.
func Un{{.N}}SR2[{{template "GenR2" .}}](fn{{range $i, $a := .Args}} func({{if eq $i $.Last}}...{{end}}T{{$i}}){{end}} (R0, R1)) func({{template "TypesS" .}}) (R0, R1) {
	return func({{ range $i, $a := .Args }}{{if $i}}, {{end}}{{$a.Name }} {{if eq $i $.Last}}...{{end}}T{{$i}}{{end}}) (R0, R1) {
		return fn{{ range $i, $a := .Args }}({{$a.Name}}{{if eq $i $.Last}}...{{end}}){{end}}
	}
}
{{end}}