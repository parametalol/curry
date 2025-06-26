{{- if ne .N 1 -}}
// region: Reverse{{.N}}

// Reverse{{.N}}R the order of the parameters of a function returning one value.
func Reverse{{.N}}R[{{ range $i, $x := .Args }}T{{$i}}, {{end}}R any](fn func({{ range $i, $a := .Args }}{{if ne $i 0}}, {{end}}T{{$a.I}}{{end}}) R) func({{ range $i, $a := .RArgs }}{{if ne $i 0}}, {{end}}T{{$a.I}}{{end}}) R {
    return func({{ range $i, $a := .RArgs }}{{if ne $i 0}}, {{end}}{{$a.Name}} T{{$a.I}}{{end}}) R {
        return fn({{ range $i, $a := .Args }}{{if ne $i 0}}, {{end}}{{$a.Name}}{{end}})
    }
}

// Reverse{{.N}}R2 the order of the parameters of a function returning two values.
func Reverse{{.N}}R2[{{ range $i, $x := .Args }}T{{$i}}, {{end}}R1, R2 any](fn func({{ range $i, $a := .Args }}{{if ne $i 0}}, {{end}}T{{$a.I}}{{end}}) (R1, R2)) func({{ range $i, $a := .RArgs }}{{if ne $i 0}}, {{end}}T{{$a.I}}{{end}}) (R1, R2) {
    return func({{ range $i, $a := .RArgs }}{{if ne $i 0}}, {{end}}{{$a.Name}} T{{$a.I}}{{end}}) (R1, R2) {
        return fn({{ range $i, $a := .Args }}{{if ne $i 0}}, {{end}}{{$a.Name}}{{end}})
    }
}

// Reverse{{.N}} the order of the parameters of a function returning no values.
func Reverse{{.N}}[{{ range $i, $x := .Args }}{{if ne $i 0}}, {{end}}T{{$i}}{{end}} any](fn func({{ range $i, $a := .Args }}{{if ne $i 0}}, {{end}}T{{$a.I}}{{end}})) func({{ range $i, $a := .RArgs }}{{if ne $i 0}}, {{end}}T{{$a.I}}{{end}}) {
    return func({{ range $i, $a := .RArgs }}{{if ne $i 0}}, {{end}}{{$a.Name}} T{{$a.I}}{{end}}) {
        fn({{ range $i, $a := .Args }}{{if ne $i 0}}, {{end}}{{$a.Name}}{{end}})
    }
}
{{ end }}
