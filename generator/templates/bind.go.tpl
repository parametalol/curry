
// region: Bind{{if ne $.N 1}}FirstOf{{end}}{{.N}}

// Bind{{if ne $.N 1}}FirstOf{{end}}{{.N}}R binds the first argument of a function of {{.N}}, returning one value.
func Bind{{if ne $.N 1}}FirstOf{{end}}{{.N}}R[{{ range $i, $x := .Args }}T{{$i}}, {{end}}R any](fn func({{ range $i, $x := .Args }}{{if ne $i 0}}, {{end}}T{{$i}}{{end}}) R, arg0 T0) func({{ range $i, $x := .Args }}{{if gt $i 0}}{{if ne $i 1}}, {{end}}T{{$i}}{{end}}{{end}}) R {
	return func({{ range $i, $x := .Args }}{{if gt $i 0}}{{if ne $i 1}}, {{end}}{{$x.Name}} T{{$i}}{{end}}{{end}}) R {
		return fn(arg0{{ range $i, $x := .Args }}{{if gt $i 0}}, {{$x.Name}}{{end}}{{end}})
	}
}

// Bind{{if ne $.N 1}}FirstOf{{end}}{{.N}} binds the first argument of a function of {{.N}}, returning no values.
func Bind{{if ne $.N 1}}FirstOf{{end}}{{.N}}[{{ range $i, $x := .Args }}{{if ne $i 0}}, {{end}}T{{$i}}{{end}} any](fn func({{ range $i, $x := .Args }}{{if ne $i 0}}, {{end}}T{{$i}}{{end}}), arg0 T0) func({{ range $i, $x := .Args }}{{if gt $i 0}}{{if ne $i 1}}, {{end}}T{{$i}}{{end}}{{end}}) {
	return func({{ range $i, $x := .Args }}{{if gt $i 0}}{{if ne $i 1}}, {{end}}{{$x.Name}} T{{$i}}{{end}}{{end}}) {
		fn(arg0{{ range $i, $x := .Args }}{{if gt $i 0}}, {{$x.Name}}{{end}}{{end}})
	}
}

// Bind{{if ne $.N 1}}FirstOf{{end}}{{.N}}R2 binds the first argument of a function of {{.N}}, returning two values.
func Bind{{if ne $.N 1}}FirstOf{{end}}{{.N}}R2[{{ range $i, $x := .Args }}T{{$i}}, {{end}}R0, R1 any](fn func({{ range $i, $x := .Args }}{{if ne $i 0}}, {{end}}T{{$i}}{{end}}) (R0, R1), arg0 T0) func({{ range $i, $x := .Args }}{{if gt $i 0}}{{if ne $i 1}}, {{end}}T{{$i}}{{end}}{{end}}) (R0, R1) {
	return func({{ range $i, $x := .Args }}{{if gt $i 0}}{{if ne $i 1}}, {{end}}{{$x.Name}} T{{$i}}{{end}}{{end}}) (R0, R1) {
		return fn(arg0{{ range $i, $x := .Args }}{{if gt $i 0}}, {{$x.Name}}{{end}}{{end}})
	}
}

// region: Bind{{if ne $.N 1}}FirstOf{{end}}{{.N}}S

// Bind{{if ne $.N 1}}FirstOf{{end}}{{.N}}SR binds the first argument of a function of {{.N}}, returning one value.
func Bind{{if ne $.N 1}}FirstOf{{end}}{{.N}}SR[
{{- range $i, $x := .Args }}T{{$i}}, {{ end}}R any](fn func(
	{{- range $i, $x := .Args }}{{if ne $i 0}}, {{end}}{{if ne $i $.Last}}T{{$i}}{{end}}
	{{- end}}...T{{.Last}}) R, arg0 {{if eq $.Last 0}}...{{end}}T0) func({{ range $i, $x := .Args }}
{{- if ne $i 0}}{{- if ne $i 1}}, {{end}}{{if eq $i $.Last}}...{{end}}T{{$i}}{{end}}
{{- end}}) R {
	return func({{ range $i, $x := .Args }}
		{{- if gt $i 0}}
			{{- if ne $i 1}}, {{end}}{{$x.Name}} {{if eq $i $.Last}}...{{end}}T{{$i}}
		{{- end}}
	{{- end}}) R {
		return fn(arg0{{ range $i, $x := .Args }}{{if gt $i 0}}, {{$x.Name}}{{end}}{{if eq $i $.Last}}...{{end}}{{end}})
	}
}

// Bind{{if ne $.N 1}}FirstOf{{end}}{{.N}}S binds the first argument of a function of {{.N}}, returning no values.
func Bind{{if ne $.N 1}}FirstOf{{end}}{{.N}}S[
{{- range $i, $x := .Args }}{{if ne $i 0}}, {{end}}T{{$i}}{{ end}} any](fn func(
	{{- range $i, $x := .Args }}{{if ne $i 0}}, {{end}}{{if ne $i $.Last}}T{{$i}}{{end}}
	{{- end}}...T{{.Last}}), arg0 {{if eq $.Last 0}}...{{end}}T0) func({{ range $i, $x := .Args }}
{{- if ne $i 0}}{{- if ne $i 1}}, {{end}}{{if eq $i $.Last}}...{{end}}T{{$i}}{{end}}
{{- end}}) {
	return func({{ range $i, $x := .Args }}
		{{- if gt $i 0}}
			{{- if ne $i 1}}, {{end}}{{$x.Name}} {{if eq $i $.Last}}...{{end}}T{{$i}}
		{{- end}}
	{{- end}}) {
		fn(arg0{{ range $i, $x := .Args }}{{if gt $i 0}}, {{$x.Name}}{{end}}{{if eq $i $.Last}}...{{end}}{{end}})
	}
}

// Bind{{if ne $.N 1}}FirstOf{{end}}{{.N}}SR2 binds the first argument of a function of {{.N}}, returning one value.
func Bind{{if ne $.N 1}}FirstOf{{end}}{{.N}}SR2[
{{- range $i, $x := .Args }}T{{$i}}, {{ end}}R0, R1 any](fn func(
	{{- range $i, $x := .Args }}{{if ne $i 0}}, {{end}}{{if ne $i $.Last}}T{{$i}}{{end}}
	{{- end}}...T{{.Last}}) (R0, R1), arg0 {{if eq $.Last 0}}...{{end}}T0) func({{ range $i, $x := .Args }}
{{- if ne $i 0}}{{- if ne $i 1}}, {{end}}{{if eq $i $.Last}}...{{end}}T{{$i}}{{end}}
{{- end}}) (R0, R1) {
	return func({{ range $i, $x := .Args }}
		{{- if gt $i 0}}
			{{- if ne $i 1}}, {{end}}{{$x.Name}} {{if eq $i $.Last}}...{{end}}T{{$i}}
		{{- end}}
	{{- end}}) (R0, R1) {
		return fn(arg0{{ range $i, $x := .Args }}{{if gt $i 0}}, {{$x.Name}}{{end}}{{if eq $i $.Last}}...{{end}}{{end}})
	}
}

{{- if ne $.N 1}}

// region: BindLastOf{{.N}}

// BindLastOf{{.N}}R binds the last argument of a function of {{.N}}, returning one value.
func BindLastOf{{.N}}R[{{ range $i, $x := .Args }}T{{$i}}, {{end}}R any](fn func({{range $i, $x := .Args}}
			{{- if ne $i 0}}, {{end}}T{{$i}}
	{{- end -}}) R, arg{{.Last}} T{{.Last}}) func({{ range $i, $x := .Args }}{{if ne $i $.Last}}{{if ne $i 0}}, {{end}}T{{$i}}{{end}}{{end}}) R {
	return func({{ range $i, $x := .Args }}{{if ne $i $.Last}}{{if ne $i 0}}, {{end}}{{$x.Name}} T{{$i}}{{end}}{{end}}) R {
		return fn({{ range $i, $x := .Args }}{{if gt $i 0}}, {{end}}{{$x.Name}}{{end}})
	}
}

// BindLastOf{{.N}} binds the last argument of a function of {{.N}}, returning no values.
func BindLastOf{{.N}}[{{ range $i, $x := .Args }}{{if ne $i 0}}, {{end}}T{{$i}}{{end}} any](fn func({{range $i, $x := .Args}}
			{{- if ne $i 0}}, {{end}}T{{$i}}
	{{- end -}}), arg{{.Last}} T{{.Last}}) func({{ range $i, $x := .Args }}{{if ne $i $.Last}}{{if ne $i 0}}, {{end}}T{{$i}}{{end}}{{end}}) {
	return func({{ range $i, $x := .Args }}{{if ne $i $.Last}}{{if ne $i 0}}, {{end}}{{$x.Name}} T{{$i}}{{end}}{{end}}) {
		fn({{ range $i, $x := .Args }}{{if gt $i 0}}, {{end}}{{$x.Name}}{{end}})
	}
}

// BindLastOf{{.N}}R2 binds the last ellipsis argument of a function of {{.N}}, returning two values.
func BindLastOf{{.N}}R2[{{ range $i, $x := .Args }}T{{$i}}, {{end}}R0, R1 any](fn func({{range $i, $x := .Args}}
			{{- if ne $i 0}}, {{end}}T{{$i}}
	{{- end -}}) (R0, R1), arg{{.Last}} T{{.Last}}) func({{ range $i, $x := .Args }}{{if ne $i $.Last}}{{if ne $i 0}}, {{end}}T{{$i}}{{end}}{{end}}) (R0, R1) {
	return func({{ range $i, $x := .Args }}{{if ne $i $.Last}}{{if ne $i 0}}, {{end}}{{$x.Name}} T{{$i}}{{end}}{{end}}) (R0, R1) {
		return fn({{ range $i, $x := .Args }}{{if gt $i 0}}, {{end}}{{$x.Name}}{{end}})
	}
}

// region: BindLastOf{{.N}}S

// BindLastOf{{.N}}SR binds the last ellipsis argument of a function of {{.N}}, returning one value.
func BindLastOf{{.N}}SR[{{ range $i, $x := .Args }}T{{$i}}, {{end}}R any](fn func({{range $i, $x := .Args -}}{{- if ne $i $.Last -}}
			{{- if ne $i 0}}, {{end}}T{{$i}}
		{{- end -}}
	{{- end -}}{{- if ne $.Last 0}}, {{end}}...T{{.Last}}) R, arg{{.Last}} ...T{{.Last}}) func({{ range $i, $x := .Args }}{{if ne $i $.Last}}{{if ne $i 0}}, {{end}}T{{$i}}{{end}}{{end}}) R {
	return func({{ range $i, $x := .Args }}{{if ne $i $.Last}}{{if ne $i 0}}, {{end}}{{$x.Name}} T{{$i}}{{end}}{{end}}) R {
		return fn({{ range $i, $x := .Args }}{{if gt $i 0}}, {{end}}{{$x.Name}}{{end}}...)
	}
}

// BindLastOf{{.N}}S binds the last ellipsis argument of a function of {{.N}}, returning no values.
func BindLastOf{{.N}}S[{{ range $i, $x := .Args }}{{if ne $i 0}}, {{end}}T{{$i}}{{end}} any](fn func({{- range $i, $x := .Args -}}{{- if ne $i $.Last -}}
			{{- if ne $i 0}}, {{end}}T{{$i}}
		{{- end -}}
	{{- end -}}{{- if ne $.Last 0}}, {{end}}...T{{.Last}}), arg{{.Last}} ...T{{.Last}}) func({{ range $i, $x := .Args }}{{if ne $i $.Last}}{{if ne $i 0}}, {{end}}T{{$i}}{{end}}{{end}}) {
	return func({{ range $i, $x := .Args }}{{if ne $i $.Last}}{{if ne $i 0}}, {{end}}{{$x.Name}} T{{$i}}{{end}}{{end}}) {
		fn({{ range $i, $x := .Args }}{{if gt $i 0}}, {{end}}{{$x.Name}}{{end}}...)
	}
}

// BindLastOf{{.N}}SR2 binds the last ellipsis argument of a function of {{.N}}, returning two values.
func BindLastOf{{.N}}SR2[{{ range $i, $x := .Args }}T{{$i}}, {{end}}R0, R1 any](fn func({{- range $i, $x := .Args -}}{{- if ne $i $.Last -}}
			{{- if ne $i 0}}, {{end}}T{{$i}}
		{{- end -}}
	{{- end -}}{{- if ne $.Last 0}}, {{end}}...T{{.Last}}) (R0, R1), arg{{.Last}} ...T{{.Last}}) func({{ range $i, $x := .Args }}{{if ne $i $.Last}}{{if ne $i 0}}, {{end}}T{{$i}}{{end}}{{end}}) (R0, R1) {
	return func({{ range $i, $x := .Args }}{{if ne $i $.Last}}{{if ne $i 0}}, {{end}}{{$x.Name}} T{{$i}}{{end}}{{end}}) (R0, R1) {
		return fn({{ range $i, $x := .Args }}{{if gt $i 0}}, {{end}}{{$x.Name}}{{end}}...)
	}
}
{{- end}}
