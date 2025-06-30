
// region {{.N}}

// Lazy{{.N}}R converts function parameters to thunk functions.
func Lazy{{.N}}R[{{template "GenR" .}}](fn func({{template "Types" .}}) R) func({{range $i, $a := .Args}}{{if $i}}, {{end}}func() T{{$i}}{{end}}) R {
	return func({{range $i, $a := .Args}}{{if $i}}, {{end}}{{$a.Name}} func() T{{$a.I}}{{end}}) R {
		return fn({{range $i, $a := .Args}}{{if $i}}, {{end}}{{$a.Name}}(){{end}})
	}
}

// Lazy{{.N}} converts function parameters to thunk functions.
func Lazy{{.N}}[{{template "Gen" .}}](fn func({{template "Types" .}})) func({{range $i, $a := .Args}}{{if $i}}, {{end}}func() T{{$i}}{{end}}) {
	return func({{range $i, $a := .Args}}{{if $i}}, {{end}}{{$a.Name}} func() T{{$a.I}}{{end}}) {
		fn({{range $i, $a := .Args}}{{if $i}}, {{end}}{{$a.Name}}(){{end}})
	}
}

// Lazy{{.N}}R2 converts function parameters to thunk functions.
func Lazy{{.N}}R2[{{template "GenR2" .}}](fn func({{template "Types" .}}) (R0, R1)) func({{range $i, $a := .Args}}{{if $i}}, {{end}}func() T{{$i}}{{end}}) (R0, R1) {
	return func({{range $i, $a := .Args}}{{if $i}}, {{end}}{{$a.Name}} func() T{{$a.I}}{{end}}) (R0, R1) {
		return fn({{range $i, $a := .Args}}{{if $i}}, {{end}}{{$a.Name}}(){{end}})
	}
}

// Lazy{{.N}}SR converts function parameters to thunk functions.
func Lazy{{.N}}SR[{{template "GenR" .}}](fn func({{template "TypesS" .}}) R) func({{range $i, $a := .Args}}{{if $i}}, {{end}}func() {{if eq $i $.Last}}[]{{end}}T{{$i}}{{end}}) R {
	return func({{range $i, $a := .Args}}{{if $i}}, {{end}}{{$a.Name}} func() {{if eq $i $.Last}}[]{{end}}T{{$a.I}}{{end}}) R {
		return fn({{range $i, $a := .Args}}{{if $i}}, {{end}}{{$a.Name}}(){{if eq $i $.Last}}...{{end}}{{end}})
	}
}

// Lazy{{.N}}S converts function parameters to thunk functions.
func Lazy{{.N}}S[{{template "Gen" .}}](fn func({{template "TypesS" .}})) func({{range $i, $a := .Args}}{{if $i}}, {{end}}func() {{if eq $i $.Last}}[]{{end}}T{{$i}}{{end}}) {
	return func({{range $i, $a := .Args}}{{if $i}}, {{end}}{{$a.Name}} func() {{if eq $i $.Last}}[]{{end}}T{{$a.I}}{{end}}) {
		fn({{range $i, $a := .Args}}{{if $i}}, {{end}}{{$a.Name}}(){{if eq $i $.Last}}...{{end}}{{end}})
	}
}

// Lazy{{.N}}SR2 converts function parameters to thunk functions.
func Lazy{{.N}}SR2[{{template "GenR2" .}}](fn func({{template "TypesS" .}}) (R0, R1)) func({{range $i, $a := .Args}}{{if $i}}, {{end}}func() {{if eq $i $.Last}}[]{{end}}T{{$i}}{{end}}) (R0, R1) {
	return func({{range $i, $a := .Args}}{{if $i}}, {{end}}{{$a.Name}} func() {{if eq $i $.Last}}[]{{end}}T{{$a.I}}{{end}}) (R0, R1) {
		return fn({{range $i, $a := .Args}}{{if $i}}, {{end}}{{$a.Name}}(){{if eq $i $.Last}}...{{end}}{{end}})
	}
}
